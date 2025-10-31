package ies

// NeedForGapNCSG-InfoEUTRA-r17 ::= SEQUENCE
type NeedforgapncsgInfoeutraR17 struct {
	NeedforncsgEutraR17 []NeedforncsgEutraR17 `lb:1,ub:maxBandsEUTRA`
}
