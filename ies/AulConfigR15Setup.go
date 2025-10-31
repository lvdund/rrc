package ies

import "rrc/utils"

// AUL-Config-r15-setup ::= SEQUENCE
type AulConfigR15Setup struct {
	AulCrntiR15                        CRnti
	AulSubframesR15                    utils.BITSTRING `lb:40,ub:40`
	AulHarqProcessesR15                utils.INTEGER   `lb:0,ub:16`
	TransmissionmodeulAulR15           AulConfigR15SetupTransmissionmodeulAulR15
	AulStartingfullbwInsidemcotR15     utils.BITSTRING `lb:5,ub:5`
	AulStartingfullbwOutsidemcotR15    utils.BITSTRING `lb:7,ub:7`
	AulStartingpartialbwInsidemcotR15  AulConfigR15SetupAulStartingpartialbwInsidemcotR15
	AulStartingpartialbwOutsidemcotR15 AulConfigR15SetupAulStartingpartialbwOutsidemcotR15
	AulRetransmissiontimerR15          AulConfigR15SetupAulRetransmissiontimerR15
	EndingsymbolaulR15                 utils.INTEGER `lb:0,ub:13`
	SubframeoffsetcotSharingR15        utils.INTEGER `lb:0,ub:4`
	ContentionwindowsizetimerR15       AulConfigR15SetupContentionwindowsizetimerR15
}
