package ies

import "rrc/utils"

// MBMS-Parameters-r11-mbms-NonServingCell-r11 ::= ENUMERATED
type MbmsParametersR11MbmsNonservingcellR11 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersR11MbmsNonservingcellR11EnumeratedNothing = iota
	MbmsParametersR11MbmsNonservingcellR11EnumeratedSupported
)
