package handler

import (
	"Ticketing/internal/service"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

// BlogHandler handles HTTP requests related to Blogs.
type SaldoHandler struct {
	saldoService service.Ceksaldo
}

func ceksaldo(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	balance, ok := balances[userID]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, balance)
}

func topUpBalance(c echo.Context) error {
	var topUpRequest TopUpRequest
	if err := c.Bind(&topUpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	userID := topUpRequest.UserID
	amount := topUpRequest.Amount

	// Validasi userID
	if _, ok := balances[userID]; !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	// Validasi amount
	if amount <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid top-up amount"})
	}

	// Melakukan top-up
	balances[userID].Amount += amount

	return c.JSON(http.StatusOK, balances[userID])
}
