package ies

// CSI-IM-Resource-csi-IM-ResourceElementPattern ::= CHOICE
const (
	CsiImResourceCsiImResourceelementpatternChoiceNothing = iota
	CsiImResourceCsiImResourceelementpatternChoicePattern0
	CsiImResourceCsiImResourceelementpatternChoicePattern1
)

type CsiImResourceCsiImResourceelementpattern struct {
	Choice   uint64
	Pattern0 *CsiImResourceCsiImResourceelementpatternPattern0
	Pattern1 *CsiImResourceCsiImResourceelementpatternPattern1
}
