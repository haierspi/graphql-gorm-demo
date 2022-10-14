package types

import "github.com/graphql-go/graphql"

type goods struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Tag         []tag  `json:"category"`
}

var GoodsType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Goods",
	Description: "商品",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "商品ID",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "商品名称",
		},
		"description": &graphql.Field{
			Type:        graphql.String,
			Description: "商品描述",
		},
		"category": &graphql.Field{
			Type:        graphql.String,
			Description: "商品类型",
		},
		"tags": &graphql.Field{
			Type:        graphql.NewList(graphql.String),
			Description: "商品标签",
		},
	},
})

func GoodsResolve(p graphql.ResolveParams) (interface{}, error) {

	return []goods{goods{ID: 1, Name: "name", Description: "description", Tag: []tag{tag{ID: 1, Name: "tagName"}}}}, nil
}
