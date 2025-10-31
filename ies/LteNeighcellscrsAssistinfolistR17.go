package ies

// LTE-NeighCellsCRS-AssistInfoList-r17 ::= SEQUENCE OF LTE-NeighCellsCRS-AssistInfo-r17
// SIZE (1..maxNrofCRS-IM-InterfCell-r17)
type LteNeighcellscrsAssistinfolistR17 struct {
	Value []LteNeighcellscrsAssistinfoR17 `lb:1,ub:maxNrofCRSImInterfcellR17`
}
