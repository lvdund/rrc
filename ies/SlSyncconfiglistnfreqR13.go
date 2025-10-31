package ies

// SL-SyncConfigListNFreq-r13 ::= SEQUENCE OF SL-SyncConfigNFreq-r13
// SIZE (1..maxSL-SyncConfig-r12)
type SlSyncconfiglistnfreqR13 struct {
	Value []SlSyncconfignfreqR13 `lb:1,ub:maxSLSyncconfigR12`
}
