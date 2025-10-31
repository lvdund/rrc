package ies

import "rrc/utils"

// SHR-Cause-r17 ::= SEQUENCE
// Extensible
type ShrCauseR17 struct {
	T304CauseR17         *utils.BOOLEAN
	T310CauseR17         *utils.BOOLEAN
	T312CauseR17         *utils.BOOLEAN
	SourcedapsFailureR17 *utils.BOOLEAN
}
