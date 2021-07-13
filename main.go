package main

import (
	"fmt"
	"os"

	"github.com/yoshhiide/hue/config"
	"github.com/yoshhiide/hue/handlers"
	"github.com/yoshhiide/hue/infrastructures"
	"github.com/yoshhiide/hue/services"
)

var (
	conf          *config.Config
	lightsService *services.LightsService
	status        bool
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := setup(); err != nil {
		fmt.Errorf(`setup(): %w`, err)
		return 1
	}

	if err := cmdLightSwitch(); err != nil {
		fmt.Errorf(err.Error())
		return 1
	}

	return 0
}

func setup() error {
	var err error
	conf, err = config.Load()
	if err != nil {
		return fmt.Errorf(`config.Load(): %w`, err)
	}

	hueBridgeClient := infrastructures.NewHueBridgeClient(
		conf.BridgeURL,
		conf.Token,
	)
	lightsService = &services.LightsService{
		HueBridgeClient: hueBridgeClient,
	}

	return nil
}

func cmdLightSwitch() error {
	if len(os.Args) == 1 {
		return fmt.Errorf(`require args 1 or 2`)
	}
	if len(os.Args) == 2 {
		status = false
	}
	if len(os.Args) == 3 {
		status = true
	}

	h := &handlers.LightHandler{
		LightsService: lightsService,
	}
	if err := h.Do(string(os.Args[1]), status); err != nil {
		return fmt.Errorf(`err := h.Do(): %w`, err)
	}
	return nil
}
