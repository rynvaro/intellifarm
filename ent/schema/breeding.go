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
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.String("name").Optional(),
		field.String("earNumber"),
		field.Int("times").Comment("胎次"),
		field.Int64("breedingAt"),
		field.Int("breedingTypeId"),
		field.String("breedingTypeName"),
		field.Int64("semenFrozenTypeId"),
		field.String("semenFrozenTypeName"),
		field.String("bullId"),
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
