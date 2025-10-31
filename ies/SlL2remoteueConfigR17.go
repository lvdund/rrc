package ies

// SL-L2RemoteUE-Config-r17 ::= SEQUENCE
// Extensible
type SlL2remoteueConfigR17 struct {
	SlSrapConfigremoteR17 *SlSrapConfigR17
	SlUeidentityremoteR17 *RntiValue
}
