// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSRS-TxFrequencyHoppingRRC-Connected-r18 (line 23375).

var posSRSTxFrequencyHoppingRRCConnectedR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maximumSRS-BandwidthAcrossAllHopsFR1-r18", Optional: true},
		{Name: "maximumSRS-BandwidthAcrossAllHopsFR2-r18", Optional: true},
		{Name: "maximumTxFH-Hops-r18", Optional: true},
		{Name: "rf-TxRetuneTimeFR1-r18", Optional: true},
		{Name: "rf-TxRetuneTimeFR2-r18", Optional: true},
		{Name: "switchTimeBetweenActiveBWP-FrequencyHop-r18", Optional: true},
		{Name: "numOfOverlappingPRB-r18", Optional: true},
		{Name: "maximumSRS-ResourcePeriodic-r18", Optional: true},
		{Name: "maximumSRS-ResourceAperiodic-r18", Optional: true},
		{Name: "maximumSRS-ResourceSemipersistent-r18", Optional: true},
	},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz40  = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz50  = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz80  = 2
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz100 = 3
)

var posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSBandwidthAcrossAllHopsFR1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz40, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz50, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz80, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR1_r18_Mhz100},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR2_r18_Mhz100 = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR2_r18_Mhz200 = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR2_r18_Mhz400 = 2
)

var posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSBandwidthAcrossAllHopsFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR2_r18_Mhz100, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR2_r18_Mhz200, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_BandwidthAcrossAllHopsFR2_r18_Mhz400},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N2 = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N3 = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N4 = 2
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N5 = 3
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N6 = 4
)

var posSRSTxFrequencyHoppingRRCConnectedR18MaximumTxFHHopsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N2, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N3, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N4, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N5, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumTxFH_Hops_r18_N6},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR1_r18_N70  = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR1_r18_N140 = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR1_r18_N210 = 2
)

var posSRSTxFrequencyHoppingRRCConnectedR18RfTxRetuneTimeFR1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR1_r18_N70, PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR1_r18_N140, PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR1_r18_N210},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR2_r18_N35  = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR2_r18_N70  = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR2_r18_N140 = 2
)

var posSRSTxFrequencyHoppingRRCConnectedR18RfTxRetuneTimeFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR2_r18_N35, PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR2_r18_N70, PosSRS_TxFrequencyHoppingRRC_Connected_r18_Rf_TxRetuneTimeFR2_r18_N140},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N100 = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N140 = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N200 = 2
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N300 = 3
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N500 = 4
)

var posSRSTxFrequencyHoppingRRCConnectedR18SwitchTimeBetweenActiveBWPFrequencyHopR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N100, PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N140, PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N200, PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N300, PosSRS_TxFrequencyHoppingRRC_Connected_r18_SwitchTimeBetweenActiveBWP_FrequencyHop_r18_N500},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N0 = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N1 = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N2 = 2
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N4 = 3
)

var posSRSTxFrequencyHoppingRRCConnectedR18NumOfOverlappingPRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N0, PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N1, PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N2, PosSRS_TxFrequencyHoppingRRC_Connected_r18_NumOfOverlappingPRB_r18_N4},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N1  = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N2  = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N4  = 2
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N8  = 3
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N16 = 4
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N32 = 5
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N64 = 6
)

var posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourcePeriodicR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N1, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N2, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N4, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N8, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N16, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N32, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourcePeriodic_r18_N64},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N0  = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N1  = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N2  = 2
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N4  = 3
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N8  = 4
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N16 = 5
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N32 = 6
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N64 = 7
)

var posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourceAperiodicR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N0, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N1, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N2, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N4, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N8, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N16, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N32, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceAperiodic_r18_N64},
}

const (
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N0  = 0
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N1  = 1
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N2  = 2
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N4  = 3
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N8  = 4
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N16 = 5
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N32 = 6
	PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N64 = 7
)

var posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourceSemipersistentR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N0, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N1, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N2, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N4, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N8, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N16, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N32, PosSRS_TxFrequencyHoppingRRC_Connected_r18_MaximumSRS_ResourceSemipersistent_r18_N64},
}

type PosSRS_TxFrequencyHoppingRRC_Connected_r18 struct {
	MaximumSRS_BandwidthAcrossAllHopsFR1_r18    *int64
	MaximumSRS_BandwidthAcrossAllHopsFR2_r18    *int64
	MaximumTxFH_Hops_r18                        *int64
	Rf_TxRetuneTimeFR1_r18                      *int64
	Rf_TxRetuneTimeFR2_r18                      *int64
	SwitchTimeBetweenActiveBWP_FrequencyHop_r18 *int64
	NumOfOverlappingPRB_r18                     *int64
	MaximumSRS_ResourcePeriodic_r18             *int64
	MaximumSRS_ResourceAperiodic_r18            *int64
	MaximumSRS_ResourceSemipersistent_r18       *int64
}

func (ie *PosSRS_TxFrequencyHoppingRRC_Connected_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSRSTxFrequencyHoppingRRCConnectedR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r18 != nil, ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r18 != nil, ie.MaximumTxFH_Hops_r18 != nil, ie.Rf_TxRetuneTimeFR1_r18 != nil, ie.Rf_TxRetuneTimeFR2_r18 != nil, ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r18 != nil, ie.NumOfOverlappingPRB_r18 != nil, ie.MaximumSRS_ResourcePeriodic_r18 != nil, ie.MaximumSRS_ResourceAperiodic_r18 != nil, ie.MaximumSRS_ResourceSemipersistent_r18 != nil}); err != nil {
		return err
	}

	// 3. maximumSRS-BandwidthAcrossAllHopsFR1-r18: enumerated
	{
		if ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r18, posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSBandwidthAcrossAllHopsFR1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maximumSRS-BandwidthAcrossAllHopsFR2-r18: enumerated
	{
		if ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r18, posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSBandwidthAcrossAllHopsFR2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maximumTxFH-Hops-r18: enumerated
	{
		if ie.MaximumTxFH_Hops_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumTxFH_Hops_r18, posSRSTxFrequencyHoppingRRCConnectedR18MaximumTxFHHopsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. rf-TxRetuneTimeFR1-r18: enumerated
	{
		if ie.Rf_TxRetuneTimeFR1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rf_TxRetuneTimeFR1_r18, posSRSTxFrequencyHoppingRRCConnectedR18RfTxRetuneTimeFR1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. rf-TxRetuneTimeFR2-r18: enumerated
	{
		if ie.Rf_TxRetuneTimeFR2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rf_TxRetuneTimeFR2_r18, posSRSTxFrequencyHoppingRRCConnectedR18RfTxRetuneTimeFR2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. switchTimeBetweenActiveBWP-FrequencyHop-r18: enumerated
	{
		if ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r18, posSRSTxFrequencyHoppingRRCConnectedR18SwitchTimeBetweenActiveBWPFrequencyHopR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. numOfOverlappingPRB-r18: enumerated
	{
		if ie.NumOfOverlappingPRB_r18 != nil {
			if err := e.EncodeEnumerated(*ie.NumOfOverlappingPRB_r18, posSRSTxFrequencyHoppingRRCConnectedR18NumOfOverlappingPRBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. maximumSRS-ResourcePeriodic-r18: enumerated
	{
		if ie.MaximumSRS_ResourcePeriodic_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_ResourcePeriodic_r18, posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourcePeriodicR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. maximumSRS-ResourceAperiodic-r18: enumerated
	{
		if ie.MaximumSRS_ResourceAperiodic_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_ResourceAperiodic_r18, posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourceAperiodicR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. maximumSRS-ResourceSemipersistent-r18: enumerated
	{
		if ie.MaximumSRS_ResourceSemipersistent_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_ResourceSemipersistent_r18, posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourceSemipersistentR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PosSRS_TxFrequencyHoppingRRC_Connected_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSRSTxFrequencyHoppingRRCConnectedR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maximumSRS-BandwidthAcrossAllHopsFR1-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSBandwidthAcrossAllHopsFR1R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r18 = &idx
		}
	}

	// 4. maximumSRS-BandwidthAcrossAllHopsFR2-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSBandwidthAcrossAllHopsFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r18 = &idx
		}
	}

	// 5. maximumTxFH-Hops-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18MaximumTxFHHopsR18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumTxFH_Hops_r18 = &idx
		}
	}

	// 6. rf-TxRetuneTimeFR1-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18RfTxRetuneTimeFR1R18Constraints)
			if err != nil {
				return err
			}
			ie.Rf_TxRetuneTimeFR1_r18 = &idx
		}
	}

	// 7. rf-TxRetuneTimeFR2-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18RfTxRetuneTimeFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.Rf_TxRetuneTimeFR2_r18 = &idx
		}
	}

	// 8. switchTimeBetweenActiveBWP-FrequencyHop-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18SwitchTimeBetweenActiveBWPFrequencyHopR18Constraints)
			if err != nil {
				return err
			}
			ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r18 = &idx
		}
	}

	// 9. numOfOverlappingPRB-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18NumOfOverlappingPRBR18Constraints)
			if err != nil {
				return err
			}
			ie.NumOfOverlappingPRB_r18 = &idx
		}
	}

	// 10. maximumSRS-ResourcePeriodic-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourcePeriodicR18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_ResourcePeriodic_r18 = &idx
		}
	}

	// 11. maximumSRS-ResourceAperiodic-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourceAperiodicR18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_ResourceAperiodic_r18 = &idx
		}
	}

	// 12. maximumSRS-ResourceSemipersistent-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCConnectedR18MaximumSRSResourceSemipersistentR18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_ResourceSemipersistent_r18 = &idx
		}
	}

	return nil
}
