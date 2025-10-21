package ies

import "rrc/utils"

// FreqPriorityListUTRA-TDD ::= SEQUENCE OF FreqPriorityUTRA-TDD
// SIZE (1..maxUTRA-TDD-Carrier)
type FreqprioritylistutraTdd struct {
	Value utils.Sequence[FreqpriorityutraTdd]
}
