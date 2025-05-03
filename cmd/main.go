package cmd

import (
	"ecomm/internal/services/auth-service/handler"
	"ecomm/pkg/config"
	"ecomm/pkg/db"
	"ecomm/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	cfg := config.Config()

	url := cfg.DBaseURL()

	conn, err := db.Open(url)
	if err != nil {
		logger.Log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer conn.Close()

	h := handler.CreateUserHandler(conn)
	g := gin.New()

	{
		l := g.Group("/login")
		l.POST("/sig")
		l.POST("/", h.) 
	}
}
