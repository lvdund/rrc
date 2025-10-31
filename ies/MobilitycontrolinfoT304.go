package ies

import "rrc/utils"

// MobilityControlInfo-t304 ::= ENUMERATED
type MobilitycontrolinfoT304 struct {
	Value utils.ENUMERATED
}

const (
	MobilitycontrolinfoT304EnumeratedNothing = iota
	MobilitycontrolinfoT304EnumeratedMs50
	MobilitycontrolinfoT304EnumeratedMs100
	MobilitycontrolinfoT304EnumeratedMs150
	MobilitycontrolinfoT304EnumeratedMs200
	MobilitycontrolinfoT304EnumeratedMs500
	MobilitycontrolinfoT304EnumeratedMs1000
	MobilitycontrolinfoT304EnumeratedMs2000
	MobilitycontrolinfoT304EnumeratedMs10000_V1310
)
