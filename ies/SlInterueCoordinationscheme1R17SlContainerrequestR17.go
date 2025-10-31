package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme1-r17-sl-ContainerRequest-r17 ::= ENUMERATED
type SlInterueCoordinationscheme1R17SlContainerrequestR17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme1R17SlContainerrequestR17EnumeratedNothing = iota
	SlInterueCoordinationscheme1R17SlContainerrequestR17EnumeratedEnabled
	SlInterueCoordinationscheme1R17SlContainerrequestR17EnumeratedDisabled
)
