package ies

// SRS-PosResourceSet-r16-pathlossReferenceRS-Pos-r16 ::= CHOICE
const (
	SrsPosresourcesetR16PathlossreferencersPosR16ChoiceNothing = iota
	SrsPosresourcesetR16PathlossreferencersPosR16ChoiceSsbIndexservingR16
	SrsPosresourcesetR16PathlossreferencersPosR16ChoiceSsbNcellR16
	SrsPosresourcesetR16PathlossreferencersPosR16ChoiceDlPrsR16
)

type SrsPosresourcesetR16PathlossreferencersPosR16 struct {
	Choice             uint64
	SsbIndexservingR16 *SsbIndex
	SsbNcellR16        *SsbInfoncellR16
	DlPrsR16           *DlPrsInfoR16
}
