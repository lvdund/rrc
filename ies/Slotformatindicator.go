package ies

import "rrc/utils"

// SlotFormatIndicator ::= SEQUENCE
// Extensible
type Slotformatindicator struct {
	SfiRnti                              RntiValue
	DciPayloadsize                       utils.INTEGER                    `lb:0,ub:maxSFIDciPayloadsize`
	Slotformatcombtoaddmodlist           *[]Slotformatcombinationspercell `lb:1,ub:maxNrofAggregatedCellsPerCellGroup`
	Slotformatcombtoreleaselist          *[]Servcellindex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup`
	AvailablerbSetstoaddmodlistR16       *[]AvailablerbSetspercellR16     `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,ext`
	AvailablerbSetstoreleaselistR16      *[]Servcellindex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,ext`
	SwitchtriggertoaddmodlistR16         *[]SearchspaceswitchtriggerR16   `lb:1,ub:4,ext`
	SwitchtriggertoreleaselistR16        *[]Servcellindex                 `lb:1,ub:4,ext`
	CoDurationspercelltoaddmodlistR16    *[]CoDurationspercellR16         `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,ext`
	CoDurationspercelltoreleaselistR16   *[]Servcellindex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,ext`
	SwitchtriggertoaddmodlistsizeextR16  *[]SearchspaceswitchtriggerR16   `lb:1,ub:maxNrofAggregatedCellsPerCellGroupMinus4R16,ext`
	SwitchtriggertoreleaselistsizeextR16 *[]Servcellindex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroupMinus4R16,ext`
	CoDurationspercelltoaddmodlistR17    *[]CoDurationspercellR17         `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,ext`
}
