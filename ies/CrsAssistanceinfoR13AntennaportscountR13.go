package ies

import "rrc/utils"

// CRS-AssistanceInfo-r13-antennaPortsCount-r13 ::= ENUMERATED
type CrsAssistanceinfoR13AntennaportscountR13 struct {
	Value utils.ENUMERATED
}

const (
	CrsAssistanceinfoR13AntennaportscountR13EnumeratedNothing = iota
	CrsAssistanceinfoR13AntennaportscountR13EnumeratedAn1
	CrsAssistanceinfoR13AntennaportscountR13EnumeratedAn2
	CrsAssistanceinfoR13AntennaportscountR13EnumeratedAn4
	CrsAssistanceinfoR13AntennaportscountR13EnumeratedSpare1
)
