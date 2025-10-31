package ies

import "rrc/utils"

// NR-DL-PRS-PDC-ResourceSet-r17 ::= SEQUENCE
// Extensible
type NrDlPrsPdcResourcesetR17 struct {
	PeriodicityandoffsetR17   NrDlPrsPeriodicityAndResourcesetslotoffsetR17
	NumsymbolsR17             NrDlPrsPdcResourcesetR17NumsymbolsR17
	DlPrsResourcebandwidthR17 utils.INTEGER        `lb:0,ub:63`
	DlPrsStartprbR17          utils.INTEGER        `lb:0,ub:2176`
	ResourcelistR17           []NrDlPrsResourceR17 `lb:1,ub:maxNrofPRSResourcespersetR17`
	RepfactorandtimegapR17    *RepfactorandtimegapR17
}
