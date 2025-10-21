package ies

import "rrc/utils"

// SL-Parameters-v1310 ::= SEQUENCE
type SlParametersV1310 struct {
	DiscsysinforeportingR13 *utils.ENUMERATED
	CommmultipletxR13       *utils.ENUMERATED
	DiscinterfreqtxR13      *utils.ENUMERATED
	DiscperiodicslssR13     *utils.ENUMERATED
}
