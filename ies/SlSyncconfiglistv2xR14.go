package ies

// SL-SyncConfigListV2X-r14 ::= SEQUENCE OF SL-SyncConfig-r12
// SIZE (1.. maxSL-V2X-SyncConfig-r14)
type SlSyncconfiglistv2xR14 struct {
	Value []SlSyncconfigR12 `lb:1,ub:maxSLV2xSyncconfigR14`
}
