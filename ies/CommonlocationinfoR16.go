package ies

import "rrc/utils"

// CommonLocationInfo-r16 ::= SEQUENCE
type CommonlocationinfoR16 struct {
	GnssTodMsecR16        *utils.OCTETSTRING
	LocationtimestampR16  *utils.OCTETSTRING
	LocationcoordinateR16 *utils.OCTETSTRING
	LocationerrorR16      *utils.OCTETSTRING
	LocationsourceR16     *utils.OCTETSTRING
	VelocityestimateR16   *utils.OCTETSTRING
}
