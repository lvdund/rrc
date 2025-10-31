package ies

// CSI-RS-ConfigEMIMO-v1480-setup ::= CHOICE
const (
	CsiRsConfigemimoV1480SetupChoiceNothing = iota
	CsiRsConfigemimoV1480SetupChoiceNonprecodedV1480
	CsiRsConfigemimoV1480SetupChoiceBeamformedV1480
)

type CsiRsConfigemimoV1480Setup struct {
	Choice           uint64
	NonprecodedV1480 *CsiRsConfignonprecodedV1480
	BeamformedV1480  *CsiRsConfigbeamformedV1430
}
