package ies

import "rrc/utils"

// BandNR-maxNumber-NGSO-SatellitesWithinOneSMTC-r17 ::= ENUMERATED
type BandnrMaxnumberNgsoSatelliteswithinonesmtcR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMaxnumberNgsoSatelliteswithinonesmtcR17EnumeratedNothing = iota
	BandnrMaxnumberNgsoSatelliteswithinonesmtcR17EnumeratedN1
	BandnrMaxnumberNgsoSatelliteswithinonesmtcR17EnumeratedN2
	BandnrMaxnumberNgsoSatelliteswithinonesmtcR17EnumeratedN3
	BandnrMaxnumberNgsoSatelliteswithinonesmtcR17EnumeratedN4
)
