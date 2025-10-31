package ies

// SL-SyncConfigListNFreqV2X-r14 ::= SEQUENCE OF SL-SyncConfigNFreq-r13
// SIZE (1..maxSL-V2X-SyncConfig-r14)
type SlSyncconfiglistnfreqv2xR14 struct {
	Value []SlSyncconfignfreqR13 `lb:1,ub:maxSLV2xSyncconfigR14`
}
