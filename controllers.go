package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)
const (
	OCCUPIED_ERROR = "error, there is only %v occupied seats"
	SEAT_WRONG_NUMBER = "an error occurred, wrong number of seats"
)

func setSeat(c *gin.Context) {
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + ", invalid request body."})
		return
	}
	re := setId(&request)
	re = removeRepeated(&request)
	c.JSON(http.StatusOK, gin.H{
		"result" : re,
		"status":    "ok",
	})
}

func getSeat(c *gin.Context) {

	characters, err := getCharacters(request.SeatsPerRow)
	oc := getOccupied(request.Seats)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result":  fmt.Sprintf(OCCUPIED_ERROR, len(oc)),
		})
	} else {
		seatNumberQuery := c.Query("seatNumber")
		seatNumber, err := strconv.Atoi(seatNumberQuery)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"result": fmt.Sprintf(OCCUPIED_ERROR, len(oc)),
			})
		} else {


			var x Selection
			if seatNumber <= len(oc) {
				x = oc[seatNumber-1]
			}


			resultNumber, err := json.Marshal(x.Input)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("an error occurred, %v", err),
				})
			}
			ch := characters[x.Seat-1]
			if ch == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"result": SEAT_WRONG_NUMBER,
				})
			} else {
				result := string(resultNumber) + ch
				log.Println(result)
				if result != "0" {
					c.JSON(http.StatusOK, gin.H{
						"result": result,
					})
				} else {
					c.JSON(http.StatusNotFound, gin.H{
						"result": fmt.Sprintf(OCCUPIED_ERROR, len(oc)),
					})
				}
			}
		}
	}

}
