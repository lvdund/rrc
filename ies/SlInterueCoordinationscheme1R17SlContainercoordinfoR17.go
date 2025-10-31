package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme1-r17-sl-ContainerCoordInfo-r17 ::= ENUMERATED
type SlInterueCoordinationscheme1R17SlContainercoordinfoR17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme1R17SlContainercoordinfoR17EnumeratedNothing = iota
	SlInterueCoordinationscheme1R17SlContainercoordinfoR17EnumeratedEnabled
	SlInterueCoordinationscheme1R17SlContainercoordinfoR17EnumeratedDisabled
)
