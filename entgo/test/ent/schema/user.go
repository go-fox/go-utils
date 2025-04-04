package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键"),
		field.String("name").Comment("姓名"),
		field.Int64("age").Comment("年龄"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
