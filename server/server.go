package server

import (
	"io/ioutil"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"os"
	"log"
	"errors"
)

const configPath = "/etc/sshrun/Server.json"

type Server struct {
	Name           string `json:"name"`
	User           string `json:"user"`
	Host           string `json:"host"`
	AuthType       string `json:"authType"`
	Pass           string `json:"pass"`
	PrivateKeyPath string `json:"privateKeyPath"`
}

var Servers map[string]*Server

func init() {
	serMap, err := parseConfigFile(configPath)
	if err != nil {
		log.Fatalf("parse Server config file error: %v\n", err)
		os.Exit(1)
	}
	Servers = serMap
}

func parseConfigFile(configPath string) (map[string]*Server, error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var servers []*Server
	json.Unmarshal(raw, &servers)
	serMap := make(map[string]*Server)
	for _, server := range servers {
		serMap[server.Name] = server
	}
	return serMap, nil
}

func List() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "User", "Host", "Auth Type", "Password", "Private Key Path"})
	for _, server := range Servers {
		row := []string{
			server.Name,
			server.User,
			server.Host,
			server.AuthType,
			server.Pass,
			server.PrivateKeyPath,
		}
		table.Append(row)
	}
	table.Render()
}

func GetServer(name string) (*Server, error) {
	server, ok := Servers[name]
	if !ok {
		return nil, errors.New("no Server found")
	}
	return server, nil
}