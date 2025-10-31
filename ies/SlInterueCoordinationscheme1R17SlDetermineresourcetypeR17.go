package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme1-r17-sl-DetermineResourceType-r17 ::= ENUMERATED
type SlInterueCoordinationscheme1R17SlDetermineresourcetypeR17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme1R17SlDetermineresourcetypeR17EnumeratedNothing = iota
	SlInterueCoordinationscheme1R17SlDetermineresourcetypeR17EnumeratedUea
	SlInterueCoordinationscheme1R17SlDetermineresourcetypeR17EnumeratedUeb
)
