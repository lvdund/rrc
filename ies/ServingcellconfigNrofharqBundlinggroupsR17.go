package ies

import "rrc/utils"

// ServingCellConfig-nrofHARQ-BundlingGroups-r17 ::= ENUMERATED
type ServingcellconfigNrofharqBundlinggroupsR17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigNrofharqBundlinggroupsR17EnumeratedNothing = iota
	ServingcellconfigNrofharqBundlinggroupsR17EnumeratedN1
	ServingcellconfigNrofharqBundlinggroupsR17EnumeratedN2
	ServingcellconfigNrofharqBundlinggroupsR17EnumeratedN4
)
