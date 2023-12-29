package mixins

import (
  "entgo.io/ent"
  "entgo.io/ent/schema/field"
  "entgo.io/ent/schema/mixin"
  "time"
)

type IDMixin struct {
  mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
  return []ent.Field{
    field.Uint64("id").
      Unique().
      Immutable(),
    field.Time("created_at").
      Immutable().
      Default(time.Now),
    field.Time("updated_at").
      Default(time.Now).
      UpdateDefault(time.Now),
  }
}
