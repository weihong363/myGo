/*
@Time : 2020/6/18 9:05 下午
@Author : ironion
@File : main
@Software: GoLand
*/
// Package main
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// main function
func main() {
	// engine
	router := gin.Default()

	// Anonymous function
	//router.GET("/hello", func(context *gin.Context) {
	//	log.Println(">>>> hello gin start <<<<")
	//	// response header
	//	context.JSON(200,gin.H{
	//		"code":200,
	//		"success":true,
	//	})
	//})

	router.GET("/hello", hello)

	// define address and port
	router.Run(":9090")
}

func hello(context *gin.Context) {
	log.Println(">>>> hello gin start <<<<")
	// response header
	context.JSON(200, gin.H{
		"code":    200,
		"success": true,
	})
}
