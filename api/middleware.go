package api

import (
	"github.com/farismecinovic/bankapp/token"
	"github.com/gin-gonic/gin"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
