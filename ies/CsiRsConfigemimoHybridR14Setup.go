package ies

import "rrc/utils"

// CSI-RS-ConfigEMIMO-Hybrid-r14-setup ::= SEQUENCE
type CsiRsConfigemimoHybridR14Setup struct {
	PeriodicityoffsetindexR14 *utils.INTEGER `lb:0,ub:1023`
	EmimoType2R14             *CsiRsConfigemimo2R14
}
