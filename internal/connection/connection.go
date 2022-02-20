package connection

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
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
	// r.Use(cors.New(cors.Config{
	// 	// アクセス許可するオリジン
	// 	AllowOrigins: []string{
	// 		"http://localhost",
	// 	},
	// 	// アクセス許可するHTTPメソッド
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"OPTIONS",
	// 	},
	// 	// 許可するHTTPリクエストヘッダ
	// 	AllowHeaders: []string{
	// 		"Access-Control-Allow-Headers",
	// 		"Access-Control-Allow-Origin",
	// 		"Access-Control-Allow-Credentials",
	// 		"Content-Type",
	// 	},
	// 	// cookieなどの情報を必要とするかどうか
	// 	AllowCredentials: false,
	// 	// preflightリクエストの結果をキャッシュする時間
	// 	MaxAge: 24 * time.Hour,
	// }))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!, this is pokewordle_solver!!",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/poke", func(c *gin.Context) {
		s := c.DefaultQuery("name", "")
		l := c.DefaultQuery("less", DEFAULT_LESS_GEN)
		g := c.DefaultQuery("greater", DEFAULT_GREATER_GEN)
		// .:no hit, x:hit but position is not correct, o:hit and position is correct
		h := c.DefaultQuery("hits", DEFAULT_HITS)
		var r types.Response
		if l < DEFAULT_LESS_GEN || g > DEFAULT_GREATER_GEN {
			err := fmt.Errorf("invalid generation")
			r.Error = err.Error()
			c.JSON(400, gin.H{
				"response": r,
			})
			return
		}
		if l > g {
			err := fmt.Errorf("invalid generation")
			r.Error = err.Error()
			c.JSON(400, gin.H{
				"response": r,
			})
			return
		}
		q, err := check.ParseCorrect(s, h)
		if err != nil {
			fmt.Println(err)
			r.Error = err.Error()
			c.JSON(400, gin.H{
				"response": r,
			})
			return
		}
		p, err := check.ParsePartial(s, h)
		if err != nil {
			fmt.Println(err)
			r.Error = err.Error()
			c.JSON(400, gin.H{
				"response": r,
			})
			return
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
		r.Names = names
		c.JSON(200, gin.H{
			"response": r,
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
