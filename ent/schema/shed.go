package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Shed holds the schema definition for the Shed entity.
type Shed struct {
	ent.Schema
}

// Fields of the Shed.
func (Shed) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称").NotEmpty(),
		field.String("name").Comment("牛舍名称").NotEmpty(),
		field.Int64("shedCateId").Comment("牛舍类别ID"),
		field.String("shedCateName").Comment("牛舍类别").NotEmpty(),
		field.Int("shedTypeId").Comment("牛舍类型"),
		field.String("shedTypeName").Comment("牛舍类型").NotEmpty(),
		field.Int64("square").Comment("建筑面积（m²）"),
		field.Int64("length").Comment("长（m）"),
		field.Int64("width").Comment("宽（m）"),
		field.Int64("height").Comment("高（m）"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int("userId").Comment("负责人ID"),
		field.String("userName").Comment("负责人"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Shed.
func (Shed) Edges() []ent.Edge {
	return nil
}
