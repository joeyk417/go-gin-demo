package middlewares

import (
	"fmt"
	"go-gin-demo/internal/handlers"
	"go-gin-demo/internal/responses"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name                string
	authorizationHeader string
	authorizationType   string
	token               string
	statusCode          int
}

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

func TestAuthMiddleware(t *testing.T) {

	testCases := []testCase{
		{
			name:                "Ok",
			authorizationHeader: "Authorization",
			authorizationType:   "Bearer",
			token:               testToken,
			statusCode:          200,
		},
		{
			name:                "NoAuthorization",
			authorizationHeader: "",
			authorizationType:   "",
			token:               "",
			statusCode:          401,
		},
		{
			name:                "InvalidAuthorizationFormat",
			authorizationHeader: "Authorization",
			authorizationType:   "",
			token:               testToken,
			statusCode:          401,
		},
		{
			name:                "InvalidAuthorizationType",
			authorizationHeader: "Authorization",
			authorizationType:   "unsupport",
			token:               testToken,
			statusCode:          401,
		},
		{
			name:                "InvalidToken",
			authorizationHeader: "Authorization",
			authorizationType:   "Bearer",
			token:               "xxx",
			statusCode:          401,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			handlers.GetUserEquityPositionsService = func(testUserId string) ([]responses.EquityPosition, error) {
				return []responses.EquityPosition{mockResponse}, nil
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			authRoutes := r.Group("/").Use(AuthMiddleware())
			authRoutes.GET("/api/equityPositions", handlers.GetEquityPositionsHandler)

			req, _ := http.NewRequest("GET", "/api/equityPositions", nil)
			w := httptest.NewRecorder()
			authorizationHeaderString := fmt.Sprintf("%s %s", tc.authorizationType, tc.token)
			req.Header.Set(tc.authorizationHeader, authorizationHeaderString)
			r.ServeHTTP(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}

}
