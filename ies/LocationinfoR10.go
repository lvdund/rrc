package ies

import "rrc/utils"

// LocationInfo-r10 ::= SEQUENCE
// Extensible
type LocationinfoR10 struct {
	LocationcoordinatesR10 LocationinfoR10LocationcoordinatesR10
	HorizontalvelocityR10  *utils.OCTETSTRING
	GnssTodMsecR10         *utils.OCTETSTRING
}
