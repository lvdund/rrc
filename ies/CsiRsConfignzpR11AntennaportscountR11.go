package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-r11-antennaPortsCount-r11 ::= ENUMERATED
type CsiRsConfignzpR11AntennaportscountR11 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignzpR11AntennaportscountR11EnumeratedNothing = iota
	CsiRsConfignzpR11AntennaportscountR11EnumeratedAn1
	CsiRsConfignzpR11AntennaportscountR11EnumeratedAn2
	CsiRsConfignzpR11AntennaportscountR11EnumeratedAn4
	CsiRsConfignzpR11AntennaportscountR11EnumeratedAn8
)
