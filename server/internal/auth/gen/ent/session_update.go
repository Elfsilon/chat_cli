// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/internal/auth/gen/ent/predicate"
	"server/internal/auth/gen/ent/session"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUserID sets the "user_id" field.
func (su *SessionUpdate) SetUserID(i int) *SessionUpdate {
	su.mutation.ResetUserID()
	su.mutation.SetUserID(i)
	return su
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserID(i *int) *SessionUpdate {
	if i != nil {
		su.SetUserID(*i)
	}
	return su
}

// AddUserID adds i to the "user_id" field.
func (su *SessionUpdate) AddUserID(i int) *SessionUpdate {
	su.mutation.AddUserID(i)
	return su
}

// SetRefreshToken sets the "refresh_token" field.
func (su *SessionUpdate) SetRefreshToken(s string) *SessionUpdate {
	su.mutation.SetRefreshToken(s)
	return su
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (su *SessionUpdate) SetNillableRefreshToken(s *string) *SessionUpdate {
	if s != nil {
		su.SetRefreshToken(*s)
	}
	return su
}

// SetExpiredAt sets the "expired_at" field.
func (su *SessionUpdate) SetExpiredAt(t time.Time) *SessionUpdate {
	su.mutation.SetExpiredAt(t)
	return su
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableExpiredAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetExpiredAt(*t)
	}
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SessionUpdate) SetCreatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UserID(); ok {
		_spec.SetField(session.FieldUserID, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedUserID(); ok {
		_spec.AddField(session.FieldUserID, field.TypeInt, value)
	}
	if value, ok := su.mutation.RefreshToken(); ok {
		_spec.SetField(session.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := su.mutation.ExpiredAt(); ok {
		_spec.SetField(session.FieldExpiredAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetUserID sets the "user_id" field.
func (suo *SessionUpdateOne) SetUserID(i int) *SessionUpdateOne {
	suo.mutation.ResetUserID()
	suo.mutation.SetUserID(i)
	return suo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserID(i *int) *SessionUpdateOne {
	if i != nil {
		suo.SetUserID(*i)
	}
	return suo
}

// AddUserID adds i to the "user_id" field.
func (suo *SessionUpdateOne) AddUserID(i int) *SessionUpdateOne {
	suo.mutation.AddUserID(i)
	return suo
}

// SetRefreshToken sets the "refresh_token" field.
func (suo *SessionUpdateOne) SetRefreshToken(s string) *SessionUpdateOne {
	suo.mutation.SetRefreshToken(s)
	return suo
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableRefreshToken(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetRefreshToken(*s)
	}
	return suo
}

// SetExpiredAt sets the "expired_at" field.
func (suo *SessionUpdateOne) SetExpiredAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetExpiredAt(t)
	return suo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableExpiredAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetExpiredAt(*t)
	}
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SessionUpdateOne) SetCreatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UserID(); ok {
		_spec.SetField(session.FieldUserID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedUserID(); ok {
		_spec.AddField(session.FieldUserID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.RefreshToken(); ok {
		_spec.SetField(session.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := suo.mutation.ExpiredAt(); ok {
		_spec.SetField(session.FieldExpiredAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
