package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserTag holds the schema definition for the UserTag entity.
type UserTag struct {
	ent.Schema
}

// Fields of the UserTag.
func (UserTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			StructTag(`json:"uid,omitempty"`).
			StorageKey("uid").Annotations(entsql.Annotation{
			Size: 64,
		}).Validate(MaxRuneCount(64)).Default("").
			Comment("用户ID").
			NotEmpty(),
		field.String("tag").
			StorageKey("tag").Annotations(entsql.Annotation{
			Size: 64,
		}).Validate(MaxRuneCount(64)).Annotations(entsql.Annotation{
			Size: 10,
		}).Validate(MinRuneCount(10)).
			Default("dd").
			Comment("tag").Sensitive().
			Optional().Nillable(),
	}
}

// Indexes of the UserTag.
func (UserTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").StorageKey("uidx_user_id").Annotations(entsql.Desc()).Unique(),

		index.Fields("tag").StorageKey("uidxname").Annotations(entsql.Desc()).Unique(),
	}
}

// Edges of the UserTag.
func (UserTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.From("owner", User.Type).Ref("tags"),
	}
}

// Annotations of the UserTag.
func (UserTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "标签表"},
	}
}
