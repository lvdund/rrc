package ies

import "rrc/utils"

// PDSCH-Config-antennaPortsFieldPresenceDCI-1-2-r16 ::= ENUMERATED
type PdschConfigAntennaportsfieldpresencedci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigAntennaportsfieldpresencedci12R16EnumeratedNothing = iota
	PdschConfigAntennaportsfieldpresencedci12R16EnumeratedEnabled
)
