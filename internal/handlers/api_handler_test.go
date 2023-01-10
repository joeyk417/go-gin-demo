package handlers

import (
	"encoding/json"
	"errors"
	"go-gin-demo/internal/middlewares"
	"go-gin-demo/internal/responses"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var mockResponse = responses.EquityPosition{
	Symbol:                   "TEST",
	Name:                     "This is a test",
	OpenQty:                  "5",
	AvailableForTradingQty:   "5",
	AveragePrice:             "107.000000",
	MarketValue:              "104.630000",
	MarketPrice:              "104.630000",
	PriorClose:               "100.770000",
	DayProfitOrLoss:          "77.200000",
	DayProfitOrLossPercent:   "3.830505",
	TotalProfitOrLoss:        "-47.400000",
	TotalProfitOrLossPercent: "-2.214953",
}

func TestOK_GetEquityPositionsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetUserEquityPositionsService = func(testUserId string) ([]responses.EquityPosition, error) {
		return []responses.EquityPosition{mockResponse}, nil
	}

	GetEquityPositionsHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestError_GetEquityPositionsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetUserEquityPositionsService = func(testUserId string) ([]responses.EquityPosition, error) {
		return []responses.EquityPosition{}, errors.New("error")
	}

	GetEquityPositionsHandler(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNoDataError_GetEquityPositionsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetUserEquityPositionsService = func(testUserId string) ([]responses.EquityPosition, error) {
		return nil, nil
	}

	GetEquityPositionsHandler(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestData_GetEquityPositionsHandler(t *testing.T) {
	GetUserEquityPositionsService = func(testUserId string) ([]responses.EquityPosition, error) {
		return []responses.EquityPosition{mockResponse}, nil
	}

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	authRoutes := r.Group("/").Use(middlewares.AuthMiddleware())
	authRoutes.GET("/api/equityPositions", GetEquityPositionsHandler)

	req, _ := http.NewRequest("GET", "/api/equityPositions", nil)
	w := httptest.NewRecorder()
	authorizationHeader := "Bearer fJCoxhq8uR9GiUIgaIGfMgw7zCqxwDhQ"
	req.Header.Set("Authorization", authorizationHeader)

	r.ServeHTTP(w, req)

	data := responses.EquityPositions{}
	_ = json.Unmarshal(w.Body.Bytes(), &data)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse.Name, data.EquityPositions[0].Name)
}
