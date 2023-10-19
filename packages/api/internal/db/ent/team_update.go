// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/internal"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/predicate"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/team"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/user"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/usersteams"
	"github.com/google/uuid"
)

// TeamUpdate is the builder for updating Team entities.
type TeamUpdate struct {
	config
	hooks    []Hook
	mutation *TeamMutation
}

// Where appends a list predicates to the TeamUpdate builder.
func (tu *TeamUpdate) Where(ps ...predicate.Team) *TeamUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetIsDefault sets the "is_default" field.
func (tu *TeamUpdate) SetIsDefault(b bool) *TeamUpdate {
	tu.mutation.SetIsDefault(b)
	return tu
}

// SetName sets the "name" field.
func (tu *TeamUpdate) SetName(s string) *TeamUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetIsBlocked sets the "is_blocked" field.
func (tu *TeamUpdate) SetIsBlocked(b bool) *TeamUpdate {
	tu.mutation.SetIsBlocked(b)
	return tu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tu *TeamUpdate) AddUserIDs(ids ...uuid.UUID) *TeamUpdate {
	tu.mutation.AddUserIDs(ids...)
	return tu
}

// AddUsers adds the "users" edges to the User entity.
func (tu *TeamUpdate) AddUsers(u ...*User) *TeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUserIDs(ids...)
}

// AddUsersTeamIDs adds the "users_teams" edge to the UsersTeams entity by IDs.
func (tu *TeamUpdate) AddUsersTeamIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddUsersTeamIDs(ids...)
	return tu
}

// AddUsersTeams adds the "users_teams" edges to the UsersTeams entity.
func (tu *TeamUpdate) AddUsersTeams(u ...*UsersTeams) *TeamUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUsersTeamIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tu *TeamUpdate) Mutation() *TeamMutation {
	return tu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tu *TeamUpdate) ClearUsers() *TeamUpdate {
	tu.mutation.ClearUsers()
	return tu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tu *TeamUpdate) RemoveUserIDs(ids ...uuid.UUID) *TeamUpdate {
	tu.mutation.RemoveUserIDs(ids...)
	return tu
}

// RemoveUsers removes "users" edges to User entities.
func (tu *TeamUpdate) RemoveUsers(u ...*User) *TeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUserIDs(ids...)
}

// ClearUsersTeams clears all "users_teams" edges to the UsersTeams entity.
func (tu *TeamUpdate) ClearUsersTeams() *TeamUpdate {
	tu.mutation.ClearUsersTeams()
	return tu
}

// RemoveUsersTeamIDs removes the "users_teams" edge to UsersTeams entities by IDs.
func (tu *TeamUpdate) RemoveUsersTeamIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveUsersTeamIDs(ids...)
	return tu
}

// RemoveUsersTeams removes "users_teams" edges to UsersTeams entities.
func (tu *TeamUpdate) RemoveUsersTeams(u ...*UsersTeams) *TeamUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUsersTeamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.IsDefault(); ok {
		_spec.SetField(team.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.IsBlocked(); ok {
		_spec.SetField(team.FieldIsBlocked, field.TypeBool, value)
	}
	if tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUsersTeamsIDs(); len(nodes) > 0 && !tu.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tu.schemaConfig.Team
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeamUpdateOne is the builder for updating a single Team entity.
type TeamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamMutation
}

// SetIsDefault sets the "is_default" field.
func (tuo *TeamUpdateOne) SetIsDefault(b bool) *TeamUpdateOne {
	tuo.mutation.SetIsDefault(b)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TeamUpdateOne) SetName(s string) *TeamUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetIsBlocked sets the "is_blocked" field.
func (tuo *TeamUpdateOne) SetIsBlocked(b bool) *TeamUpdateOne {
	tuo.mutation.SetIsBlocked(b)
	return tuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tuo *TeamUpdateOne) AddUserIDs(ids ...uuid.UUID) *TeamUpdateOne {
	tuo.mutation.AddUserIDs(ids...)
	return tuo
}

// AddUsers adds the "users" edges to the User entity.
func (tuo *TeamUpdateOne) AddUsers(u ...*User) *TeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUserIDs(ids...)
}

// AddUsersTeamIDs adds the "users_teams" edge to the UsersTeams entity by IDs.
func (tuo *TeamUpdateOne) AddUsersTeamIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddUsersTeamIDs(ids...)
	return tuo
}

// AddUsersTeams adds the "users_teams" edges to the UsersTeams entity.
func (tuo *TeamUpdateOne) AddUsersTeams(u ...*UsersTeams) *TeamUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUsersTeamIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tuo *TeamUpdateOne) Mutation() *TeamMutation {
	return tuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tuo *TeamUpdateOne) ClearUsers() *TeamUpdateOne {
	tuo.mutation.ClearUsers()
	return tuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tuo *TeamUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *TeamUpdateOne {
	tuo.mutation.RemoveUserIDs(ids...)
	return tuo
}

// RemoveUsers removes "users" edges to User entities.
func (tuo *TeamUpdateOne) RemoveUsers(u ...*User) *TeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUserIDs(ids...)
}

// ClearUsersTeams clears all "users_teams" edges to the UsersTeams entity.
func (tuo *TeamUpdateOne) ClearUsersTeams() *TeamUpdateOne {
	tuo.mutation.ClearUsersTeams()
	return tuo
}

// RemoveUsersTeamIDs removes the "users_teams" edge to UsersTeams entities by IDs.
func (tuo *TeamUpdateOne) RemoveUsersTeamIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveUsersTeamIDs(ids...)
	return tuo
}

// RemoveUsersTeams removes "users_teams" edges to UsersTeams entities.
func (tuo *TeamUpdateOne) RemoveUsersTeams(u ...*UsersTeams) *TeamUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUsersTeamIDs(ids...)
}

// Where appends a list predicates to the TeamUpdate builder.
func (tuo *TeamUpdateOne) Where(ps ...predicate.Team) *TeamUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamUpdateOne) Select(field string, fields ...string) *TeamUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Team entity.
func (tuo *TeamUpdateOne) Save(ctx context.Context) (*Team, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamUpdateOne) SaveX(ctx context.Context) *Team {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TeamUpdateOne) sqlSave(ctx context.Context) (_node *Team, err error) {
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Team.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for _, f := range fields {
			if !team.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.IsDefault(); ok {
		_spec.SetField(team.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.IsBlocked(); ok {
		_spec.SetField(team.FieldIsBlocked, field.TypeBool, value)
	}
	if tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUsersTeamsIDs(); len(nodes) > 0 && !tuo.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tuo.schemaConfig.Team
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_node = &Team{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}