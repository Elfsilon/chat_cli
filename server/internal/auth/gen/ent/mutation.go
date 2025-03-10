// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/internal/auth/gen/ent/predicate"
	"server/internal/auth/gen/ent/session"
	"sync"
	"time"

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
	TypeSession = "Session"
)

// SessionMutation represents an operation that mutates the Session nodes in the graph.
type SessionMutation struct {
	config
	op            Op
	typ           string
	id            *int
	user_id       *int
	adduser_id    *int
	refresh_token *string
	expired_at    *time.Time
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Session, error)
	predicates    []predicate.Session
}

var _ ent.Mutation = (*SessionMutation)(nil)

// sessionOption allows management of the mutation configuration using functional options.
type sessionOption func(*SessionMutation)

// newSessionMutation creates new mutation for the Session entity.
func newSessionMutation(c config, op Op, opts ...sessionOption) *SessionMutation {
	m := &SessionMutation{
		config:        c,
		op:            op,
		typ:           TypeSession,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSessionID sets the ID field of the mutation.
func withSessionID(id int) sessionOption {
	return func(m *SessionMutation) {
		var (
			err   error
			once  sync.Once
			value *Session
		)
		m.oldValue = func(ctx context.Context) (*Session, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Session.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSession sets the old Session of the mutation.
func withSession(node *Session) sessionOption {
	return func(m *SessionMutation) {
		m.oldValue = func(context.Context) (*Session, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SessionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SessionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SessionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SessionMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Session.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserID sets the "user_id" field.
func (m *SessionMutation) SetUserID(i int) {
	m.user_id = &i
	m.adduser_id = nil
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *SessionMutation) UserID() (r int, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldUserID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// AddUserID adds i to the "user_id" field.
func (m *SessionMutation) AddUserID(i int) {
	if m.adduser_id != nil {
		*m.adduser_id += i
	} else {
		m.adduser_id = &i
	}
}

// AddedUserID returns the value that was added to the "user_id" field in this mutation.
func (m *SessionMutation) AddedUserID() (r int, exists bool) {
	v := m.adduser_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserID resets all changes to the "user_id" field.
func (m *SessionMutation) ResetUserID() {
	m.user_id = nil
	m.adduser_id = nil
}

// SetRefreshToken sets the "refresh_token" field.
func (m *SessionMutation) SetRefreshToken(s string) {
	m.refresh_token = &s
}

// RefreshToken returns the value of the "refresh_token" field in the mutation.
func (m *SessionMutation) RefreshToken() (r string, exists bool) {
	v := m.refresh_token
	if v == nil {
		return
	}
	return *v, true
}

// OldRefreshToken returns the old "refresh_token" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldRefreshToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRefreshToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRefreshToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRefreshToken: %w", err)
	}
	return oldValue.RefreshToken, nil
}

// ResetRefreshToken resets all changes to the "refresh_token" field.
func (m *SessionMutation) ResetRefreshToken() {
	m.refresh_token = nil
}

// SetExpiredAt sets the "expired_at" field.
func (m *SessionMutation) SetExpiredAt(t time.Time) {
	m.expired_at = &t
}

// ExpiredAt returns the value of the "expired_at" field in the mutation.
func (m *SessionMutation) ExpiredAt() (r time.Time, exists bool) {
	v := m.expired_at
	if v == nil {
		return
	}
	return *v, true
}

// OldExpiredAt returns the old "expired_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldExpiredAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExpiredAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExpiredAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpiredAt: %w", err)
	}
	return oldValue.ExpiredAt, nil
}

// ResetExpiredAt resets all changes to the "expired_at" field.
func (m *SessionMutation) ResetExpiredAt() {
	m.expired_at = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *SessionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SessionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SessionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the SessionMutation builder.
func (m *SessionMutation) Where(ps ...predicate.Session) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SessionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SessionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Session, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SessionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SessionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Session).
func (m *SessionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SessionMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.user_id != nil {
		fields = append(fields, session.FieldUserID)
	}
	if m.refresh_token != nil {
		fields = append(fields, session.FieldRefreshToken)
	}
	if m.expired_at != nil {
		fields = append(fields, session.FieldExpiredAt)
	}
	if m.created_at != nil {
		fields = append(fields, session.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SessionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case session.FieldUserID:
		return m.UserID()
	case session.FieldRefreshToken:
		return m.RefreshToken()
	case session.FieldExpiredAt:
		return m.ExpiredAt()
	case session.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SessionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case session.FieldUserID:
		return m.OldUserID(ctx)
	case session.FieldRefreshToken:
		return m.OldRefreshToken(ctx)
	case session.FieldExpiredAt:
		return m.OldExpiredAt(ctx)
	case session.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Session field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case session.FieldUserID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case session.FieldRefreshToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRefreshToken(v)
		return nil
	case session.FieldExpiredAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpiredAt(v)
		return nil
	case session.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SessionMutation) AddedFields() []string {
	var fields []string
	if m.adduser_id != nil {
		fields = append(fields, session.FieldUserID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SessionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case session.FieldUserID:
		return m.AddedUserID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) AddField(name string, value ent.Value) error {
	switch name {
	case session.FieldUserID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserID(v)
		return nil
	}
	return fmt.Errorf("unknown Session numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SessionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SessionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SessionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Session nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SessionMutation) ResetField(name string) error {
	switch name {
	case session.FieldUserID:
		m.ResetUserID()
		return nil
	case session.FieldRefreshToken:
		m.ResetRefreshToken()
		return nil
	case session.FieldExpiredAt:
		m.ResetExpiredAt()
		return nil
	case session.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SessionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SessionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SessionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SessionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SessionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SessionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SessionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Session unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SessionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Session edge %s", name)
}
