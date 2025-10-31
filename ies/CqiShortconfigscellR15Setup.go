package ies

import "rrc/utils"

// CQI-ShortConfigSCell-r15-setup ::= SEQUENCE
type CqiShortconfigscellR15Setup struct {
	CqiPmiConfigindexshortR15  utils.INTEGER  `lb:0,ub:1023`
	RiConfigindexshortR15      *utils.INTEGER `lb:0,ub:1023`
	CqiFormatindicatorshortR15 *CqiShortconfigscellR15SetupCqiFormatindicatorshortR15
}
