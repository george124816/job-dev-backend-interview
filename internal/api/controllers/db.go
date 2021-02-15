package controllers

import (
	"database/sql"

	"github.com/george124816/job-dev-backend-interview/internal/pkg/db"
	"github.com/george124816/job-dev-backend-interview/internal/pkg/models"
)

var database *sql.DB
var handle *models.BaseHandler

func init() {
	database = db.ConnectDB()
	migration(database)
	handle = models.NewBaseHandler(database)
}
