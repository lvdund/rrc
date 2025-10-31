package ies

// PCI-ARFCN-NB-r14 ::= SEQUENCE
type PciArfcnNbR14 struct {
	PhyscellidR14  Physcellid
	CarrierfreqR14 *CarrierfreqNbR13
}
