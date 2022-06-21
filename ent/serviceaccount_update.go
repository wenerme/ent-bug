// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wenerme/ent-demo/ent/predicate"
	"github.com/wenerme/ent-demo/ent/serviceaccount"
)

// ServiceAccountUpdate is the builder for updating ServiceAccount entities.
type ServiceAccountUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceAccountMutation
}

// Where appends a list predicates to the ServiceAccountUpdate builder.
func (sau *ServiceAccountUpdate) Where(ps ...predicate.ServiceAccount) *ServiceAccountUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetDisplayName sets the "displayName" field.
func (sau *ServiceAccountUpdate) SetDisplayName(s string) *ServiceAccountUpdate {
	sau.mutation.SetDisplayName(s)
	return sau
}

// SetDisabled sets the "disabled" field.
func (sau *ServiceAccountUpdate) SetDisabled(b bool) *ServiceAccountUpdate {
	sau.mutation.SetDisabled(b)
	return sau
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (sau *ServiceAccountUpdate) SetNillableDisabled(b *bool) *ServiceAccountUpdate {
	if b != nil {
		sau.SetDisabled(*b)
	}
	return sau
}

// SetUsername sets the "username" field.
func (sau *ServiceAccountUpdate) SetUsername(s string) *ServiceAccountUpdate {
	sau.mutation.SetUsername(s)
	return sau
}

// SetPassword sets the "password" field.
func (sau *ServiceAccountUpdate) SetPassword(s string) *ServiceAccountUpdate {
	sau.mutation.SetPassword(s)
	return sau
}

// Mutation returns the ServiceAccountMutation object of the builder.
func (sau *ServiceAccountUpdate) Mutation() *ServiceAccountMutation {
	return sau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *ServiceAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sau.hooks) == 0 {
		if err = sau.check(); err != nil {
			return 0, err
		}
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sau.check(); err != nil {
				return 0, err
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *ServiceAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *ServiceAccountUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *ServiceAccountUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sau *ServiceAccountUpdate) check() error {
	if v, ok := sau.mutation.DisplayName(); ok {
		if err := serviceaccount.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "displayName", err: fmt.Errorf(`ent: validator failed for field "ServiceAccount.displayName": %w`, err)}
		}
	}
	return nil
}

func (sau *ServiceAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceaccount.Table,
			Columns: serviceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: serviceaccount.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldDisplayName,
		})
	}
	if value, ok := sau.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: serviceaccount.FieldDisabled,
		})
	}
	if value, ok := sau.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldUsername,
		})
	}
	if value, ok := sau.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldPassword,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ServiceAccountUpdateOne is the builder for updating a single ServiceAccount entity.
type ServiceAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceAccountMutation
}

// SetDisplayName sets the "displayName" field.
func (sauo *ServiceAccountUpdateOne) SetDisplayName(s string) *ServiceAccountUpdateOne {
	sauo.mutation.SetDisplayName(s)
	return sauo
}

// SetDisabled sets the "disabled" field.
func (sauo *ServiceAccountUpdateOne) SetDisabled(b bool) *ServiceAccountUpdateOne {
	sauo.mutation.SetDisabled(b)
	return sauo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (sauo *ServiceAccountUpdateOne) SetNillableDisabled(b *bool) *ServiceAccountUpdateOne {
	if b != nil {
		sauo.SetDisabled(*b)
	}
	return sauo
}

// SetUsername sets the "username" field.
func (sauo *ServiceAccountUpdateOne) SetUsername(s string) *ServiceAccountUpdateOne {
	sauo.mutation.SetUsername(s)
	return sauo
}

// SetPassword sets the "password" field.
func (sauo *ServiceAccountUpdateOne) SetPassword(s string) *ServiceAccountUpdateOne {
	sauo.mutation.SetPassword(s)
	return sauo
}

// Mutation returns the ServiceAccountMutation object of the builder.
func (sauo *ServiceAccountUpdateOne) Mutation() *ServiceAccountMutation {
	return sauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *ServiceAccountUpdateOne) Select(field string, fields ...string) *ServiceAccountUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated ServiceAccount entity.
func (sauo *ServiceAccountUpdateOne) Save(ctx context.Context) (*ServiceAccount, error) {
	var (
		err  error
		node *ServiceAccount
	)
	if len(sauo.hooks) == 0 {
		if err = sauo.check(); err != nil {
			return nil, err
		}
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sauo.check(); err != nil {
				return nil, err
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ServiceAccount)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ServiceAccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *ServiceAccountUpdateOne) SaveX(ctx context.Context) *ServiceAccount {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *ServiceAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *ServiceAccountUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sauo *ServiceAccountUpdateOne) check() error {
	if v, ok := sauo.mutation.DisplayName(); ok {
		if err := serviceaccount.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "displayName", err: fmt.Errorf(`ent: validator failed for field "ServiceAccount.displayName": %w`, err)}
		}
	}
	return nil
}

func (sauo *ServiceAccountUpdateOne) sqlSave(ctx context.Context) (_node *ServiceAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceaccount.Table,
			Columns: serviceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: serviceaccount.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ServiceAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceaccount.FieldID)
		for _, f := range fields {
			if !serviceaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != serviceaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldDisplayName,
		})
	}
	if value, ok := sauo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: serviceaccount.FieldDisabled,
		})
	}
	if value, ok := sauo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldUsername,
		})
	}
	if value, ok := sauo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldPassword,
		})
	}
	_node = &ServiceAccount{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}