package ies

// PCI-ARFCN-r13 ::= SEQUENCE
type PciArfcnR13 struct {
	PhyscellidR13  Physcellid
	CarrierfreqR13 *ArfcnValueeutraR9
}
