package ies

// OverheatingAssistance ::= SEQUENCE
type Overheatingassistance struct {
	Reducedmaxccs           *ReducedmaxccsR16
	ReducedmaxbwFr1         *ReducedmaxbwFrxR16
	ReducedmaxbwFr2         *ReducedmaxbwFrxR16
	ReducedmaxmimoLayersfr1 *OverheatingassistanceReducedmaxmimoLayersfr1
	ReducedmaxmimoLayersfr2 *OverheatingassistanceReducedmaxmimoLayersfr2
}
