package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("company").Comment("机构所属公司"),
		field.String("name").Comment("租户组织名称"),
		field.String("code").Comment("租户组织代码"),
		field.Int("enabled").Comment("是否启用"),
		field.String("region").Comment("区域"),
		field.String("address").Comment("地址"),
		field.String("userName").Comment("联系人"),
		field.String("phone").Comment("联系电话"),
		field.String("remarks").Comment("备注"),
		field.Int("deleted").Comment("是否已删除").Default(0),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}
