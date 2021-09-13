package models

type UplinkMessage struct {
	meta     Meta
	battery  Battery
	gnss     GNSS
	position *Position
}

//UplinkMessageService interface for UplinkMessage struct
type UplinkMessageService interface {
	// create a new breed
	UplinkMessageToJson(message UplinkMessage) error
	// get breed based on category identifiers
	UpdateDevice(message UplinkMessage) error
}
