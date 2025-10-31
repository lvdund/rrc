package ies

// NTN-NeighCellConfig-r17 ::= SEQUENCE
type NtnNeighcellconfigR17 struct {
	NtnConfigR17   *NtnConfigR17
	CarrierfreqR17 *ArfcnValuenr
	PhyscellidR17  *Physcellid
}
