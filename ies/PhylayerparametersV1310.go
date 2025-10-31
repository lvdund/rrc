package ies

import "rrc/utils"

// PhyLayerParameters-v1310 ::= SEQUENCE
type PhylayerparametersV1310 struct {
	AperiodiccsiReportingR13     *utils.BITSTRING `lb:2,ub:2`
	CodebookHarqAckR13           *utils.BITSTRING `lb:2,ub:2`
	CrosscarrierschedulingB5cR13 *PhylayerparametersV1310CrosscarrierschedulingB5cR13
	FddHarqTimingtddR13          *PhylayerparametersV1310FddHarqTimingtddR13
	MaxnumberupdatedcsiProcR13   *utils.INTEGER `lb:0,ub:32`
	PucchFormat4R13              *PhylayerparametersV1310PucchFormat4R13
	PucchFormat5R13              *PhylayerparametersV1310PucchFormat5R13
	PucchScellR13                *PhylayerparametersV1310PucchScellR13
	SpatialbundlingHarqAckR13    *PhylayerparametersV1310SpatialbundlingHarqAckR13
	SupportedblinddecodingR13    *PhylayerparametersV1310SupportedblinddecodingR13
	UciPuschExtR13               *PhylayerparametersV1310UciPuschExtR13
	CrsInterfmitigationtm10R13   *PhylayerparametersV1310CrsInterfmitigationtm10R13
	PdschCollisionhandlingR13    *PhylayerparametersV1310PdschCollisionhandlingR13
}
