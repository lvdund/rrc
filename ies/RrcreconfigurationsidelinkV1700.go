package ies

import "rrc/utils"

// RRCReconfigurationSidelink-v1700-IEs ::= SEQUENCE
type RrcreconfigurationsidelinkV1700 struct {
	SlDrxConfigucPc5R17             *utils.Setuprelease[SlDrxConfigucR17]
	SlLatencyboundiucReportR17      *utils.Setuprelease[SlLatencyboundiucReportR17]
	SlRlcChanneltoreleaselistpc5R17 *[]SlRlcChannelidR17        `lb:1,ub:maxSLLcidR16`
	SlRlcChanneltoaddmodlistpc5R17  *[]SlRlcChannelconfigpc5R17 `lb:1,ub:maxSLLcidR16`
	Noncriticalextension            *RrcreconfigurationsidelinkV1700IesNoncriticalextension
}
