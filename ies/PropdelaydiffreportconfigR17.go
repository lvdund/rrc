package ies

// PropDelayDiffReportConfig-r17 ::= SEQUENCE
type PropdelaydiffreportconfigR17 struct {
	ThreshpropdelaydiffR17 *PropdelaydiffreportconfigR17ThreshpropdelaydiffR17
	NeighcellinfolistR17   *[]NeighbourcellinfoR17 `lb:1,ub:maxCellNTNR17`
}
