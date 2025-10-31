package ies

// GuardbandTDD-NB-r15-sib-GuardbandInfo-r15 ::= CHOICE
const (
	GuardbandtddNbR15SibGuardbandinfoR15ChoiceNothing = iota
	GuardbandtddNbR15SibGuardbandinfoR15ChoiceSibGuardbandanchorR15
	GuardbandtddNbR15SibGuardbandinfoR15ChoiceSibGuardbandguardbandR15
	GuardbandtddNbR15SibGuardbandinfoR15ChoiceSibGuardbandinbandsamepciR15
	GuardbandtddNbR15SibGuardbandinfoR15ChoiceSibGuardbandinbanddiffpciR15
)

type GuardbandtddNbR15SibGuardbandinfoR15 struct {
	Choice                       uint64
	SibGuardbandanchorR15        *SibGuardbandanchortddNbR15
	SibGuardbandguardbandR15     *SibGuardbandguardbandtddNbR15
	SibGuardbandinbandsamepciR15 *SibGuardbandinbandsamepciTddNbR15
	SibGuardbandinbanddiffpciR15 *SibGuardbandinbanddiffpciTddNbR15
}
