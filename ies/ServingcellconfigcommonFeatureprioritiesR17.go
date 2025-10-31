package ies

// ServingCellConfigCommon-featurePriorities-r17 ::= SEQUENCE
type ServingcellconfigcommonFeatureprioritiesR17 struct {
	RedcappriorityR17          *FeaturepriorityR17
	SlicingpriorityR17         *FeaturepriorityR17
	Msg3RepetitionsPriorityR17 *FeaturepriorityR17
	SdtPriorityR17             *FeaturepriorityR17
}
