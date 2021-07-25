package main

import (
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type configMain struct {
	Cmd string
}

type configServers struct {
	Name     string `xml:"Name,attr"`
	Host     string
	Port     int
	UserName string
	PassWord string
	SshKey   string
}

type config struct {
	Main    configMain
	Servers []configServers
}

// LoadConfig to struct
func LoadConfig() *config {
	// load xml file
	content, err := ioutil.ReadFile(DirRoot + "config.xml")
	if err != nil {
		logrus.Fatalln(err)
	}

	// parse xml to object
	xmlConfig := new(config)
	err = xml.Unmarshal(content, xmlConfig)
	if err != nil {
		logrus.Fatalln(err)
	}

	return xmlConfig
}
