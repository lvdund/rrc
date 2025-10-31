package ies

import "rrc/utils"

// UplinkTxDirectCurrentBWP ::= SEQUENCE
type Uplinktxdirectcurrentbwp struct {
	BwpId                   BwpId
	Shift7dot5khz           utils.BOOLEAN
	Txdirectcurrentlocation utils.INTEGER `lb:0,ub:3301`
}
