package ies

// Enable256QAM-r14-setup ::= CHOICE
const (
	Enable256qamR14SetupChoiceNothing = iota
	Enable256qamR14SetupChoiceTpcSubframesetConfiguredR14
	Enable256qamR14SetupChoiceTpcSubframesetNotconfiguredR14
)

type Enable256qamR14Setup struct {
	Choice                         uint64
	TpcSubframesetConfiguredR14    *Enable256qamR14SetupTpcSubframesetConfiguredR14
	TpcSubframesetNotconfiguredR14 *Enable256qamR14SetupTpcSubframesetNotconfiguredR14
}
