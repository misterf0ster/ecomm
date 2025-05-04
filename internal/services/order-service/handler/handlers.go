package handler

import (
	"context"
	m "ecomm/internal/services/order-service/model"
	"ecomm/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetOrderList(c *gin.Context){
	sql:="SELECT id"

	list, err := h.DB.Pslq.Query(context.Background(), sql)
	if err != nil{
		logger.LogError("Database connect folse", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Database connect fals",
			"Status": "Error",
		})
		return
	}
	defer list.Close()

	var orders []m.Order
	
	for list.Next() {
		var u m.Order
		if err := list.Scan(&u.ID); err == nil {
			orders = append(orders, u)
		}
		return
	}

	logger.LogInfo("Users get successfully")
	c.JSON(http.StatusOK, orders)
}

func GetOrderID(c *gin.Context) {
	id := c.Param("id")

	if _,err := 
}
