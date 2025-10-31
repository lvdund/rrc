package ies

// PowSav-ParametersCommon-r16 ::= SEQUENCE
// Extensible
type PowsavParameterscommonR16 struct {
	DrxPreferenceR16                 *PowsavParameterscommonR16DrxPreferenceR16
	MaxccPreferenceR16               *PowsavParameterscommonR16MaxccPreferenceR16
	ReleasepreferenceR16             *PowsavParameterscommonR16ReleasepreferenceR16
	MinschedulingoffsetpreferenceR16 *PowsavParameterscommonR16MinschedulingoffsetpreferenceR16
}
