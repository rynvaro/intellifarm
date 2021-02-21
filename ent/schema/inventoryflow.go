package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// InventoryFlow holds the schema definition for the InventoryFlow entity.
type InventoryFlow struct {
	ent.Schema
}

// Fields of the InventoryFlow.
func (InventoryFlow) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("materialID").Comment("物料ID"),
		field.String("materialName").Comment("物料名称"),
		field.String("materialCode").Comment("物料编号"),
		field.String("seqNumber").Comment("单号"),
		field.Int64("date").Comment("日期"),
		field.Int("type").Comment("1: 入库，2: 出库"),
		field.Int("status").Comment("1 已入库 2 已结算 3 已出库"),
		field.Int("count").Comment("出入库的量"),
		field.String("unit").Comment("量的单位"),
		field.String("userName"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the InventoryFlow.
func (InventoryFlow) Edges() []ent.Edge {
	return nil
}
