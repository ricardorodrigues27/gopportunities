package handler

import (
	"github.com/ricardorodrigues27/gopportunities.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
