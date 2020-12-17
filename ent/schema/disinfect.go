package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Disinfect holds the schema definition for the Disinfect entity.
type Disinfect struct {
	ent.Schema
}

// Fields of the Disinfect.
func (Disinfect) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.Int64("date").Comment("日期"),
		field.Int("typeId").Comment("类型ID"),
		field.String("typeName").Comment("类型"),
		field.Int("methodId"),
		field.String("methodName"),
		field.Int("wayId"),
		field.String("wayName"),
		field.String("drug").Comment("用药"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Disinfect.
func (Disinfect) Edges() []ent.Edge {
	return nil
}
