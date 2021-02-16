package api

import (
	"github.com/george124816/job-dev-backend-interview/internal/api/router"
	"github.com/george124816/job-dev-backend-interview/internal/pkg/util"
)

func Run() {
	web := router.Setup()
	web.Run(util.LoadConfigPort("."))
}
