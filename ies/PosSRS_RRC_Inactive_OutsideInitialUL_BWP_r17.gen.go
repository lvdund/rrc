// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSRS-RRC-Inactive-OutsideInitialUL-BWP-r17 (line 23352).

var posSRSRRCInactiveOutsideInitialULBWPR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxSRSposBandwidthForEachSCS-withinCC-FR1-r17", Optional: true},
		{Name: "maxSRSposBandwidthForEachSCS-withinCC-FR2-r17", Optional: true},
		{Name: "maxNumOfSRSposResourceSets-r17", Optional: true},
		{Name: "maxNumOfPeriodicSRSposResources-r17", Optional: true},
		{Name: "maxNumOfPeriodicSRSposResourcesPerSlot-r17", Optional: true},
		{Name: "differentNumerologyBetweenSRSposAndInitialBWP-r17", Optional: true},
		{Name: "srsPosWithoutRestrictionOnBWP-r17", Optional: true},
		{Name: "maxNumOfPeriodicAndSemipersistentSRSposResources-r17", Optional: true},
		{Name: "maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot-r17", Optional: true},
		{Name: "differentCenterFreqBetweenSRSposAndInitialBWP-r17", Optional: true},
		{Name: "switchingTimeSRS-TX-OtherTX-r17", Optional: true},
		{Name: "maxNumOfSemiPersistentSRSposResources-r17", Optional: true},
		{Name: "maxNumOfSemiPersistentSRSposResourcesPerSlot-r17", Optional: true},
	},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz5   = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz10  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz15  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz20  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz25  = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz30  = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz35  = 6
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz40  = 7
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz45  = 8
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz50  = 9
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz60  = 10
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz70  = 11
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz80  = 12
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz90  = 13
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz100 = 14
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxSRSposBandwidthForEachSCSWithinCCFR1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz5, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz10, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz15, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz20, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz25, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz30, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz35, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz40, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz45, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz50, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz60, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz70, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz80, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz90, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17_Mhz100},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz50  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz100 = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz200 = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz400 = 3
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxSRSposBandwidthForEachSCSWithinCCFR2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz50, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz100, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz200, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17_Mhz400},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N1  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N2  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N4  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N8  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N12 = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N16 = 5
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSRSposResourceSetsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N1, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N2, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N4, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N8, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N12, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSRSposResourceSets_r17_N16},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N1  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N2  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N4  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N8  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N16 = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N32 = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N64 = 6
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicSRSposResourcesR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N1, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N2, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N4, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N8, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N16, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N32, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResources_r17_N64},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N1  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N2  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N3  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N4  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N5  = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N6  = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N8  = 6
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N10 = 7
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N12 = 8
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N14 = 9
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicSRSposResourcesPerSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N1, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N2, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N3, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N4, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N5, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N6, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N8, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N10, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N12, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicSRSposResourcesPerSlot_r17_N14},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_DifferentNumerologyBetweenSRSposAndInitialBWP_r17_Supported = 0
)

var posSRSRRCInactiveOutsideInitialULBWPR17DifferentNumerologyBetweenSRSposAndInitialBWPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_DifferentNumerologyBetweenSRSposAndInitialBWP_r17_Supported},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SrsPosWithoutRestrictionOnBWP_r17_Supported = 0
)

var posSRSRRCInactiveOutsideInitialULBWPR17SrsPosWithoutRestrictionOnBWPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SrsPosWithoutRestrictionOnBWP_r17_Supported},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N1  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N2  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N4  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N8  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N16 = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N32 = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N64 = 6
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicAndSemipersistentSRSposResourcesR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N1, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N2, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N4, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N8, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N16, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N32, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResources_r17_N64},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N1  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N2  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N3  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N4  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N5  = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N6  = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N8  = 6
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N10 = 7
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N12 = 8
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N14 = 9
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N1, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N2, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N3, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N4, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N5, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N6, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N8, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N10, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N12, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17_N14},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_DifferentCenterFreqBetweenSRSposAndInitialBWP_r17_Supported = 0
)

var posSRSRRCInactiveOutsideInitialULBWPR17DifferentCenterFreqBetweenSRSposAndInitialBWPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_DifferentCenterFreqBetweenSRSposAndInitialBWP_r17_Supported},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us100 = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us140 = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us200 = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us300 = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us500 = 4
)

var posSRSRRCInactiveOutsideInitialULBWPR17SwitchingTimeSRSTXOtherTXR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us100, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us140, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us200, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us300, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_SwitchingTimeSRS_TX_OtherTX_r17_Us500},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N1  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N2  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N4  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N8  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N16 = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N32 = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N64 = 6
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSemiPersistentSRSposResourcesR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N1, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N2, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N4, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N8, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N16, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N32, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResources_r17_N64},
}

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N1  = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N2  = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N3  = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N4  = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N5  = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N6  = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N8  = 6
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N10 = 7
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N12 = 8
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N14 = 9
)

var posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSemiPersistentSRSposResourcesPerSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N1, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N2, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N3, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N4, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N5, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N6, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N8, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N10, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N12, PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N14},
}

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 struct {
	MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17               *int64
	MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17               *int64
	MaxNumOfSRSposResourceSets_r17                              *int64
	MaxNumOfPeriodicSRSposResources_r17                         *int64
	MaxNumOfPeriodicSRSposResourcesPerSlot_r17                  *int64
	DifferentNumerologyBetweenSRSposAndInitialBWP_r17           *int64
	SrsPosWithoutRestrictionOnBWP_r17                           *int64
	MaxNumOfPeriodicAndSemipersistentSRSposResources_r17        *int64
	MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 *int64
	DifferentCenterFreqBetweenSRSposAndInitialBWP_r17           *int64
	SwitchingTimeSRS_TX_OtherTX_r17                             *int64
	MaxNumOfSemiPersistentSRSposResources_r17                   *int64
	MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17            *int64
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSRSRRCInactiveOutsideInitialULBWPR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17 != nil, ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17 != nil, ie.MaxNumOfSRSposResourceSets_r17 != nil, ie.MaxNumOfPeriodicSRSposResources_r17 != nil, ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17 != nil, ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17 != nil, ie.SrsPosWithoutRestrictionOnBWP_r17 != nil, ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17 != nil, ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 != nil, ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17 != nil, ie.SwitchingTimeSRS_TX_OtherTX_r17 != nil, ie.MaxNumOfSemiPersistentSRSposResources_r17 != nil, ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 != nil}); err != nil {
		return err
	}

	// 3. maxSRSposBandwidthForEachSCS-withinCC-FR1-r17: enumerated
	{
		if ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxSRSposBandwidthForEachSCSWithinCCFR1R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maxSRSposBandwidthForEachSCS-withinCC-FR2-r17: enumerated
	{
		if ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxSRSposBandwidthForEachSCSWithinCCFR2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maxNumOfSRSposResourceSets-r17: enumerated
	{
		if ie.MaxNumOfSRSposResourceSets_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumOfSRSposResourceSets_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSRSposResourceSetsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. maxNumOfPeriodicSRSposResources-r17: enumerated
	{
		if ie.MaxNumOfPeriodicSRSposResources_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumOfPeriodicSRSposResources_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicSRSposResourcesR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. maxNumOfPeriodicSRSposResourcesPerSlot-r17: enumerated
	{
		if ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicSRSposResourcesPerSlotR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. differentNumerologyBetweenSRSposAndInitialBWP-r17: enumerated
	{
		if ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17, posSRSRRCInactiveOutsideInitialULBWPR17DifferentNumerologyBetweenSRSposAndInitialBWPR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. srsPosWithoutRestrictionOnBWP-r17: enumerated
	{
		if ie.SrsPosWithoutRestrictionOnBWP_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SrsPosWithoutRestrictionOnBWP_r17, posSRSRRCInactiveOutsideInitialULBWPR17SrsPosWithoutRestrictionOnBWPR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. maxNumOfPeriodicAndSemipersistentSRSposResources-r17: enumerated
	{
		if ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicAndSemipersistentSRSposResourcesR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot-r17: enumerated
	{
		if ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlotR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. differentCenterFreqBetweenSRSposAndInitialBWP-r17: enumerated
	{
		if ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17, posSRSRRCInactiveOutsideInitialULBWPR17DifferentCenterFreqBetweenSRSposAndInitialBWPR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. switchingTimeSRS-TX-OtherTX-r17: enumerated
	{
		if ie.SwitchingTimeSRS_TX_OtherTX_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingTimeSRS_TX_OtherTX_r17, posSRSRRCInactiveOutsideInitialULBWPR17SwitchingTimeSRSTXOtherTXR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. maxNumOfSemiPersistentSRSposResources-r17: enumerated
	{
		if ie.MaxNumOfSemiPersistentSRSposResources_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumOfSemiPersistentSRSposResources_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSemiPersistentSRSposResourcesR17Constraints); err != nil {
				return err
			}
		}
	}

	// 15. maxNumOfSemiPersistentSRSposResourcesPerSlot-r17: enumerated
	{
		if ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17, posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSemiPersistentSRSposResourcesPerSlotR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSRSRRCInactiveOutsideInitialULBWPR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maxSRSposBandwidthForEachSCS-withinCC-FR1-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxSRSposBandwidthForEachSCSWithinCCFR1R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR1_r17 = &idx
		}
	}

	// 4. maxSRSposBandwidthForEachSCS-withinCC-FR2-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxSRSposBandwidthForEachSCSWithinCCFR2R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxSRSposBandwidthForEachSCS_WithinCC_FR2_r17 = &idx
		}
	}

	// 5. maxNumOfSRSposResourceSets-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSRSposResourceSetsR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumOfSRSposResourceSets_r17 = &idx
		}
	}

	// 6. maxNumOfPeriodicSRSposResources-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicSRSposResourcesR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumOfPeriodicSRSposResources_r17 = &idx
		}
	}

	// 7. maxNumOfPeriodicSRSposResourcesPerSlot-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicSRSposResourcesPerSlotR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumOfPeriodicSRSposResourcesPerSlot_r17 = &idx
		}
	}

	// 8. differentNumerologyBetweenSRSposAndInitialBWP-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17DifferentNumerologyBetweenSRSposAndInitialBWPR17Constraints)
			if err != nil {
				return err
			}
			ie.DifferentNumerologyBetweenSRSposAndInitialBWP_r17 = &idx
		}
	}

	// 9. srsPosWithoutRestrictionOnBWP-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17SrsPosWithoutRestrictionOnBWPR17Constraints)
			if err != nil {
				return err
			}
			ie.SrsPosWithoutRestrictionOnBWP_r17 = &idx
		}
	}

	// 10. maxNumOfPeriodicAndSemipersistentSRSposResources-r17: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicAndSemipersistentSRSposResourcesR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumOfPeriodicAndSemipersistentSRSposResources_r17 = &idx
		}
	}

	// 11. maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlotR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 = &idx
		}
	}

	// 12. differentCenterFreqBetweenSRSposAndInitialBWP-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17DifferentCenterFreqBetweenSRSposAndInitialBWPR17Constraints)
			if err != nil {
				return err
			}
			ie.DifferentCenterFreqBetweenSRSposAndInitialBWP_r17 = &idx
		}
	}

	// 13. switchingTimeSRS-TX-OtherTX-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17SwitchingTimeSRSTXOtherTXR17Constraints)
			if err != nil {
				return err
			}
			ie.SwitchingTimeSRS_TX_OtherTX_r17 = &idx
		}
	}

	// 14. maxNumOfSemiPersistentSRSposResources-r17: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSemiPersistentSRSposResourcesR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumOfSemiPersistentSRSposResources_r17 = &idx
		}
	}

	// 15. maxNumOfSemiPersistentSRSposResourcesPerSlot-r17: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(posSRSRRCInactiveOutsideInitialULBWPR17MaxNumOfSemiPersistentSRSposResourcesPerSlotR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 = &idx
		}
	}

	return nil
}
