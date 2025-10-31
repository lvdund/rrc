package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1620-supportFDM-SchemeB-r16 ::= ENUMERATED
type FeaturesetdownlinkperccV1620SupportfdmSchemebR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccV1620SupportfdmSchemebR16EnumeratedNothing = iota
	FeaturesetdownlinkperccV1620SupportfdmSchemebR16EnumeratedSupported
)
