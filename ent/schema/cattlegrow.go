package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CattleGrow holds the schema definition for the CattleGrow entity.
type CattleGrow struct {
	ent.Schema
}

// Fields of the CattleGrow.
func (CattleGrow) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("earNumber").Comment("牛耳号").NotEmpty(),
		field.String("stage").Comment("生长阶段").NotEmpty(),
		field.Int64("dateStart").Comment("开始日期"),
		field.Int64("dateEnd").Comment("结束日期"),
		field.Float32("weightStart").Comment("初始体重(kg)"),
		field.Float32("weightEnd").Comment("最终体重(kg)"),
		field.Float32("dailyWeight").Comment("日增重(kg)"),
		field.Float32("feedWeight").Comment("总耗料量(kg)"),
		field.Float32("dailyFeedWeight").Comment("平均日耗料量(kg)"),
		field.Float32("conversionRate").Comment("饲料转化效率"),
		field.String("userName").Comment("记录人"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleGrow.
func (CattleGrow) Edges() []ent.Edge {
	return nil
}
