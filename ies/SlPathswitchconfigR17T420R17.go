package ies

import "rrc/utils"

// SL-PathSwitchConfig-r17-t420-r17 ::= ENUMERATED
type SlPathswitchconfigR17T420R17 struct {
	Value utils.ENUMERATED
}

const (
	SlPathswitchconfigR17T420R17EnumeratedNothing = iota
	SlPathswitchconfigR17T420R17EnumeratedMs50
	SlPathswitchconfigR17T420R17EnumeratedMs100
	SlPathswitchconfigR17T420R17EnumeratedMs150
	SlPathswitchconfigR17T420R17EnumeratedMs200
	SlPathswitchconfigR17T420R17EnumeratedMs500
	SlPathswitchconfigR17T420R17EnumeratedMs1000
	SlPathswitchconfigR17T420R17EnumeratedMs2000
	SlPathswitchconfigR17T420R17EnumeratedMs10000
)
