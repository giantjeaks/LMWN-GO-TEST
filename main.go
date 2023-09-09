package main

import (
	"covid_sumary/handle"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	url := "https://static.wongnai.com/devinterview/covid-cases.json"

	h := handle.NewCaseData()
	s := h.FetchDataForm(url)
	allCases := h.ConvertToModel(s)

	p := handle.ProvinceCounting(allCases)
	a := handle.AgeCounting(allCases)

	covid := router.Group("/covid")
	{
		covid.GET("/summary", func(c *gin.Context) {
			c.JSON(200, gin.H{"Provinces": p, "AgeGroup": a})
		})
	}

	router.Run(":8080")
}
