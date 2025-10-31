package ies

// GuardbandTDD-NB-r15 ::= SEQUENCE
type GuardbandtddNbR15 struct {
	RasteroffsetR15     ChannelrasteroffsetNbR13
	SibGuardbandinfoR15 GuardbandtddNbR15SibGuardbandinfoR15
	EutraBandwitdhR15   GuardbandtddNbR15EutraBandwitdhR15
}
