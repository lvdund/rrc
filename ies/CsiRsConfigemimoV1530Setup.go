package ies

// CSI-RS-ConfigEMIMO-v1530-setup ::= CHOICE
const (
	CsiRsConfigemimoV1530SetupChoiceNothing = iota
	CsiRsConfigemimoV1530SetupChoiceNonprecodedV1530
)

type CsiRsConfigemimoV1530Setup struct {
	Choice           uint64
	NonprecodedV1530 *CsiRsConfignonprecodedV1530
}
