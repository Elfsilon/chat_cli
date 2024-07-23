package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("email").Unique(),
		field.String("password").Sensitive(),
		field.Int("role"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
