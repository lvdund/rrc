package ies

import "rrc/utils"

// UEAssistanceInformationSidelink-r17-IEs ::= SEQUENCE
type UeassistanceinformationsidelinkR17 struct {
	SlPreferreddrxConfiglistR17 *[]SlDrxConfigucSemistaticR17 `lb:1,ub:maxNrofSLRxinfosetR17`
	Latenoncriticalextension    *utils.OCTETSTRING
	Noncriticalextension        *UeassistanceinformationsidelinkR17IesNoncriticalextension
}
