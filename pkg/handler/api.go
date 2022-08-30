package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	order "http_server"
	"net/http"
	"strconv"
)

// Receiving order data from the channel and transfer it to the service
func (h *Handler) createOrder(order order.Order) {
	id, err := h.services.Order.CreateOrder(order)
	if err != nil {
		logrus.Error(http.StatusInternalServerError)
	}

	logrus.Println(http.StatusOK, map[string]int{
		"id": id})
}

//Validation of message format
func validation(msg []byte) (order.Order, error) {
	var order order.Order
	if err := json.Unmarshal(msg, &order); err != nil {
		fmt.Println(err)
		return order, fmt.Errorf("unappropriate format: %w", err)
	}

	fmt.Println("from valid")
	return order, nil
}

// API for getting Order by Id (will be used in web app)
func (h *Handler) getOrderByID(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	order, err := h.services.Order.GetOrderById(orderId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

//var tpl *template.Template
//
//func templates() error {
//	tpl, err := template.ParseGlob("/home/valeriya/Документы/GitHub/http-server/templates/*.templates")
//	if err != nil {
//		return fmt.Errorf("error while parse templates template: %w", err)
//	}
//
//	tpl.ExecuteTemplate()
//
//}

func (h *Handler) searchHandler(c *gin.Context) {
	fmt.Println("***searchHandler is running***")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
