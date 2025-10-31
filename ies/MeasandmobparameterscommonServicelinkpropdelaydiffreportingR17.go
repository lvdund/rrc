package ies

import "rrc/utils"

// MeasAndMobParametersCommon-serviceLinkPropDelayDiffReporting-r17 ::= ENUMERATED
type MeasandmobparameterscommonServicelinkpropdelaydiffreportingR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonServicelinkpropdelaydiffreportingR17EnumeratedNothing = iota
	MeasandmobparameterscommonServicelinkpropdelaydiffreportingR17EnumeratedSupported
)
