package monitor

func Boot(cp string) {
	conf, err := Setup(cp)
	if err != nil {
		panic(err)
	}

	monitor, err := BuildMonitor(conf)
	if err != nil {
		panic(err)
	}

	go monitor.Loop()
	ServeRPC(monitor, conf)
}
