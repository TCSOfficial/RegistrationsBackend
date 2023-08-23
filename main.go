package main

import (
	_ "github.com/joho/godotenv/autoload"
  d "github.com/TCSOfficial/RegistrationsBackend/database"
  re "github.com/TCSOfficial/RegistrationsBackend/registrations"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
	"os"
  "time"
)

func main() {
	d.ConnectToDB()
	re.Init() // shitty

  r := gin.New()

  r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{"GET, POST"},
    AllowHeaders: []string{"Origin"},
    ExposeHeaders: []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return true // origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
  }))

  r.LoadHTMLGlob("templates/*")

	re.Routes(r)

  r.Run(":" + os.Getenv("PORT"))
}
