package ies

import "rrc/utils"

// FeatureSetDownlink-ue-SpecificUL-DL-Assignment ::= ENUMERATED
type FeaturesetdownlinkUeSpecificulDlAssignment struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkUeSpecificulDlAssignmentEnumeratedNothing = iota
	FeaturesetdownlinkUeSpecificulDlAssignmentEnumeratedSupported
)
