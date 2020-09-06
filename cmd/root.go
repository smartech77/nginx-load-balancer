package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/stephenafamo/orchestra"
	"github.com/stephenafamo/warden/internal"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(settings internal.Settings) {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "warden",
		Short: "Setup and manage a reverse proxy",
		Long:  "Setup and manage a reverse proxy",
		RunE: func(cmd *cobra.Command, args []string) error {

			log.Println("Cleaning up...")
			c := exec.Command(
				"/bin/sh",
				"-c",
				fmt.Sprintf(
					"rm -rf %s %s %s",
					settings.DB_PATH,
					filepath.Join(settings.CONFIG_OUTPUT_DIR, "/http/*"),
					filepath.Join(settings.CONFIG_OUTPUT_DIR, "/streams/*"),
				),
			)

			output, err := c.CombinedOutput()
			if err != nil {
				return fmt.Errorf(
					"Error cleaning up: %s: %s",
					err,
					output,
				)
			}

			log.Println("Connecting to DB...")
			db, err := sql.Open("sqlite3", settings.DB_PATH+"?_fk=1")
			if err != nil {
				return err
			}
			defer db.Close()

			log.Println("Creating tables...")
			err = createTables(db)
			if err != nil {
				return err
			}

			conductor := &orchestra.Conductor{
				Timeout: 15 * time.Second,
				Players: make(map[string]orchestra.Player),
			}

			hub, err := getMonitor(settings)
			if err != nil {
				return fmt.Errorf("could not get monitor: %w", err)
			}
			defer hub.Flush(time.Second * 5)

			allPlayers, err := setPlayers(db, settings, hub)
			if err != nil {
				return fmt.Errorf("could not get players: %w", err)
			}

			// Start all if no args were given
			if len(args) == 0 {
				conductor.Players = allPlayers
			}

			for _, pl := range args {
				player, ok := allPlayers[pl]
				if ok {
					conductor.Players[pl] = player
				}
			}

			return orchestra.PlayUntilSignal(
				conductor,
				os.Interrupt, syscall.SIGTERM,
			)
		},
	}

	rootCmd.Execute()
}
