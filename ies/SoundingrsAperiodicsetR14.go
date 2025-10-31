package ies

// SoundingRS-AperiodicSet-r14 ::= SEQUENCE
type SoundingrsAperiodicsetR14 struct {
	SrsCcSetindexlistR14                    *[]SrsCcSetindexR14 `lb:1,ub:4`
	SoundingrsUlConfigdedicatedaperiodicR14 SoundingrsUlConfigdedicatedaperiodicR10
}
