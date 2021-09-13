package models

// META
type Meta struct {
	id               string
	protocol_version string
	time             string
	date             string
	temperature      float64
	sequence         int32
	firmware         string
}

// BAT
type Battery struct {
	volatage    int32
	average     int32
	total       int32
	temperature float64
}

// GSTAT
type GNSS struct {
	hdop     float64
	no_sat   int32
	accuracy float64
	time_fix int32
}

// Position POS (optional)
type Position struct {
	latitude_nmea       float64
	latitude_direction  byte
	longitude_nmea      float64
	longitude_direction byte
	altitude            int32
	speed               float64
}

// Network Status N_STAT (optional)
type NetworkStatus struct {
	base_station_id string
	rssi            uint8
	last_sequence   uint32
}

// Network Status N_STAT (optional)
type Orientation struct {
	pitch           int8
	roll            int16
	heading_compass uint16
}

// Error Message ERR (optional)
type ErrorMessage string

//Event Shock EVT_S (optional)

type Event_Shock struct {
	magnitude float64
	duration  uint32
	date      string
	time      string
	state     string
}

const (
	SHOCK = iota
	LEAVE
	ARRIVE
	STAP
	DTAP
)

// Event Reset EVT_R (optional)

type Event_Reset struct {
	date   string
	time   string
	reason string
}

/// Message Authentication Code MAC (required)
type MAC string
