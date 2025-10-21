package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1530 ::= SEQUENCE
type PhylayerparametersNbV1530 struct {
	MixedoperationmodeR15         *utils.ENUMERATED
	SrWithharqAckR15              *utils.ENUMERATED
	SrWithoutharqAckR15           *utils.ENUMERATED
	NprachFormat2R15              *utils.ENUMERATED
	Additionaltransmissionsib1R15 *utils.ENUMERATED
	Npusch3dot75khzScsTddR15      *utils.ENUMERATED
}
