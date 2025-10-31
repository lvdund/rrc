package ies

import "rrc/utils"

// NPDCCH-ConfigDedicated-NB-r13-npdcch-Offset-USS-r13 ::= ENUMERATED
type NpdcchConfigdedicatedNbR13NpdcchOffsetUssR13 struct {
	Value utils.ENUMERATED
}

const (
	NpdcchConfigdedicatedNbR13NpdcchOffsetUssR13EnumeratedNothing = iota
	NpdcchConfigdedicatedNbR13NpdcchOffsetUssR13EnumeratedZero
	NpdcchConfigdedicatedNbR13NpdcchOffsetUssR13EnumeratedOneeighth
	NpdcchConfigdedicatedNbR13NpdcchOffsetUssR13EnumeratedOnefourth
	NpdcchConfigdedicatedNbR13NpdcchOffsetUssR13EnumeratedThreeeighth
)
