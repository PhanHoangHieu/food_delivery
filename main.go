package main

import (
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/middleware"
	"g05-food-delivery/module/restaurant/transport/ginrestaurant"
	"g05-food-delivery/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func init() {
	err := gotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
}

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	r.Static("/static", "static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//POST restaurants
	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.UploadImage(appContext))

	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	restaurants.GET("/:id", ginrestaurant.FindRestaurant(appContext))

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
