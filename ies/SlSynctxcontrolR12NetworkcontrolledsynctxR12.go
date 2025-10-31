package ies

import "rrc/utils"

// SL-SyncTxControl-r12-networkControlledSyncTx-r12 ::= ENUMERATED
type SlSynctxcontrolR12NetworkcontrolledsynctxR12 struct {
	Value utils.ENUMERATED
}

const (
	SlSynctxcontrolR12NetworkcontrolledsynctxR12EnumeratedNothing = iota
	SlSynctxcontrolR12NetworkcontrolledsynctxR12EnumeratedOn
	SlSynctxcontrolR12NetworkcontrolledsynctxR12EnumeratedOff
)
