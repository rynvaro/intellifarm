package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Farm holds the schema definition for the Farm entity.
type Farm struct {
	ent.Schema
}

// Fields of the Farm.
func (Farm) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("牛场名称").StructTag(`form:"name" required:"true"`),
		field.String("code").NotEmpty().Comment("牛场编号").Unique().StructTag(`form:"code" required:"true"`),
		field.String("constructionDate").Comment("建场日期").StructTag(`form:"constructionDate"`),
		field.Int64("feedingScale").Comment("饲养规模").StructTag(`form:"feedingScale"`),
		field.String("contactUser").Comment("联系人").StructTag(`form:"contactUser"`),
		field.String("contactPhone").Comment("联系电话").StructTag(`form:"contactPhone"`),
		field.String("contactAddress").Comment("联系地址").StructTag(`form:"contactAddress"`),
		field.Int64("square").Comment("建筑面积（m²）").StructTag(`form:"square"`),
		field.Int64("shedCount").Comment("畜舍输了（栋）").StructTag(`form:"shedCount"`),
		field.Int64("categoryId").Comment("牛场类型ID").StructTag(`form:"categoryId" required:"true"`),
		field.String("categoryName").Comment("牛场类型").StructTag(`form:"categoryName" required:"true"`),
		field.Int64("varietyId").Comment("生产品种ID").StructTag(`form:"variety" required:"true"`),
		field.String("varietyName").Comment("生产品种").StructTag(`form:"varietyName" required:"true"`),
		field.String("districtCode").Comment("地区编码").StructTag(`form:"districtCode" required:"true"`),
		field.String("districtName").Comment("地区").StructTag(`form:"districtName" required:"true"`),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks").Comment("备注信息").StructTag(`form:"remarks"`),
		field.Int64("createdAt").Comment("插入时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Default(0).Comment("是否已删除"),
	}
}

// Edges of the Farm.
func (Farm) Edges() []ent.Edge {
	return nil
}
