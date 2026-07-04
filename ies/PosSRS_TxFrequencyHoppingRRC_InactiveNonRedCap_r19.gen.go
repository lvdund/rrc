// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSRS-TxFrequencyHoppingRRC-InactiveNonRedCap-r19 (line 23425).

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maximumSRS-BandwidthAcrossAllHopsFR1-r19", Optional: true},
		{Name: "maximumSRS-BandwidthAcrossAllHopsFR2-r19", Optional: true},
		{Name: "maximumTxFH-Hops-r19", Optional: true},
		{Name: "rf-TxRetuneTimeFR1-r19", Optional: true},
		{Name: "rf-TxRetuneTimeFR2-r19", Optional: true},
		{Name: "switchTimeBetweenActiveBWP-FrequencyHop-r19", Optional: true},
		{Name: "numOfOverlappingPRB-r19", Optional: true},
		{Name: "maximumSRS-ResourcePeriodic-r19", Optional: true},
		{Name: "maximumSRS-ResourceSemipersistent-r19", Optional: true},
	},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz40  = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz50  = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz80  = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz100 = 3
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSBandwidthAcrossAllHopsFR1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz40, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz50, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz80, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR1_r19_Mhz100},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR2_r19_Mhz100 = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR2_r19_Mhz200 = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR2_r19_Mhz400 = 2
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSBandwidthAcrossAllHopsFR2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR2_r19_Mhz100, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR2_r19_Mhz200, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_BandwidthAcrossAllHopsFR2_r19_Mhz400},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N2 = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N3 = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N4 = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N5 = 3
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N6 = 4
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumTxFHHopsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N2, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N3, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N4, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N5, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumTxFH_Hops_r19_N6},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us0   = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us70  = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us140 = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us210 = 3
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19RfTxRetuneTimeFR1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us0, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us70, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us140, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR1_r19_Us210},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us0   = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us35  = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us70  = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us140 = 3
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19RfTxRetuneTimeFR2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us0, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us35, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us70, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_Rf_TxRetuneTimeFR2_r19_Us140},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us0   = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us100 = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us140 = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us200 = 3
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us300 = 4
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us500 = 5
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19SwitchTimeBetweenActiveBWPFrequencyHopR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us0, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us100, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us140, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us200, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us300, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_SwitchTimeBetweenActiveBWP_FrequencyHop_r19_Us500},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N0 = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N1 = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N2 = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N4 = 3
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19NumOfOverlappingPRBR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N0, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N1, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N2, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_NumOfOverlappingPRB_r19_N4},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N1  = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N2  = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N4  = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N8  = 3
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N16 = 4
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N32 = 5
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N64 = 6
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSResourcePeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N1, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N2, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N4, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N8, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N16, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N32, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourcePeriodic_r19_N64},
}

const (
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N0  = 0
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N1  = 1
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N2  = 2
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N4  = 3
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N8  = 4
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N16 = 5
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N32 = 6
	PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N64 = 7
)

var posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSResourceSemipersistentR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N0, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N1, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N2, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N4, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N8, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N16, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N32, PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19_MaximumSRS_ResourceSemipersistent_r19_N64},
}

type PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19 struct {
	MaximumSRS_BandwidthAcrossAllHopsFR1_r19    *int64
	MaximumSRS_BandwidthAcrossAllHopsFR2_r19    *int64
	MaximumTxFH_Hops_r19                        *int64
	Rf_TxRetuneTimeFR1_r19                      *int64
	Rf_TxRetuneTimeFR2_r19                      *int64
	SwitchTimeBetweenActiveBWP_FrequencyHop_r19 *int64
	NumOfOverlappingPRB_r19                     *int64
	MaximumSRS_ResourcePeriodic_r19             *int64
	MaximumSRS_ResourceSemipersistent_r19       *int64
}

func (ie *PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r19 != nil, ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r19 != nil, ie.MaximumTxFH_Hops_r19 != nil, ie.Rf_TxRetuneTimeFR1_r19 != nil, ie.Rf_TxRetuneTimeFR2_r19 != nil, ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r19 != nil, ie.NumOfOverlappingPRB_r19 != nil, ie.MaximumSRS_ResourcePeriodic_r19 != nil, ie.MaximumSRS_ResourceSemipersistent_r19 != nil}); err != nil {
		return err
	}

	// 3. maximumSRS-BandwidthAcrossAllHopsFR1-r19: enumerated
	{
		if ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSBandwidthAcrossAllHopsFR1R19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maximumSRS-BandwidthAcrossAllHopsFR2-r19: enumerated
	{
		if ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSBandwidthAcrossAllHopsFR2R19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maximumTxFH-Hops-r19: enumerated
	{
		if ie.MaximumTxFH_Hops_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumTxFH_Hops_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumTxFHHopsR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. rf-TxRetuneTimeFR1-r19: enumerated
	{
		if ie.Rf_TxRetuneTimeFR1_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Rf_TxRetuneTimeFR1_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19RfTxRetuneTimeFR1R19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. rf-TxRetuneTimeFR2-r19: enumerated
	{
		if ie.Rf_TxRetuneTimeFR2_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Rf_TxRetuneTimeFR2_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19RfTxRetuneTimeFR2R19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. switchTimeBetweenActiveBWP-FrequencyHop-r19: enumerated
	{
		if ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r19 != nil {
			if err := e.EncodeEnumerated(*ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19SwitchTimeBetweenActiveBWPFrequencyHopR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. numOfOverlappingPRB-r19: enumerated
	{
		if ie.NumOfOverlappingPRB_r19 != nil {
			if err := e.EncodeEnumerated(*ie.NumOfOverlappingPRB_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19NumOfOverlappingPRBR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. maximumSRS-ResourcePeriodic-r19: enumerated
	{
		if ie.MaximumSRS_ResourcePeriodic_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_ResourcePeriodic_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSResourcePeriodicR19Constraints); err != nil {
				return err
			}
		}
	}

	// 11. maximumSRS-ResourceSemipersistent-r19: enumerated
	{
		if ie.MaximumSRS_ResourceSemipersistent_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumSRS_ResourceSemipersistent_r19, posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSResourceSemipersistentR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maximumSRS-BandwidthAcrossAllHopsFR1-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSBandwidthAcrossAllHopsFR1R19Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_BandwidthAcrossAllHopsFR1_r19 = &idx
		}
	}

	// 4. maximumSRS-BandwidthAcrossAllHopsFR2-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSBandwidthAcrossAllHopsFR2R19Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_BandwidthAcrossAllHopsFR2_r19 = &idx
		}
	}

	// 5. maximumTxFH-Hops-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumTxFHHopsR19Constraints)
			if err != nil {
				return err
			}
			ie.MaximumTxFH_Hops_r19 = &idx
		}
	}

	// 6. rf-TxRetuneTimeFR1-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19RfTxRetuneTimeFR1R19Constraints)
			if err != nil {
				return err
			}
			ie.Rf_TxRetuneTimeFR1_r19 = &idx
		}
	}

	// 7. rf-TxRetuneTimeFR2-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19RfTxRetuneTimeFR2R19Constraints)
			if err != nil {
				return err
			}
			ie.Rf_TxRetuneTimeFR2_r19 = &idx
		}
	}

	// 8. switchTimeBetweenActiveBWP-FrequencyHop-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19SwitchTimeBetweenActiveBWPFrequencyHopR19Constraints)
			if err != nil {
				return err
			}
			ie.SwitchTimeBetweenActiveBWP_FrequencyHop_r19 = &idx
		}
	}

	// 9. numOfOverlappingPRB-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19NumOfOverlappingPRBR19Constraints)
			if err != nil {
				return err
			}
			ie.NumOfOverlappingPRB_r19 = &idx
		}
	}

	// 10. maximumSRS-ResourcePeriodic-r19: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSResourcePeriodicR19Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_ResourcePeriodic_r19 = &idx
		}
	}

	// 11. maximumSRS-ResourceSemipersistent-r19: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(posSRSTxFrequencyHoppingRRCInactiveNonRedCapR19MaximumSRSResourceSemipersistentR19Constraints)
			if err != nil {
				return err
			}
			ie.MaximumSRS_ResourceSemipersistent_r19 = &idx
		}
	}

	return nil
}
