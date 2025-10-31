package ies

// EUTRA-PhysCellIdRange ::= SEQUENCE
type EutraPhyscellidrange struct {
	Start EutraPhyscellid
	Range *EutraPhyscellidrangeRange
}
