package handler

import (
	"github.com/chorshik/testToimi"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

type getAllResponse struct {
	Title       string
	Photo       string
	Price       float64
}

type getById struct {
	Title       string
	Photo       string
	Price       float64
}

func (h *Handler) createAd(c *gin.Context) {
	var input testToimi.Ad
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"code": http.StatusOK,
	})
}

func (h *Handler) getAllAds(c *gin.Context) {
	list, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	respList := make([]getAllResponse, 0)
	println(len(list))
	for _, v := range list {
		respList = append(respList, getAllResponse{
			Title: v.Title,
			Photo: v.Photo[0],
			Price: v.Price,
		})

	}

	sort.Slice(respList, func(i, j int) (less bool) {
		return respList[i].Price < respList[j].Price
	})

	//sort.Slice(respList, func(i, j int) (less bool) {
	//	return respList[i].Price > respList[j].Price
	//})

	c.JSON(http.StatusOK, respList)
}



func (h *Handler) getAdById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	list, err := h.services.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp:= getById{
		Title: list.Title,
		Photo: list.Photo[0],
		Price: list.Price,
	}

	_, ok := c.GetQuery("fields")
	if ok == true{
		c.JSON(http.StatusOK, map[string]interface{}{
			"Title": list.Title,
			"Photo": list.Photo[0],
			"Price": list.Price,
			"Description": list.Description,
			"Other_photo": [2]string{list.Photo[1], list.Photo[2]},
		})
	}else {c.JSON(http.StatusOK, resp)}


}