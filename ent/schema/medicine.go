package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Medicine holds the schema definition for the Medicine entity.
type Medicine struct {
	ent.Schema
}

// Fields of the Medicine.
func (Medicine) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.Int64("epid").Comment("发病ID"),
		field.String("earNumber"),
		field.String("medicineName"),
		field.Int64("dose").Comment("剂量"),
		field.String("unit").Comment("剂量单位"),
		field.Int64("dateStart").Comment("用药开始时间"),
		field.Int64("dateEnd").Comment("休药时间"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间时间"),
		field.Int("deleted").Comment("是否已删除"),
		field.String("remarks"),
	}
}

// Edges of the Medicine.
func (Medicine) Edges() []ent.Edge {
	return nil
}
