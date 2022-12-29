package schema

import (
	"cattleai/confs"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Conf holds the schema definition for the Conf entity.
type Conf struct {
	ent.Schema
}

// Fields of the Conf.
func (Conf) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("confs", confs.AppConfs),
	}
}

// Edges of the Conf.
func (Conf) Edges() []ent.Edge {
	return nil
}
