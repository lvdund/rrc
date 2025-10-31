package ies

import "rrc/utils"

// AntennaInfoCommon-antennaPortsCount ::= ENUMERATED
type AntennainfocommonAntennaportscount struct {
	Value utils.ENUMERATED
}

const (
	AntennainfocommonAntennaportscountEnumeratedNothing = iota
	AntennainfocommonAntennaportscountEnumeratedAn1
	AntennainfocommonAntennaportscountEnumeratedAn2
	AntennainfocommonAntennaportscountEnumeratedAn4
	AntennainfocommonAntennaportscountEnumeratedSpare1
)
