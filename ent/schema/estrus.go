package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Estrus holds the schema definition for the Estrus entity.
type Estrus struct {
	ent.Schema
}

// Fields of the Estrus.
func (Estrus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("earNumber"),
		field.Int("times").Comment("胎次"),
		field.String("reproductiveState"),
		field.String("shedName"),
		field.Int64("estrusAt"),
		field.Int("estrusTypeId"),
		field.String("estrusTypeName"),
		field.String("userName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Estrus.
func (Estrus) Edges() []ent.Edge {
	return nil
}
