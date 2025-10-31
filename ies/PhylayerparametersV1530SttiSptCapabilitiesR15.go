package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15 ::= SEQUENCE
type PhylayerparametersV1530SttiSptCapabilitiesR15 struct {
	AperiodiccsiReportingsttiR15             *PhylayerparametersV1530SttiSptCapabilitiesR15AperiodiccsiReportingsttiR15
	DmrsBasedspdcchMbsfnR15                  *PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchMbsfnR15
	DmrsBasedspdcchNonmbsfnR15               *PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchNonmbsfnR15
	DmrsPositionpatternR15                   *PhylayerparametersV1530SttiSptCapabilitiesR15DmrsPositionpatternR15
	DmrsSharingsubslotpdschR15               *PhylayerparametersV1530SttiSptCapabilitiesR15DmrsSharingsubslotpdschR15
	DmrsRepetitionsubslotpdschR15            *PhylayerparametersV1530SttiSptCapabilitiesR15DmrsRepetitionsubslotpdschR15
	EpdcchSptDifferentcellsR15               *PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSptDifferentcellsR15
	EpdcchSttiDifferentcellsR15              *PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSttiDifferentcellsR15
	MaxlayersslotorsubslotpuschR15           *PhylayerparametersV1530SttiSptCapabilitiesR15MaxlayersslotorsubslotpuschR15
	MaxnumberupdatedcsiProcSptR15            *utils.INTEGER `lb:0,ub:32`
	MaxnumberupdatedcsiProcSttiComb77R15     *utils.INTEGER `lb:0,ub:32`
	MaxnumberupdatedcsiProcSttiComb27R15     *utils.INTEGER `lb:0,ub:32`
	MaxnumberupdatedcsiProcSttiComb22Set1R15 *utils.INTEGER `lb:0,ub:32`
	MaxnumberupdatedcsiProcSttiComb22Set2R15 *utils.INTEGER `lb:0,ub:32`
	MimoUeParameterssttiR15                  *MimoUeParametersR13
	MimoUeParameterssttiV1530                *MimoUeParametersV1430
	NumberofblinddecodesussR15               *utils.INTEGER `lb:0,ub:32`
	PdschSlotsubslotpdschDecodingR15         *PhylayerparametersV1530SttiSptCapabilitiesR15PdschSlotsubslotpdschDecodingR15
	PoweruciSlotpusch                        *PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSlotpusch
	PoweruciSubslotpusch                     *PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSubslotpusch
	SlotpdschTxdivTm9and10                   *PhylayerparametersV1530SttiSptCapabilitiesR15SlotpdschTxdivTm9and10
	SubslotpdschTxdivTm9and10                *PhylayerparametersV1530SttiSptCapabilitiesR15SubslotpdschTxdivTm9and10
	SpdcchDifferentrsTypesR15                *PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchDifferentrsTypesR15
	SrsDci7Triggeringfs2R15                  *PhylayerparametersV1530SttiSptCapabilitiesR15SrsDci7Triggeringfs2R15
	SpsCyclicshiftR15                        *PhylayerparametersV1530SttiSptCapabilitiesR15SpsCyclicshiftR15
	SpdcchReuseR15                           *PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchReuseR15
	SpsSttiR15                               *PhylayerparametersV1530SttiSptCapabilitiesR15SpsSttiR15
	Tm8SlotpdschR15                          *PhylayerparametersV1530SttiSptCapabilitiesR15Tm8SlotpdschR15
	Tm9SlotsubslotR15                        *PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotR15
	Tm9SlotsubslotmbsfnR15                   *PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotmbsfnR15
	Tm10SlotsubslotR15                       *PhylayerparametersV1530SttiSptCapabilitiesR15Tm10SlotsubslotR15
	Tm10SlotsubslotmbsfnR15                  *PhylayerparametersV1530SttiSptCapabilitiesR15Tm10SlotsubslotmbsfnR15
	TxdivSpucchR15                           *PhylayerparametersV1530SttiSptCapabilitiesR15TxdivSpucchR15
	UlAsyncharqsharingdiffTtiLengthsR15      *PhylayerparametersV1530SttiSptCapabilitiesR15UlAsyncharqsharingdiffTtiLengthsR15
}
