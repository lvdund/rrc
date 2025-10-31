package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-v1130-cdma2000-NW-Sharing-r11 ::= ENUMERATED
type IratParameterscdma2000V1130Cdma2000NwSharingR11 struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma2000V1130Cdma2000NwSharingR11EnumeratedNothing = iota
	IratParameterscdma2000V1130Cdma2000NwSharingR11EnumeratedSupported
)
