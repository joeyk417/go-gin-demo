package services

import (
	"errors"
	"fmt"
	"go-gin-demo/internal/responses"
	"go-gin-demo/internal/utils"
	"strings"

	log "github.com/sirupsen/logrus"
)

var mockPositions = utils.ReadMockPositionData() // mockPositions
var mockPrices = utils.ReadMockPriceData()       // mockPrices

func GetUserEquityPositions(_ string) ([]responses.EquityPosition, error) {
	/* Transforms, conversions and calculations
	 *  - transform mockdata/mockpositions to mockdata/mockprices
	 *  - calculate all four profitOrLoss values
	 *  - return response/positions
	 */
	var stakePositions responses.EquityPositions
	for _, pos := range mockPositions.Positions {
		// pos.Security: "APT.ASX"
		for _, pri := range mockPrices.Prices {
			security := strings.Split(pos.Security, ".")

			if len(security) != 2 {
				log.Error("invalid security string: ", pos.Security)
				return nil, errors.New("invalid data")
			}

			if pos.AvailableUnits <= 0 {
				log.Error("invalid AvailableUnits: ", pos.AvailableUnits)
				return nil, errors.New("invalid data")
			}

			if pos.PortfolioUnits <= 0 {
				log.Error("invalid PortfolioUnits: ", pos.PortfolioUnits)
				return nil, errors.New("invalid data")
			}

			if security[0] == pri.Symbol {
				// create a EquityPosition
				var sp responses.EquityPosition
				sp.Symbol = pri.Symbol
				sp.Name = pos.SecurityDescription
				sp.OpenQty = fmt.Sprintf("%d", pos.AvailableUnits)
				sp.AvailableForTradingQty = fmt.Sprintf("%d", pos.AvailableUnits)
				sp.AveragePrice = fmt.Sprintf("%f", pos.AveragePrice)
				sp.MarketValue = fmt.Sprintf("%f", pri.LastTrade)
				sp.MarketPrice = fmt.Sprintf("%f", pri.LastTrade)
				sp.PriorClose = fmt.Sprintf("%f", pri.PriorClose)
				dayPL := getDayProfitOrLoss(pri.LastTrade, pri.PriorClose, pos.PortfolioUnits)
				sp.DayProfitOrLoss = fmt.Sprintf("%f", dayPL)
				dayPLPercent := getDayProfitOrLossPercent(dayPL, pri.PriorClose, pos.PortfolioUnits)
				sp.DayProfitOrLossPercent = fmt.Sprintf("%f", dayPLPercent)
				totalPL := getTotalProfitOrLoss(pri.LastTrade, pos.AveragePrice, pos.AvailableUnits)
				sp.TotalProfitOrLoss = fmt.Sprintf("%f", totalPL)
				totalPLPercent := getTotalProfitOrLossPercent(totalPL, pos.AveragePrice, pos.AvailableUnits)
				sp.TotalProfitOrLossPercent = fmt.Sprintf("%f", totalPLPercent)

				// append to equityPositions
				stakePositions.EquityPositions = append(stakePositions.EquityPositions, sp)
			}
		}
	}

	return stakePositions.EquityPositions, nil
}

func getDayProfitOrLoss(lastTrade, priorClose float64, unit int) float64 {
	return (lastTrade - priorClose) * float64(unit)
}

func getDayProfitOrLossPercent(dayProfitOrLoss, priorClose float64, unit int) float64 {
	return dayProfitOrLoss / (float64(unit) * priorClose) * 100
}

func getTotalProfitOrLoss(lastTrade, averagePrice float64, unit int) float64 {
	return (lastTrade - averagePrice) * float64(unit)
}

func getTotalProfitOrLossPercent(totalProfitOrLoss, averagePrice float64, unit int) float64 {
	return totalProfitOrLoss / (float64(unit) * averagePrice) * 100
}
