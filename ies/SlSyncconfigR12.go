package ies

// SL-SyncConfig-r12 ::= SEQUENCE
// Extensible
type SlSyncconfigR12 struct {
	SynccpLenR12           SlCpLenR12
	SyncoffsetindicatorR12 SlOffsetindicatorsyncR12
	SlssidR12              SlssidR12
	TxparametersR12        *SlSyncconfigR12TxparametersR12
	RxparamsncellR12       *SlSyncconfigR12RxparamsncellR12
}
