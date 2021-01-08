package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/jusso-dev/Muruwari-Language/src/server/dictionary-search-api/search-api"
)

// PHRASE struct for POST request /search-phrase
type PHRASE struct {
	WORD string `json:"word" binding:"required"`
}

// APIRES struct for return object to be return as JSON
// USERWORD = word/phrase user passed into the API
// WORD = The Muruwari word
// TRANSLATION = you guessed it, the English translation of the phrase or word
type APIRES struct {
	USERWORD    string
	WORD        string
	TRANSLATION string
}

func main() {
	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/search-phrase", func(c *gin.Context) {

		var phrase PHRASE
		c.BindJSON(&phrase)

		phraseRes, err := api.SearchPhrase(phrase.WORD)

		if err != nil {
			c.JSON(500, err)
		} else {
			res := &APIRES{}
			res.TRANSLATION = phraseRes.Hits[0].Translation
			res.WORD = phraseRes.Hits[0].WordPhrase
			res.USERWORD = phrase.WORD

			c.JSON(200, *res)
		}
	})

	router.Run(":8000")
}
