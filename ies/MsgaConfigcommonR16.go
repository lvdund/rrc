package ies

// MsgA-ConfigCommon-r16 ::= SEQUENCE
type MsgaConfigcommonR16 struct {
	RachConfigcommontwostepraR16 RachConfigcommontwostepraR16
	MsgaPuschConfigR16           *MsgaPuschConfigR16
}
