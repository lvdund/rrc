package ies

import "rrc/utils"

// FreqPriorityListUTRA-FDD ::= SEQUENCE OF FreqPriorityUTRA-FDD
// SIZE (1..maxUTRA-FDD-Carrier)
type FreqprioritylistutraFdd struct {
	Value []FreqpriorityutraFdd
}
