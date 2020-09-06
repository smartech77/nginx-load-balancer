// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// NginxConfig is an object representing the database table.
type NginxConfig struct {
	ID           int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ServiceID    null.Int64 `boil:"service_id" json:"service_id,omitempty" toml:"service_id" yaml:"service_id,omitempty"`
	Type         string     `boil:"type" json:"type" toml:"type" yaml:"type"`
	Path         string     `boil:"path" json:"path" toml:"path" yaml:"path"`
	LastModified time.Time  `boil:"last_modified" json:"last_modified" toml:"last_modified" yaml:"last_modified"`

	R *nginxConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L nginxConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NginxConfigColumns = struct {
	ID           string
	ServiceID    string
	Type         string
	Path         string
	LastModified string
}{
	ID:           "id",
	ServiceID:    "service_id",
	Type:         "type",
	Path:         "path",
	LastModified: "last_modified",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var NginxConfigWhere = struct {
	ID           whereHelperint64
	ServiceID    whereHelpernull_Int64
	Type         whereHelperstring
	Path         whereHelperstring
	LastModified whereHelpertime_Time
}{
	ID:           whereHelperint64{field: "\"nginx_configs\".\"id\""},
	ServiceID:    whereHelpernull_Int64{field: "\"nginx_configs\".\"service_id\""},
	Type:         whereHelperstring{field: "\"nginx_configs\".\"type\""},
	Path:         whereHelperstring{field: "\"nginx_configs\".\"path\""},
	LastModified: whereHelpertime_Time{field: "\"nginx_configs\".\"last_modified\""},
}

// NginxConfigRels is where relationship names are stored.
var NginxConfigRels = struct {
	Service string
}{
	Service: "Service",
}

// nginxConfigR is where relationships are stored.
type nginxConfigR struct {
	Service *Service `boil:"Service" json:"Service" toml:"Service" yaml:"Service"`
}

// NewStruct creates a new relationship struct
func (*nginxConfigR) NewStruct() *nginxConfigR {
	return &nginxConfigR{}
}

// nginxConfigL is where Load methods for each relationship are stored.
type nginxConfigL struct{}

var (
	nginxConfigAllColumns            = []string{"id", "service_id", "type", "path", "last_modified"}
	nginxConfigColumnsWithoutDefault = []string{"service_id", "type", "path", "last_modified"}
	nginxConfigColumnsWithDefault    = []string{"id"}
	nginxConfigPrimaryKeyColumns     = []string{"id"}
)

type (
	// NginxConfigSlice is an alias for a slice of pointers to NginxConfig.
	// This should generally be used opposed to []NginxConfig.
	NginxConfigSlice []*NginxConfig
	// NginxConfigHook is the signature for custom NginxConfig hook methods
	NginxConfigHook func(context.Context, boil.ContextExecutor, *NginxConfig) error

	nginxConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	nginxConfigType                 = reflect.TypeOf(&NginxConfig{})
	nginxConfigMapping              = queries.MakeStructMapping(nginxConfigType)
	nginxConfigPrimaryKeyMapping, _ = queries.BindMapping(nginxConfigType, nginxConfigMapping, nginxConfigPrimaryKeyColumns)
	nginxConfigInsertCacheMut       sync.RWMutex
	nginxConfigInsertCache          = make(map[string]insertCache)
	nginxConfigUpdateCacheMut       sync.RWMutex
	nginxConfigUpdateCache          = make(map[string]updateCache)
	nginxConfigUpsertCacheMut       sync.RWMutex
	nginxConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var nginxConfigBeforeInsertHooks []NginxConfigHook
var nginxConfigBeforeUpdateHooks []NginxConfigHook
var nginxConfigBeforeDeleteHooks []NginxConfigHook
var nginxConfigBeforeUpsertHooks []NginxConfigHook

var nginxConfigAfterInsertHooks []NginxConfigHook
var nginxConfigAfterSelectHooks []NginxConfigHook
var nginxConfigAfterUpdateHooks []NginxConfigHook
var nginxConfigAfterDeleteHooks []NginxConfigHook
var nginxConfigAfterUpsertHooks []NginxConfigHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NginxConfig) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NginxConfig) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NginxConfig) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NginxConfig) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NginxConfig) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NginxConfig) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NginxConfig) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NginxConfig) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NginxConfig) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nginxConfigAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNginxConfigHook registers your hook function for all future operations.
func AddNginxConfigHook(hookPoint boil.HookPoint, nginxConfigHook NginxConfigHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		nginxConfigBeforeInsertHooks = append(nginxConfigBeforeInsertHooks, nginxConfigHook)
	case boil.BeforeUpdateHook:
		nginxConfigBeforeUpdateHooks = append(nginxConfigBeforeUpdateHooks, nginxConfigHook)
	case boil.BeforeDeleteHook:
		nginxConfigBeforeDeleteHooks = append(nginxConfigBeforeDeleteHooks, nginxConfigHook)
	case boil.BeforeUpsertHook:
		nginxConfigBeforeUpsertHooks = append(nginxConfigBeforeUpsertHooks, nginxConfigHook)
	case boil.AfterInsertHook:
		nginxConfigAfterInsertHooks = append(nginxConfigAfterInsertHooks, nginxConfigHook)
	case boil.AfterSelectHook:
		nginxConfigAfterSelectHooks = append(nginxConfigAfterSelectHooks, nginxConfigHook)
	case boil.AfterUpdateHook:
		nginxConfigAfterUpdateHooks = append(nginxConfigAfterUpdateHooks, nginxConfigHook)
	case boil.AfterDeleteHook:
		nginxConfigAfterDeleteHooks = append(nginxConfigAfterDeleteHooks, nginxConfigHook)
	case boil.AfterUpsertHook:
		nginxConfigAfterUpsertHooks = append(nginxConfigAfterUpsertHooks, nginxConfigHook)
	}
}

// One returns a single nginxConfig record from the query.
func (q nginxConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NginxConfig, error) {
	o := &NginxConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for nginx_configs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NginxConfig records from the query.
func (q nginxConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (NginxConfigSlice, error) {
	var o []*NginxConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NginxConfig slice")
	}

	if len(nginxConfigAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NginxConfig records in the query.
func (q nginxConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count nginx_configs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q nginxConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if nginx_configs exists")
	}

	return count > 0, nil
}

// Service pointed to by the foreign key.
func (o *NginxConfig) Service(mods ...qm.QueryMod) serviceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ServiceID),
	}

	queryMods = append(queryMods, mods...)

	query := Services(queryMods...)
	queries.SetFrom(query.Query, "\"services\"")

	return query
}

// LoadService allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (nginxConfigL) LoadService(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNginxConfig interface{}, mods queries.Applicator) error {
	var slice []*NginxConfig
	var object *NginxConfig

	if singular {
		object = maybeNginxConfig.(*NginxConfig)
	} else {
		slice = *maybeNginxConfig.(*[]*NginxConfig)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &nginxConfigR{}
		}
		if !queries.IsNil(object.ServiceID) {
			args = append(args, object.ServiceID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &nginxConfigR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ServiceID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ServiceID) {
				args = append(args, obj.ServiceID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`services`),
		qm.WhereIn(`services.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Service")
	}

	var resultSlice []*Service
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Service")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for services")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for services")
	}

	if len(nginxConfigAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Service = foreign
		if foreign.R == nil {
			foreign.R = &serviceR{}
		}
		foreign.R.NginxConfigs = append(foreign.R.NginxConfigs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ServiceID, foreign.ID) {
				local.R.Service = foreign
				if foreign.R == nil {
					foreign.R = &serviceR{}
				}
				foreign.R.NginxConfigs = append(foreign.R.NginxConfigs, local)
				break
			}
		}
	}

	return nil
}

// SetService of the nginxConfig to the related item.
// Sets o.R.Service to related.
// Adds o to related.R.NginxConfigs.
func (o *NginxConfig) SetService(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Service) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"nginx_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"service_id"}),
		strmangle.WhereClause("\"", "\"", 0, nginxConfigPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ServiceID, related.ID)
	if o.R == nil {
		o.R = &nginxConfigR{
			Service: related,
		}
	} else {
		o.R.Service = related
	}

	if related.R == nil {
		related.R = &serviceR{
			NginxConfigs: NginxConfigSlice{o},
		}
	} else {
		related.R.NginxConfigs = append(related.R.NginxConfigs, o)
	}

	return nil
}

// RemoveService relationship.
// Sets o.R.Service to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *NginxConfig) RemoveService(ctx context.Context, exec boil.ContextExecutor, related *Service) error {
	var err error

	queries.SetScanner(&o.ServiceID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("service_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Service = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.NginxConfigs {
		if queries.Equal(o.ServiceID, ri.ServiceID) {
			continue
		}

		ln := len(related.R.NginxConfigs)
		if ln > 1 && i < ln-1 {
			related.R.NginxConfigs[i] = related.R.NginxConfigs[ln-1]
		}
		related.R.NginxConfigs = related.R.NginxConfigs[:ln-1]
		break
	}
	return nil
}

// NginxConfigs retrieves all the records using an executor.
func NginxConfigs(mods ...qm.QueryMod) nginxConfigQuery {
	mods = append(mods, qm.From("\"nginx_configs\""))
	return nginxConfigQuery{NewQuery(mods...)}
}

// FindNginxConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNginxConfig(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*NginxConfig, error) {
	nginxConfigObj := &NginxConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"nginx_configs\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, nginxConfigObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from nginx_configs")
	}

	return nginxConfigObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NginxConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no nginx_configs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(nginxConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	nginxConfigInsertCacheMut.RLock()
	cache, cached := nginxConfigInsertCache[key]
	nginxConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			nginxConfigAllColumns,
			nginxConfigColumnsWithDefault,
			nginxConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(nginxConfigType, nginxConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(nginxConfigType, nginxConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"nginx_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"nginx_configs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"nginx_configs\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, nginxConfigPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into nginx_configs")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == nginxConfigMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for nginx_configs")
	}

CacheNoHooks:
	if !cached {
		nginxConfigInsertCacheMut.Lock()
		nginxConfigInsertCache[key] = cache
		nginxConfigInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NginxConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NginxConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	nginxConfigUpdateCacheMut.RLock()
	cache, cached := nginxConfigUpdateCache[key]
	nginxConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			nginxConfigAllColumns,
			nginxConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update nginx_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"nginx_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, nginxConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(nginxConfigType, nginxConfigMapping, append(wl, nginxConfigPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update nginx_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for nginx_configs")
	}

	if !cached {
		nginxConfigUpdateCacheMut.Lock()
		nginxConfigUpdateCache[key] = cache
		nginxConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q nginxConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for nginx_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for nginx_configs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NginxConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nginxConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"nginx_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nginxConfigPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in nginxConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all nginxConfig")
	}
	return rowsAff, nil
}

// Delete deletes a single NginxConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NginxConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NginxConfig provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), nginxConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"nginx_configs\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from nginx_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for nginx_configs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q nginxConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no nginxConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from nginx_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for nginx_configs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NginxConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(nginxConfigBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nginxConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"nginx_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nginxConfigPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from nginxConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for nginx_configs")
	}

	if len(nginxConfigAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NginxConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNginxConfig(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NginxConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NginxConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nginxConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"nginx_configs\".* FROM \"nginx_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nginxConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NginxConfigSlice")
	}

	*o = slice

	return nil
}

// NginxConfigExists checks if the NginxConfig row exists.
func NginxConfigExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"nginx_configs\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if nginx_configs exists")
	}

	return exists, nil
}