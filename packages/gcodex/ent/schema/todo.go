package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("").Values()
		field.String("name").MaxLen(12).Unique().Sensitive().Comment("名称").Immutable().NotEmpty().Default(""),
		field.Enum("ee").Values([]string{"11", "22", "333"}...),
		field.Float("amount").
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)", // Override MySQL.
				dialect.Postgres: "numeric",      // Override Postgres.
			}).Default(0),
		// Add a new field with CURRENT_TIMESTAMP
		// as a default value to all previous rows.
		field.Time("created_at").Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}

// Indexes 写明索引
func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一的联合索引
		index.Fields("value", "name"),
		// 非唯一的普通索引
		index.Fields("name"),
		// 唯一索引
		index.Fields("value").Unique(),
	}
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ToDO"},
	}
}
