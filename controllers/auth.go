package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/alexandercrosson/gingorm/middlewares"
)

const userId string = "abc123"

func Auth(c *gin.Context) {
    token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), userId)
    if err != nil {

    }
    c.JSON(200, token)
}

