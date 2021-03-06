package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Commute holds the schema definition for the Commute entity.
type Commute struct {
	ent.Schema
}

// Fields of the Commute.
func (Commute) Fields() []ent.Field {
	return []ent.Field {
		field.String("channel_id"),
		field.String("author_id"),
		field.Time("go_to_work_at").Default(time.Now),
		field.Time("get_off_at").Default(time.Now),
	}
}

// Edges of the Commute.
func (Commute) Edges() []ent.Edge {
	return nil
}
