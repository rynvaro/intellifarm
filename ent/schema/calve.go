package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Calve holds the schema definition for the Calve entity.
type Calve struct {
	ent.Schema
}

// Fields of the Calve.
func (Calve) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("earNumber"),
		field.Int("times").Comment("胎次"),
		field.Int64("pregnantAt").Comment("怀孕日期"),
		field.Int64("calveAt"),
		field.Int("calveTypeId"),
		field.String("calveTypeName"),
		field.Int("calveCate").Comment("产犊类别： 1 正常 2 早产"),
		field.Int("calveCountId"),
		field.String("calveCountName"),
		field.Int("complexity"),
		field.String("userName"),
		field.Int("babyStatus").Comment("胎儿状况： 1，正常 2 死胎"),
		field.Int("babyGender").Comment("1 公 2 母"),
		field.Int("babyEarNumber"),
		field.Int("babyBreedId"),
		field.String("babyBreedName"),
		field.Int("babyHairColorId"),
		field.String("babyHairColorName"),
		field.Float32("babyWeight"),
		field.Int("babyShedId"),
		field.String("babyShedName"),
		field.Int64("tenantId").Comment("租户ID"),
		field.String("tenantName").Comment("租户组织名称"),
		field.String("remarks"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Calve.
func (Calve) Edges() []ent.Edge {
	return nil
}
