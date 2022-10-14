package graphql

import "github.com/graphql-go/graphql"

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createGoods": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "创建商品",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "商品名称",
				},
				"Description": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
					Description:  "商品描述",
				},
				"category": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
					Description:  "商品种类",
				},
				"tags": &graphql.ArgumentConfig{
					Type:         graphql.NewList(graphql.NewNonNull(graphql.String)),
					DefaultValue: []string{},
					Description:  "商品标签",
				},
			},
		},
	},
})
