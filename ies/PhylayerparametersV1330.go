package ies

import "rrc/utils"

// PhyLayerParameters-v1330 ::= SEQUENCE
type PhylayerparametersV1330 struct {
	CchInterfmitigationRefrectypeaR13 *utils.ENUMERATED
	CchInterfmitigationRefrectypebR13 *utils.ENUMERATED
	CchInterfmitigationMaxnumccsR13   *utils.INTEGER
	CrsInterfmitigationtm1totm9R13    *utils.INTEGER
}
