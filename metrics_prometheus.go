package main

import (
	"fmt"
)

func promDesc(name string, help string,  typ string, metricStr string) string {

	helpStr := promHelp(name, help)
	typeStr := promType(name, typ)
	rep := fmt.Sprintf("%s%s%s", helpStr, typeStr, metricStr)
	return rep
}

func promType(name string, typ string) string {
	rep := fmt.Sprintf("# TYPE %s %s\n", name, typ)
	return rep
}

func promHelp(name string, help string) string {
	rep := fmt.Sprintf("# HELP %s %s\n", name, help)
	return rep
}

func promMetrics(name string, site string, value interface{}) string {
	rep := fmt.Sprintf("%s{site=\"%s\"} %v\n", name, site, value)
	return rep
}