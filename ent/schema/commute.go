package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/aid95/zaksim-discord-bot/utils/tiktok"
)

// Commute holds the schema definition for the Commute entity.
type Commute struct {
	ent.Schema
}

// Fields of the Commute.
func (Commute) Fields() []ent.Field {
	tz := func() time.Time {
		t, _ := tiktok.LocalNow("Asia/Seoul")
		return t
	}
	return []ent.Field{
		field.String("channel_id"),
		field.String("author_id"),
		field.Time("go_to_work_at").Default(tz),
		field.Time("get_off_at").Default(tz),
	}
}

// Edges of the Commute.
func (Commute) Edges() []ent.Edge {
	return nil
}
