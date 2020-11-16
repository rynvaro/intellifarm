package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Abortion holds the schema definition for the Abortion entity.
type Abortion struct {
	ent.Schema
}

// Fields of the Abortion.
func (Abortion) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("earNumber"),
		field.Int("times").Comment("胎次"),
		field.String("reproductiveState"),
		field.String("shedName"),
		field.Int64("pregnantAt"),
		field.Int64("abortionAt"),
		field.Int("abortionTypeId"),
		field.String("abortionTypeName"),
		field.String("userName"),
		field.String("remarks").Comment("备注"),
		field.Int64("createdAt").Comment("创建时间"),
		field.Int64("updatedAt").Comment("更新时间"),
		field.Int("deleted").Comment("是否已删除"),
	}
}

// Edges of the Abortion.
func (Abortion) Edges() []ent.Edge {
	return nil
}
