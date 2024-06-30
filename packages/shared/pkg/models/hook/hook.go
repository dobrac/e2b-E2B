// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/e2b-dev/infra/packages/shared/pkg/models"
)

// The AccessTokenFunc type is an adapter to allow the use of ordinary
// function as AccessToken mutator.
type AccessTokenFunc func(context.Context, *models.AccessTokenMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f AccessTokenFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.AccessTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.AccessTokenMutation", m)
}

// The EnvFunc type is an adapter to allow the use of ordinary
// function as Env mutator.
type EnvFunc func(context.Context, *models.EnvMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f EnvFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.EnvMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.EnvMutation", m)
}

// The EnvAliasFunc type is an adapter to allow the use of ordinary
// function as EnvAlias mutator.
type EnvAliasFunc func(context.Context, *models.EnvAliasMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f EnvAliasFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.EnvAliasMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.EnvAliasMutation", m)
}

// The EnvBuildFunc type is an adapter to allow the use of ordinary
// function as EnvBuild mutator.
type EnvBuildFunc func(context.Context, *models.EnvBuildMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f EnvBuildFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.EnvBuildMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.EnvBuildMutation", m)
}

// The TeamFunc type is an adapter to allow the use of ordinary
// function as Team mutator.
type TeamFunc func(context.Context, *models.TeamMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f TeamFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.TeamMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.TeamMutation", m)
}

// The TeamAPIKeyFunc type is an adapter to allow the use of ordinary
// function as TeamAPIKey mutator.
type TeamAPIKeyFunc func(context.Context, *models.TeamAPIKeyMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f TeamAPIKeyFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.TeamAPIKeyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.TeamAPIKeyMutation", m)
}

// The TierFunc type is an adapter to allow the use of ordinary
// function as Tier mutator.
type TierFunc func(context.Context, *models.TierMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f TierFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.TierMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.TierMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *models.UserMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.UserMutation", m)
}

// The UsersTeamsFunc type is an adapter to allow the use of ordinary
// function as UsersTeams mutator.
type UsersTeamsFunc func(context.Context, *models.UsersTeamsMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f UsersTeamsFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	if mv, ok := m.(*models.UsersTeamsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *models.UsersTeamsMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, models.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op models.Op) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk models.Hook, cond Condition) models.Hook {
	return func(next models.Mutator) models.Mutator {
		return models.MutateFunc(func(ctx context.Context, m models.Mutation) (models.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, models.Delete|models.Create)
func On(hk models.Hook, op models.Op) models.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, models.Update|models.UpdateOne)
func Unless(hk models.Hook, op models.Op) models.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) models.Hook {
	return func(models.Mutator) models.Mutator {
		return models.MutateFunc(func(context.Context, models.Mutation) (models.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []models.Hook {
//		return []models.Hook{
//			Reject(models.Delete|models.Update),
//		}
//	}
func Reject(op models.Op) models.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []models.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...models.Hook) Chain {
	return Chain{append([]models.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() models.Hook {
	return func(mutator models.Mutator) models.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...models.Hook) Chain {
	newHooks := make([]models.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}