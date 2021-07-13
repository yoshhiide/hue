package services

import "fmt"

type LightsService struct {
	HueBridgeClient HueBridgeClient
}

func (l *LightsService) LightSwitch(lightNo string, status bool) (string, error) {
	b, err := l.HueBridgeClient.PutChangeLightStatus(lightNo, status)
	if err != nil {
		return "", fmt.Errorf(`b, err := l.HueBridgeClient.PutChangeLightStatus(lightNo, status):%w`, err)
	}

	return string(b), nil
}
