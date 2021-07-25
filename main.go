package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"syscall"
)

func main()  {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()



	l := widgets.NewList()
	l.Title = "Servers"

	conf := LoadConfig()
	l.Rows = make([]string, len(conf.Servers))
	for idx, server := range conf.Servers {
		l.Rows[idx] = server.Name
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 25, 15)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<Enter>":
			ui.Close()

			args := []string{"ssh"}
			idx := l.SelectedRow
			server := conf.Servers[idx]

			// append arg port
			if server.Port != 0 {
				args = append(args, "-p", strconv.Itoa(server.Port))
			}

			// append arg host
			host := server.Host
			if server.UserName != "" {
				host = server.UserName + "@" + host
			}
			args = append(args, host)

			err := syscall.Exec(conf.Main.Cmd, args, os.Environ())
			if err != nil {
				log.Errorln(err)
			}

			return
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}
