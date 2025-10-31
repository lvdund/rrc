package ies

import "rrc/utils"

// MBMS-Parameters-r11-mbms-SCell-r11 ::= ENUMERATED
type MbmsParametersR11MbmsScellR11 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersR11MbmsScellR11EnumeratedNothing = iota
	MbmsParametersR11MbmsScellR11EnumeratedSupported
)
