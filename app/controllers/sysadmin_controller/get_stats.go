package sysadmin_controller

import "github.com/gin-gonic/gin"

type UserData struct {
}

type PostData struct {
}

type StatsResp struct {
	UserStats UserData `json:"user_stats"`
	PostStats PostData `json:"post_stats"`
}

func GetStats(c *gin.Context) {

}
