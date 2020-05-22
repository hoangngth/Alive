package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthorizationRequired(ctx *gin.Context) {

	session := sessions.Default(ctx)
	auth := session.Get("auth")

	if auth != true {
		ctx.JSON(401, gin.H{"error": "profile needs to be signed in to access this service"})
		ctx.Abort()
		return
	} else {
		ctx.Next()
	}
}
