package ies

import "rrc/utils"

// UEPagingCoverageInformation-r13-IEs ::= SEQUENCE
type UepagingcoverageinformationR13 struct {
	MpdcchNumrepetitionR13 *utils.INTEGER `lb:0,ub:256`
	Noncriticalextension   *UepagingcoverageinformationR13IesNoncriticalextension
}
