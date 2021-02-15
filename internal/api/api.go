package api

import "github.com/george124816/job-dev-backend-interview/internal/api/router"

func Run() {
	web := router.Setup()
	web.Run(":3000")
}
