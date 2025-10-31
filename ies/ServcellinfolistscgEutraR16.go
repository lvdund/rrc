package ies

// ServCellInfoListSCG-EUTRA-r16 ::= SEQUENCE OF ServCellInfoXCG-EUTRA-r16
// SIZE (1.. maxNrofServingCellsEUTRA)
type ServcellinfolistscgEutraR16 struct {
	Value []ServcellinfoxcgEutraR16 `lb:1,ub:maxNrofServingCellsEUTRA`
}
