package ies

import "rrc/utils"

// NPDCCH-ConfigDedicated-NB-v1530-npdcch-StartSF-USS-v1530 ::= ENUMERATED
type NpdcchConfigdedicatedNbV1530NpdcchStartsfUssV1530 struct {
	Value utils.ENUMERATED
}

const (
	NpdcchConfigdedicatedNbV1530NpdcchStartsfUssV1530EnumeratedNothing = iota
	NpdcchConfigdedicatedNbV1530NpdcchStartsfUssV1530EnumeratedV96
	NpdcchConfigdedicatedNbV1530NpdcchStartsfUssV1530EnumeratedV128
)
