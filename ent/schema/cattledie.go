package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CattleDie holds the schema definition for the CattleDie entity.
type CattleDie struct {
	ent.Schema
}

// Fields of the CattleDie.
func (CattleDie) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cattleId").Comment("牛只ID"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称"),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称"),
		field.String("earNumber"),
		field.Int64("date").Comment("日期"),
		field.String("userName"),
		field.Int("reasonId"),
		field.String("reasonName"),
		field.Int("insured").Comment("是否参保"),
		field.String("insuredCode").Comment("参保单号"),
		field.String("insuredCompany").Comment("投保公司"),
		field.Int64("weight"),
		field.String("handleMethod"),
		field.Int("declared").Comment("是否申报"),
		field.String("dUserName").Comment("诊断兽医"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the CattleDie.
func (CattleDie) Edges() []ent.Edge {
	return nil
}
