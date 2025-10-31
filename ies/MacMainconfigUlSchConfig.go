package ies

import "rrc/utils"

// MAC-MainConfig-ul-SCH-Config ::= SEQUENCE
type MacMainconfigUlSchConfig struct {
	MaxharqTx        *MacMainconfigUlSchConfigMaxharqTx
	PeriodicbsrTimer *PeriodicbsrTimerR12
	RetxbsrTimer     RetxbsrTimerR12
	Ttibundling      utils.BOOLEAN
}
