package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FeedInfo holds the schema definition for the FeedInfo entity.
type FeedInfo struct {
	ent.Schema
}

// Fields of the FeedInfo.
func (FeedInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("code").Comment("编码"),
		field.Float32("moisture").Comment("水分比重"),
		field.Float32("dryMatter").Comment("干物质比重"),
		field.Float32("ndf").Comment("中性洗涤纤维"),
		field.Float32("adf").Comment("酸性洗涤纤维"),
		field.Float32("endf").Comment("有效中性洗涤纤维"),
		field.Float32("lacticAcid").Comment("乳酸"),
		field.Float32("wsc").Comment("糖 (WSC）"),
		field.Float32("starch").Comment("淀粉（B1）"),
		field.Float32("solubleFiber").Comment("可溶性纤维（B2）"),
		field.Float32("totalProtein").Comment("总蛋白"),
		field.Float32("solubleProtein").Comment("可溶性蛋白"),
		field.Float32("rdp").Comment("瘤胃降解蛋白 3倍水平"),
		field.Float32("me").Comment("代谢能"),
		field.Float32("nel").Comment("泌乳净能 3倍NRC"),
		field.Float32("crudeFat").Comment("粗脂肪"),
		field.Float32("totalFttyAcid").Comment("总脂肪酸"),
		field.Float32("ash").Comment("灰分"),
		field.Float32("ca").Comment("钙"),
		field.Float32("p").Comment("磷"),
		field.Float32("mg").Comment("镁"),
		field.Float32("k").Comment("钾"),
		field.Float32("mn").Comment("锰 - 总"),
		field.Float32("cu").Comment("铜 - 总"),
		field.Float32("fe").Comment("铁 - 总"),
		field.Float32("zn").Comment("锌 - 总"),
		field.Float32("methionine").Comment("蛋氨酸"),
		field.Float32("lysine").Comment("赖氨酸"),
		field.Float32("vitaminA").Comment("维生素 A"),
		field.Float32("vitaminD3").Comment("维生素 D3"),
		field.Float32("vitaminE").Comment("维生素 E"),
		field.Float32("choline").Comment("胆碱"),
		field.Float32("biotin").Comment("生物素"),
		field.String("description").Comment("描述").Optional(),
		field.Int64("tenantId").Comment("租户ID").Optional(),
		field.String("tenantName").Comment("租户组织名称").Optional(),
		field.Int64("farmId").Comment("牛场ID").Optional(),
		field.String("farmName").Comment("牛场名称").Optional(),
		field.String("remarks").Comment("备注").Optional(),
		field.Int64("createdAt").Comment("创建时间").Optional().Default(time.Now().Unix()),
		field.Int64("updatedAt").Comment("更新时间").Optional().Default(time.Now().Unix()),
		field.Int("deleted").Comment("是否已删除").Optional().Default(0),
	}
}

// Edges of the FeedInfo.
func (FeedInfo) Edges() []ent.Edge {
	return nil
}
