package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Articles holds the schema definition for the Articles entity.
type Articles struct {
	ent.Schema
}

// Fields of the Articles.
func (Articles) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255),
		field.String("content").Optional().Nillable(),
		field.String("image_url").Optional().Nillable(),
		field.Int("like_count").Default(0),
		field.Int("views").Default(0),
		field.Int("user_id"),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Articles.
func (Articles) Edges() []ent.Edge {
	return nil
}
