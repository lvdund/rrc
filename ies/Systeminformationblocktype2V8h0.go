package ies

// SystemInformationBlockType2-v8h0-IEs ::= SEQUENCE
type Systeminformationblocktype2V8h0 struct {
	Multibandinfolist    *[]Additionalspectrumemission `lb:1,ub:maxMultiBands`
	Noncriticalextension *Systeminformationblocktype2V9e0
}
