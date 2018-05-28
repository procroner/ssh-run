package cmdJob

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"github.com/procroner/ssh-run/core/coreServer"
	"github.com/procroner/ssh-run/core/coreJob"
	"strconv"
)

func List() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"ID", "Title", "Server Host", "Command", "Status"})

	jobs := coreJob.All()
	for _, job := range jobs {
		server := coreServer.Query(job.ServerId)
		row := []string{
			strconv.Itoa(int(job.ID)),
			job.Title,
			server.Host,
			job.Command,
			coreJob.StatusDesc[job.Status],
		}
		table.Append(row)
	}
	table.Render()
}
