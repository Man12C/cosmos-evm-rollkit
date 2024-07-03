// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omni-network/omni/explorer/db/ent/predicate"
	"github.com/omni-network/omni/explorer/db/ent/xprovidercursor"
)

// XProviderCursorUpdate is the builder for updating XProviderCursor entities.
type XProviderCursorUpdate struct {
	config
	hooks    []Hook
	mutation *XProviderCursorMutation
}

// Where appends a list predicates to the XProviderCursorUpdate builder.
func (xcu *XProviderCursorUpdate) Where(ps ...predicate.XProviderCursor) *XProviderCursorUpdate {
	xcu.mutation.Where(ps...)
	return xcu
}

// SetChainID sets the "chain_id" field.
func (xcu *XProviderCursorUpdate) SetChainID(u uint64) *XProviderCursorUpdate {
	xcu.mutation.ResetChainID()
	xcu.mutation.SetChainID(u)
	return xcu
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (xcu *XProviderCursorUpdate) SetNillableChainID(u *uint64) *XProviderCursorUpdate {
	if u != nil {
		xcu.SetChainID(*u)
	}
	return xcu
}

// AddChainID adds u to the "chain_id" field.
func (xcu *XProviderCursorUpdate) AddChainID(u int64) *XProviderCursorUpdate {
	xcu.mutation.AddChainID(u)
	return xcu
}

// SetHeight sets the "height" field.
func (xcu *XProviderCursorUpdate) SetHeight(u uint64) *XProviderCursorUpdate {
	xcu.mutation.ResetHeight()
	xcu.mutation.SetHeight(u)
	return xcu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (xcu *XProviderCursorUpdate) SetNillableHeight(u *uint64) *XProviderCursorUpdate {
	if u != nil {
		xcu.SetHeight(*u)
	}
	return xcu
}

// AddHeight adds u to the "height" field.
func (xcu *XProviderCursorUpdate) AddHeight(u int64) *XProviderCursorUpdate {
	xcu.mutation.AddHeight(u)
	return xcu
}

// SetOffset sets the "offset" field.
func (xcu *XProviderCursorUpdate) SetOffset(u uint64) *XProviderCursorUpdate {
	xcu.mutation.ResetOffset()
	xcu.mutation.SetOffset(u)
	return xcu
}

// SetNillableOffset sets the "offset" field if the given value is not nil.
func (xcu *XProviderCursorUpdate) SetNillableOffset(u *uint64) *XProviderCursorUpdate {
	if u != nil {
		xcu.SetOffset(*u)
	}
	return xcu
}

// AddOffset adds u to the "offset" field.
func (xcu *XProviderCursorUpdate) AddOffset(u int64) *XProviderCursorUpdate {
	xcu.mutation.AddOffset(u)
	return xcu
}

// SetCreatedAt sets the "created_at" field.
func (xcu *XProviderCursorUpdate) SetCreatedAt(t time.Time) *XProviderCursorUpdate {
	xcu.mutation.SetCreatedAt(t)
	return xcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (xcu *XProviderCursorUpdate) SetNillableCreatedAt(t *time.Time) *XProviderCursorUpdate {
	if t != nil {
		xcu.SetCreatedAt(*t)
	}
	return xcu
}

// SetUpdatedAt sets the "updated_at" field.
func (xcu *XProviderCursorUpdate) SetUpdatedAt(t time.Time) *XProviderCursorUpdate {
	xcu.mutation.SetUpdatedAt(t)
	return xcu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (xcu *XProviderCursorUpdate) SetNillableUpdatedAt(t *time.Time) *XProviderCursorUpdate {
	if t != nil {
		xcu.SetUpdatedAt(*t)
	}
	return xcu
}

// Mutation returns the XProviderCursorMutation object of the builder.
func (xcu *XProviderCursorUpdate) Mutation() *XProviderCursorMutation {
	return xcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (xcu *XProviderCursorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, xcu.sqlSave, xcu.mutation, xcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (xcu *XProviderCursorUpdate) SaveX(ctx context.Context) int {
	affected, err := xcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (xcu *XProviderCursorUpdate) Exec(ctx context.Context) error {
	_, err := xcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xcu *XProviderCursorUpdate) ExecX(ctx context.Context) {
	if err := xcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (xcu *XProviderCursorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(xprovidercursor.Table, xprovidercursor.Columns, sqlgraph.NewFieldSpec(xprovidercursor.FieldID, field.TypeUUID))
	if ps := xcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xcu.mutation.ChainID(); ok {
		_spec.SetField(xprovidercursor.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := xcu.mutation.AddedChainID(); ok {
		_spec.AddField(xprovidercursor.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := xcu.mutation.Height(); ok {
		_spec.SetField(xprovidercursor.FieldHeight, field.TypeUint64, value)
	}
	if value, ok := xcu.mutation.AddedHeight(); ok {
		_spec.AddField(xprovidercursor.FieldHeight, field.TypeUint64, value)
	}
	if value, ok := xcu.mutation.Offset(); ok {
		_spec.SetField(xprovidercursor.FieldOffset, field.TypeUint64, value)
	}
	if value, ok := xcu.mutation.AddedOffset(); ok {
		_spec.AddField(xprovidercursor.FieldOffset, field.TypeUint64, value)
	}
	if value, ok := xcu.mutation.CreatedAt(); ok {
		_spec.SetField(xprovidercursor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := xcu.mutation.UpdatedAt(); ok {
		_spec.SetField(xprovidercursor.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, xcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xprovidercursor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	xcu.mutation.done = true
	return n, nil
}

// XProviderCursorUpdateOne is the builder for updating a single XProviderCursor entity.
type XProviderCursorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *XProviderCursorMutation
}

// SetChainID sets the "chain_id" field.
func (xcuo *XProviderCursorUpdateOne) SetChainID(u uint64) *XProviderCursorUpdateOne {
	xcuo.mutation.ResetChainID()
	xcuo.mutation.SetChainID(u)
	return xcuo
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (xcuo *XProviderCursorUpdateOne) SetNillableChainID(u *uint64) *XProviderCursorUpdateOne {
	if u != nil {
		xcuo.SetChainID(*u)
	}
	return xcuo
}

// AddChainID adds u to the "chain_id" field.
func (xcuo *XProviderCursorUpdateOne) AddChainID(u int64) *XProviderCursorUpdateOne {
	xcuo.mutation.AddChainID(u)
	return xcuo
}

// SetHeight sets the "height" field.
func (xcuo *XProviderCursorUpdateOne) SetHeight(u uint64) *XProviderCursorUpdateOne {
	xcuo.mutation.ResetHeight()
	xcuo.mutation.SetHeight(u)
	return xcuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (xcuo *XProviderCursorUpdateOne) SetNillableHeight(u *uint64) *XProviderCursorUpdateOne {
	if u != nil {
		xcuo.SetHeight(*u)
	}
	return xcuo
}

// AddHeight adds u to the "height" field.
func (xcuo *XProviderCursorUpdateOne) AddHeight(u int64) *XProviderCursorUpdateOne {
	xcuo.mutation.AddHeight(u)
	return xcuo
}

// SetOffset sets the "offset" field.
func (xcuo *XProviderCursorUpdateOne) SetOffset(u uint64) *XProviderCursorUpdateOne {
	xcuo.mutation.ResetOffset()
	xcuo.mutation.SetOffset(u)
	return xcuo
}

// SetNillableOffset sets the "offset" field if the given value is not nil.
func (xcuo *XProviderCursorUpdateOne) SetNillableOffset(u *uint64) *XProviderCursorUpdateOne {
	if u != nil {
		xcuo.SetOffset(*u)
	}
	return xcuo
}

// AddOffset adds u to the "offset" field.
func (xcuo *XProviderCursorUpdateOne) AddOffset(u int64) *XProviderCursorUpdateOne {
	xcuo.mutation.AddOffset(u)
	return xcuo
}

// SetCreatedAt sets the "created_at" field.
func (xcuo *XProviderCursorUpdateOne) SetCreatedAt(t time.Time) *XProviderCursorUpdateOne {
	xcuo.mutation.SetCreatedAt(t)
	return xcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (xcuo *XProviderCursorUpdateOne) SetNillableCreatedAt(t *time.Time) *XProviderCursorUpdateOne {
	if t != nil {
		xcuo.SetCreatedAt(*t)
	}
	return xcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (xcuo *XProviderCursorUpdateOne) SetUpdatedAt(t time.Time) *XProviderCursorUpdateOne {
	xcuo.mutation.SetUpdatedAt(t)
	return xcuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (xcuo *XProviderCursorUpdateOne) SetNillableUpdatedAt(t *time.Time) *XProviderCursorUpdateOne {
	if t != nil {
		xcuo.SetUpdatedAt(*t)
	}
	return xcuo
}

// Mutation returns the XProviderCursorMutation object of the builder.
func (xcuo *XProviderCursorUpdateOne) Mutation() *XProviderCursorMutation {
	return xcuo.mutation
}

// Where appends a list predicates to the XProviderCursorUpdate builder.
func (xcuo *XProviderCursorUpdateOne) Where(ps ...predicate.XProviderCursor) *XProviderCursorUpdateOne {
	xcuo.mutation.Where(ps...)
	return xcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (xcuo *XProviderCursorUpdateOne) Select(field string, fields ...string) *XProviderCursorUpdateOne {
	xcuo.fields = append([]string{field}, fields...)
	return xcuo
}

// Save executes the query and returns the updated XProviderCursor entity.
func (xcuo *XProviderCursorUpdateOne) Save(ctx context.Context) (*XProviderCursor, error) {
	return withHooks(ctx, xcuo.sqlSave, xcuo.mutation, xcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (xcuo *XProviderCursorUpdateOne) SaveX(ctx context.Context) *XProviderCursor {
	node, err := xcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (xcuo *XProviderCursorUpdateOne) Exec(ctx context.Context) error {
	_, err := xcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xcuo *XProviderCursorUpdateOne) ExecX(ctx context.Context) {
	if err := xcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (xcuo *XProviderCursorUpdateOne) sqlSave(ctx context.Context) (_node *XProviderCursor, err error) {
	_spec := sqlgraph.NewUpdateSpec(xprovidercursor.Table, xprovidercursor.Columns, sqlgraph.NewFieldSpec(xprovidercursor.FieldID, field.TypeUUID))
	id, ok := xcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "XProviderCursor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := xcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xprovidercursor.FieldID)
		for _, f := range fields {
			if !xprovidercursor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != xprovidercursor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := xcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xcuo.mutation.ChainID(); ok {
		_spec.SetField(xprovidercursor.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := xcuo.mutation.AddedChainID(); ok {
		_spec.AddField(xprovidercursor.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := xcuo.mutation.Height(); ok {
		_spec.SetField(xprovidercursor.FieldHeight, field.TypeUint64, value)
	}
	if value, ok := xcuo.mutation.AddedHeight(); ok {
		_spec.AddField(xprovidercursor.FieldHeight, field.TypeUint64, value)
	}
	if value, ok := xcuo.mutation.Offset(); ok {
		_spec.SetField(xprovidercursor.FieldOffset, field.TypeUint64, value)
	}
	if value, ok := xcuo.mutation.AddedOffset(); ok {
		_spec.AddField(xprovidercursor.FieldOffset, field.TypeUint64, value)
	}
	if value, ok := xcuo.mutation.CreatedAt(); ok {
		_spec.SetField(xprovidercursor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := xcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(xprovidercursor.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &XProviderCursor{config: xcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, xcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xprovidercursor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	xcuo.mutation.done = true
	return _node, nil
}