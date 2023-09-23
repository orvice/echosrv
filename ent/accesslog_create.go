// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.orx.me/echosrv/ent/accesslog"
)

// AccessLogCreate is the builder for creating a AccessLog entity.
type AccessLogCreate struct {
	config
	mutation *AccessLogMutation
	hooks    []Hook
}

// SetCreatedUnix sets the "created_unix" field.
func (alc *AccessLogCreate) SetCreatedUnix(i int) *AccessLogCreate {
	alc.mutation.SetCreatedUnix(i)
	return alc
}

// SetPath sets the "path" field.
func (alc *AccessLogCreate) SetPath(s string) *AccessLogCreate {
	alc.mutation.SetPath(s)
	return alc
}

// SetMethod sets the "method" field.
func (alc *AccessLogCreate) SetMethod(s string) *AccessLogCreate {
	alc.mutation.SetMethod(s)
	return alc
}

// SetIP sets the "ip" field.
func (alc *AccessLogCreate) SetIP(s string) *AccessLogCreate {
	alc.mutation.SetIP(s)
	return alc
}

// SetUa sets the "ua" field.
func (alc *AccessLogCreate) SetUa(s string) *AccessLogCreate {
	alc.mutation.SetUa(s)
	return alc
}

// Mutation returns the AccessLogMutation object of the builder.
func (alc *AccessLogCreate) Mutation() *AccessLogMutation {
	return alc.mutation
}

// Save creates the AccessLog in the database.
func (alc *AccessLogCreate) Save(ctx context.Context) (*AccessLog, error) {
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AccessLogCreate) SaveX(ctx context.Context) *AccessLog {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AccessLogCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AccessLogCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AccessLogCreate) check() error {
	if _, ok := alc.mutation.CreatedUnix(); !ok {
		return &ValidationError{Name: "created_unix", err: errors.New(`ent: missing required field "AccessLog.created_unix"`)}
	}
	if _, ok := alc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "AccessLog.path"`)}
	}
	if _, ok := alc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "AccessLog.method"`)}
	}
	if _, ok := alc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "AccessLog.ip"`)}
	}
	if _, ok := alc.mutation.Ua(); !ok {
		return &ValidationError{Name: "ua", err: errors.New(`ent: missing required field "AccessLog.ua"`)}
	}
	return nil
}

func (alc *AccessLogCreate) sqlSave(ctx context.Context) (*AccessLog, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *AccessLogCreate) createSpec() (*AccessLog, *sqlgraph.CreateSpec) {
	var (
		_node = &AccessLog{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(accesslog.Table, sqlgraph.NewFieldSpec(accesslog.FieldID, field.TypeInt))
	)
	if value, ok := alc.mutation.CreatedUnix(); ok {
		_spec.SetField(accesslog.FieldCreatedUnix, field.TypeInt, value)
		_node.CreatedUnix = value
	}
	if value, ok := alc.mutation.Path(); ok {
		_spec.SetField(accesslog.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := alc.mutation.Method(); ok {
		_spec.SetField(accesslog.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := alc.mutation.IP(); ok {
		_spec.SetField(accesslog.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := alc.mutation.Ua(); ok {
		_spec.SetField(accesslog.FieldUa, field.TypeString, value)
		_node.Ua = value
	}
	return _node, _spec
}

// AccessLogCreateBulk is the builder for creating many AccessLog entities in bulk.
type AccessLogCreateBulk struct {
	config
	err      error
	builders []*AccessLogCreate
}

// Save creates the AccessLog entities in the database.
func (alcb *AccessLogCreateBulk) Save(ctx context.Context) ([]*AccessLog, error) {
	if alcb.err != nil {
		return nil, alcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AccessLog, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccessLogMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AccessLogCreateBulk) SaveX(ctx context.Context) []*AccessLog {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AccessLogCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AccessLogCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}
