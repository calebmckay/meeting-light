package main

import (
	mediaDevices "github.com/antonfisher/go-media-devices-state"
	gin "github.com/gin-gonic/gin"
	cron "github.com/robfig/cron"
)

var isCameraOn bool = false
var isMicrophoneOn bool = false

func updateDevices() {
	isCameraOn, _ = mediaDevices.IsCameraOn()
	isMicrophoneOn, _ = mediaDevices.IsMicrophoneOn()
}

func main() {
	updateDevices()
	c := cron.New()
	c.AddFunc("*/2 * * * * *", updateDevices)
	c.Start()

	r := gin.Default()
	r.GET("/meetingstate", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"camera":     isCameraOn,
			"microphone": isMicrophoneOn,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
