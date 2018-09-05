/*
1. The program should allow users to define the dimension of the
matrix that will represent seating arrangement.
Example: Input: Rows: 10, SeatsPerRow: 6

2. The program should allow users to select or input a seating assignment
randomly in the matrix or in sequence for a passenger:
Example: Input: Row: 3, Seat: 2, Occupied: true

3. The program should allow users to search for a seat number
by inputing the sequence of seat assigned.
Example:

Input: 3 (means 3rd seat assigned)
Output: 6F (depending on what part/coordinate in the matrix{seating arrangement}
the 3rd passenger was assigned.)
 */

package main

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

// Main data
var request Request

func getOccupied(s []Selection) []Selection {
	var res []Selection
	for _, v := range s {
		if v.Occupied == true {
			slc := Selection{
				Input: v.Input,
				Seat: v.Seat,
				Occupied: true,
			}
			res = append(res, slc)
		}
	}
	return res
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func setId(r *Request) *Request {
	for k, v := range r.Seats {
		input, _ := json.Marshal(v.Input)
		seat, _ := json.Marshal(v.Seat)
		v.Id =  string(input) + string(seat) + strconv.FormatBool(v.Occupied)
		r.Seats[k] = v
	}
	return r
}

func stringInSlice(a string, list []string) bool {

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func removeRepeated(r *Request) *Request {
	var s []string
	//var result []Selection
	for _, v := range r.Seats {
		s = append(s, v.Id)
	}
	seats := r.Seats
	var items []Selection
	checkList := []string{"32false", "75true", "105true"}
	for _, v := range seats {
		var count = true
		for _, v1 := range seats {

			if v.Id == v1.Id  && count && stringInSlice(v.Id, checkList) {
				count = false
				checkList = remove(checkList, v.Id)
				items = append(items, v1)
			}
		}
	}
	response := Request{Rows:r.Rows, SeatsPerRow: r.SeatsPerRow, Seats: items}
	log.Println(response.Seats)
	return &response
}

func getCharacters(inputRows int) (map[int]string, error) {
	result := make(map[int]string)
	if inputRows == 4{
		result[0] = "B"
		result[1] = "C"
		result[2] = "A"
		result[3] = "D"
		return result, nil
	} else if inputRows == 6 {
		result[0] = "B"
		result[1] = "C"
		result[2] = "D"
		result[3] = "E"
		result[4] = "A"
		result[5] = "F"
		return result, nil

	} else if inputRows == 10 {
		result[0] = "C"
		result[1] = "D"
		result[2] = "G"
		result[3] = "H"
		result[4] = "A"
		result[5] = "J"
		result[6] = "B"
		result[7] = "E"
		result[8] = "F"
		result[9] = "I"
		return result, nil

	} else {
		return result, errors.New("invalid input row number")
	}
}


func main() {
	port := os.Getenv("PORT")
	if port =="" {
		port = ":8000"
	}
	router := gin.Default()
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")
	v1.Use(RequestValidator())
	{
		v1.POST("/setSeat", setSeat)
		v1.GET("/getSeat", getSeat)
	}
	router.Run(port)
}
