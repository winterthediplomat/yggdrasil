//This file is part of Yggdrasil
//
//Yggdrasil is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//Yggdrasil is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with Yggdrasil.  If not, see <http://www.gnu.org/licenses/>.

package main

type Config struct {
	Services []Service `json:"services"`
	Irc      IrcConfig `json:"irc"`
	Interval int       `json:"interval"`
}

type IrcConfig struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Nick     string `json:"nick"`
	RealName string `json:"realname"`
	Channel  string `json:"channel"`
}

type Service struct {
	Host  string      `json:"host"`
	Port  int         `json:"port"`
	Proto string      `json:"proto"`
	Type  ServiceType `json:"type"`
	Name  string      `json:"name"`
}

type ServiceType string

const (
	HTTPService ServiceType = "http"
	FTPService  ServiceType = "ftp"
	SSHService  ServiceType = "ssh"
	SMTPService ServiceType = "smtp"
)
