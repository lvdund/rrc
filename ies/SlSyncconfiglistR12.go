package ies

// SL-SyncConfigList-r12 ::= SEQUENCE OF SL-SyncConfig-r12
// SIZE (1..maxSL-SyncConfig-r12)
type SlSyncconfiglistR12 struct {
	Value []SlSyncconfigR12 `lb:1,ub:maxSLSyncconfigR12`
}
