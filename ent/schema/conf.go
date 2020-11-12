package schema

import (
	"cattleai/confs"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
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
