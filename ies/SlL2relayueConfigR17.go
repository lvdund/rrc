package ies

// SL-L2RelayUE-Config-r17 ::= SEQUENCE
// Extensible
type SlL2relayueConfigR17 struct {
	SlRemoteueToaddmodlistR17  *[]SlRemoteueToaddmodR17    `lb:1,ub:maxNrofRemoteUER17`
	SlRemoteueToreleaselistR17 *[]SlDestinationidentityR16 `lb:1,ub:maxNrofRemoteUER17`
}
