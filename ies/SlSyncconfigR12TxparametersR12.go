package ies

import "rrc/utils"

// SL-SyncConfig-r12-txParameters-r12 ::= SEQUENCE
type SlSyncconfigR12TxparametersR12 struct {
	SynctxparametersR12 SlTxparametersR12
	SynctxthreshicR12   RsrpRangeslR12
	SyncinforeservedR12 *utils.BITSTRING `lb:19,ub:19`
}
