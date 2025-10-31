package ies

// CSI-RS-ConfigEMIMO-r13-setup ::= CHOICE
const (
	CsiRsConfigemimoR13SetupChoiceNothing = iota
	CsiRsConfigemimoR13SetupChoiceNonprecodedR13
	CsiRsConfigemimoR13SetupChoiceBeamformedR13
)

type CsiRsConfigemimoR13Setup struct {
	Choice         uint64
	NonprecodedR13 *CsiRsConfignonprecodedR13
	BeamformedR13  *CsiRsConfigbeamformedR13
}
