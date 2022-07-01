// ent generate --feature sql/upsert ./ent/schema
// The above feature flag has been added to ent/generate.go
// and you can now use go generate ./...
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("screen_name").Unique(),
		field.Int64("twitter_user_id").Unique(),
		field.String("description").Optional(),
		field.String("token"),
		field.String("token_secret"),
		field.String("twitter_profile_image_url").Optional(),
		field.Int32("min").Optional(),
		field.Int32("max").Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Stripe.Type),
	}
}
