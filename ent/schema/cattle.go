package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Cattle holds the schema definition for the Cattle entity.
type Cattle struct {
	ent.Schema
}

// Fields of the Cattle.
func (Cattle) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称").NotEmpty(),
		field.Int64("shedId").Comment("牛舍ID"),
		field.String("shedName").Comment("牛舍名称").NotEmpty(),
		field.Int64("shedTypeId").Comment("牛舍类型ID"),
		field.String("shedTypeName").Comment("牛舍类型").NotEmpty(),
		field.String("earNumber").Comment("牛耳号").NotEmpty(),
		field.String("electronicEarNumber").Comment("电子耳号"),
		field.String("pedometer").Comment("计步器"),
		field.Int("genderId").Comment("性别ID"),
		field.String("genderName").Comment("性别名称"),
		field.Int64("birthday").Comment("出生日期"),
		field.Int64("joinedAt").Comment("入群日期"),
		field.Int("cateId").Comment("用途ID"),
		field.String("cateName").Comment("用途"),
		field.Int("type").Comment("牛只类型ID"),
		field.String("typeName").Comment("牛只类型"),
		field.Int32("weight").Comment("入群体重"),
		field.String("from").Comment("来源"),
		field.String("breed").Comment("品种"),
		field.Int("joinedTypeId").Comment("入群类型ID"),
		field.String("joinedTypeName").Comment("入群类型"),
		field.String("mother").Comment("母亲号"),
		field.String("father").Comment("父亲号"),
		field.String("grandfather").Comment("外祖父号"),
		field.Int64("ownerId").Comment("拥有者ID"),
		field.String("ownerName").Comment("拥有者"),
		field.Int64("hairColorId").Comment("毛色ID"),
		field.String("hairColorName").Comment("毛色"),
		field.Int("reproductiveStateId").Comment("繁育状态ID"),
		field.String("reproductiveStateName").Comment("繁育状态"),
		field.Int("pregnantTimes").Comment("胎次"),
		field.Int64("lastCalvingAt").Comment("最后产犊日期"),
		field.Int64("breedingAt").Comment("配种日期"),
		field.Int("breedingTypeId").Comment("配种类型ID"),
		field.String("breedingTypeName").Comment("配种类型"),
		field.Int64("bullId").Comment("公牛号"),
		field.Int64("pregnancyCheckAt").Comment("妊检日期"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Cattle.
func (Cattle) Edges() []ent.Edge {
	return nil
}
