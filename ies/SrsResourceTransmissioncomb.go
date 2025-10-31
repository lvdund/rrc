package ies

// SRS-Resource-transmissionComb ::= CHOICE
const (
	SrsResourceTransmissioncombChoiceNothing = iota
	SrsResourceTransmissioncombChoiceN2
	SrsResourceTransmissioncombChoiceN4
)

type SrsResourceTransmissioncomb struct {
	Choice uint64
	N2     *SrsResourceTransmissioncombN2
	N4     *SrsResourceTransmissioncombN4
}
