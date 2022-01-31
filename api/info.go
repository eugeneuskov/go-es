package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) info(c *gin.Context) {
	res, err := h.services.Info.Info()
	defer res.Body.Close()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if res.IsError() {
		newErrorResponse(c, http.StatusInternalServerError, res.String())
		return
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error parsing the response body: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, r)
}
