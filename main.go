package main

import (
	"fmt"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	// init ui
	if err := termui.Init(); err != nil {
		logrus.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	// draw server list
	list := widgets.NewList()
	list.Title = "Servers"

	conf := LoadConfig()
	list.Rows = make([]string, len(conf.Servers))
	for idx, server := range conf.Servers {
		list.Rows[idx] = fmt.Sprintf("[%d] %s", idx, server.Name)
	}
	list.TextStyle = termui.NewStyle(termui.ColorYellow)
	list.WrapText = false
	list.SetRect(0, 0, 25, len(conf.Servers)+2)
	termui.Render(list)

	previousKey := ""
	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<Enter>":
			termui.Close()

			// fetch data
			idx := list.SelectedRow
			server := conf.Servers[idx]

			// init shell arg
			sshShell := ""
			args := []string{conf.Main.Cmd, "--shell"}

			// append shell password
			if server.PassWord != "" {
				sshShell += fmt.Sprintf("sshpass -p %s ", server.PassWord)
			}

			// append shell ssh command
			host := server.Host
			if server.UserName != "" {
				host = server.UserName + "@" + host
			}
			sshShell += fmt.Sprintf("ssh %s ", host)

			// append arg port
			if server.Port != 0 {
				sshShell += fmt.Sprintf("-p %d ", server.Port)
			}

			// append ssh key
			if server.SshKey != "" {
				sshShell += fmt.Sprintf("-i %s ", server.SshKey)
			}

			// format arg, exec
			sshShell = strings.Trim(sshShell, " ")
			args = append(args, sshShell)
			err := syscall.Exec(conf.Main.Cmd, args, os.Environ())
			if err != nil {
				logrus.Errorln(err)
			}
			return
		case  "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			idx, err := strconv.Atoi(e.ID)
			if err != nil {
				logrus.Panicln(err)
			}

			if idx < len(list.Rows) {
				list.SelectedRow = idx
			}


		case "q", "<C-c>":
			return
		case "j", "<Down>":
			list.ScrollDown()
		case "k", "<Up>":
			list.ScrollUp()
		case "<C-d>":
			list.ScrollHalfPageDown()
		case "<C-u>":
			list.ScrollHalfPageUp()
		case "<C-f>":
			list.ScrollPageDown()
		case "<C-b>":
			list.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				list.ScrollTop()
			}
		case "<Home>":
			list.ScrollTop()
		case "G", "<End>":
			list.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		termui.Render(list)
	}
}
