// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.orx.me/echosrv/ent/accesslog"
	"go.orx.me/echosrv/ent/predicate"
)

// AccessLogUpdate is the builder for updating AccessLog entities.
type AccessLogUpdate struct {
	config
	hooks    []Hook
	mutation *AccessLogMutation
}

// Where appends a list predicates to the AccessLogUpdate builder.
func (alu *AccessLogUpdate) Where(ps ...predicate.AccessLog) *AccessLogUpdate {
	alu.mutation.Where(ps...)
	return alu
}

// SetCreatedUnix sets the "created_unix" field.
func (alu *AccessLogUpdate) SetCreatedUnix(i int) *AccessLogUpdate {
	alu.mutation.ResetCreatedUnix()
	alu.mutation.SetCreatedUnix(i)
	return alu
}

// AddCreatedUnix adds i to the "created_unix" field.
func (alu *AccessLogUpdate) AddCreatedUnix(i int) *AccessLogUpdate {
	alu.mutation.AddCreatedUnix(i)
	return alu
}

// SetPath sets the "path" field.
func (alu *AccessLogUpdate) SetPath(s string) *AccessLogUpdate {
	alu.mutation.SetPath(s)
	return alu
}

// SetMethod sets the "method" field.
func (alu *AccessLogUpdate) SetMethod(s string) *AccessLogUpdate {
	alu.mutation.SetMethod(s)
	return alu
}

// SetIP sets the "ip" field.
func (alu *AccessLogUpdate) SetIP(s string) *AccessLogUpdate {
	alu.mutation.SetIP(s)
	return alu
}

// SetUa sets the "ua" field.
func (alu *AccessLogUpdate) SetUa(s string) *AccessLogUpdate {
	alu.mutation.SetUa(s)
	return alu
}

// SetTrace sets the "trace" field.
func (alu *AccessLogUpdate) SetTrace(s string) *AccessLogUpdate {
	alu.mutation.SetTrace(s)
	return alu
}

// Mutation returns the AccessLogMutation object of the builder.
func (alu *AccessLogUpdate) Mutation() *AccessLogMutation {
	return alu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (alu *AccessLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, alu.sqlSave, alu.mutation, alu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (alu *AccessLogUpdate) SaveX(ctx context.Context) int {
	affected, err := alu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (alu *AccessLogUpdate) Exec(ctx context.Context) error {
	_, err := alu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alu *AccessLogUpdate) ExecX(ctx context.Context) {
	if err := alu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (alu *AccessLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesslog.Table, accesslog.Columns, sqlgraph.NewFieldSpec(accesslog.FieldID, field.TypeInt))
	if ps := alu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := alu.mutation.CreatedUnix(); ok {
		_spec.SetField(accesslog.FieldCreatedUnix, field.TypeInt, value)
	}
	if value, ok := alu.mutation.AddedCreatedUnix(); ok {
		_spec.AddField(accesslog.FieldCreatedUnix, field.TypeInt, value)
	}
	if value, ok := alu.mutation.Path(); ok {
		_spec.SetField(accesslog.FieldPath, field.TypeString, value)
	}
	if value, ok := alu.mutation.Method(); ok {
		_spec.SetField(accesslog.FieldMethod, field.TypeString, value)
	}
	if value, ok := alu.mutation.IP(); ok {
		_spec.SetField(accesslog.FieldIP, field.TypeString, value)
	}
	if value, ok := alu.mutation.Ua(); ok {
		_spec.SetField(accesslog.FieldUa, field.TypeString, value)
	}
	if value, ok := alu.mutation.Trace(); ok {
		_spec.SetField(accesslog.FieldTrace, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, alu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesslog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	alu.mutation.done = true
	return n, nil
}

// AccessLogUpdateOne is the builder for updating a single AccessLog entity.
type AccessLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessLogMutation
}

// SetCreatedUnix sets the "created_unix" field.
func (aluo *AccessLogUpdateOne) SetCreatedUnix(i int) *AccessLogUpdateOne {
	aluo.mutation.ResetCreatedUnix()
	aluo.mutation.SetCreatedUnix(i)
	return aluo
}

// AddCreatedUnix adds i to the "created_unix" field.
func (aluo *AccessLogUpdateOne) AddCreatedUnix(i int) *AccessLogUpdateOne {
	aluo.mutation.AddCreatedUnix(i)
	return aluo
}

// SetPath sets the "path" field.
func (aluo *AccessLogUpdateOne) SetPath(s string) *AccessLogUpdateOne {
	aluo.mutation.SetPath(s)
	return aluo
}

// SetMethod sets the "method" field.
func (aluo *AccessLogUpdateOne) SetMethod(s string) *AccessLogUpdateOne {
	aluo.mutation.SetMethod(s)
	return aluo
}

// SetIP sets the "ip" field.
func (aluo *AccessLogUpdateOne) SetIP(s string) *AccessLogUpdateOne {
	aluo.mutation.SetIP(s)
	return aluo
}

// SetUa sets the "ua" field.
func (aluo *AccessLogUpdateOne) SetUa(s string) *AccessLogUpdateOne {
	aluo.mutation.SetUa(s)
	return aluo
}

// SetTrace sets the "trace" field.
func (aluo *AccessLogUpdateOne) SetTrace(s string) *AccessLogUpdateOne {
	aluo.mutation.SetTrace(s)
	return aluo
}

// Mutation returns the AccessLogMutation object of the builder.
func (aluo *AccessLogUpdateOne) Mutation() *AccessLogMutation {
	return aluo.mutation
}

// Where appends a list predicates to the AccessLogUpdate builder.
func (aluo *AccessLogUpdateOne) Where(ps ...predicate.AccessLog) *AccessLogUpdateOne {
	aluo.mutation.Where(ps...)
	return aluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aluo *AccessLogUpdateOne) Select(field string, fields ...string) *AccessLogUpdateOne {
	aluo.fields = append([]string{field}, fields...)
	return aluo
}

// Save executes the query and returns the updated AccessLog entity.
func (aluo *AccessLogUpdateOne) Save(ctx context.Context) (*AccessLog, error) {
	return withHooks(ctx, aluo.sqlSave, aluo.mutation, aluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aluo *AccessLogUpdateOne) SaveX(ctx context.Context) *AccessLog {
	node, err := aluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aluo *AccessLogUpdateOne) Exec(ctx context.Context) error {
	_, err := aluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aluo *AccessLogUpdateOne) ExecX(ctx context.Context) {
	if err := aluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aluo *AccessLogUpdateOne) sqlSave(ctx context.Context) (_node *AccessLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesslog.Table, accesslog.Columns, sqlgraph.NewFieldSpec(accesslog.FieldID, field.TypeInt))
	id, ok := aluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccessLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesslog.FieldID)
		for _, f := range fields {
			if !accesslog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accesslog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aluo.mutation.CreatedUnix(); ok {
		_spec.SetField(accesslog.FieldCreatedUnix, field.TypeInt, value)
	}
	if value, ok := aluo.mutation.AddedCreatedUnix(); ok {
		_spec.AddField(accesslog.FieldCreatedUnix, field.TypeInt, value)
	}
	if value, ok := aluo.mutation.Path(); ok {
		_spec.SetField(accesslog.FieldPath, field.TypeString, value)
	}
	if value, ok := aluo.mutation.Method(); ok {
		_spec.SetField(accesslog.FieldMethod, field.TypeString, value)
	}
	if value, ok := aluo.mutation.IP(); ok {
		_spec.SetField(accesslog.FieldIP, field.TypeString, value)
	}
	if value, ok := aluo.mutation.Ua(); ok {
		_spec.SetField(accesslog.FieldUa, field.TypeString, value)
	}
	if value, ok := aluo.mutation.Trace(); ok {
		_spec.SetField(accesslog.FieldTrace, field.TypeString, value)
	}
	_node = &AccessLog{config: aluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesslog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aluo.mutation.done = true
	return _node, nil
}