package ies

import "rrc/utils"

// CRS-AssistanceInfo-r11-antennaPortsCount-r11 ::= ENUMERATED
type CrsAssistanceinfoR11AntennaportscountR11 struct {
	Value utils.ENUMERATED
}

const (
	CrsAssistanceinfoR11AntennaportscountR11EnumeratedNothing = iota
	CrsAssistanceinfoR11AntennaportscountR11EnumeratedAn1
	CrsAssistanceinfoR11AntennaportscountR11EnumeratedAn2
	CrsAssistanceinfoR11AntennaportscountR11EnumeratedAn4
	CrsAssistanceinfoR11AntennaportscountR11EnumeratedSpare1
)
