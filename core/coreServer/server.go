package coreServer

import (
	"github.com/procroner/ssh-run/core/coreDb"
	"github.com/jinzhu/gorm"
)

type Server struct {
	gorm.Model
	Name           string
	User           string
	Host           string
	AuthType       string
	Pass           string
	PrivateKeyPath string
	ProxyServerId  int
}

func Migrate() {
	db := coreDb.Connect()
	defer db.Close()
	db.AutoMigrate(&Server{})
}

func Query(serverId int) Server {
	db := coreDb.Connect()
	defer db.Close()
	var server Server
	db.First(&server, serverId)
	return server
}

func All() []Server {
	db := coreDb.Connect()
	defer db.Close()
	var servers []Server
	db.Find(&servers)
	return servers
}