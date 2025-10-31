package ies

import "rrc/utils"

// LWA-Config-r13 ::= SEQUENCE
// Extensible
type LwaConfigR13 struct {
	LwaMobilityconfigR13 *WlanMobilityconfigR13
	LwaWtCounterR13      *utils.INTEGER `lb:0,ub:65535`
}
