// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.orx.me/echosrv/ent/accesslog"
	"go.orx.me/echosrv/ent/predicate"
)

// AccessLogDelete is the builder for deleting a AccessLog entity.
type AccessLogDelete struct {
	config
	hooks    []Hook
	mutation *AccessLogMutation
}

// Where appends a list predicates to the AccessLogDelete builder.
func (ald *AccessLogDelete) Where(ps ...predicate.AccessLog) *AccessLogDelete {
	ald.mutation.Where(ps...)
	return ald
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ald *AccessLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ald.sqlExec, ald.mutation, ald.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ald *AccessLogDelete) ExecX(ctx context.Context) int {
	n, err := ald.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ald *AccessLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(accesslog.Table, sqlgraph.NewFieldSpec(accesslog.FieldID, field.TypeInt))
	if ps := ald.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ald.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ald.mutation.done = true
	return affected, err
}

// AccessLogDeleteOne is the builder for deleting a single AccessLog entity.
type AccessLogDeleteOne struct {
	ald *AccessLogDelete
}

// Where appends a list predicates to the AccessLogDelete builder.
func (aldo *AccessLogDeleteOne) Where(ps ...predicate.AccessLog) *AccessLogDeleteOne {
	aldo.ald.mutation.Where(ps...)
	return aldo
}

// Exec executes the deletion query.
func (aldo *AccessLogDeleteOne) Exec(ctx context.Context) error {
	n, err := aldo.ald.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accesslog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aldo *AccessLogDeleteOne) ExecX(ctx context.Context) {
	if err := aldo.Exec(ctx); err != nil {
		panic(err)
	}
}
