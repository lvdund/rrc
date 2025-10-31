package ies

import "rrc/utils"

// GuardbandTDD-NB-r15-eutra-Bandwitdh-r15 ::= ENUMERATED
type GuardbandtddNbR15EutraBandwitdhR15 struct {
	Value utils.ENUMERATED
}

const (
	GuardbandtddNbR15EutraBandwitdhR15EnumeratedNothing = iota
	GuardbandtddNbR15EutraBandwitdhR15EnumeratedBw5or10
	GuardbandtddNbR15EutraBandwitdhR15EnumeratedBw15or20
)
