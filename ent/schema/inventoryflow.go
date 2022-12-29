package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// InventoryFlow holds the schema definition for the InventoryFlow entity.
type InventoryFlow struct {
	ent.Schema
}

// Fields of the InventoryFlow.
func (InventoryFlow) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("sysMaterialId").Comment("基础资料物料Id"),
		field.Int64("materialId").Comment("物料Id"),
		field.String("materialName").Comment("物料名称"),
		field.String("materialCode").Comment("物料编号"),
		field.String("seqNumber").Comment("单号"),
		field.Int64("date").Comment("日期"),
		field.Int("type").Comment("1: 入库，2: 出库"),
		field.String("status").Comment("状态"),
		field.Int("count").Comment("出入库的量"),
		field.String("unit").Comment("量的单位"),
		field.Int64("before").Comment("出入库前的库存"),
		field.Int64("after").Comment("出入库后的库存"),
		field.String("userName"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.String("remarks").Comment("备注"),
		field.Bool("isChecked").Comment("是否已上传检测报告"),
		field.String("reportFileAddress").Comment("检测报告文件地址"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the InventoryFlow.
func (InventoryFlow) Edges() []ent.Edge {
	return nil
}
