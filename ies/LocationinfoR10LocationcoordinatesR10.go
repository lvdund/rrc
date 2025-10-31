package ies

import "rrc/utils"

// LocationInfo-r10-locationCoordinates-r10 ::= CHOICE
// Extensible
const (
	LocationinfoR10LocationcoordinatesR10ChoiceNothing = iota
	LocationinfoR10LocationcoordinatesR10ChoiceEllipsoidPointR10
	LocationinfoR10LocationcoordinatesR10ChoiceEllipsoidpointwithaltitudeR10
	LocationinfoR10LocationcoordinatesR10ChoiceEllipsoidpointwithuncertaintycircleR11
	LocationinfoR10LocationcoordinatesR10ChoiceEllipsoidpointwithuncertaintyellipseR11
	LocationinfoR10LocationcoordinatesR10ChoiceEllipsoidpointwithaltitudeanduncertaintyellipsoidR11
	LocationinfoR10LocationcoordinatesR10ChoiceEllipsoidarcR11
	LocationinfoR10LocationcoordinatesR10ChoicePolygonR11
)

type LocationinfoR10LocationcoordinatesR10 struct {
	Choice                                               uint64
	EllipsoidPointR10                                    *utils.OCTETSTRING
	EllipsoidpointwithaltitudeR10                        *utils.OCTETSTRING
	EllipsoidpointwithuncertaintycircleR11               *utils.OCTETSTRING
	EllipsoidpointwithuncertaintyellipseR11              *utils.OCTETSTRING
	EllipsoidpointwithaltitudeanduncertaintyellipsoidR11 *utils.OCTETSTRING
	EllipsoidarcR11                                      *utils.OCTETSTRING
	PolygonR11                                           *utils.OCTETSTRING
}
