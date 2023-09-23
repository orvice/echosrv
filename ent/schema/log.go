package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AccessLog holds the schema definition for the AccessLog entity.
type AccessLog struct {
	ent.Schema
}

// Fields of the AccessLog.
func (AccessLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_unix"),
		field.String("path"),
		field.String("method"),
		field.String("ip"),
		field.String("ua").
			Unique(),
	}
}

// Edges of the AccessLog.
func (AccessLog) Edges() []ent.Edge {
	return nil
}
