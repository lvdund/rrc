package ies

// PhyLayerParameters-NB-v1530 ::= SEQUENCE
type PhylayerparametersNbV1530 struct {
	MixedoperationmodeR15         *PhylayerparametersNbV1530MixedoperationmodeR15
	SrWithharqAckR15              *PhylayerparametersNbV1530SrWithharqAckR15
	SrWithoutharqAckR15           *PhylayerparametersNbV1530SrWithoutharqAckR15
	NprachFormat2R15              *PhylayerparametersNbV1530NprachFormat2R15
	Additionaltransmissionsib1R15 *PhylayerparametersNbV1530Additionaltransmissionsib1R15
	Npusch3dot75khzScsTddR15      *PhylayerparametersNbV1530Npusch3dot75khzScsTddR15
}
