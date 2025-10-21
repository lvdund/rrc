package ies

import "rrc/utils"

// SL-Parameters-v1430 ::= SEQUENCE
type SlParametersV1430 struct {
	ZonebasedpoolselectionR14          *utils.ENUMERATED
	UeAutonomouswithfullsensingR14     *utils.ENUMERATED
	UeAutonomouswithpartialsensingR14  *utils.ENUMERATED
	SlCongestioncontrolR14             *utils.ENUMERATED
	V2xTxwithshortresvintervalR14      *utils.ENUMERATED
	V2xNumbertxrxtimingR14             *utils.INTEGER
	V2xNonadjacentpscchPsschR14        *utils.ENUMERATED
	SlssTxrxR14                        *utils.ENUMERATED
	V2xSupportedbandcombinationlistR14 *V2xSupportedbandcombinationR14
}
