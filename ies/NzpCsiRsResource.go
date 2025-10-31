package ies

import "rrc/utils"

// NZP-CSI-RS-Resource ::= SEQUENCE
// Extensible
type NzpCsiRsResource struct {
	NzpCsiRsResourceid   NzpCsiRsResourceid
	Resourcemapping      CsiRsResourcemapping
	Powercontroloffset   utils.INTEGER `lb:0,ub:15`
	Powercontroloffsetss *NzpCsiRsResourcePowercontroloffsetss
	Scramblingid         Scramblingid
	Periodicityandoffset *CsiResourceperiodicityandoffset
	QclInfoperiodiccsiRs *TciStateid
}
