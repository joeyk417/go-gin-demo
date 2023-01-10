package services

import (
	"go-gin-demo/internal/models"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestResult_GetUserEquityPositions(t *testing.T) {
	pos := models.Position{
		Security: "TEST.ASX", SecurityDescription: "This is a test",
		Cost:           1000.000000,
		AveragePrice:   100.000000,
		AvailableUnits: 10,
		PortfolioUnits: 10,
	}
	pri := models.Price{
		Symbol:       "TEST",
		MarketStatus: "Active",
		LastTrade:    110.000000,
		Bid:          102.000000,
		Ask:          102.000000,
		PriorClose:   110.0000000,
	}
	mockPositions = models.Positions{Positions: []models.Position{pos}}
	mockPrices = models.Prices{Prices: []models.Price{pri}}

	result, _ := GetUserEquityPositions("test")

	require.Equal(t, pos.SecurityDescription, result[0].Name)
	require.Equal(t, "0.000000", result[0].DayProfitOrLoss)
	require.Equal(t, "0.000000", result[0].DayProfitOrLossPercent)
	require.Equal(t, "100.000000", result[0].TotalProfitOrLoss)
	require.Equal(t, "10.000000", result[0].TotalProfitOrLossPercent)
}

func TestInvalidSecurityString_GetUserEquityPositions(t *testing.T) {
	pos := models.Position{
		Security: "TESTASX", SecurityDescription: "This is a test",
		Cost:           1000.000000,
		AveragePrice:   100.000000,
		AvailableUnits: 10,
		PortfolioUnits: 10,
	}
	pri := models.Price{
		Symbol:       "TEST",
		MarketStatus: "Active",
		LastTrade:    110.000000,
		Bid:          102.000000,
		Ask:          102.000000,
		PriorClose:   110.0000000,
	}
	mockPositions = models.Positions{Positions: []models.Position{pos}}
	mockPrices = models.Prices{Prices: []models.Price{pri}}

	_, err := GetUserEquityPositions("test")

	assert.Equal(t, "invalid data", err.Error())
}

func TestInvalidAvailableUnits_GetUserEquityPositions(t *testing.T) {
	pos := models.Position{
		Security: "TEST.ASX", SecurityDescription: "This is a test",
		Cost:           1000.000000,
		AveragePrice:   100.000000,
		AvailableUnits: 0,
		PortfolioUnits: 10,
	}
	pri := models.Price{
		Symbol:       "TEST",
		MarketStatus: "Active",
		LastTrade:    110.000000,
		Bid:          102.000000,
		Ask:          102.000000,
		PriorClose:   110.0000000,
	}
	mockPositions = models.Positions{Positions: []models.Position{pos}}
	mockPrices = models.Prices{Prices: []models.Price{pri}}

	_, err := GetUserEquityPositions("test")

	assert.Equal(t, "invalid data", err.Error())
}

func TestInvalidPortfolioUnits_GetUserEquityPositions(t *testing.T) {
	pos := models.Position{
		Security: "TEST.ASX", SecurityDescription: "This is a test",
		Cost:           1000.000000,
		AveragePrice:   100.000000,
		AvailableUnits: 10,
		PortfolioUnits: 0,
	}
	pri := models.Price{
		Symbol:       "TEST",
		MarketStatus: "Active",
		LastTrade:    110.000000,
		Bid:          102.000000,
		Ask:          102.000000,
		PriorClose:   110.0000000,
	}
	mockPositions = models.Positions{Positions: []models.Position{pos}}
	mockPrices = models.Prices{Prices: []models.Price{pri}}

	_, err := GetUserEquityPositions("test")

	assert.Equal(t, "invalid data", err.Error())
}
