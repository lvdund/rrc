package ies

// DL-PRS-QCL-Info-r17 ::= CHOICE
// Extensible
const (
	DlPrsQclInfoR17ChoiceNothing = iota
	DlPrsQclInfoR17ChoiceSsbR17
	DlPrsQclInfoR17ChoiceDlPrsR17
)

type DlPrsQclInfoR17 struct {
	Choice   uint64
	SsbR17   *DlPrsQclInfoR17SsbR17
	DlPrsR17 *DlPrsQclInfoR17DlPrsR17
}
