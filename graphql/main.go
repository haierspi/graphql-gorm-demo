package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/labstack/gommon/log"
	"net/http"
)

type graphBody struct {
	Query     string
	Variables map[string]interface{}
}

var schema graphql.Schema

func init() {
	var err error
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatalf("graphql schema initial failed: %v \n", err)
	}
}

func Handler(c *gin.Context) {

	var body graphBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Warnf("graqhql bindJson failed: %v \n", err)
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: body.Query,
	})
	if len(result.Errors) > 0 {
		log.Warnf("wrong result: %v \n", result.Errors)
		c.JSON(http.StatusBadRequest, result.Errors)
		return
	}

	c.JSON(http.StatusOK, result.Data)
}
