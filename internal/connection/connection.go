package connection

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otyamura/pokewordle_solver/internal/check"
	"github.com/otyamura/pokewordle_solver/internal/search"
	"github.com/otyamura/pokewordle_solver/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DEFAULT_LESS_GEN = "1"
const DEFAULT_GREATER_GEN = "8"
const DEFAULT_HITS = "xxxxx"

func CreateConnection() (*gorm.DB, *gin.Engine) {
	db := CreateDBConnection()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!, this is pokewordle_solver!!",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/poke", func(c *gin.Context) {
		s := c.Query("name")
		l := c.DefaultQuery("less", DEFAULT_LESS_GEN)
		g := c.DefaultQuery("greater", DEFAULT_LESS_GEN)
		// .:no hit, x:hit but position is not correct, o:hit and position is correct
		h := c.DefaultQuery("hits", DEFAULT_HITS)
		q, err := check.ParseCorrect(s, h)
		if err != nil {
			log.Fatal(err)
		}
		p, err := check.ParsePartial(s, h)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(q)
		fmt.Println(p)
		var pokes []types.Pokemon
		tmp := search.SearchGeneration(db, l, g)
		if q == "_____" {
			tmp = search.SearchPartial(tmp, p)
		} else {
			tmp = search.SearchCorrect(tmp, q)
			tmp = search.SearchPartial(tmp, p)
		}
		tmp.Find(&pokes)
		fmt.Println(pokes)
		var names []string
		for _, poke := range pokes {
			names = append(names, poke.Name)
		}
		c.JSON(200, gin.H{
			"names": names,
		})
	})
	return db, r
}

func CreateDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db/pokemon.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
