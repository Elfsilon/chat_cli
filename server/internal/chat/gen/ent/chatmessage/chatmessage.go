// Code generated by ent, DO NOT EDIT.

package chatmessage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the chatmessage type in the database.
	Label = "chat_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserid holds the string denoting the userid field in the database.
	FieldUserid = "userid"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the chatmessage in the database.
	Table = "chat_messages"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "chat_messages"
	// OwnerInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	OwnerInverseTable = "chats"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "chat_chat_message"
)

// Columns holds all SQL columns for chatmessage fields.
var Columns = []string{
	FieldID,
	FieldUserid,
	FieldUsername,
	FieldMessage,
	FieldTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "chat_messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"chat_chat_message",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTime holds the default value on creation for the "time" field.
	DefaultTime time.Time
)

// OrderOption defines the ordering options for the ChatMessage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserid orders the results by the userid field.
func ByUserid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserid, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
