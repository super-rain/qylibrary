// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 8},
		{Name: "value", Type: field.TypeString, Default: ""},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "todo_value_name",
				Unique:  false,
				Columns: []*schema.Column{TodosColumns[2], TodosColumns[1]},
			},
			{
				Name:    "todo_name",
				Unique:  false,
				Columns: []*schema.Column{TodosColumns[1]},
			},
			{
				Name:    "todo_value",
				Unique:  true,
				Columns: []*schema.Column{TodosColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
	}
)

func init() {
}
