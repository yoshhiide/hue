package handlers

type LightsService interface {
	LightSwitch(lightNo string, status bool) (string, error)
}
