// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sample/ent/counter"
	"sample/ent/predicate"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCounter = "Counter"
)

// CounterMutation represents an operation that mutates the Counter nodes in the graph.
type CounterMutation struct {
	config
	op            Op
	typ           string
	id            *int
	count         *int
	addcount      *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Counter, error)
	predicates    []predicate.Counter
}

var _ ent.Mutation = (*CounterMutation)(nil)

// counterOption allows management of the mutation configuration using functional options.
type counterOption func(*CounterMutation)

// newCounterMutation creates new mutation for the Counter entity.
func newCounterMutation(c config, op Op, opts ...counterOption) *CounterMutation {
	m := &CounterMutation{
		config:        c,
		op:            op,
		typ:           TypeCounter,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCounterID sets the ID field of the mutation.
func withCounterID(id int) counterOption {
	return func(m *CounterMutation) {
		var (
			err   error
			once  sync.Once
			value *Counter
		)
		m.oldValue = func(ctx context.Context) (*Counter, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Counter.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCounter sets the old Counter of the mutation.
func withCounter(node *Counter) counterOption {
	return func(m *CounterMutation) {
		m.oldValue = func(context.Context) (*Counter, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CounterMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CounterMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CounterMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CounterMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Counter.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCount sets the "count" field.
func (m *CounterMutation) SetCount(i int) {
	m.count = &i
	m.addcount = nil
}

// Count returns the value of the "count" field in the mutation.
func (m *CounterMutation) Count() (r int, exists bool) {
	v := m.count
	if v == nil {
		return
	}
	return *v, true
}

// OldCount returns the old "count" field's value of the Counter entity.
// If the Counter object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CounterMutation) OldCount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCount: %w", err)
	}
	return oldValue.Count, nil
}

// AddCount adds i to the "count" field.
func (m *CounterMutation) AddCount(i int) {
	if m.addcount != nil {
		*m.addcount += i
	} else {
		m.addcount = &i
	}
}

// AddedCount returns the value that was added to the "count" field in this mutation.
func (m *CounterMutation) AddedCount() (r int, exists bool) {
	v := m.addcount
	if v == nil {
		return
	}
	return *v, true
}

// ResetCount resets all changes to the "count" field.
func (m *CounterMutation) ResetCount() {
	m.count = nil
	m.addcount = nil
}

// Where appends a list predicates to the CounterMutation builder.
func (m *CounterMutation) Where(ps ...predicate.Counter) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CounterMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CounterMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Counter, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CounterMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CounterMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Counter).
func (m *CounterMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CounterMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.count != nil {
		fields = append(fields, counter.FieldCount)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CounterMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case counter.FieldCount:
		return m.Count()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CounterMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case counter.FieldCount:
		return m.OldCount(ctx)
	}
	return nil, fmt.Errorf("unknown Counter field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CounterMutation) SetField(name string, value ent.Value) error {
	switch name {
	case counter.FieldCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCount(v)
		return nil
	}
	return fmt.Errorf("unknown Counter field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CounterMutation) AddedFields() []string {
	var fields []string
	if m.addcount != nil {
		fields = append(fields, counter.FieldCount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CounterMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case counter.FieldCount:
		return m.AddedCount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CounterMutation) AddField(name string, value ent.Value) error {
	switch name {
	case counter.FieldCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCount(v)
		return nil
	}
	return fmt.Errorf("unknown Counter numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CounterMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CounterMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CounterMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Counter nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CounterMutation) ResetField(name string) error {
	switch name {
	case counter.FieldCount:
		m.ResetCount()
		return nil
	}
	return fmt.Errorf("unknown Counter field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CounterMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CounterMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CounterMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CounterMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CounterMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CounterMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CounterMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Counter unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CounterMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Counter edge %s", name)
}
