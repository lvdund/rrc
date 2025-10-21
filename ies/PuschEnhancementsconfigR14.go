package ies

import "rrc/utils"

// PUSCH-EnhancementsConfig-r14 ::= CHOICE
type PuschEnhancementsconfigR14 interface {
	isPuschEnhancementsconfigR14()
}

type PuschEnhancementsconfigR14Release struct {
	Value struct{}
}

func (*PuschEnhancementsconfigR14Release) isPuschEnhancementsconfigR14() {}

type PuschEnhancementsconfigR14Setup struct {
	Value interface{}
}

func (*PuschEnhancementsconfigR14Setup) isPuschEnhancementsconfigR14() {}
