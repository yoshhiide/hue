package services

type HueBridgeClient interface {
	PutChangeLightStatus(lightNo string, status bool) ([]byte, error)
}
