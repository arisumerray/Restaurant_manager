package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var r struct {
		SpecialRequests string  `json:"special_requests"`
		DishIds         []int64 `json:"dish_ids"`
		Quantities      []int64 `json:"quantities"`
	}
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(r.DishIds) != len(r.Quantities) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ids and quantities have different sizes"})
		return
	}
	req := CreateOrderReq{
		UserId:          c.MustGet("user_id").(int64),
		SpecialRequests: r.SpecialRequests,
		DishIds:         r.DishIds,
		Quantities:      r.Quantities,
	}
	res, err := h.Service.Order(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go h.SetStatus(&UpdateStatusReq{
		Id:     res.Id,
		Status: "done",
	})
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	o, err := h.Service.GetOrder(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, o)
}
