package ies

// ServCellInfoListSCG-NR-r16 ::= SEQUENCE OF ServCellInfoXCG-NR-r16
// SIZE (1.. maxNrofServingCells)
type ServcellinfolistscgNrR16 struct {
	Value []ServcellinfoxcgNrR16 `lb:1,ub:maxNrofServingCells`
}
