package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Chat struct {
	ent.Schema
}

func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Ints("users"),
	}
}

func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat_message", ChatMessage.Type),
	}
}
