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
		field.Int("level").Comment("1 系统用户 2 集团用户 3 集团下属用户"),
		field.Int64("farmId").Comment("牛场ID").Optional(),
		field.String("farmName").Comment("牛场名称").Optional(),
		field.Int64("positionId").Comment("职务ID").Optional(),
		field.String("positionName").Comment("职务").Optional(),
		field.String("dutyName").Comment("职责，斜线分割").Optional(),
		field.String("name").Comment("员工姓名").NotEmpty(),
		field.Int("gender").Comment("员工性别： 1：男，2: 女").Optional(),
		field.Int("age").Comment("员工年龄").Optional(),
		field.String("education").Comment("学历").Optional(),
		field.String("major").Comment("专业").Optional(),
		field.String("jobTitle").Comment("职称").Optional(),
		field.String("phone").Comment("联系电话").Optional(),
		field.String("idCard").Comment("身份证号").Optional(),
		field.String("address").Comment("联系地址").Optional(),
		field.Int("onJobState").Comment("在职状态： 1： 在职， 0: 已离职").Default(1),
		field.Int64("joinedAt").Comment("入职时间").Optional(),
		field.Int64("tenantId").Comment("租户ID").Optional(),
		field.String("tenantName").Comment("租户组织名称").Optional(),
		field.String("password").Comment("用户密码"),
		field.String("remarks").Comment("备注").Optional(),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
