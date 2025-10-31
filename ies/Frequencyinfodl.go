package ies

// FrequencyInfoDL ::= SEQUENCE
// Extensible
type Frequencyinfodl struct {
	Absolutefrequencyssb    *ArfcnValuenr
	Frequencybandlist       Multifrequencybandlistnr
	Absolutefrequencypointa ArfcnValuenr
	ScsSpecificcarrierlist  []ScsSpecificcarrier `lb:1,ub:maxSCSs`
}
