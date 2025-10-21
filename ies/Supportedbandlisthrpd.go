package ies

import "rrc/utils"

// SupportedBandListHRPD ::= SEQUENCE OF BandclassCDMA2000
// SIZE (1..maxCDMA-BandClass)
type Supportedbandlisthrpd struct {
	Value []Bandclasscdma2000
}
