package ies

// ServCellInfoXCG-EUTRA-r16 ::= SEQUENCE
// Extensible
type ServcellinfoxcgEutraR16 struct {
	DlCarrierfreqEutraR16         *ArfcnValueeutra
	UlCarrierfreqEutraR16         *ArfcnValueeutra
	TransmissionbandwidthEutraR16 *TransmissionbandwidthEutraR16
}
