package ies

// SRS-TPC-PDCCH-Config ::= SEQUENCE
type SrsTpcPdcchConfig struct {
	SrsCcSetindexlist *[]SrsCcSetindex `lb:1,ub:4`
}
