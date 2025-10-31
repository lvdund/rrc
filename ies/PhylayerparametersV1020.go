package ies

// PhyLayerParameters-v1020 ::= SEQUENCE
type PhylayerparametersV1020 struct {
	TwoantennaportsforpucchR10       *PhylayerparametersV1020TwoantennaportsforpucchR10
	Tm9With8txFddR10                 *PhylayerparametersV1020Tm9With8txFddR10
	PmiDisablingR10                  *PhylayerparametersV1020PmiDisablingR10
	CrosscarrierschedulingR10        *PhylayerparametersV1020CrosscarrierschedulingR10
	SimultaneouspucchPuschR10        *PhylayerparametersV1020SimultaneouspucchPuschR10
	MulticlusterpuschWithinccR10     *PhylayerparametersV1020MulticlusterpuschWithinccR10
	NoncontiguousulRaWithinccListR10 *NoncontiguousulRaWithinccListR10
}
