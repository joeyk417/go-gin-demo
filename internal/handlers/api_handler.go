package handlers

import (
	"go-gin-demo/internal/responses"
	"go-gin-demo/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var testUserId = "f9e27aa9-727c-4fd8-907e-46522421089d"
var testToken = "fJCoxhq8uR9GiUIgaIGfMgw7zCqxwDhQ"

var GetUserEquityPositionsService = func(testUserId string) ([]responses.EquityPosition, error) {
	return services.GetUserEquityPositions(testUserId)
}

func GetEquityPositionsHandler(c *gin.Context) {
	stakePositions, err := GetUserEquityPositionsService(testUserId)

	// handle case where invalid/null equityPositions are returned by the service class
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if stakePositions == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"equityPositions": stakePositions,
	})
}
