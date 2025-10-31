package ies

import "rrc/utils"

// MeasDS-Config-r12-setup-ds-OccasionDuration-r12 ::= CHOICE
const (
	MeasdsConfigR12SetupDsOccasiondurationR12ChoiceNothing = iota
	MeasdsConfigR12SetupDsOccasiondurationR12ChoiceDurationfddR12
	MeasdsConfigR12SetupDsOccasiondurationR12ChoiceDurationtddR12
)

type MeasdsConfigR12SetupDsOccasiondurationR12 struct {
	Choice         uint64
	DurationfddR12 *utils.INTEGER `lb:1,ub:maxDSDurationR12`
	DurationtddR12 *utils.INTEGER `lb:2,ub:maxDSDurationR12`
}
