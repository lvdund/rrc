package ies

// SL-SyncConfigList-r16 ::= SEQUENCE OF SL-SyncConfig-r16
// SIZE (1..maxSL-SyncConfig-r16)
type SlSyncconfiglistR16 struct {
	Value []SlSyncconfigR16 `lb:1,ub:maxSLSyncconfigR16`
}
