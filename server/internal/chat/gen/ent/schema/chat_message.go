package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ChatMessage struct {
	ent.Schema
}

func (ChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("userid"),
		field.String("username"),
		field.String("message"),
		field.Time("time").Default(time.Now()),
	}
}

func (ChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Chat.Type).
			Ref("chat_message").
			Unique(),
	}
}
