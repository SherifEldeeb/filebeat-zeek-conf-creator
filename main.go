package main

import (
	"fmt"
	"log"
	"os"
)

var logs = []string{
	"conn",
	"dce_rpc",
	"dhcp",
	"dnp3",
	"dns",
	"ftp",
	"http",
	"irc",
	"kerberos",
	"modbus",
	"modbus_register_change",
	"mysql",
	"ntlm",
	"radius",
	"rdp",
	"rfb",
	"sip",
	"smb_cmd",
	"smb_files",
	"smb_mapping",
	"smtp",
	"snmp",
	"socks",
	"ssh",
	"ssl",
	"syslog",
	"tunnel",
	"files",
	"ocsp",
	"pe",
	"x509",
	"netcontrol",
	"netcontrol_drop",
	"netcontrol_shunt",
	"netcontrol_catch_release",
	"openflow",
	"intel",
	"notice",
	"notice_alarm",
	"signatures",
	"traceroute",
	"known_certs",
	"known_hosts",
	"known_modbus",
	"known_services",
	"software",
	"barnyard2",
	"dpd",
	"unified2",
	"weird",
}

var conf = `filebeat.inputs:
%s
output.kafka:
  enabled: true
  hosts: ["localhost:9092"]
  topic: zeek
  codec.json:
    pretty: false
    escape_html: false
`

var logitem = `- type: log
  enabled: true
  paths:
    - /opt/zeek/logs/current/%[1]s.log
  json.keys_under_root: true
  fields_under_root: true
  fields.logtype: zeek-%[1]s
`

func main() {
	var AllLogs string

	for _, l := range logs {
		AllLogs = AllLogs + fmt.Sprintf(logitem, l)
	}

	config := fmt.Sprintf(conf, AllLogs)

	ofile, err := os.Create("filebeat.yml")
	if err != nil {
		log.Fatal("error creating file:", err)
	}
	defer ofile.Close()

	ofile.WriteString(config)

}
