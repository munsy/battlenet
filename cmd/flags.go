package main

import (
	"flag"
)

var (
	tokenFlag  = flag.String("t", "", "Set the OAuth2 token for Battle.net profile API access.")
	keyFlag    = flag.String("k", "", "Set the Battle.net client key for API access.")
	writeFlag  = flag.Bool("o", false, "Write current key/token settings in memory to config file.")
	configFlag = flag.String("f", "config.toml", "Configuration file to load key/token values.")
)
