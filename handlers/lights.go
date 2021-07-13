package handlers

import "fmt"

type LightHandler struct {
	LightsService LightsService
}

func (l *LightHandler) Do(lightNo string, status bool) error {
	result, err := l.LightsService.LightSwitch(lightNo, status)
	if err != nil {
		return fmt.Errorf(`result, err := l.LightsService.LightSwitch(lightNo, status): %w`, err)
	}

	fmt.Println(result)
	return nil
}
