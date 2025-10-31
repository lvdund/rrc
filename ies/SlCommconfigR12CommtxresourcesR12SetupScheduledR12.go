package ies

import "rrc/utils"

// SL-CommConfig-r12-commTxResources-r12-setup-scheduled-r12 ::= SEQUENCE
type SlCommconfigR12CommtxresourcesR12SetupScheduledR12 struct {
	SlRntiR12         CRnti
	MacMainconfigR12  MacMainconfigslR12
	ScCommtxconfigR12 SlCommresourcepoolR12
	McsR12            *utils.INTEGER `lb:0,ub:28`
}
