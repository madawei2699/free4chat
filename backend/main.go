package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/MixinNetwork/kraken/engine"
	"github.com/MixinNetwork/kraken/monitor"
	"github.com/MixinNetwork/mixin/logger"
)

func main() {
	cp := flag.String("c", "~/.kraken/engine.toml", "configuration file path")
	sr := flag.String("s", "engine", "service engine or monitor")
	flag.Parse()

	if strings.HasPrefix(*cp, "~/") {
		usr, _ := user.Current()
		*cp = filepath.Join(usr.HomeDir, (*cp)[2:])
	}

	logger.SetLevel(logger.VERBOSE)

	go func() {
		http.ListenAndServe(":9000", http.DefaultServeMux)
	}()

	switch *sr {
	case "engine":
		engine.Boot(*cp)
	case "monitor":
		monitor.Boot(*cp)
	}
}
