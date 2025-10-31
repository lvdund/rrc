package ies

import "rrc/utils"

// SystemInformationBlockType1-cellAccessRelatedInfo ::= SEQUENCE
type Systeminformationblocktype1Cellaccessrelatedinfo struct {
	PlmnIdentitylist     PlmnIdentitylist
	Trackingareacode     Trackingareacode
	Cellidentity         Cellidentity
	Cellbarred           Systeminformationblocktype1CellaccessrelatedinfoCellbarred
	Intrafreqreselection Systeminformationblocktype1CellaccessrelatedinfoIntrafreqreselection
	CsgIndication        utils.BOOLEAN
	CsgIdentity          *CsgIdentity
}
