package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/haierspi/graphql-gorm-demo/graphql/types"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"goods": &graphql.Field{
			Type:        graphql.NewList(types.GoodsType),
			Description: "商品列表",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
					Description:  "商品名称",
				},
				"pn": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "当前第几页",
				},
				"ps": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "每页显示多少数据",
				},
			},
			Resolve: types.GoodsResolve,
		},
	},
})
