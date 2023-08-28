package registrations

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
  mysql "github.com/go-mysql/errors"
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

	r := route.Group("/api/registrations/recruitment")
	{
		r.GET("/", func(ctx *gin.Context) {
			var registrations []RecruitmentRegistration
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

		r.POST("/", func(ctx *gin.Context) {
			var r RecruitmentRegistration
			ctx.Bind(&r)

			res := db.Create(&r)
			if res.Error != nil {
				_, myErr := mysql.Error(res.Error)
				if myErr == mysql.ErrDupeKey {

					split := strings.Split(res.Error.Error(), " ")
					k := split[len(split)-1]
					key := k[1:len(k)-1]

				  ctx.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("%s already in use", key)})
			    log.Printf("Conflict while saving new registration to DB: %v\n", res.Error)
				} else {
				  ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			    log.Printf("Unexpected error while saving new registration to DB: %v\n", res.Error)
				}
				return
			}

			log.Printf("Saved new registration to DB: %d\n", r.ID)
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
}
