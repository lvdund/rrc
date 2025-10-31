package ies

import "rrc/utils"

// SL-SyncConfig-r16-txParameters-r16 ::= SEQUENCE
type SlSyncconfigR16TxparametersR16 struct {
	SynctxthreshicR16   *SlRsrpRangeR16
	SynctxthreshoocR16  *SlRsrpRangeR16
	SyncinforeservedR16 *utils.BITSTRING `lb:2,ub:2`
}
