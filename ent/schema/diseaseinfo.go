package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// DiseaseInfo holds the schema definition for the DiseaseInfo entity.
type DiseaseInfo struct {
	ent.Schema
}

// Fields of the DiseaseInfo.
func (DiseaseInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("code").Comment("编码"),
		field.String("type").Comment("类型"),
		field.String("description").Comment("描述"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the DiseaseInfo.
func (DiseaseInfo) Edges() []ent.Edge {
	return nil
}
