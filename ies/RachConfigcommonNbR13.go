package ies

import "rrc/utils"

// RACH-ConfigCommon-NB-r13 ::= SEQUENCE
// Extensible
type RachConfigcommonNbR13 struct {
	PreambletransmaxCeR13     Preambletransmax
	PowerrampingparametersR13 Powerrampingparameters
	RachInfolistR13           RachInfolistNbR13
	ConnestfailoffsetR13      *utils.INTEGER
}
