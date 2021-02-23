package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("type").Comment("类型"),
		field.String("gender").Comment("性别"),
		field.String("province").Comment("省"),
		field.String("phone").Comment("电话"),
		field.String("address").Comment("联系地址"),
		field.String("email").Comment("邮箱"),
		field.String("bank").Comment("开户行"),
		field.String("account").Comment("银行账号"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
