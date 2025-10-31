package ies

// CSI-ResourceConfig-csi-RS-ResourceSetList ::= CHOICE
const (
	CsiResourceconfigCsiRsResourcesetlistChoiceNothing = iota
	CsiResourceconfigCsiRsResourcesetlistChoiceNzpCsiRsSsb
	CsiResourceconfigCsiRsResourcesetlistChoiceCsiImResourcesetlist
)

type CsiResourceconfigCsiRsResourcesetlist struct {
	Choice               uint64
	NzpCsiRsSsb          *CsiResourceconfigCsiRsResourcesetlistNzpCsiRsSsb
	CsiImResourcesetlist *[]CsiImResourcesetid `lb:1,ub:maxNrofCSIImResourcesetsperconfig`
}
