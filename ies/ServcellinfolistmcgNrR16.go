package ies

// ServCellInfoListMCG-NR-r16 ::= SEQUENCE OF ServCellInfoXCG-NR-r16
// SIZE (1.. maxNrofServingCells)
type ServcellinfolistmcgNrR16 struct {
	Value []ServcellinfoxcgNrR16 `lb:1,ub:maxNrofServingCells`
}
