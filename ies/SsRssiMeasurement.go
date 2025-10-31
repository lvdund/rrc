package ies

import "rrc/utils"

// SS-RSSI-Measurement ::= SEQUENCE
type SsRssiMeasurement struct {
	Measurementslots utils.BITSTRING `lb:1,ub:80`
	Endsymbol        utils.INTEGER   `lb:0,ub:3`
}
