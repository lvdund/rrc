package ies

import "rrc/utils"

// SupportedCSI-RS-Resource ::= SEQUENCE
type SupportedcsiRsResource struct {
	Maxnumbertxportsperresource SupportedcsiRsResourceMaxnumbertxportsperresource
	Maxnumberresourcesperband   utils.INTEGER `lb:0,ub:64`
	Totalnumbertxportsperband   utils.INTEGER `lb:0,ub:256`
}
