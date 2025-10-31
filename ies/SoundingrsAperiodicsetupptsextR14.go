package ies

// SoundingRS-AperiodicSetUpPTsExt-r14 ::= SEQUENCE
type SoundingrsAperiodicsetupptsextR14 struct {
	SrsCcSetindexlistR14                            *[]SrsCcSetindexR14 `lb:1,ub:4`
	SoundingrsUlConfigdedicatedaperiodicupptsextR14 SoundingrsUlConfigdedicatedaperiodicupptsextR13
}
