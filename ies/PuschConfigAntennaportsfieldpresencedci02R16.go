package ies

import "rrc/utils"

// PUSCH-Config-antennaPortsFieldPresenceDCI-0-2-r16 ::= ENUMERATED
type PuschConfigAntennaportsfieldpresencedci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigAntennaportsfieldpresencedci02R16EnumeratedNothing = iota
	PuschConfigAntennaportsfieldpresencedci02R16EnumeratedEnabled
)
