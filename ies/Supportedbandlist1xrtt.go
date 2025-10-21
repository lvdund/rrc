package ies

import "rrc/utils"

// SupportedBandList1XRTT ::= SEQUENCE OF BandclassCDMA2000
// SIZE (1..maxCDMA-BandClass)
type Supportedbandlist1xrtt struct {
	Value utils.Sequence[Bandclasscdma2000]
}
