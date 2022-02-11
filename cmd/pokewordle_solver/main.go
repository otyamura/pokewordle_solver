package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/otyamura/pokewordle_solver/internal/check"
	"github.com/otyamura/pokewordle_solver/internal/load"
	"github.com/otyamura/pokewordle_solver/internal/search"
	"github.com/otyamura/pokewordle_solver/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DEFAULT_LESS_GEN = "1"
const DEFAULT_GREATER_GEN = "8"
const DEFAULT_HITS = "xxxxx"

func main() {
	dsn := "host=db user=admin password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=pokes port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&types.Pokemon{})
	pokes := load.LoadPokes()
	db.Create(&pokes)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/poke", func(c *gin.Context) {
		s := c.Query("poke")
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
