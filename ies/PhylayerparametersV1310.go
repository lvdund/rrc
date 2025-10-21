package ies

import "rrc/utils"

// PhyLayerParameters-v1310 ::= SEQUENCE
type PhylayerparametersV1310 struct {
	AperiodiccsiReportingR13     *utils.BITSTRING
	CodebookHarqAckR13           *utils.BITSTRING
	CrosscarrierschedulingB5cR13 *utils.ENUMERATED
	FddHarqTimingtddR13          *utils.ENUMERATED
	MaxnumberupdatedcsiProcR13   *utils.INTEGER
	PucchFormat4R13              *utils.ENUMERATED
	PucchFormat5R13              *utils.ENUMERATED
	PucchScellR13                *utils.ENUMERATED
	SpatialbundlingHarqAckR13    *utils.ENUMERATED
	SupportedblinddecodingR13    *interface{}
	UciPuschExtR13               *utils.ENUMERATED
	CrsInterfmitigationtm10R13   *utils.ENUMERATED
	PdschCollisionhandlingR13    *utils.ENUMERATED
}
