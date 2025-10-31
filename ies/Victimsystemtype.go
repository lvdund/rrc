package ies

import "rrc/utils"

// VictimSystemType ::= SEQUENCE
type Victimsystemtype struct {
	Gps       *utils.BOOLEAN
	Glonass   *utils.BOOLEAN
	Bds       *utils.BOOLEAN
	Galileo   *utils.BOOLEAN
	Wlan      *utils.BOOLEAN
	Bluetooth *utils.BOOLEAN
}
