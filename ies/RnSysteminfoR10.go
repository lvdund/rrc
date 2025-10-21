package ies

import "rrc/utils"

// RN-SystemInfo-r10 ::= SEQUENCE
// Extensible
type RnSysteminfoR10 struct {
	Systeminformationblocktype1R10 *utils.OCTETSTRING
	Systeminformationblocktype2R10 *Systeminformationblocktype2
}
