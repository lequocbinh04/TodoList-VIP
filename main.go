package main

import (
	"TodoList/controller"
	"TodoList/database"
	"TodoList/graph/generated"
	"TodoList/graph/resolver"
	"TodoList/middleware"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
    err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
    database.InitDatabase()


	r := gin.Default()
	r.POST("/get-jwt", controller.GetJWT)
	r.GET("/", playgroundHandler())
    protectedRoute := r.Group("/").Use(middleware.GinContextToContextMiddleware()).Use(middleware.Authenticate())
    {
        protectedRoute.POST("/graphql", graphqlHandler())
    }
	r.Run(":8080")
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}