package ies

import "rrc/utils"

// SL-Parameters-v1430 ::= SEQUENCE
type SlParametersV1430 struct {
	ZonebasedpoolselectionR14          *SlParametersV1430ZonebasedpoolselectionR14
	UeAutonomouswithfullsensingR14     *SlParametersV1430UeAutonomouswithfullsensingR14
	UeAutonomouswithpartialsensingR14  *SlParametersV1430UeAutonomouswithpartialsensingR14
	SlCongestioncontrolR14             *SlParametersV1430SlCongestioncontrolR14
	V2xTxwithshortresvintervalR14      *SlParametersV1430V2xTxwithshortresvintervalR14
	V2xNumbertxrxtimingR14             *utils.INTEGER `lb:0,ub:16`
	V2xNonadjacentpscchPsschR14        *SlParametersV1430V2xNonadjacentpscchPsschR14
	SlssTxrxR14                        *SlParametersV1430SlssTxrxR14
	V2xSupportedbandcombinationlistR14 *V2xSupportedbandcombinationR14
}
