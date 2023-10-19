// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/internal"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/predicate"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/teamapikey"
)

// TeamApiKeyDelete is the builder for deleting a TeamApiKey entity.
type TeamApiKeyDelete struct {
	config
	hooks    []Hook
	mutation *TeamApiKeyMutation
}

// Where appends a list predicates to the TeamApiKeyDelete builder.
func (takd *TeamApiKeyDelete) Where(ps ...predicate.TeamApiKey) *TeamApiKeyDelete {
	takd.mutation.Where(ps...)
	return takd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (takd *TeamApiKeyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, takd.sqlExec, takd.mutation, takd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (takd *TeamApiKeyDelete) ExecX(ctx context.Context) int {
	n, err := takd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (takd *TeamApiKeyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(teamapikey.Table, sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString))
	_spec.Node.Schema = takd.schemaConfig.TeamApiKey
	ctx = internal.NewSchemaConfigContext(ctx, takd.schemaConfig)
	if ps := takd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, takd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	takd.mutation.done = true
	return affected, err
}

// TeamApiKeyDeleteOne is the builder for deleting a single TeamApiKey entity.
type TeamApiKeyDeleteOne struct {
	takd *TeamApiKeyDelete
}

// Where appends a list predicates to the TeamApiKeyDelete builder.
func (takdo *TeamApiKeyDeleteOne) Where(ps ...predicate.TeamApiKey) *TeamApiKeyDeleteOne {
	takdo.takd.mutation.Where(ps...)
	return takdo
}

// Exec executes the deletion query.
func (takdo *TeamApiKeyDeleteOne) Exec(ctx context.Context) error {
	n, err := takdo.takd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{teamapikey.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (takdo *TeamApiKeyDeleteOne) ExecX(ctx context.Context) {
	if err := takdo.Exec(ctx); err != nil {
		panic(err)
	}
}