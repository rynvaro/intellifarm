package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Breeding holds the schema definition for the Breeding entity.
type Breeding struct {
	ent.Schema
}

// Fields of the Breeding.
func (Breeding) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("earNumber"),
		field.Int("times").Comment("胎次"),
		field.Int64("breedingAt"),
		field.Int("breedingTypeId"),
		field.String("breedingTypeName"),
		field.Int64("semenFrozenTypeId"),
		field.String("semenFrozenTypeName"),
		field.String("bullId"),
		field.String("shedName"),
		field.Int("count").Comment("使用数量"),
		field.String("userName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Breeding.
func (Breeding) Edges() []ent.Edge {
	return nil
}
