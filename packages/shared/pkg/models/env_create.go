// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/envalias"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/team"
	"github.com/google/uuid"
)

// EnvCreate is the builder for creating a Env entity.
type EnvCreate struct {
	config
	mutation *EnvMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ec *EnvCreate) SetCreatedAt(t time.Time) *EnvCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EnvCreate) SetNillableCreatedAt(t *time.Time) *EnvCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EnvCreate) SetUpdatedAt(t time.Time) *EnvCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EnvCreate) SetNillableUpdatedAt(t *time.Time) *EnvCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetTeamID sets the "team_id" field.
func (ec *EnvCreate) SetTeamID(u uuid.UUID) *EnvCreate {
	ec.mutation.SetTeamID(u)
	return ec
}

// SetDockerfile sets the "dockerfile" field.
func (ec *EnvCreate) SetDockerfile(s string) *EnvCreate {
	ec.mutation.SetDockerfile(s)
	return ec
}

// SetPublic sets the "public" field.
func (ec *EnvCreate) SetPublic(b bool) *EnvCreate {
	ec.mutation.SetPublic(b)
	return ec
}

// SetBuildID sets the "build_id" field.
func (ec *EnvCreate) SetBuildID(u uuid.UUID) *EnvCreate {
	ec.mutation.SetBuildID(u)
	return ec
}

// SetBuildCount sets the "build_count" field.
func (ec *EnvCreate) SetBuildCount(i int32) *EnvCreate {
	ec.mutation.SetBuildCount(i)
	return ec
}

// SetNillableBuildCount sets the "build_count" field if the given value is not nil.
func (ec *EnvCreate) SetNillableBuildCount(i *int32) *EnvCreate {
	if i != nil {
		ec.SetBuildCount(*i)
	}
	return ec
}

// SetSpawnCount sets the "spawn_count" field.
func (ec *EnvCreate) SetSpawnCount(i int32) *EnvCreate {
	ec.mutation.SetSpawnCount(i)
	return ec
}

// SetNillableSpawnCount sets the "spawn_count" field if the given value is not nil.
func (ec *EnvCreate) SetNillableSpawnCount(i *int32) *EnvCreate {
	if i != nil {
		ec.SetSpawnCount(*i)
	}
	return ec
}

// SetLastSpawnedAt sets the "last_spawned_at" field.
func (ec *EnvCreate) SetLastSpawnedAt(t time.Time) *EnvCreate {
	ec.mutation.SetLastSpawnedAt(t)
	return ec
}

// SetNillableLastSpawnedAt sets the "last_spawned_at" field if the given value is not nil.
func (ec *EnvCreate) SetNillableLastSpawnedAt(t *time.Time) *EnvCreate {
	if t != nil {
		ec.SetLastSpawnedAt(*t)
	}
	return ec
}

// SetVcpu sets the "vcpu" field.
func (ec *EnvCreate) SetVcpu(i int64) *EnvCreate {
	ec.mutation.SetVcpu(i)
	return ec
}

// SetRAMMB sets the "ram_mb" field.
func (ec *EnvCreate) SetRAMMB(i int64) *EnvCreate {
	ec.mutation.SetRAMMB(i)
	return ec
}

// SetFreeDiskSizeMB sets the "free_disk_size_mb" field.
func (ec *EnvCreate) SetFreeDiskSizeMB(i int64) *EnvCreate {
	ec.mutation.SetFreeDiskSizeMB(i)
	return ec
}

// SetTotalDiskSizeMB sets the "total_disk_size_mb" field.
func (ec *EnvCreate) SetTotalDiskSizeMB(i int64) *EnvCreate {
	ec.mutation.SetTotalDiskSizeMB(i)
	return ec
}

// SetID sets the "id" field.
func (ec *EnvCreate) SetID(s string) *EnvCreate {
	ec.mutation.SetID(s)
	return ec
}

// SetTeam sets the "team" edge to the Team entity.
func (ec *EnvCreate) SetTeam(t *Team) *EnvCreate {
	return ec.SetTeamID(t.ID)
}

// AddEnvAliasIDs adds the "env_aliases" edge to the EnvAlias entity by IDs.
func (ec *EnvCreate) AddEnvAliasIDs(ids ...string) *EnvCreate {
	ec.mutation.AddEnvAliasIDs(ids...)
	return ec
}

// AddEnvAliases adds the "env_aliases" edges to the EnvAlias entity.
func (ec *EnvCreate) AddEnvAliases(e ...*EnvAlias) *EnvCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEnvAliasIDs(ids...)
}

// Mutation returns the EnvMutation object of the builder.
func (ec *EnvCreate) Mutation() *EnvMutation {
	return ec.mutation
}

// Save creates the Env in the database.
func (ec *EnvCreate) Save(ctx context.Context) (*Env, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EnvCreate) SaveX(ctx context.Context) *Env {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EnvCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EnvCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EnvCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := env.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := env.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.BuildCount(); !ok {
		v := env.DefaultBuildCount
		ec.mutation.SetBuildCount(v)
	}
	if _, ok := ec.mutation.SpawnCount(); !ok {
		v := env.DefaultSpawnCount
		ec.mutation.SetSpawnCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EnvCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`models: missing required field "Env.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`models: missing required field "Env.updated_at"`)}
	}
	if _, ok := ec.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team_id", err: errors.New(`models: missing required field "Env.team_id"`)}
	}
	if _, ok := ec.mutation.Dockerfile(); !ok {
		return &ValidationError{Name: "dockerfile", err: errors.New(`models: missing required field "Env.dockerfile"`)}
	}
	if _, ok := ec.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`models: missing required field "Env.public"`)}
	}
	if _, ok := ec.mutation.BuildID(); !ok {
		return &ValidationError{Name: "build_id", err: errors.New(`models: missing required field "Env.build_id"`)}
	}
	if _, ok := ec.mutation.BuildCount(); !ok {
		return &ValidationError{Name: "build_count", err: errors.New(`models: missing required field "Env.build_count"`)}
	}
	if _, ok := ec.mutation.SpawnCount(); !ok {
		return &ValidationError{Name: "spawn_count", err: errors.New(`models: missing required field "Env.spawn_count"`)}
	}
	if _, ok := ec.mutation.Vcpu(); !ok {
		return &ValidationError{Name: "vcpu", err: errors.New(`models: missing required field "Env.vcpu"`)}
	}
	if _, ok := ec.mutation.RAMMB(); !ok {
		return &ValidationError{Name: "ram_mb", err: errors.New(`models: missing required field "Env.ram_mb"`)}
	}
	if _, ok := ec.mutation.FreeDiskSizeMB(); !ok {
		return &ValidationError{Name: "free_disk_size_mb", err: errors.New(`models: missing required field "Env.free_disk_size_mb"`)}
	}
	if _, ok := ec.mutation.TotalDiskSizeMB(); !ok {
		return &ValidationError{Name: "total_disk_size_mb", err: errors.New(`models: missing required field "Env.total_disk_size_mb"`)}
	}
	if _, ok := ec.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`models: missing required edge "Env.team"`)}
	}
	return nil
}

func (ec *EnvCreate) sqlSave(ctx context.Context) (*Env, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Env.ID type: %T", _spec.ID.Value)
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EnvCreate) createSpec() (*Env, *sqlgraph.CreateSpec) {
	var (
		_node = &Env{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(env.Table, sqlgraph.NewFieldSpec(env.FieldID, field.TypeString))
	)
	_spec.Schema = ec.schemaConfig.Env
	_spec.OnConflict = ec.conflict
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(env.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(env.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.Dockerfile(); ok {
		_spec.SetField(env.FieldDockerfile, field.TypeString, value)
		_node.Dockerfile = value
	}
	if value, ok := ec.mutation.Public(); ok {
		_spec.SetField(env.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := ec.mutation.BuildID(); ok {
		_spec.SetField(env.FieldBuildID, field.TypeUUID, value)
		_node.BuildID = value
	}
	if value, ok := ec.mutation.BuildCount(); ok {
		_spec.SetField(env.FieldBuildCount, field.TypeInt32, value)
		_node.BuildCount = value
	}
	if value, ok := ec.mutation.SpawnCount(); ok {
		_spec.SetField(env.FieldSpawnCount, field.TypeInt32, value)
		_node.SpawnCount = value
	}
	if value, ok := ec.mutation.LastSpawnedAt(); ok {
		_spec.SetField(env.FieldLastSpawnedAt, field.TypeTime, value)
		_node.LastSpawnedAt = value
	}
	if value, ok := ec.mutation.Vcpu(); ok {
		_spec.SetField(env.FieldVcpu, field.TypeInt64, value)
		_node.Vcpu = value
	}
	if value, ok := ec.mutation.RAMMB(); ok {
		_spec.SetField(env.FieldRAMMB, field.TypeInt64, value)
		_node.RAMMB = value
	}
	if value, ok := ec.mutation.FreeDiskSizeMB(); ok {
		_spec.SetField(env.FieldFreeDiskSizeMB, field.TypeInt64, value)
		_node.FreeDiskSizeMB = value
	}
	if value, ok := ec.mutation.TotalDiskSizeMB(); ok {
		_spec.SetField(env.FieldTotalDiskSizeMB, field.TypeInt64, value)
		_node.TotalDiskSizeMB = value
	}
	if nodes := ec.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   env.TeamTable,
			Columns: []string{env.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = ec.schemaConfig.Env
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EnvAliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   env.EnvAliasesTable,
			Columns: []string{env.EnvAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeString),
			},
		}
		edge.Schema = ec.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Env.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EnvUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ec *EnvCreate) OnConflict(opts ...sql.ConflictOption) *EnvUpsertOne {
	ec.conflict = opts
	return &EnvUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Env.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *EnvCreate) OnConflictColumns(columns ...string) *EnvUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EnvUpsertOne{
		create: ec,
	}
}

type (
	// EnvUpsertOne is the builder for "upsert"-ing
	//  one Env node.
	EnvUpsertOne struct {
		create *EnvCreate
	}

	// EnvUpsert is the "OnConflict" setter.
	EnvUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *EnvUpsert) SetUpdatedAt(v time.Time) *EnvUpsert {
	u.Set(env.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EnvUpsert) UpdateUpdatedAt() *EnvUpsert {
	u.SetExcluded(env.FieldUpdatedAt)
	return u
}

// SetTeamID sets the "team_id" field.
func (u *EnvUpsert) SetTeamID(v uuid.UUID) *EnvUpsert {
	u.Set(env.FieldTeamID, v)
	return u
}

// UpdateTeamID sets the "team_id" field to the value that was provided on create.
func (u *EnvUpsert) UpdateTeamID() *EnvUpsert {
	u.SetExcluded(env.FieldTeamID)
	return u
}

// SetDockerfile sets the "dockerfile" field.
func (u *EnvUpsert) SetDockerfile(v string) *EnvUpsert {
	u.Set(env.FieldDockerfile, v)
	return u
}

// UpdateDockerfile sets the "dockerfile" field to the value that was provided on create.
func (u *EnvUpsert) UpdateDockerfile() *EnvUpsert {
	u.SetExcluded(env.FieldDockerfile)
	return u
}

// SetPublic sets the "public" field.
func (u *EnvUpsert) SetPublic(v bool) *EnvUpsert {
	u.Set(env.FieldPublic, v)
	return u
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *EnvUpsert) UpdatePublic() *EnvUpsert {
	u.SetExcluded(env.FieldPublic)
	return u
}

// SetBuildID sets the "build_id" field.
func (u *EnvUpsert) SetBuildID(v uuid.UUID) *EnvUpsert {
	u.Set(env.FieldBuildID, v)
	return u
}

// UpdateBuildID sets the "build_id" field to the value that was provided on create.
func (u *EnvUpsert) UpdateBuildID() *EnvUpsert {
	u.SetExcluded(env.FieldBuildID)
	return u
}

// SetBuildCount sets the "build_count" field.
func (u *EnvUpsert) SetBuildCount(v int32) *EnvUpsert {
	u.Set(env.FieldBuildCount, v)
	return u
}

// UpdateBuildCount sets the "build_count" field to the value that was provided on create.
func (u *EnvUpsert) UpdateBuildCount() *EnvUpsert {
	u.SetExcluded(env.FieldBuildCount)
	return u
}

// AddBuildCount adds v to the "build_count" field.
func (u *EnvUpsert) AddBuildCount(v int32) *EnvUpsert {
	u.Add(env.FieldBuildCount, v)
	return u
}

// SetSpawnCount sets the "spawn_count" field.
func (u *EnvUpsert) SetSpawnCount(v int32) *EnvUpsert {
	u.Set(env.FieldSpawnCount, v)
	return u
}

// UpdateSpawnCount sets the "spawn_count" field to the value that was provided on create.
func (u *EnvUpsert) UpdateSpawnCount() *EnvUpsert {
	u.SetExcluded(env.FieldSpawnCount)
	return u
}

// AddSpawnCount adds v to the "spawn_count" field.
func (u *EnvUpsert) AddSpawnCount(v int32) *EnvUpsert {
	u.Add(env.FieldSpawnCount, v)
	return u
}

// SetLastSpawnedAt sets the "last_spawned_at" field.
func (u *EnvUpsert) SetLastSpawnedAt(v time.Time) *EnvUpsert {
	u.Set(env.FieldLastSpawnedAt, v)
	return u
}

// UpdateLastSpawnedAt sets the "last_spawned_at" field to the value that was provided on create.
func (u *EnvUpsert) UpdateLastSpawnedAt() *EnvUpsert {
	u.SetExcluded(env.FieldLastSpawnedAt)
	return u
}

// ClearLastSpawnedAt clears the value of the "last_spawned_at" field.
func (u *EnvUpsert) ClearLastSpawnedAt() *EnvUpsert {
	u.SetNull(env.FieldLastSpawnedAt)
	return u
}

// SetVcpu sets the "vcpu" field.
func (u *EnvUpsert) SetVcpu(v int64) *EnvUpsert {
	u.Set(env.FieldVcpu, v)
	return u
}

// UpdateVcpu sets the "vcpu" field to the value that was provided on create.
func (u *EnvUpsert) UpdateVcpu() *EnvUpsert {
	u.SetExcluded(env.FieldVcpu)
	return u
}

// AddVcpu adds v to the "vcpu" field.
func (u *EnvUpsert) AddVcpu(v int64) *EnvUpsert {
	u.Add(env.FieldVcpu, v)
	return u
}

// SetRAMMB sets the "ram_mb" field.
func (u *EnvUpsert) SetRAMMB(v int64) *EnvUpsert {
	u.Set(env.FieldRAMMB, v)
	return u
}

// UpdateRAMMB sets the "ram_mb" field to the value that was provided on create.
func (u *EnvUpsert) UpdateRAMMB() *EnvUpsert {
	u.SetExcluded(env.FieldRAMMB)
	return u
}

// AddRAMMB adds v to the "ram_mb" field.
func (u *EnvUpsert) AddRAMMB(v int64) *EnvUpsert {
	u.Add(env.FieldRAMMB, v)
	return u
}

// SetFreeDiskSizeMB sets the "free_disk_size_mb" field.
func (u *EnvUpsert) SetFreeDiskSizeMB(v int64) *EnvUpsert {
	u.Set(env.FieldFreeDiskSizeMB, v)
	return u
}

// UpdateFreeDiskSizeMB sets the "free_disk_size_mb" field to the value that was provided on create.
func (u *EnvUpsert) UpdateFreeDiskSizeMB() *EnvUpsert {
	u.SetExcluded(env.FieldFreeDiskSizeMB)
	return u
}

// AddFreeDiskSizeMB adds v to the "free_disk_size_mb" field.
func (u *EnvUpsert) AddFreeDiskSizeMB(v int64) *EnvUpsert {
	u.Add(env.FieldFreeDiskSizeMB, v)
	return u
}

// SetTotalDiskSizeMB sets the "total_disk_size_mb" field.
func (u *EnvUpsert) SetTotalDiskSizeMB(v int64) *EnvUpsert {
	u.Set(env.FieldTotalDiskSizeMB, v)
	return u
}

// UpdateTotalDiskSizeMB sets the "total_disk_size_mb" field to the value that was provided on create.
func (u *EnvUpsert) UpdateTotalDiskSizeMB() *EnvUpsert {
	u.SetExcluded(env.FieldTotalDiskSizeMB)
	return u
}

// AddTotalDiskSizeMB adds v to the "total_disk_size_mb" field.
func (u *EnvUpsert) AddTotalDiskSizeMB(v int64) *EnvUpsert {
	u.Add(env.FieldTotalDiskSizeMB, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Env.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(env.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EnvUpsertOne) UpdateNewValues() *EnvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(env.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(env.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Env.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EnvUpsertOne) Ignore() *EnvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EnvUpsertOne) DoNothing() *EnvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EnvCreate.OnConflict
// documentation for more info.
func (u *EnvUpsertOne) Update(set func(*EnvUpsert)) *EnvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EnvUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EnvUpsertOne) SetUpdatedAt(v time.Time) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateUpdatedAt() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTeamID sets the "team_id" field.
func (u *EnvUpsertOne) SetTeamID(v uuid.UUID) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetTeamID(v)
	})
}

// UpdateTeamID sets the "team_id" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateTeamID() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateTeamID()
	})
}

// SetDockerfile sets the "dockerfile" field.
func (u *EnvUpsertOne) SetDockerfile(v string) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetDockerfile(v)
	})
}

// UpdateDockerfile sets the "dockerfile" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateDockerfile() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateDockerfile()
	})
}

// SetPublic sets the "public" field.
func (u *EnvUpsertOne) SetPublic(v bool) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdatePublic() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdatePublic()
	})
}

// SetBuildID sets the "build_id" field.
func (u *EnvUpsertOne) SetBuildID(v uuid.UUID) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetBuildID(v)
	})
}

// UpdateBuildID sets the "build_id" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateBuildID() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateBuildID()
	})
}

// SetBuildCount sets the "build_count" field.
func (u *EnvUpsertOne) SetBuildCount(v int32) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetBuildCount(v)
	})
}

// AddBuildCount adds v to the "build_count" field.
func (u *EnvUpsertOne) AddBuildCount(v int32) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.AddBuildCount(v)
	})
}

// UpdateBuildCount sets the "build_count" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateBuildCount() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateBuildCount()
	})
}

// SetSpawnCount sets the "spawn_count" field.
func (u *EnvUpsertOne) SetSpawnCount(v int32) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetSpawnCount(v)
	})
}

// AddSpawnCount adds v to the "spawn_count" field.
func (u *EnvUpsertOne) AddSpawnCount(v int32) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.AddSpawnCount(v)
	})
}

// UpdateSpawnCount sets the "spawn_count" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateSpawnCount() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateSpawnCount()
	})
}

// SetLastSpawnedAt sets the "last_spawned_at" field.
func (u *EnvUpsertOne) SetLastSpawnedAt(v time.Time) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetLastSpawnedAt(v)
	})
}

// UpdateLastSpawnedAt sets the "last_spawned_at" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateLastSpawnedAt() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateLastSpawnedAt()
	})
}

// ClearLastSpawnedAt clears the value of the "last_spawned_at" field.
func (u *EnvUpsertOne) ClearLastSpawnedAt() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.ClearLastSpawnedAt()
	})
}

// SetVcpu sets the "vcpu" field.
func (u *EnvUpsertOne) SetVcpu(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetVcpu(v)
	})
}

// AddVcpu adds v to the "vcpu" field.
func (u *EnvUpsertOne) AddVcpu(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.AddVcpu(v)
	})
}

// UpdateVcpu sets the "vcpu" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateVcpu() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateVcpu()
	})
}

// SetRAMMB sets the "ram_mb" field.
func (u *EnvUpsertOne) SetRAMMB(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetRAMMB(v)
	})
}

// AddRAMMB adds v to the "ram_mb" field.
func (u *EnvUpsertOne) AddRAMMB(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.AddRAMMB(v)
	})
}

// UpdateRAMMB sets the "ram_mb" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateRAMMB() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateRAMMB()
	})
}

// SetFreeDiskSizeMB sets the "free_disk_size_mb" field.
func (u *EnvUpsertOne) SetFreeDiskSizeMB(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetFreeDiskSizeMB(v)
	})
}

// AddFreeDiskSizeMB adds v to the "free_disk_size_mb" field.
func (u *EnvUpsertOne) AddFreeDiskSizeMB(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.AddFreeDiskSizeMB(v)
	})
}

// UpdateFreeDiskSizeMB sets the "free_disk_size_mb" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateFreeDiskSizeMB() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateFreeDiskSizeMB()
	})
}

// SetTotalDiskSizeMB sets the "total_disk_size_mb" field.
func (u *EnvUpsertOne) SetTotalDiskSizeMB(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.SetTotalDiskSizeMB(v)
	})
}

// AddTotalDiskSizeMB adds v to the "total_disk_size_mb" field.
func (u *EnvUpsertOne) AddTotalDiskSizeMB(v int64) *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.AddTotalDiskSizeMB(v)
	})
}

// UpdateTotalDiskSizeMB sets the "total_disk_size_mb" field to the value that was provided on create.
func (u *EnvUpsertOne) UpdateTotalDiskSizeMB() *EnvUpsertOne {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateTotalDiskSizeMB()
	})
}

// Exec executes the query.
func (u *EnvUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("models: missing options for EnvCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EnvUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EnvUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("models: EnvUpsertOne.ID is not supported by MySQL driver. Use EnvUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EnvUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EnvCreateBulk is the builder for creating many Env entities in bulk.
type EnvCreateBulk struct {
	config
	err      error
	builders []*EnvCreate
	conflict []sql.ConflictOption
}

// Save creates the Env entities in the database.
func (ecb *EnvCreateBulk) Save(ctx context.Context) ([]*Env, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Env, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EnvMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EnvCreateBulk) SaveX(ctx context.Context) []*Env {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EnvCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EnvCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Env.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EnvUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ecb *EnvCreateBulk) OnConflict(opts ...sql.ConflictOption) *EnvUpsertBulk {
	ecb.conflict = opts
	return &EnvUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Env.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *EnvCreateBulk) OnConflictColumns(columns ...string) *EnvUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EnvUpsertBulk{
		create: ecb,
	}
}

// EnvUpsertBulk is the builder for "upsert"-ing
// a bulk of Env nodes.
type EnvUpsertBulk struct {
	create *EnvCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Env.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(env.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EnvUpsertBulk) UpdateNewValues() *EnvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(env.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(env.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Env.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EnvUpsertBulk) Ignore() *EnvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EnvUpsertBulk) DoNothing() *EnvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EnvCreateBulk.OnConflict
// documentation for more info.
func (u *EnvUpsertBulk) Update(set func(*EnvUpsert)) *EnvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EnvUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EnvUpsertBulk) SetUpdatedAt(v time.Time) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateUpdatedAt() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTeamID sets the "team_id" field.
func (u *EnvUpsertBulk) SetTeamID(v uuid.UUID) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetTeamID(v)
	})
}

// UpdateTeamID sets the "team_id" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateTeamID() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateTeamID()
	})
}

// SetDockerfile sets the "dockerfile" field.
func (u *EnvUpsertBulk) SetDockerfile(v string) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetDockerfile(v)
	})
}

// UpdateDockerfile sets the "dockerfile" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateDockerfile() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateDockerfile()
	})
}

// SetPublic sets the "public" field.
func (u *EnvUpsertBulk) SetPublic(v bool) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdatePublic() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdatePublic()
	})
}

// SetBuildID sets the "build_id" field.
func (u *EnvUpsertBulk) SetBuildID(v uuid.UUID) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetBuildID(v)
	})
}

// UpdateBuildID sets the "build_id" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateBuildID() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateBuildID()
	})
}

// SetBuildCount sets the "build_count" field.
func (u *EnvUpsertBulk) SetBuildCount(v int32) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetBuildCount(v)
	})
}

// AddBuildCount adds v to the "build_count" field.
func (u *EnvUpsertBulk) AddBuildCount(v int32) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.AddBuildCount(v)
	})
}

// UpdateBuildCount sets the "build_count" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateBuildCount() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateBuildCount()
	})
}

// SetSpawnCount sets the "spawn_count" field.
func (u *EnvUpsertBulk) SetSpawnCount(v int32) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetSpawnCount(v)
	})
}

// AddSpawnCount adds v to the "spawn_count" field.
func (u *EnvUpsertBulk) AddSpawnCount(v int32) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.AddSpawnCount(v)
	})
}

// UpdateSpawnCount sets the "spawn_count" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateSpawnCount() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateSpawnCount()
	})
}

// SetLastSpawnedAt sets the "last_spawned_at" field.
func (u *EnvUpsertBulk) SetLastSpawnedAt(v time.Time) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetLastSpawnedAt(v)
	})
}

// UpdateLastSpawnedAt sets the "last_spawned_at" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateLastSpawnedAt() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateLastSpawnedAt()
	})
}

// ClearLastSpawnedAt clears the value of the "last_spawned_at" field.
func (u *EnvUpsertBulk) ClearLastSpawnedAt() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.ClearLastSpawnedAt()
	})
}

// SetVcpu sets the "vcpu" field.
func (u *EnvUpsertBulk) SetVcpu(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetVcpu(v)
	})
}

// AddVcpu adds v to the "vcpu" field.
func (u *EnvUpsertBulk) AddVcpu(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.AddVcpu(v)
	})
}

// UpdateVcpu sets the "vcpu" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateVcpu() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateVcpu()
	})
}

// SetRAMMB sets the "ram_mb" field.
func (u *EnvUpsertBulk) SetRAMMB(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetRAMMB(v)
	})
}

// AddRAMMB adds v to the "ram_mb" field.
func (u *EnvUpsertBulk) AddRAMMB(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.AddRAMMB(v)
	})
}

// UpdateRAMMB sets the "ram_mb" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateRAMMB() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateRAMMB()
	})
}

// SetFreeDiskSizeMB sets the "free_disk_size_mb" field.
func (u *EnvUpsertBulk) SetFreeDiskSizeMB(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetFreeDiskSizeMB(v)
	})
}

// AddFreeDiskSizeMB adds v to the "free_disk_size_mb" field.
func (u *EnvUpsertBulk) AddFreeDiskSizeMB(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.AddFreeDiskSizeMB(v)
	})
}

// UpdateFreeDiskSizeMB sets the "free_disk_size_mb" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateFreeDiskSizeMB() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateFreeDiskSizeMB()
	})
}

// SetTotalDiskSizeMB sets the "total_disk_size_mb" field.
func (u *EnvUpsertBulk) SetTotalDiskSizeMB(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.SetTotalDiskSizeMB(v)
	})
}

// AddTotalDiskSizeMB adds v to the "total_disk_size_mb" field.
func (u *EnvUpsertBulk) AddTotalDiskSizeMB(v int64) *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.AddTotalDiskSizeMB(v)
	})
}

// UpdateTotalDiskSizeMB sets the "total_disk_size_mb" field to the value that was provided on create.
func (u *EnvUpsertBulk) UpdateTotalDiskSizeMB() *EnvUpsertBulk {
	return u.Update(func(s *EnvUpsert) {
		s.UpdateTotalDiskSizeMB()
	})
}

// Exec executes the query.
func (u *EnvUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("models: OnConflict was set for builder %d. Set it on the EnvCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("models: missing options for EnvCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EnvUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
