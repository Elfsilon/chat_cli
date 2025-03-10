// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/internal/chat/gen/ent/chat"
	"server/internal/chat/gen/ent/chatmessage"
	"server/internal/chat/gen/ent/predicate"
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
	TypeChat        = "Chat"
	TypeChatMessage = "ChatMessage"
)

// ChatMutation represents an operation that mutates the Chat nodes in the graph.
type ChatMutation struct {
	config
	op                  Op
	typ                 string
	id                  *int
	name                *string
	users               *[]int
	appendusers         []int
	clearedFields       map[string]struct{}
	chat_message        map[int]struct{}
	removedchat_message map[int]struct{}
	clearedchat_message bool
	done                bool
	oldValue            func(context.Context) (*Chat, error)
	predicates          []predicate.Chat
}

var _ ent.Mutation = (*ChatMutation)(nil)

// chatOption allows management of the mutation configuration using functional options.
type chatOption func(*ChatMutation)

// newChatMutation creates new mutation for the Chat entity.
func newChatMutation(c config, op Op, opts ...chatOption) *ChatMutation {
	m := &ChatMutation{
		config:        c,
		op:            op,
		typ:           TypeChat,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatID sets the ID field of the mutation.
func withChatID(id int) chatOption {
	return func(m *ChatMutation) {
		var (
			err   error
			once  sync.Once
			value *Chat
		)
		m.oldValue = func(ctx context.Context) (*Chat, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Chat.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChat sets the old Chat of the mutation.
func withChat(node *Chat) chatOption {
	return func(m *ChatMutation) {
		m.oldValue = func(context.Context) (*Chat, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Chat.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *ChatMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ChatMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ChatMutation) ResetName() {
	m.name = nil
}

// SetUsers sets the "users" field.
func (m *ChatMutation) SetUsers(i []int) {
	m.users = &i
	m.appendusers = nil
}

// Users returns the value of the "users" field in the mutation.
func (m *ChatMutation) Users() (r []int, exists bool) {
	v := m.users
	if v == nil {
		return
	}
	return *v, true
}

// OldUsers returns the old "users" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldUsers(ctx context.Context) (v []int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsers is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsers requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsers: %w", err)
	}
	return oldValue.Users, nil
}

// AppendUsers adds i to the "users" field.
func (m *ChatMutation) AppendUsers(i []int) {
	m.appendusers = append(m.appendusers, i...)
}

// AppendedUsers returns the list of values that were appended to the "users" field in this mutation.
func (m *ChatMutation) AppendedUsers() ([]int, bool) {
	if len(m.appendusers) == 0 {
		return nil, false
	}
	return m.appendusers, true
}

// ResetUsers resets all changes to the "users" field.
func (m *ChatMutation) ResetUsers() {
	m.users = nil
	m.appendusers = nil
}

// AddChatMessageIDs adds the "chat_message" edge to the ChatMessage entity by ids.
func (m *ChatMutation) AddChatMessageIDs(ids ...int) {
	if m.chat_message == nil {
		m.chat_message = make(map[int]struct{})
	}
	for i := range ids {
		m.chat_message[ids[i]] = struct{}{}
	}
}

// ClearChatMessage clears the "chat_message" edge to the ChatMessage entity.
func (m *ChatMutation) ClearChatMessage() {
	m.clearedchat_message = true
}

// ChatMessageCleared reports if the "chat_message" edge to the ChatMessage entity was cleared.
func (m *ChatMutation) ChatMessageCleared() bool {
	return m.clearedchat_message
}

// RemoveChatMessageIDs removes the "chat_message" edge to the ChatMessage entity by IDs.
func (m *ChatMutation) RemoveChatMessageIDs(ids ...int) {
	if m.removedchat_message == nil {
		m.removedchat_message = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.chat_message, ids[i])
		m.removedchat_message[ids[i]] = struct{}{}
	}
}

// RemovedChatMessage returns the removed IDs of the "chat_message" edge to the ChatMessage entity.
func (m *ChatMutation) RemovedChatMessageIDs() (ids []int) {
	for id := range m.removedchat_message {
		ids = append(ids, id)
	}
	return
}

// ChatMessageIDs returns the "chat_message" edge IDs in the mutation.
func (m *ChatMutation) ChatMessageIDs() (ids []int) {
	for id := range m.chat_message {
		ids = append(ids, id)
	}
	return
}

// ResetChatMessage resets all changes to the "chat_message" edge.
func (m *ChatMutation) ResetChatMessage() {
	m.chat_message = nil
	m.clearedchat_message = false
	m.removedchat_message = nil
}

// Where appends a list predicates to the ChatMutation builder.
func (m *ChatMutation) Where(ps ...predicate.Chat) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Chat, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Chat).
func (m *ChatMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, chat.FieldName)
	}
	if m.users != nil {
		fields = append(fields, chat.FieldUsers)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chat.FieldName:
		return m.Name()
	case chat.FieldUsers:
		return m.Users()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chat.FieldName:
		return m.OldName(ctx)
	case chat.FieldUsers:
		return m.OldUsers(ctx)
	}
	return nil, fmt.Errorf("unknown Chat field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chat.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case chat.FieldUsers:
		v, ok := value.([]int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsers(v)
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Chat numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Chat nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatMutation) ResetField(name string) error {
	switch name {
	case chat.FieldName:
		m.ResetName()
		return nil
	case chat.FieldUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.chat_message != nil {
		edges = append(edges, chat.EdgeChatMessage)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case chat.EdgeChatMessage:
		ids := make([]ent.Value, 0, len(m.chat_message))
		for id := range m.chat_message {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedchat_message != nil {
		edges = append(edges, chat.EdgeChatMessage)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case chat.EdgeChatMessage:
		ids := make([]ent.Value, 0, len(m.removedchat_message))
		for id := range m.removedchat_message {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedchat_message {
		edges = append(edges, chat.EdgeChatMessage)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatMutation) EdgeCleared(name string) bool {
	switch name {
	case chat.EdgeChatMessage:
		return m.clearedchat_message
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Chat unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatMutation) ResetEdge(name string) error {
	switch name {
	case chat.EdgeChatMessage:
		m.ResetChatMessage()
		return nil
	}
	return fmt.Errorf("unknown Chat edge %s", name)
}

// ChatMessageMutation represents an operation that mutates the ChatMessage nodes in the graph.
type ChatMessageMutation struct {
	config
	op            Op
	typ           string
	id            *int
	userid        *int
	adduserid     *int
	username      *string
	message       *string
	time          *time.Time
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*ChatMessage, error)
	predicates    []predicate.ChatMessage
}

var _ ent.Mutation = (*ChatMessageMutation)(nil)

// chatmessageOption allows management of the mutation configuration using functional options.
type chatmessageOption func(*ChatMessageMutation)

// newChatMessageMutation creates new mutation for the ChatMessage entity.
func newChatMessageMutation(c config, op Op, opts ...chatmessageOption) *ChatMessageMutation {
	m := &ChatMessageMutation{
		config:        c,
		op:            op,
		typ:           TypeChatMessage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatMessageID sets the ID field of the mutation.
func withChatMessageID(id int) chatmessageOption {
	return func(m *ChatMessageMutation) {
		var (
			err   error
			once  sync.Once
			value *ChatMessage
		)
		m.oldValue = func(ctx context.Context) (*ChatMessage, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ChatMessage.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChatMessage sets the old ChatMessage of the mutation.
func withChatMessage(node *ChatMessage) chatmessageOption {
	return func(m *ChatMessageMutation) {
		m.oldValue = func(context.Context) (*ChatMessage, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatMessageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatMessageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatMessageMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatMessageMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ChatMessage.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserid sets the "userid" field.
func (m *ChatMessageMutation) SetUserid(i int) {
	m.userid = &i
	m.adduserid = nil
}

// Userid returns the value of the "userid" field in the mutation.
func (m *ChatMessageMutation) Userid() (r int, exists bool) {
	v := m.userid
	if v == nil {
		return
	}
	return *v, true
}

// OldUserid returns the old "userid" field's value of the ChatMessage entity.
// If the ChatMessage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMessageMutation) OldUserid(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserid: %w", err)
	}
	return oldValue.Userid, nil
}

// AddUserid adds i to the "userid" field.
func (m *ChatMessageMutation) AddUserid(i int) {
	if m.adduserid != nil {
		*m.adduserid += i
	} else {
		m.adduserid = &i
	}
}

// AddedUserid returns the value that was added to the "userid" field in this mutation.
func (m *ChatMessageMutation) AddedUserid() (r int, exists bool) {
	v := m.adduserid
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserid resets all changes to the "userid" field.
func (m *ChatMessageMutation) ResetUserid() {
	m.userid = nil
	m.adduserid = nil
}

// SetUsername sets the "username" field.
func (m *ChatMessageMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *ChatMessageMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the ChatMessage entity.
// If the ChatMessage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMessageMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *ChatMessageMutation) ResetUsername() {
	m.username = nil
}

// SetMessage sets the "message" field.
func (m *ChatMessageMutation) SetMessage(s string) {
	m.message = &s
}

// Message returns the value of the "message" field in the mutation.
func (m *ChatMessageMutation) Message() (r string, exists bool) {
	v := m.message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "message" field's value of the ChatMessage entity.
// If the ChatMessage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMessageMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "message" field.
func (m *ChatMessageMutation) ResetMessage() {
	m.message = nil
}

// SetTime sets the "time" field.
func (m *ChatMessageMutation) SetTime(t time.Time) {
	m.time = &t
}

// Time returns the value of the "time" field in the mutation.
func (m *ChatMessageMutation) Time() (r time.Time, exists bool) {
	v := m.time
	if v == nil {
		return
	}
	return *v, true
}

// OldTime returns the old "time" field's value of the ChatMessage entity.
// If the ChatMessage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMessageMutation) OldTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTime: %w", err)
	}
	return oldValue.Time, nil
}

// ResetTime resets all changes to the "time" field.
func (m *ChatMessageMutation) ResetTime() {
	m.time = nil
}

// SetOwnerID sets the "owner" edge to the Chat entity by id.
func (m *ChatMessageMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the Chat entity.
func (m *ChatMessageMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the Chat entity was cleared.
func (m *ChatMessageMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *ChatMessageMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *ChatMessageMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *ChatMessageMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the ChatMessageMutation builder.
func (m *ChatMessageMutation) Where(ps ...predicate.ChatMessage) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatMessageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatMessageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ChatMessage, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatMessageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatMessageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ChatMessage).
func (m *ChatMessageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatMessageMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.userid != nil {
		fields = append(fields, chatmessage.FieldUserid)
	}
	if m.username != nil {
		fields = append(fields, chatmessage.FieldUsername)
	}
	if m.message != nil {
		fields = append(fields, chatmessage.FieldMessage)
	}
	if m.time != nil {
		fields = append(fields, chatmessage.FieldTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatMessageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chatmessage.FieldUserid:
		return m.Userid()
	case chatmessage.FieldUsername:
		return m.Username()
	case chatmessage.FieldMessage:
		return m.Message()
	case chatmessage.FieldTime:
		return m.Time()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatMessageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chatmessage.FieldUserid:
		return m.OldUserid(ctx)
	case chatmessage.FieldUsername:
		return m.OldUsername(ctx)
	case chatmessage.FieldMessage:
		return m.OldMessage(ctx)
	case chatmessage.FieldTime:
		return m.OldTime(ctx)
	}
	return nil, fmt.Errorf("unknown ChatMessage field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMessageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chatmessage.FieldUserid:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserid(v)
		return nil
	case chatmessage.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case chatmessage.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	case chatmessage.FieldTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTime(v)
		return nil
	}
	return fmt.Errorf("unknown ChatMessage field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatMessageMutation) AddedFields() []string {
	var fields []string
	if m.adduserid != nil {
		fields = append(fields, chatmessage.FieldUserid)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatMessageMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case chatmessage.FieldUserid:
		return m.AddedUserid()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMessageMutation) AddField(name string, value ent.Value) error {
	switch name {
	case chatmessage.FieldUserid:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserid(v)
		return nil
	}
	return fmt.Errorf("unknown ChatMessage numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatMessageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatMessageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatMessageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ChatMessage nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatMessageMutation) ResetField(name string) error {
	switch name {
	case chatmessage.FieldUserid:
		m.ResetUserid()
		return nil
	case chatmessage.FieldUsername:
		m.ResetUsername()
		return nil
	case chatmessage.FieldMessage:
		m.ResetMessage()
		return nil
	case chatmessage.FieldTime:
		m.ResetTime()
		return nil
	}
	return fmt.Errorf("unknown ChatMessage field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatMessageMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, chatmessage.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatMessageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case chatmessage.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatMessageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatMessageMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatMessageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, chatmessage.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatMessageMutation) EdgeCleared(name string) bool {
	switch name {
	case chatmessage.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatMessageMutation) ClearEdge(name string) error {
	switch name {
	case chatmessage.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown ChatMessage unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatMessageMutation) ResetEdge(name string) error {
	switch name {
	case chatmessage.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown ChatMessage edge %s", name)
}
