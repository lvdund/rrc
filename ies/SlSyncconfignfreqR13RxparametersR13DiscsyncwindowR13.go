package ies

import "rrc/utils"

// SL-SyncConfigNFreq-r13-rxParameters-r13-discSyncWindow-r13 ::= ENUMERATED
type SlSyncconfignfreqR13RxparametersR13DiscsyncwindowR13 struct {
	Value utils.ENUMERATED
}

const (
	SlSyncconfignfreqR13RxparametersR13DiscsyncwindowR13EnumeratedNothing = iota
	SlSyncconfignfreqR13RxparametersR13DiscsyncwindowR13EnumeratedW1
	SlSyncconfignfreqR13RxparametersR13DiscsyncwindowR13EnumeratedW2
)
