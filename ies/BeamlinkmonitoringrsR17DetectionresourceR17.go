package ies

// BeamLinkMonitoringRS-r17-detectionResource-r17 ::= CHOICE
const (
	BeamlinkmonitoringrsR17DetectionresourceR17ChoiceNothing = iota
	BeamlinkmonitoringrsR17DetectionresourceR17ChoiceSsbIndex
	BeamlinkmonitoringrsR17DetectionresourceR17ChoiceCsiRsIndex
)

type BeamlinkmonitoringrsR17DetectionresourceR17 struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
}
