package ies

// MobilityParameters-v1610 ::= SEQUENCE
type MobilityparametersV1610 struct {
	ChoR16                 *MobilityparametersV1610ChoR16
	ChoFddTddR16           *MobilityparametersV1610ChoFddTddR16
	ChoFailureR16          *MobilityparametersV1610ChoFailureR16
	ChoTwotriggereventsR16 *MobilityparametersV1610ChoTwotriggereventsR16
}
