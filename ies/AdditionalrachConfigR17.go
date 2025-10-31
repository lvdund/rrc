package ies

// AdditionalRACH-Config-r17 ::= SEQUENCE
// Extensible
type AdditionalrachConfigR17 struct {
	RachConfigcommonR17 *RachConfigcommon
	MsgaConfigcommonR17 *MsgaConfigcommonR16
}
