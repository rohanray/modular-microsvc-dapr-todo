package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("task").
			NotEmpty().
			Annotations(
				entproto.Field(2),
			),
		field.Time("created_at").
			Immutable().
			Annotations(
				entproto.Field(3),
			).
			Default(time.Now),
		field.Enum("status").
			Values("NEW", "IN_PROGRESS", "COMPLETED").
			Annotations(
				entproto.Field(4),
				entproto.Enum(map[string]int32{
					"NEW":         0,
					"IN_PROGRESS": 1,
					"COMPLETED":   2,
				}),
			).
			Default("NEW"),
	}
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
