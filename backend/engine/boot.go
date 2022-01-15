package engine

import "github.com/MixinNetwork/mixin/logger"

func Boot(cp string) {
	conf, err := Setup(cp)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(conf.Engine.LogLevel)

	engine, err := BuildEngine(conf)
	if err != nil {
		panic(err)
	}

	go engine.Loop()
	ServeRPC(engine, conf)
}
