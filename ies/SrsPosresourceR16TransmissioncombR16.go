package ies

// SRS-PosResource-r16-transmissionComb-r16 ::= CHOICE
// Extensible
const (
	SrsPosresourceR16TransmissioncombR16ChoiceNothing = iota
	SrsPosresourceR16TransmissioncombR16ChoiceN2R16
	SrsPosresourceR16TransmissioncombR16ChoiceN4R16
	SrsPosresourceR16TransmissioncombR16ChoiceN8R16
)

type SrsPosresourceR16TransmissioncombR16 struct {
	Choice uint64
	N2R16  *SrsPosresourceR16TransmissioncombR16N2R16
	N4R16  *SrsPosresourceR16TransmissioncombR16N4R16
	N8R16  *SrsPosresourceR16TransmissioncombR16N8R16
}
