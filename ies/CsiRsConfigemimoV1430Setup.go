package ies

// CSI-RS-ConfigEMIMO-v1430-setup ::= CHOICE
const (
	CsiRsConfigemimoV1430SetupChoiceNothing = iota
	CsiRsConfigemimoV1430SetupChoiceNonprecodedV1430
	CsiRsConfigemimoV1430SetupChoiceBeamformedV1430
)

type CsiRsConfigemimoV1430Setup struct {
	Choice           uint64
	NonprecodedV1430 *CsiRsConfignonprecodedV1430
	BeamformedV1430  *CsiRsConfigbeamformedV1430
}
