package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Stripe holds the schema definition for the Stripe entity.
type Stripe struct {
	ent.Schema
}

// Fields of the Stripe.
func (Stripe) Fields() []ent.Field {
	return []ent.Field{
		field.String("api_key").Unique(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Stripe.
func (Stripe) Edges() []ent.Edge {
	return nil
}
