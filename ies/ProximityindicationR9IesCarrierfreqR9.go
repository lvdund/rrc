package ies

// ProximityIndication-r9-IEs-carrierFreq-r9 ::= CHOICE
// Extensible
const (
	ProximityindicationR9IesCarrierfreqR9ChoiceNothing = iota
	ProximityindicationR9IesCarrierfreqR9ChoiceEutraR9
	ProximityindicationR9IesCarrierfreqR9ChoiceUtraR9
	ProximityindicationR9IesCarrierfreqR9ChoiceEutra2V9e0
)

type ProximityindicationR9IesCarrierfreqR9 struct {
	Choice     uint64
	EutraR9    *ArfcnValueeutra
	UtraR9     *ArfcnValueutra
	Eutra2V9e0 *ArfcnValueeutraV9e0
}
