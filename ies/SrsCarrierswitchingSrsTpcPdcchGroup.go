package ies

// SRS-CarrierSwitching-srs-TPC-PDCCH-Group ::= CHOICE
const (
	SrsCarrierswitchingSrsTpcPdcchGroupChoiceNothing = iota
	SrsCarrierswitchingSrsTpcPdcchGroupChoiceTypea
	SrsCarrierswitchingSrsTpcPdcchGroupChoiceTypeb
)

type SrsCarrierswitchingSrsTpcPdcchGroup struct {
	Choice uint64
	Typea  *[]SrsTpcPdcchConfig `lb:1,ub:32`
	Typeb  *SrsTpcPdcchConfig
}
