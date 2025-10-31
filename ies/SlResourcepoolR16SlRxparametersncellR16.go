package ies

import "rrc/utils"

// SL-ResourcePool-r16-sl-RxParametersNcell-r16 ::= SEQUENCE
type SlResourcepoolR16SlRxparametersncellR16 struct {
	SlTddConfigurationR16 *TddUlDlConfigcommon
	SlSyncconfigindexR16  utils.INTEGER `lb:0,ub:15`
}
