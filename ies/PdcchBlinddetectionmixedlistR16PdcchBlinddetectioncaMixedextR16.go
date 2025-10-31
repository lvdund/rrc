package ies

// PDCCH-BlindDetectionMixedList-r16-pdcch-BlindDetectionCA-MixedExt-r16 ::= CHOICE
const (
	PdcchBlinddetectionmixedlistR16PdcchBlinddetectioncaMixedextR16ChoiceNothing = iota
	PdcchBlinddetectionmixedlistR16PdcchBlinddetectioncaMixedextR16ChoicePdcchBlinddetectioncaMixedV16a0
	PdcchBlinddetectionmixedlistR16PdcchBlinddetectioncaMixedextR16ChoicePdcchBlinddetectioncaMixedNonalignedspanV16a0
)

type PdcchBlinddetectionmixedlistR16PdcchBlinddetectioncaMixedextR16 struct {
	Choice                                        uint64
	PdcchBlinddetectioncaMixedV16a0               *PdcchBlinddetectioncaMixedextR16
	PdcchBlinddetectioncaMixedNonalignedspanV16a0 *PdcchBlinddetectioncaMixedextR16
}
