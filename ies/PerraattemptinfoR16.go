package ies

import "rrc/utils"

// PerRAAttemptInfo-r16 ::= SEQUENCE
// Extensible
type PerraattemptinfoR16 struct {
	ContentiondetectedR16   *utils.BOOLEAN
	DlrsrpabovethresholdR16 *utils.BOOLEAN
	FallbacktofourstepraR17 *utils.BOOLEAN `ext`
}
