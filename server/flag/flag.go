package parseflag

import "flag"

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config", "", "Server config")
}
