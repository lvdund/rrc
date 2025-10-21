package ies

import "rrc/utils"

// RadioResourceConfigDedicatedSCG-r12 ::= SEQUENCE
// Extensible
type RadioresourceconfigdedicatedscgR12 struct {
	DrbToaddmodlistscgR12       *DrbToaddmodlistscgR12
	MacMainconfigscgR12         *MacMainconfig
	RlfTimersandconstantsscgR12 *RlfTimersandconstantsscgR12
}
