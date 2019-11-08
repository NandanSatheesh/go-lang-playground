package main

import (
	"fmt"
	"github.com/NandanSatheesh/go-lang-playground/startup-name-generator/models"
	"github.com/NandanSatheesh/go-lang-playground/startup-name-generator/score"
	"github.com/NandanSatheesh/go-lang-playground/startup-name-generator/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	str "strings"
)

func main() {

	g := gin.Default()

	g.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Running Fine"})
	})

	g.GET("/:words", func(context *gin.Context) {
		parameters := context.Params.ByName("words")
		var words []string = str.Split(parameters, " ")
		var namesList []string = utils.Permutate(words)

		scoreMap := make(models.ScoreMap, len(namesList))

		for index, nameGenerated := range namesList {
			scoreMap[index] = models.Score{
				Name:  utils.Normalize(nameGenerated),
				Value: score.Score(nameGenerated),
			}
		}

		sort.Sort(sort.Reverse(scoreMap))

		fmt.Println(scoreMap)
		var namesSuggestion []string

		namesSuggestion = make([]string, len(scoreMap))

		for index, eachNameWithScore := range scoreMap {
			namesSuggestion[index] = eachNameWithScore.Name
		}

		var response = models.NamesResponse{
			Data:    models.Data{Names: namesSuggestion},
			Status:  200,
			Message: "Success",
		}

		context.JSON(http.StatusOK, response)
	})

	g.Run(":8080")
}
