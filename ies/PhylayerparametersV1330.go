package ies

import "rrc/utils"

// PhyLayerParameters-v1330 ::= SEQUENCE
type PhylayerparametersV1330 struct {
	CchInterfmitigationRefrectypeaR13 *PhylayerparametersV1330CchInterfmitigationRefrectypeaR13
	CchInterfmitigationRefrectypebR13 *PhylayerparametersV1330CchInterfmitigationRefrectypebR13
	CchInterfmitigationMaxnumccsR13   *utils.INTEGER `lb:0,ub:maxServCellR13`
	CrsInterfmitigationtm1totm9R13    *utils.INTEGER `lb:0,ub:maxServCellR13`
}
