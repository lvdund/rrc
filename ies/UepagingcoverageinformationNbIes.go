package ies

import "rrc/utils"

// UEPagingCoverageInformation-NB-IEs ::= SEQUENCE
type UepagingcoverageinformationNbIes struct {
	NpdcchNumrepetitionpagingR13 *utils.INTEGER
	Noncriticalextension         *interface{}
}
