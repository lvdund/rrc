package ies

import "rrc/utils"

// UEPagingCoverageInformation-r13-IEs ::= SEQUENCE
type UepagingcoverageinformationR13Ies struct {
	MpdcchNumrepetitionR13 *utils.INTEGER
	Noncriticalextension   *interface{}
}
