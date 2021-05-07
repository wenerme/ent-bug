// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wenerme/ent-demo/ent/pet"
	"github.com/wenerme/ent-demo/ent/predicate"
	"github.com/wenerme/ent-demo/ent/user"
	"github.com/wenerme/ent-demo/models"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	hooks    []Hook
	mutation *PetMutation
}

// Where adds a new predicate for the PetUpdate builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetOwnerID sets the "ownerID" field.
func (pu *PetUpdate) SetOwnerID(m models.ID) *PetUpdate {
	pu.mutation.SetOwnerID(m)
	return pu
}

// SetNillableOwnerID sets the "ownerID" field if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(m *models.ID) *PetUpdate {
	if m != nil {
		pu.SetOwnerID(*m)
	}
	return pu
}

// ClearOwnerID clears the value of the "ownerID" field.
func (pu *PetUpdate) ClearOwnerID() *PetUpdate {
	pu.mutation.ClearOwnerID()
	return pu
}

// SetOwnerType sets the "ownerType" field.
func (pu *PetUpdate) SetOwnerType(s string) *PetUpdate {
	pu.mutation.SetOwnerType(s)
	return pu
}

// SetNillableOwnerType sets the "ownerType" field if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerType(s *string) *PetUpdate {
	if s != nil {
		pu.SetOwnerType(*s)
	}
	return pu
}

// ClearOwnerType clears the value of the "ownerType" field.
func (pu *PetUpdate) ClearOwnerType() *PetUpdate {
	pu.mutation.ClearOwnerType()
	return pu
}

// SetOwningUserID sets the "owningUserID" field.
func (pu *PetUpdate) SetOwningUserID(m models.ID) *PetUpdate {
	pu.mutation.SetOwningUserID(m)
	return pu
}

// SetNillableOwningUserID sets the "owningUserID" field if the given value is not nil.
func (pu *PetUpdate) SetNillableOwningUserID(m *models.ID) *PetUpdate {
	if m != nil {
		pu.SetOwningUserID(*m)
	}
	return pu
}

// ClearOwningUserID clears the value of the "owningUserID" field.
func (pu *PetUpdate) ClearOwningUserID() *PetUpdate {
	pu.mutation.ClearOwningUserID()
	return pu
}

// SetOwnerUID sets the "ownerUID" field.
func (pu *PetUpdate) SetOwnerUID(u uuid.UUID) *PetUpdate {
	pu.mutation.SetOwnerUID(u)
	return pu
}

// ClearOwnerUID clears the value of the "ownerUID" field.
func (pu *PetUpdate) ClearOwnerUID() *PetUpdate {
	pu.mutation.ClearOwnerUID()
	return pu
}

// SetName sets the "name" field.
func (pu *PetUpdate) SetName(s string) *PetUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetOwningUser sets the "owningUser" edge to the User entity.
func (pu *PetUpdate) SetOwningUser(u *User) *PetUpdate {
	return pu.SetOwningUserID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pu *PetUpdate) Mutation() *PetMutation {
	return pu.mutation
}

// ClearOwningUser clears the "owningUser" edge to the User entity.
func (pu *PetUpdate) ClearOwningUser() *PetUpdate {
	pu.mutation.ClearOwningUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pet.Table,
			Columns: pet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: pet.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.OwnerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldOwnerID,
		})
	}
	if pu.mutation.OwnerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pet.FieldOwnerID,
		})
	}
	if value, ok := pu.mutation.OwnerType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldOwnerType,
		})
	}
	if pu.mutation.OwnerTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pet.FieldOwnerType,
		})
	}
	if value, ok := pu.mutation.OwnerUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pet.FieldOwnerUID,
		})
	}
	if pu.mutation.OwnerUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: pet.FieldOwnerUID,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldName,
		})
	}
	if pu.mutation.OwningUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pet.OwningUserTable,
			Columns: []string{pet.OwningUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwningUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pet.OwningUserTable,
			Columns: []string{pet.OwningUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PetMutation
}

// SetOwnerID sets the "ownerID" field.
func (puo *PetUpdateOne) SetOwnerID(m models.ID) *PetUpdateOne {
	puo.mutation.SetOwnerID(m)
	return puo
}

// SetNillableOwnerID sets the "ownerID" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(m *models.ID) *PetUpdateOne {
	if m != nil {
		puo.SetOwnerID(*m)
	}
	return puo
}

// ClearOwnerID clears the value of the "ownerID" field.
func (puo *PetUpdateOne) ClearOwnerID() *PetUpdateOne {
	puo.mutation.ClearOwnerID()
	return puo
}

// SetOwnerType sets the "ownerType" field.
func (puo *PetUpdateOne) SetOwnerType(s string) *PetUpdateOne {
	puo.mutation.SetOwnerType(s)
	return puo
}

// SetNillableOwnerType sets the "ownerType" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerType(s *string) *PetUpdateOne {
	if s != nil {
		puo.SetOwnerType(*s)
	}
	return puo
}

// ClearOwnerType clears the value of the "ownerType" field.
func (puo *PetUpdateOne) ClearOwnerType() *PetUpdateOne {
	puo.mutation.ClearOwnerType()
	return puo
}

// SetOwningUserID sets the "owningUserID" field.
func (puo *PetUpdateOne) SetOwningUserID(m models.ID) *PetUpdateOne {
	puo.mutation.SetOwningUserID(m)
	return puo
}

// SetNillableOwningUserID sets the "owningUserID" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwningUserID(m *models.ID) *PetUpdateOne {
	if m != nil {
		puo.SetOwningUserID(*m)
	}
	return puo
}

// ClearOwningUserID clears the value of the "owningUserID" field.
func (puo *PetUpdateOne) ClearOwningUserID() *PetUpdateOne {
	puo.mutation.ClearOwningUserID()
	return puo
}

// SetOwnerUID sets the "ownerUID" field.
func (puo *PetUpdateOne) SetOwnerUID(u uuid.UUID) *PetUpdateOne {
	puo.mutation.SetOwnerUID(u)
	return puo
}

// ClearOwnerUID clears the value of the "ownerUID" field.
func (puo *PetUpdateOne) ClearOwnerUID() *PetUpdateOne {
	puo.mutation.ClearOwnerUID()
	return puo
}

// SetName sets the "name" field.
func (puo *PetUpdateOne) SetName(s string) *PetUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetOwningUser sets the "owningUser" edge to the User entity.
func (puo *PetUpdateOne) SetOwningUser(u *User) *PetUpdateOne {
	return puo.SetOwningUserID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (puo *PetUpdateOne) Mutation() *PetMutation {
	return puo.mutation
}

// ClearOwningUser clears the "owningUser" edge to the User entity.
func (puo *PetUpdateOne) ClearOwningUser() *PetUpdateOne {
	puo.mutation.ClearOwningUser()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PetUpdateOne) Select(field string, fields ...string) *PetUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pet entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	var (
		err  error
		node *Pet
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) sqlSave(ctx context.Context) (_node *Pet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pet.Table,
			Columns: pet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: pet.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Pet.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pet.FieldID)
		for _, f := range fields {
			if !pet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.OwnerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldOwnerID,
		})
	}
	if puo.mutation.OwnerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pet.FieldOwnerID,
		})
	}
	if value, ok := puo.mutation.OwnerType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldOwnerType,
		})
	}
	if puo.mutation.OwnerTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pet.FieldOwnerType,
		})
	}
	if value, ok := puo.mutation.OwnerUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pet.FieldOwnerUID,
		})
	}
	if puo.mutation.OwnerUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: pet.FieldOwnerUID,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldName,
		})
	}
	if puo.mutation.OwningUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pet.OwningUserTable,
			Columns: []string{pet.OwningUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwningUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pet.OwningUserTable,
			Columns: []string{pet.OwningUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Pet{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
