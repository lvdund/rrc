package ies

import "rrc/utils"

// Other-Parameters-r11-powerPrefInd-r11 ::= ENUMERATED
type OtherParametersR11PowerprefindR11 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersR11PowerprefindR11EnumeratedNothing = iota
	OtherParametersR11PowerprefindR11EnumeratedSupported
)
