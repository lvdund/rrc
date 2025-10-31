package ies

import "rrc/utils"

// SL-SyncConfigNFreq-r13-txParameters-r13 ::= SEQUENCE
type SlSyncconfignfreqR13TxparametersR13 struct {
	SynctxparametersR13 SlTxparametersR12
	SynctxthreshicR13   RsrpRangeslR12
	SyncinforeservedR13 *utils.BITSTRING `lb:19,ub:19`
	SynctxperiodicR13   *bool
}
