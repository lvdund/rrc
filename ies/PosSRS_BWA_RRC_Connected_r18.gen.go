// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSRS-BWA-RRC-Connected-r18 (line 20495).

var posSRSBWARRCConnectedR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "numOfCarriersIntraBandContiguous-r18"},
		{Name: "maximumAggregatedBW-TwoCarriersFR1-r18", Optional: true},
		{Name: "maximumAggregatedBW-TwoCarriersFR2-r18", Optional: true},
		{Name: "maximumAggregatedBW-ThreeCarriersFR1-r18", Optional: true},
		{Name: "maximumAggregatedBW-ThreeCarriersFR2-r18", Optional: true},
		{Name: "maximumAggregatedResourceSet-r18"},
		{Name: "maximumAggregatedResourcePeriodic-r18"},
		{Name: "maximumAggregatedResourceAperiodic-r18"},
		{Name: "maximumAggregatedResourceSemi-r18"},
		{Name: "maximumAggregatedResourcePeriodicPerSlot-r18"},
		{Name: "maximumAggregatedResourceAperiodicPerSlot-r18"},
		{Name: "maximumAggregatedResourceSemiPerSlot-r18"},
	},
}

const (
	PosSRS_BWA_RRC_Connected_r18_NumOfCarriersIntraBandContiguous_r18_Two         = 0
	PosSRS_BWA_RRC_Connected_r18_NumOfCarriersIntraBandContiguous_r18_Three       = 1
	PosSRS_BWA_RRC_Connected_r18_NumOfCarriersIntraBandContiguous_r18_Twoandthree = 2
)

var posSRSBWARRCConnectedR18NumOfCarriersIntraBandContiguousR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_NumOfCarriersIntraBandContiguous_r18_Two, PosSRS_BWA_RRC_Connected_r18_NumOfCarriersIntraBandContiguous_r18_Three, PosSRS_BWA_RRC_Connected_r18_NumOfCarriersIntraBandContiguous_r18_Twoandthree},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz20  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz40  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz50  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz80  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz100 = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz160 = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz180 = 6
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz190 = 7
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz200 = 8
)

var posSRSBWARRCConnectedR18MaximumAggregatedBWTwoCarriersFR1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz20, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz40, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz50, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz80, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz100, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz160, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz180, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz190, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR1_r18_Mhz200},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz50  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz100 = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz200 = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz400 = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz600 = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz800 = 5
)

var posSRSBWARRCConnectedR18MaximumAggregatedBWTwoCarriersFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz50, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz100, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz200, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz400, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz600, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_TwoCarriersFR2_r18_Mhz800},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz80  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz100 = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz160 = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz200 = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz240 = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz300 = 5
)

var posSRSBWARRCConnectedR18MaximumAggregatedBWThreeCarriersFR1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz80, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz100, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz160, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz200, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz240, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR1_r18_Mhz300},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz50   = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz100  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz200  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz300  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz400  = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz600  = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz800  = 6
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz1000 = 7
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz1200 = 8
)

var posSRSBWARRCConnectedR18MaximumAggregatedBWThreeCarriersFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz50, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz100, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz200, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz300, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz400, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz600, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz800, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz1000, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedBW_ThreeCarriersFR2_r18_Mhz1200},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N1  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N2  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N4  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N8  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N12 = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N16 = 5
)

var posSRSBWARRCConnectedR18MaximumAggregatedResourceSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N1, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N2, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N4, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N8, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N12, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSet_r18_N16},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N1  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N2  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N4  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N8  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N16 = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N32 = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N64 = 6
)

var posSRSBWARRCConnectedR18MaximumAggregatedResourcePeriodicR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N1, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N2, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N4, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N8, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N16, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N32, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodic_r18_N64},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N0  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N1  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N2  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N4  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N8  = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N16 = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N32 = 6
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N64 = 7
)

var posSRSBWARRCConnectedR18MaximumAggregatedResourceAperiodicR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N0, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N1, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N2, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N4, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N8, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N16, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N32, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodic_r18_N64},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N0  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N1  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N2  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N4  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N8  = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N16 = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N32 = 6
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N64 = 7
)

var posSRSBWARRCConnectedR18MaximumAggregatedResourceSemiR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N0, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N1, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N2, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N4, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N8, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N16, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N32, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemi_r18_N64},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N1  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N2  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N3  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N4  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N5  = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N6  = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N8  = 6
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N10 = 7
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N12 = 8
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N14 = 9
)

var posSRSBWARRCConnectedR18MaximumAggregatedResourcePeriodicPerSlotR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N1, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N2, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N3, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N4, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N5, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N6, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N8, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N10, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N12, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourcePeriodicPerSlot_r18_N14},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N0  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N1  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N2  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N3  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N4  = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N5  = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N6  = 6
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N8  = 7
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N10 = 8
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N12 = 9
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N14 = 10
)

var posSRSBWARRCConnectedR18MaximumAggregatedResourceAperiodicPerSlotR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N0, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N1, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N2, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N3, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N4, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N5, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N6, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N8, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N10, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N12, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceAperiodicPerSlot_r18_N14},
}

const (
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N0  = 0
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N1  = 1
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N2  = 2
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N3  = 3
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N4  = 4
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N5  = 5
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N6  = 6
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N8  = 7
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N10 = 8
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N12 = 9
	PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N14 = 10
)

var posSRSBWARRCConnectedR18MaximumAggregatedResourceSemiPerSlotR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N0, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N1, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N2, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N3, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N4, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N5, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N6, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N8, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N10, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N12, PosSRS_BWA_RRC_Connected_r18_MaximumAggregatedResourceSemiPerSlot_r18_N14},
}

type PosSRS_BWA_RRC_Connected_r18 struct {
	NumOfCarriersIntraBandContiguous_r18          int64
	MaximumAggregatedBW_TwoCarriersFR1_r18        *int64
	MaximumAggregatedBW_TwoCarriersFR2_r18        *int64
	MaximumAggregatedBW_ThreeCarriersFR1_r18      *int64
	MaximumAggregatedBW_ThreeCarriersFR2_r18      *int64
	MaximumAggregatedResourceSet_r18              int64
	MaximumAggregatedResourcePeriodic_r18         int64
	MaximumAggregatedResourceAperiodic_r18        int64
	MaximumAggregatedResourceSemi_r18             int64
	MaximumAggregatedResourcePeriodicPerSlot_r18  int64
	MaximumAggregatedResourceAperiodicPerSlot_r18 int64
	MaximumAggregatedResourceSemiPerSlot_r18      int64
}

func (ie *PosSRS_BWA_RRC_Connected_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSRSBWARRCConnectedR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaximumAggregatedBW_TwoCarriersFR1_r18 != nil, ie.MaximumAggregatedBW_TwoCarriersFR2_r18 != nil, ie.MaximumAggregatedBW_ThreeCarriersFR1_r18 != nil, ie.MaximumAggregatedBW_ThreeCarriersFR2_r18 != nil}); err != nil {
		return err
	}

	// 3. numOfCarriersIntraBandContiguous-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.NumOfCarriersIntraBandContiguous_r18, posSRSBWARRCConnectedR18NumOfCarriersIntraBandContiguousR18Constraints); err != nil {
			return err
		}
	}

	// 4. maximumAggregatedBW-TwoCarriersFR1-r18: enumerated
	{
		if ie.MaximumAggregatedBW_TwoCarriersFR1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumAggregatedBW_TwoCarriersFR1_r18, posSRSBWARRCConnectedR18MaximumAggregatedBWTwoCarriersFR1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maximumAggregatedBW-TwoCarriersFR2-r18: enumerated
	{
		if ie.MaximumAggregatedBW_TwoCarriersFR2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumAggregatedBW_TwoCarriersFR2_r18, posSRSBWARRCConnectedR18MaximumAggregatedBWTwoCarriersFR2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. maximumAggregatedBW-ThreeCarriersFR1-r18: enumerated
	{
		if ie.MaximumAggregatedBW_ThreeCarriersFR1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumAggregatedBW_ThreeCarriersFR1_r18, posSRSBWARRCConnectedR18MaximumAggregatedBWThreeCarriersFR1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. maximumAggregatedBW-ThreeCarriersFR2-r18: enumerated
	{
		if ie.MaximumAggregatedBW_ThreeCarriersFR2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumAggregatedBW_ThreeCarriersFR2_r18, posSRSBWARRCConnectedR18MaximumAggregatedBWThreeCarriersFR2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. maximumAggregatedResourceSet-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaximumAggregatedResourceSet_r18, posSRSBWARRCConnectedR18MaximumAggregatedResourceSetR18Constraints); err != nil {
			return err
		}
	}

	// 9. maximumAggregatedResourcePeriodic-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaximumAggregatedResourcePeriodic_r18, posSRSBWARRCConnectedR18MaximumAggregatedResourcePeriodicR18Constraints); err != nil {
			return err
		}
	}

	// 10. maximumAggregatedResourceAperiodic-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaximumAggregatedResourceAperiodic_r18, posSRSBWARRCConnectedR18MaximumAggregatedResourceAperiodicR18Constraints); err != nil {
			return err
		}
	}

	// 11. maximumAggregatedResourceSemi-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaximumAggregatedResourceSemi_r18, posSRSBWARRCConnectedR18MaximumAggregatedResourceSemiR18Constraints); err != nil {
			return err
		}
	}

	// 12. maximumAggregatedResourcePeriodicPerSlot-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaximumAggregatedResourcePeriodicPerSlot_r18, posSRSBWARRCConnectedR18MaximumAggregatedResourcePeriodicPerSlotR18Constraints); err != nil {
			return err
		}
	}

	// 13. maximumAggregatedResourceAperiodicPerSlot-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaximumAggregatedResourceAperiodicPerSlot_r18, posSRSBWARRCConnectedR18MaximumAggregatedResourceAperiodicPerSlotR18Constraints); err != nil {
			return err
		}
	}

	// 14. maximumAggregatedResourceSemiPerSlot-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaximumAggregatedResourceSemiPerSlot_r18, posSRSBWARRCConnectedR18MaximumAggregatedResourceSemiPerSlotR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PosSRS_BWA_RRC_Connected_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSRSBWARRCConnectedR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. numOfCarriersIntraBandContiguous-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18NumOfCarriersIntraBandContiguousR18Constraints)
		if err != nil {
			return err
		}
		ie.NumOfCarriersIntraBandContiguous_r18 = v0
	}

	// 4. maximumAggregatedBW-TwoCarriersFR1-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedBWTwoCarriersFR1R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumAggregatedBW_TwoCarriersFR1_r18 = &idx
		}
	}

	// 5. maximumAggregatedBW-TwoCarriersFR2-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedBWTwoCarriersFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumAggregatedBW_TwoCarriersFR2_r18 = &idx
		}
	}

	// 6. maximumAggregatedBW-ThreeCarriersFR1-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedBWThreeCarriersFR1R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumAggregatedBW_ThreeCarriersFR1_r18 = &idx
		}
	}

	// 7. maximumAggregatedBW-ThreeCarriersFR2-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedBWThreeCarriersFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumAggregatedBW_ThreeCarriersFR2_r18 = &idx
		}
	}

	// 8. maximumAggregatedResourceSet-r18: enumerated
	{
		v5, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedResourceSetR18Constraints)
		if err != nil {
			return err
		}
		ie.MaximumAggregatedResourceSet_r18 = v5
	}

	// 9. maximumAggregatedResourcePeriodic-r18: enumerated
	{
		v6, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedResourcePeriodicR18Constraints)
		if err != nil {
			return err
		}
		ie.MaximumAggregatedResourcePeriodic_r18 = v6
	}

	// 10. maximumAggregatedResourceAperiodic-r18: enumerated
	{
		v7, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedResourceAperiodicR18Constraints)
		if err != nil {
			return err
		}
		ie.MaximumAggregatedResourceAperiodic_r18 = v7
	}

	// 11. maximumAggregatedResourceSemi-r18: enumerated
	{
		v8, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedResourceSemiR18Constraints)
		if err != nil {
			return err
		}
		ie.MaximumAggregatedResourceSemi_r18 = v8
	}

	// 12. maximumAggregatedResourcePeriodicPerSlot-r18: enumerated
	{
		v9, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedResourcePeriodicPerSlotR18Constraints)
		if err != nil {
			return err
		}
		ie.MaximumAggregatedResourcePeriodicPerSlot_r18 = v9
	}

	// 13. maximumAggregatedResourceAperiodicPerSlot-r18: enumerated
	{
		v10, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedResourceAperiodicPerSlotR18Constraints)
		if err != nil {
			return err
		}
		ie.MaximumAggregatedResourceAperiodicPerSlot_r18 = v10
	}

	// 14. maximumAggregatedResourceSemiPerSlot-r18: enumerated
	{
		v11, err := d.DecodeEnumerated(posSRSBWARRCConnectedR18MaximumAggregatedResourceSemiPerSlotR18Constraints)
		if err != nil {
			return err
		}
		ie.MaximumAggregatedResourceSemiPerSlot_r18 = v11
	}

	return nil
}
