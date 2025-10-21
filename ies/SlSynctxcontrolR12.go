package ies

import "rrc/utils"

// SL-SyncTxControl-r12 ::= SEQUENCE
type SlSynctxcontrolR12 struct {
	NetworkcontrolledsynctxR12 *utils.ENUMERATED
}
