package ies

import "rrc/utils"

// MeasAndMobParametersCommon-ncsg-SymbolLevelScheduleRestrictionInter-r17 ::= ENUMERATED
type MeasandmobparameterscommonNcsgSymbollevelschedulerestrictioninterR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNcsgSymbollevelschedulerestrictioninterR17EnumeratedNothing = iota
	MeasandmobparameterscommonNcsgSymbollevelschedulerestrictioninterR17EnumeratedSupported
)
