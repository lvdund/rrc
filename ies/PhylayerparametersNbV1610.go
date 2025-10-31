package ies

// PhyLayerParameters-NB-v1610 ::= SEQUENCE
type PhylayerparametersNbV1610 struct {
	NpdschMultitbR16             *PhylayerparametersNbV1610NpdschMultitbR16
	NpdschMultitbInterleavingR16 *PhylayerparametersNbV1610NpdschMultitbInterleavingR16
	NpuschMultitbR16             *PhylayerparametersNbV1610NpuschMultitbR16
	NpuschMultitbInterleavingR16 *PhylayerparametersNbV1610NpuschMultitbInterleavingR16
	MultitbHarqAckbundlingR16    *PhylayerparametersNbV1610MultitbHarqAckbundlingR16
	SlotsymbolresourceresvdlR16  *PhylayerparametersNbV1610SlotsymbolresourceresvdlR16
	SlotsymbolresourceresvulR16  *PhylayerparametersNbV1610SlotsymbolresourceresvulR16
	SubframeresourceresvdlR16    *PhylayerparametersNbV1610SubframeresourceresvdlR16
	SubframeresourceresvulR16    *PhylayerparametersNbV1610SubframeresourceresvulR16
}
