package ies

import "rrc/utils"

// UEPagingCoverageInformation-NB-IEs ::= SEQUENCE
type UepagingcoverageinformationNb struct {
	NpdcchNumrepetitionpagingR13 *utils.INTEGER `lb:0,ub:2048`
	Noncriticalextension         *UepagingcoverageinformationNbIesNoncriticalextension
}
