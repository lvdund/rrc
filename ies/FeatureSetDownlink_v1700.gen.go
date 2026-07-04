// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v1700 (line 19582).

var featureSetDownlinkV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scalingFactor-1024QAM-FR1-r17", Optional: true},
		{Name: "timeDurationForQCL-v1710", Optional: true},
		{Name: "sfn-SchemeA-r17", Optional: true},
		{Name: "sfn-SchemeA-PDCCH-only-r17", Optional: true},
		{Name: "sfn-SchemeA-DynamicSwitching-r17", Optional: true},
		{Name: "sfn-SchemeA-PDSCH-only-r17", Optional: true},
		{Name: "sfn-SchemeB-r17", Optional: true},
		{Name: "sfn-SchemeB-DynamicSwitching-r17", Optional: true},
		{Name: "sfn-SchemeB-PDSCH-only-r17", Optional: true},
		{Name: "mTRP-PDCCH-Case2-1SpanGap-r17", Optional: true},
		{Name: "mTRP-PDCCH-legacyMonitoring-r17", Optional: true},
		{Name: "mTRP-PDCCH-multiDCI-multiTRP-r17", Optional: true},
		{Name: "dynamicMulticastPCell-r17", Optional: true},
		{Name: "mTRP-PDCCH-Repetition-r17", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1700_ScalingFactor_1024QAM_FR1_r17_F0p4  = 0
	FeatureSetDownlink_v1700_ScalingFactor_1024QAM_FR1_r17_F0p75 = 1
	FeatureSetDownlink_v1700_ScalingFactor_1024QAM_FR1_r17_F0p8  = 2
)

var featureSetDownlinkV1700ScalingFactor1024QAMFR1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_ScalingFactor_1024QAM_FR1_r17_F0p4, FeatureSetDownlink_v1700_ScalingFactor_1024QAM_FR1_r17_F0p75, FeatureSetDownlink_v1700_ScalingFactor_1024QAM_FR1_r17_F0p8},
}

const (
	FeatureSetDownlink_v1700_Sfn_SchemeA_r17_Supported = 0
)

var featureSetDownlinkV1700SfnSchemeAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_Sfn_SchemeA_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_Sfn_SchemeA_PDCCH_Only_r17_Supported = 0
)

var featureSetDownlinkV1700SfnSchemeAPDCCHOnlyR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_Sfn_SchemeA_PDCCH_Only_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_Sfn_SchemeA_DynamicSwitching_r17_Supported = 0
)

var featureSetDownlinkV1700SfnSchemeADynamicSwitchingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_Sfn_SchemeA_DynamicSwitching_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_Sfn_SchemeA_PDSCH_Only_r17_Supported = 0
)

var featureSetDownlinkV1700SfnSchemeAPDSCHOnlyR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_Sfn_SchemeA_PDSCH_Only_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_Sfn_SchemeB_r17_Supported = 0
)

var featureSetDownlinkV1700SfnSchemeBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_Sfn_SchemeB_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_Sfn_SchemeB_DynamicSwitching_r17_Supported = 0
)

var featureSetDownlinkV1700SfnSchemeBDynamicSwitchingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_Sfn_SchemeB_DynamicSwitching_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_Sfn_SchemeB_PDSCH_Only_r17_Supported = 0
)

var featureSetDownlinkV1700SfnSchemeBPDSCHOnlyR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_Sfn_SchemeB_PDSCH_Only_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_MTRP_PDCCH_MultiDCI_MultiTRP_r17_Supported = 0
)

var featureSetDownlinkV1700MTRPPDCCHMultiDCIMultiTRPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_MTRP_PDCCH_MultiDCI_MultiTRP_r17_Supported},
}

const (
	FeatureSetDownlink_v1700_DynamicMulticastPCell_r17_Supported = 0
)

var featureSetDownlinkV1700DynamicMulticastPCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_DynamicMulticastPCell_r17_Supported},
}

var featureSetDownlinkV1700TimeDurationForQCLV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-480kHz", Optional: true},
		{Name: "scs-960kHz", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_480kHz_S56  = 0
	FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_480kHz_S112 = 1
)

var featureSetDownlinkV1700TimeDurationForQCLV1710Scs480kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_480kHz_S56, FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_480kHz_S112},
}

const (
	FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_960kHz_S112 = 0
	FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_960kHz_S224 = 1
)

var featureSetDownlinkV1700TimeDurationForQCLV1710Scs960kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_960kHz_S112, FeatureSetDownlink_v1700_TimeDurationForQCL_v1710_Scs_960kHz_S224},
}

var featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r17", Optional: true},
		{Name: "scs-30kHz-r17", Optional: true},
		{Name: "scs-60kHz-r17", Optional: true},
		{Name: "scs-120kHz-r17", Optional: true},
	},
}

var featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r17", Optional: true},
		{Name: "scs-30kHz-r17", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N1  = 0
	FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N2  = 1
	FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N3  = 2
	FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N5  = 3
	FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N10 = 4
	FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N20 = 5
	FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N40 = 6
)

var featureSetDownlinkV1700MTRPPDCCHRepetitionR17MaxNumOverlapsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N1, FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N2, FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N3, FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N5, FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N10, FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N20, FeatureSetDownlink_v1700_MTRP_PDCCH_Repetition_r17_MaxNumOverlaps_r17_N40},
}

type FeatureSetDownlink_v1700 struct {
	ScalingFactor_1024QAM_FR1_r17 *int64
	TimeDurationForQCL_v1710      *struct {
		Scs_480kHz *int64
		Scs_960kHz *int64
	}
	Sfn_SchemeA_r17                  *int64
	Sfn_SchemeA_PDCCH_Only_r17       *int64
	Sfn_SchemeA_DynamicSwitching_r17 *int64
	Sfn_SchemeA_PDSCH_Only_r17       *int64
	Sfn_SchemeB_r17                  *int64
	Sfn_SchemeB_DynamicSwitching_r17 *int64
	Sfn_SchemeB_PDSCH_Only_r17       *int64
	MTRP_PDCCH_Case2_1SpanGap_r17    *struct {
		Scs_15kHz_r17  *PDCCH_RepetitionParameters_r17
		Scs_30kHz_r17  *PDCCH_RepetitionParameters_r17
		Scs_60kHz_r17  *PDCCH_RepetitionParameters_r17
		Scs_120kHz_r17 *PDCCH_RepetitionParameters_r17
	}
	MTRP_PDCCH_LegacyMonitoring_r17 *struct {
		Scs_15kHz_r17 *PDCCH_RepetitionParameters_r17
		Scs_30kHz_r17 *PDCCH_RepetitionParameters_r17
	}
	MTRP_PDCCH_MultiDCI_MultiTRP_r17 *int64
	DynamicMulticastPCell_r17        *int64
	MTRP_PDCCH_Repetition_r17        *struct {
		NumBD_TwoPDCCH_r17 int64
		MaxNumOverlaps_r17 int64
	}
}

func (ie *FeatureSetDownlink_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ScalingFactor_1024QAM_FR1_r17 != nil, ie.TimeDurationForQCL_v1710 != nil, ie.Sfn_SchemeA_r17 != nil, ie.Sfn_SchemeA_PDCCH_Only_r17 != nil, ie.Sfn_SchemeA_DynamicSwitching_r17 != nil, ie.Sfn_SchemeA_PDSCH_Only_r17 != nil, ie.Sfn_SchemeB_r17 != nil, ie.Sfn_SchemeB_DynamicSwitching_r17 != nil, ie.Sfn_SchemeB_PDSCH_Only_r17 != nil, ie.MTRP_PDCCH_Case2_1SpanGap_r17 != nil, ie.MTRP_PDCCH_LegacyMonitoring_r17 != nil, ie.MTRP_PDCCH_MultiDCI_MultiTRP_r17 != nil, ie.DynamicMulticastPCell_r17 != nil, ie.MTRP_PDCCH_Repetition_r17 != nil}); err != nil {
		return err
	}

	// 2. scalingFactor-1024QAM-FR1-r17: enumerated
	{
		if ie.ScalingFactor_1024QAM_FR1_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ScalingFactor_1024QAM_FR1_r17, featureSetDownlinkV1700ScalingFactor1024QAMFR1R17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. timeDurationForQCL-v1710: sequence
	{
		if ie.TimeDurationForQCL_v1710 != nil {
			c := ie.TimeDurationForQCL_v1710
			featureSetDownlinkV1700TimeDurationForQCLV1710Seq := e.NewSequenceEncoder(featureSetDownlinkV1700TimeDurationForQCLV1710Constraints)
			if err := featureSetDownlinkV1700TimeDurationForQCLV1710Seq.EncodePreamble([]bool{c.Scs_480kHz != nil, c.Scs_960kHz != nil}); err != nil {
				return err
			}
			if c.Scs_480kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_480kHz), featureSetDownlinkV1700TimeDurationForQCLV1710Scs480kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_960kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_960kHz), featureSetDownlinkV1700TimeDurationForQCLV1710Scs960kHzConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. sfn-SchemeA-r17: enumerated
	{
		if ie.Sfn_SchemeA_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SchemeA_r17, featureSetDownlinkV1700SfnSchemeAR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sfn-SchemeA-PDCCH-only-r17: enumerated
	{
		if ie.Sfn_SchemeA_PDCCH_Only_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SchemeA_PDCCH_Only_r17, featureSetDownlinkV1700SfnSchemeAPDCCHOnlyR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sfn-SchemeA-DynamicSwitching-r17: enumerated
	{
		if ie.Sfn_SchemeA_DynamicSwitching_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SchemeA_DynamicSwitching_r17, featureSetDownlinkV1700SfnSchemeADynamicSwitchingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sfn-SchemeA-PDSCH-only-r17: enumerated
	{
		if ie.Sfn_SchemeA_PDSCH_Only_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SchemeA_PDSCH_Only_r17, featureSetDownlinkV1700SfnSchemeAPDSCHOnlyR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sfn-SchemeB-r17: enumerated
	{
		if ie.Sfn_SchemeB_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SchemeB_r17, featureSetDownlinkV1700SfnSchemeBR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sfn-SchemeB-DynamicSwitching-r17: enumerated
	{
		if ie.Sfn_SchemeB_DynamicSwitching_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SchemeB_DynamicSwitching_r17, featureSetDownlinkV1700SfnSchemeBDynamicSwitchingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sfn-SchemeB-PDSCH-only-r17: enumerated
	{
		if ie.Sfn_SchemeB_PDSCH_Only_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sfn_SchemeB_PDSCH_Only_r17, featureSetDownlinkV1700SfnSchemeBPDSCHOnlyR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. mTRP-PDCCH-Case2-1SpanGap-r17: sequence
	{
		if ie.MTRP_PDCCH_Case2_1SpanGap_r17 != nil {
			c := ie.MTRP_PDCCH_Case2_1SpanGap_r17
			featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq := e.NewSequenceEncoder(featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Constraints)
			if err := featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq.EncodePreamble([]bool{c.Scs_15kHz_r17 != nil, c.Scs_30kHz_r17 != nil, c.Scs_60kHz_r17 != nil, c.Scs_120kHz_r17 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r17 != nil {
				if err := c.Scs_15kHz_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r17 != nil {
				if err := c.Scs_30kHz_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r17 != nil {
				if err := c.Scs_60kHz_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_120kHz_r17 != nil {
				if err := c.Scs_120kHz_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 12. mTRP-PDCCH-legacyMonitoring-r17: sequence
	{
		if ie.MTRP_PDCCH_LegacyMonitoring_r17 != nil {
			c := ie.MTRP_PDCCH_LegacyMonitoring_r17
			featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Seq := e.NewSequenceEncoder(featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Constraints)
			if err := featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Seq.EncodePreamble([]bool{c.Scs_15kHz_r17 != nil, c.Scs_30kHz_r17 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r17 != nil {
				if err := c.Scs_15kHz_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r17 != nil {
				if err := c.Scs_30kHz_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. mTRP-PDCCH-multiDCI-multiTRP-r17: enumerated
	{
		if ie.MTRP_PDCCH_MultiDCI_MultiTRP_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MTRP_PDCCH_MultiDCI_MultiTRP_r17, featureSetDownlinkV1700MTRPPDCCHMultiDCIMultiTRPR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. dynamicMulticastPCell-r17: enumerated
	{
		if ie.DynamicMulticastPCell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DynamicMulticastPCell_r17, featureSetDownlinkV1700DynamicMulticastPCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 15. mTRP-PDCCH-Repetition-r17: sequence
	{
		if ie.MTRP_PDCCH_Repetition_r17 != nil {
			c := ie.MTRP_PDCCH_Repetition_r17
			if err := e.EncodeInteger(c.NumBD_TwoPDCCH_r17, per.Constrained(2, 3)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumOverlaps_r17, featureSetDownlinkV1700MTRPPDCCHRepetitionR17MaxNumOverlapsR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. scalingFactor-1024QAM-FR1-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700ScalingFactor1024QAMFR1R17Constraints)
			if err != nil {
				return err
			}
			ie.ScalingFactor_1024QAM_FR1_r17 = &idx
		}
	}

	// 3. timeDurationForQCL-v1710: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.TimeDurationForQCL_v1710 = &struct {
				Scs_480kHz *int64
				Scs_960kHz *int64
			}{}
			c := ie.TimeDurationForQCL_v1710
			featureSetDownlinkV1700TimeDurationForQCLV1710Seq := d.NewSequenceDecoder(featureSetDownlinkV1700TimeDurationForQCLV1710Constraints)
			if err := featureSetDownlinkV1700TimeDurationForQCLV1710Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1700TimeDurationForQCLV1710Seq.IsComponentPresent(0) {
				c.Scs_480kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1700TimeDurationForQCLV1710Scs480kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz) = v
			}
			if featureSetDownlinkV1700TimeDurationForQCLV1710Seq.IsComponentPresent(1) {
				c.Scs_960kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1700TimeDurationForQCLV1710Scs960kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz) = v
			}
		}
	}

	// 4. sfn-SchemeA-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700SfnSchemeAR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SchemeA_r17 = &idx
		}
	}

	// 5. sfn-SchemeA-PDCCH-only-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700SfnSchemeAPDCCHOnlyR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SchemeA_PDCCH_Only_r17 = &idx
		}
	}

	// 6. sfn-SchemeA-DynamicSwitching-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700SfnSchemeADynamicSwitchingR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SchemeA_DynamicSwitching_r17 = &idx
		}
	}

	// 7. sfn-SchemeA-PDSCH-only-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700SfnSchemeAPDSCHOnlyR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SchemeA_PDSCH_Only_r17 = &idx
		}
	}

	// 8. sfn-SchemeB-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700SfnSchemeBR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SchemeB_r17 = &idx
		}
	}

	// 9. sfn-SchemeB-DynamicSwitching-r17: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700SfnSchemeBDynamicSwitchingR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SchemeB_DynamicSwitching_r17 = &idx
		}
	}

	// 10. sfn-SchemeB-PDSCH-only-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700SfnSchemeBPDSCHOnlyR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SchemeB_PDSCH_Only_r17 = &idx
		}
	}

	// 11. mTRP-PDCCH-Case2-1SpanGap-r17: sequence
	{
		if seq.IsComponentPresent(9) {
			ie.MTRP_PDCCH_Case2_1SpanGap_r17 = &struct {
				Scs_15kHz_r17  *PDCCH_RepetitionParameters_r17
				Scs_30kHz_r17  *PDCCH_RepetitionParameters_r17
				Scs_60kHz_r17  *PDCCH_RepetitionParameters_r17
				Scs_120kHz_r17 *PDCCH_RepetitionParameters_r17
			}{}
			c := ie.MTRP_PDCCH_Case2_1SpanGap_r17
			featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq := d.NewSequenceDecoder(featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Constraints)
			if err := featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r17 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_15kHz_r17).Decode(d); err != nil {
					return err
				}
			}
			if featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r17 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_30kHz_r17).Decode(d); err != nil {
					return err
				}
			}
			if featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r17 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_60kHz_r17).Decode(d); err != nil {
					return err
				}
			}
			if featureSetDownlinkV1700MTRPPDCCHCase21SpanGapR17Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r17 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_120kHz_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. mTRP-PDCCH-legacyMonitoring-r17: sequence
	{
		if seq.IsComponentPresent(10) {
			ie.MTRP_PDCCH_LegacyMonitoring_r17 = &struct {
				Scs_15kHz_r17 *PDCCH_RepetitionParameters_r17
				Scs_30kHz_r17 *PDCCH_RepetitionParameters_r17
			}{}
			c := ie.MTRP_PDCCH_LegacyMonitoring_r17
			featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Seq := d.NewSequenceDecoder(featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Constraints)
			if err := featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r17 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_15kHz_r17).Decode(d); err != nil {
					return err
				}
			}
			if featureSetDownlinkV1700MTRPPDCCHLegacyMonitoringR17Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r17 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_30kHz_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. mTRP-PDCCH-multiDCI-multiTRP-r17: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700MTRPPDCCHMultiDCIMultiTRPR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PDCCH_MultiDCI_MultiTRP_r17 = &idx
		}
	}

	// 14. dynamicMulticastPCell-r17: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1700DynamicMulticastPCellR17Constraints)
			if err != nil {
				return err
			}
			ie.DynamicMulticastPCell_r17 = &idx
		}
	}

	// 15. mTRP-PDCCH-Repetition-r17: sequence
	{
		if seq.IsComponentPresent(13) {
			ie.MTRP_PDCCH_Repetition_r17 = &struct {
				NumBD_TwoPDCCH_r17 int64
				MaxNumOverlaps_r17 int64
			}{}
			c := ie.MTRP_PDCCH_Repetition_r17
			{
				v, err := d.DecodeInteger(per.Constrained(2, 3))
				if err != nil {
					return err
				}
				c.NumBD_TwoPDCCH_r17 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1700MTRPPDCCHRepetitionR17MaxNumOverlapsR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOverlaps_r17 = v
			}
		}
	}

	return nil
}
