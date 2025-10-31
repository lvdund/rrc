package ies

import "rrc/utils"

// FreqBandInformationNR ::= SEQUENCE
type Freqbandinformationnr struct {
	Bandnr                  Freqbandindicatornr
	Maxbandwidthrequesteddl *Aggregatedbandwidth
	Maxbandwidthrequestedul *Aggregatedbandwidth
	Maxcarriersrequesteddl  *utils.INTEGER `lb:0,ub:maxNrofServingCells`
	Maxcarriersrequestedul  *utils.INTEGER `lb:0,ub:maxNrofServingCells`
}
