package ies

// PhysCellIdRange ::= SEQUENCE
type Physcellidrange struct {
	Start Physcellid
	Range *PhyscellidrangeRange
}
