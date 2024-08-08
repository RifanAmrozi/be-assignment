package config

import (
	"github.com/rifanamrozi/payment-service/db"
)

var PrismaClient *db.PrismaClient

func InitPrisma() {
	PrismaClient = db.NewClient()
}
