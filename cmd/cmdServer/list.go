package cmdServer

import (
	"github.com/olekukonko/tablewriter"
	"github.com/procroner/ssh-run/core/coreServer"
	"os"
	"fmt"
	"strconv"
)

func List() {
	servers := coreServer.All()
	if len(servers) == 0 {
		fmt.Println("No servers")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "User", "Host", "Auth Type", "Password", "Private Key Path", "Proxy Server ID"})
	for _, server := range servers {
		row := []string{
			strconv.Itoa(int(server.ID)),
			server.Name,
			server.User,
			server.Host,
			server.AuthType,
			server.Pass,
			server.PrivateKeyPath,
			strconv.Itoa(int(server.ProxyServerId)),
		}
		table.Append(row)
	}
	table.Render()
}
