package ies

import "rrc/utils"

// SL-SyncConfig-r12-rxParamsNCell-r12-discSyncWindow-r12 ::= ENUMERATED
type SlSyncconfigR12RxparamsncellR12DiscsyncwindowR12 struct {
	Value utils.ENUMERATED
}

const (
	SlSyncconfigR12RxparamsncellR12DiscsyncwindowR12EnumeratedNothing = iota
	SlSyncconfigR12RxparamsncellR12DiscsyncwindowR12EnumeratedW1
	SlSyncconfigR12RxparamsncellR12DiscsyncwindowR12EnumeratedW2
)
