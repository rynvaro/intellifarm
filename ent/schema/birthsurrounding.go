package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// BirthSurrounding holds the schema definition for the BirthSurrounding entity.
type BirthSurrounding struct {
	ent.Schema
}

// Fields of the BirthSurrounding.
func (BirthSurrounding) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.Int64("farmId").Comment("牛场ID"),
		field.String("farmName").Comment("牛场名称").NotEmpty(),
		field.Int64("recordTime").Comment("记录时间戳"),
		field.Int64("userId").Comment("记录人ID"),
		field.String("userName").Comment("记录人"),
		field.Int64("temperature").Comment("温度（℃）"),
		field.Int64("humidity").Comment("相对湿度（%）"),
		field.Int64("breathRateId").Comment("呼吸速率ID"),
		field.String("breathRateName").Comment("呼吸速率"),
		field.Int64("windSpeed").Comment("风速"),
		field.Int("windDirectionId").Comment("风向ID"),
		field.String("windDirection").Comment("风向"),
		field.Int64("locationChanges").Comment("位置改变（次）"),
		field.Int("hairStateId").Comment("被毛状况ID"),
		field.String("hairStateName").Comment("被毛状况"),
		field.Int("soilDepth").Comment("泥深（cm）"),
		field.Int("sunExposure").Comment("日照时间（h）"),
		field.Int("walkDistance").Comment("行走距离（m）"),
		field.Int("rained").Comment("是否淋雨，1: 是，0: 否").Default(0),
		field.Float32("thIndex").Comment("温湿指数(temperature and humidity index)"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the BirthSurrounding.
func (BirthSurrounding) Edges() []ent.Edge {
	return nil
}
