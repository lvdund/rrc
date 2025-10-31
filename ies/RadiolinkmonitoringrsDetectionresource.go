package ies

// RadioLinkMonitoringRS-detectionResource ::= CHOICE
const (
	RadiolinkmonitoringrsDetectionresourceChoiceNothing = iota
	RadiolinkmonitoringrsDetectionresourceChoiceSsbIndex
	RadiolinkmonitoringrsDetectionresourceChoiceCsiRsIndex
)

type RadiolinkmonitoringrsDetectionresource struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
}
