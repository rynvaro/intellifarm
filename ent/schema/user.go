package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称").NotEmpty(),
		field.Int64("positionId").Comment("职务ID"),
		field.String("positionName").Comment("职务").NotEmpty(),
		field.String("dutyName").Comment("职责，斜线分割").NotEmpty(),
		field.String("name").Comment("员工姓名").NotEmpty(),
		field.Int("gender").Comment("员工性别： 1：男，2: 女"),
		field.Int("age").Comment("员工年龄"),
		field.String("education").Comment("学历"),
		field.String("major").Comment("专业"),
		field.String("jobTitle").Comment("职称"),
		field.String("phone").Comment("联系电话"),
		field.String("idCard").Comment("身份证号"),
		field.String("address").Comment("联系地址"),
		field.Int("onJobState").Comment("在职状态： 1： 在职， 0: 已离职").Default(1),
		field.Int64("joinedAt").Comment("入职时间"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("password").Comment("用户密码"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
