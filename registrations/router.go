package registrations

import (
	"log"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	q := route.Group("/api/registrations/quates")
	{
		q.GET("/", func(ctx *gin.Context) {
			var registrations []Registration
  	  res := db.Find(&registrations)

			pass, ok := ctx.GetQuery("pass")
			if !ok || pass != os.Getenv("SECRET") {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				ctx.Abort()
				return
			}

			if res.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				ctx.Abort()
			  log.Printf("Unexpected error while reading subects from DB: %v\n", res.Error)
				return
			}

			//ctx.JSON(http.StatusOK, gin.H{"data": registrations})
      ctx.HTML(http.StatusOK, "registrations.tmpl", gin.H{
        "total": len(registrations),
        "registrations": registrations,
      })
		})

		q.POST("/", func(ctx *gin.Context) {
			var r Registration
			ctx.Bind(&r)

			res := db.Create(&r)
			if res.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			  log.Printf("Unexpected error while saving new registration to DB: %v\n", res.Error)
				return
			}

			log.Printf("Saved new registration to DB: %d\n", r.ID)
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
}
