package ies

import "rrc/utils"

// SystemTimeInfoCDMA2000-cdma-SystemTime ::= CHOICE
const (
	Systemtimeinfocdma2000CdmaSystemtimeChoiceNothing = iota
	Systemtimeinfocdma2000CdmaSystemtimeChoiceSynchronoussystemtime
	Systemtimeinfocdma2000CdmaSystemtimeChoiceAsynchronoussystemtime
)

type Systemtimeinfocdma2000CdmaSystemtime struct {
	Choice                 uint64
	Synchronoussystemtime  *utils.BITSTRING `lb:39,ub:39`
	Asynchronoussystemtime *utils.BITSTRING `lb:49,ub:49`
}
