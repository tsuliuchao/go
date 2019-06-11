package main

import (
 "gopkg.in/gin-gonic/gin.v1"
 . "github.com/liuchao/project/apis"
)

func initRouter() *gin.Engine {
 router := gin.Default()

 router.GET("/", IndexApi)

 router.POST("/person", AddPersonApi)

 router.GET("/person", AddPersonApi)

 router.GET("/persons", GetPersonApi)

 router.GET("/person/:id", GetOnePersonApi)

// router.PUT("/person/:id", ModPersonApi)

// router.DELETE("/person/:id", DelPersonApi)

 return router
}