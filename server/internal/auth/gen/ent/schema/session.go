package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("refresh_token").Unique(),
		field.Time("expired_at"),
		field.Time("created_at").Default(time.Now),
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{}
}
