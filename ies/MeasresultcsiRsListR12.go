package ies

import "rrc/utils"

// MeasResultCSI-RS-List-r12 ::= SEQUENCE OF MeasResultCSI-RS-r12
// SIZE (1..maxCellReport)
type MeasresultcsiRsListR12 struct {
	Value utils.Sequence[MeasresultcsiRsR12]
}
