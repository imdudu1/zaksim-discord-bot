// Code generated by entc, DO NOT EDIT.

package commute

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/aid95/zaksim-discord-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorID), v))
	})
}

// GoToWorkAt applies equality check predicate on the "go_to_work_at" field. It's identical to GoToWorkAtEQ.
func GoToWorkAt(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoToWorkAt), v))
	})
}

// GetOffAt applies equality check predicate on the "get_off_at" field. It's identical to GetOffAtEQ.
func GetOffAt(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGetOffAt), v))
	})
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...string) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChannelID), v...))
	})
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...string) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChannelID), v...))
	})
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChannelID), v))
	})
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChannelID), v))
	})
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChannelID), v))
	})
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChannelID), v))
	})
}

// ChannelIDContains applies the Contains predicate on the "channel_id" field.
func ChannelIDContains(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChannelID), v))
	})
}

// ChannelIDHasPrefix applies the HasPrefix predicate on the "channel_id" field.
func ChannelIDHasPrefix(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChannelID), v))
	})
}

// ChannelIDHasSuffix applies the HasSuffix predicate on the "channel_id" field.
func ChannelIDHasSuffix(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChannelID), v))
	})
}

// ChannelIDEqualFold applies the EqualFold predicate on the "channel_id" field.
func ChannelIDEqualFold(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChannelID), v))
	})
}

// ChannelIDContainsFold applies the ContainsFold predicate on the "channel_id" field.
func ChannelIDContainsFold(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChannelID), v))
	})
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorID), v))
	})
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthorID), v))
	})
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...string) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAuthorID), v...))
	})
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...string) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAuthorID), v...))
	})
}

// AuthorIDGT applies the GT predicate on the "author_id" field.
func AuthorIDGT(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthorID), v))
	})
}

// AuthorIDGTE applies the GTE predicate on the "author_id" field.
func AuthorIDGTE(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthorID), v))
	})
}

// AuthorIDLT applies the LT predicate on the "author_id" field.
func AuthorIDLT(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthorID), v))
	})
}

// AuthorIDLTE applies the LTE predicate on the "author_id" field.
func AuthorIDLTE(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthorID), v))
	})
}

// AuthorIDContains applies the Contains predicate on the "author_id" field.
func AuthorIDContains(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthorID), v))
	})
}

// AuthorIDHasPrefix applies the HasPrefix predicate on the "author_id" field.
func AuthorIDHasPrefix(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthorID), v))
	})
}

// AuthorIDHasSuffix applies the HasSuffix predicate on the "author_id" field.
func AuthorIDHasSuffix(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthorID), v))
	})
}

// AuthorIDEqualFold applies the EqualFold predicate on the "author_id" field.
func AuthorIDEqualFold(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthorID), v))
	})
}

// AuthorIDContainsFold applies the ContainsFold predicate on the "author_id" field.
func AuthorIDContainsFold(v string) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthorID), v))
	})
}

// GoToWorkAtEQ applies the EQ predicate on the "go_to_work_at" field.
func GoToWorkAtEQ(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoToWorkAt), v))
	})
}

// GoToWorkAtNEQ applies the NEQ predicate on the "go_to_work_at" field.
func GoToWorkAtNEQ(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoToWorkAt), v))
	})
}

// GoToWorkAtIn applies the In predicate on the "go_to_work_at" field.
func GoToWorkAtIn(vs ...time.Time) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGoToWorkAt), v...))
	})
}

// GoToWorkAtNotIn applies the NotIn predicate on the "go_to_work_at" field.
func GoToWorkAtNotIn(vs ...time.Time) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGoToWorkAt), v...))
	})
}

// GoToWorkAtGT applies the GT predicate on the "go_to_work_at" field.
func GoToWorkAtGT(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoToWorkAt), v))
	})
}

// GoToWorkAtGTE applies the GTE predicate on the "go_to_work_at" field.
func GoToWorkAtGTE(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoToWorkAt), v))
	})
}

// GoToWorkAtLT applies the LT predicate on the "go_to_work_at" field.
func GoToWorkAtLT(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoToWorkAt), v))
	})
}

// GoToWorkAtLTE applies the LTE predicate on the "go_to_work_at" field.
func GoToWorkAtLTE(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoToWorkAt), v))
	})
}

// GetOffAtEQ applies the EQ predicate on the "get_off_at" field.
func GetOffAtEQ(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGetOffAt), v))
	})
}

// GetOffAtNEQ applies the NEQ predicate on the "get_off_at" field.
func GetOffAtNEQ(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGetOffAt), v))
	})
}

// GetOffAtIn applies the In predicate on the "get_off_at" field.
func GetOffAtIn(vs ...time.Time) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGetOffAt), v...))
	})
}

// GetOffAtNotIn applies the NotIn predicate on the "get_off_at" field.
func GetOffAtNotIn(vs ...time.Time) predicate.Commute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGetOffAt), v...))
	})
}

// GetOffAtGT applies the GT predicate on the "get_off_at" field.
func GetOffAtGT(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGetOffAt), v))
	})
}

// GetOffAtGTE applies the GTE predicate on the "get_off_at" field.
func GetOffAtGTE(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGetOffAt), v))
	})
}

// GetOffAtLT applies the LT predicate on the "get_off_at" field.
func GetOffAtLT(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGetOffAt), v))
	})
}

// GetOffAtLTE applies the LTE predicate on the "get_off_at" field.
func GetOffAtLTE(v time.Time) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGetOffAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Commute) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Commute) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Commute) predicate.Commute {
	return predicate.Commute(func(s *sql.Selector) {
		p(s.Not())
	})
}
