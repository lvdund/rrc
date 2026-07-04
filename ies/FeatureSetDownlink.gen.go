// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetDownlink (line 19447).

var featureSetDownlinkConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSetListPerDownlinkCC"},
		{Name: "intraBandFreqSeparationDL", Optional: true},
		{Name: "scalingFactor", Optional: true},
		{Name: "dummy8", Optional: true},
		{Name: "scellWithoutSSB", Optional: true},
		{Name: "csi-RS-MeasSCellWithoutSSB", Optional: true},
		{Name: "dummy1", Optional: true},
		{Name: "type1-3-CSS", Optional: true},
		{Name: "pdcch-MonitoringAnyOccasions", Optional: true},
		{Name: "dummy2", Optional: true},
		{Name: "ue-SpecificUL-DL-Assignment", Optional: true},
		{Name: "searchSpaceSharingCA-DL", Optional: true},
		{Name: "timeDurationForQCL", Optional: true},
		{Name: "pdsch-ProcessingType1-DifferentTB-PerSlot", Optional: true},
		{Name: "dummy3", Optional: true},
		{Name: "dummy4", Optional: true},
		{Name: "dummy5", Optional: true},
		{Name: "dummy6", Optional: true},
		{Name: "dummy7", Optional: true},
	},
}

var featureSetDownlinkFeatureSetListPerDownlinkCCConstraints = per.SizeRange(1, common.MaxNrofServingCells)

const (
	FeatureSetDownlink_ScalingFactor_F0p4  = 0
	FeatureSetDownlink_ScalingFactor_F0p75 = 1
	FeatureSetDownlink_ScalingFactor_F0p8  = 2
)

var featureSetDownlinkScalingFactorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_ScalingFactor_F0p4, FeatureSetDownlink_ScalingFactor_F0p75, FeatureSetDownlink_ScalingFactor_F0p8},
}

const (
	FeatureSetDownlink_Dummy8_Supported = 0
)

var featureSetDownlinkDummy8Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Dummy8_Supported},
}

const (
	FeatureSetDownlink_ScellWithoutSSB_Supported = 0
)

var featureSetDownlinkScellWithoutSSBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_ScellWithoutSSB_Supported},
}

const (
	FeatureSetDownlink_Csi_RS_MeasSCellWithoutSSB_Supported = 0
)

var featureSetDownlinkCsiRSMeasSCellWithoutSSBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Csi_RS_MeasSCellWithoutSSB_Supported},
}

const (
	FeatureSetDownlink_Dummy1_Supported = 0
)

var featureSetDownlinkDummy1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Dummy1_Supported},
}

const (
	FeatureSetDownlink_Type1_3_CSS_Supported = 0
)

var featureSetDownlinkType13CSSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Type1_3_CSS_Supported},
}

const (
	FeatureSetDownlink_Pdcch_MonitoringAnyOccasions_WithoutDCI_Gap = 0
	FeatureSetDownlink_Pdcch_MonitoringAnyOccasions_WithDCI_Gap    = 1
)

var featureSetDownlinkPdcchMonitoringAnyOccasionsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Pdcch_MonitoringAnyOccasions_WithoutDCI_Gap, FeatureSetDownlink_Pdcch_MonitoringAnyOccasions_WithDCI_Gap},
}

const (
	FeatureSetDownlink_Dummy2_Supported = 0
)

var featureSetDownlinkDummy2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Dummy2_Supported},
}

const (
	FeatureSetDownlink_Ue_SpecificUL_DL_Assignment_Supported = 0
)

var featureSetDownlinkUeSpecificULDLAssignmentConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Ue_SpecificUL_DL_Assignment_Supported},
}

const (
	FeatureSetDownlink_SearchSpaceSharingCA_DL_Supported = 0
)

var featureSetDownlinkSearchSpaceSharingCADLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_SearchSpaceSharingCA_DL_Supported},
}

var featureSetDownlinkDummy4Constraints = per.SizeRange(1, common.MaxNrofCodebooks)

var featureSetDownlinkDummy5Constraints = per.SizeRange(1, common.MaxNrofCodebooks)

var featureSetDownlinkDummy6Constraints = per.SizeRange(1, common.MaxNrofCodebooks)

var featureSetDownlinkDummy7Constraints = per.SizeRange(1, common.MaxNrofCodebooks)

var featureSetDownlinkTimeDurationForQCLConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	FeatureSetDownlink_TimeDurationForQCL_Scs_60kHz_S7  = 0
	FeatureSetDownlink_TimeDurationForQCL_Scs_60kHz_S14 = 1
	FeatureSetDownlink_TimeDurationForQCL_Scs_60kHz_S28 = 2
)

var featureSetDownlinkTimeDurationForQCLScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_TimeDurationForQCL_Scs_60kHz_S7, FeatureSetDownlink_TimeDurationForQCL_Scs_60kHz_S14, FeatureSetDownlink_TimeDurationForQCL_Scs_60kHz_S28},
}

const (
	FeatureSetDownlink_TimeDurationForQCL_Scs_120kHz_S14 = 0
	FeatureSetDownlink_TimeDurationForQCL_Scs_120kHz_S28 = 1
)

var featureSetDownlinkTimeDurationForQCLScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_TimeDurationForQCL_Scs_120kHz_S14, FeatureSetDownlink_TimeDurationForQCL_Scs_120kHz_S28},
}

var featureSetDownlinkPdschProcessingType1DifferentTBPerSlotConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto2 = 0
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto4 = 1
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto7 = 2
)

var featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto2, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto4, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto7},
}

const (
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto2 = 0
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto4 = 1
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto7 = 2
)

var featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto2, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto4, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto7},
}

const (
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto2 = 0
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto4 = 1
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto7 = 2
)

var featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto2, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto4, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto7},
}

const (
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto2 = 0
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto4 = 1
	FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto7 = 2
)

var featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto2, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto4, FeatureSetDownlink_Pdsch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto7},
}

type FeatureSetDownlink struct {
	FeatureSetListPerDownlinkCC  []FeatureSetDownlinkPerCC_Id
	IntraBandFreqSeparationDL    *FreqSeparationClass
	ScalingFactor                *int64
	Dummy8                       *int64
	ScellWithoutSSB              *int64
	Csi_RS_MeasSCellWithoutSSB   *int64
	Dummy1                       *int64
	Type1_3_CSS                  *int64
	Pdcch_MonitoringAnyOccasions *int64
	Dummy2                       *int64
	Ue_SpecificUL_DL_Assignment  *int64
	SearchSpaceSharingCA_DL      *int64
	TimeDurationForQCL           *struct {
		Scs_60kHz  *int64
		Scs_120kHz *int64
	}
	Pdsch_ProcessingType1_DifferentTB_PerSlot *struct {
		Scs_15kHz  *int64
		Scs_30kHz  *int64
		Scs_60kHz  *int64
		Scs_120kHz *int64
	}
	Dummy3 *DummyA
	Dummy4 []DummyB
	Dummy5 []DummyC
	Dummy6 []DummyD
	Dummy7 []DummyE
}

func (ie *FeatureSetDownlink) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraBandFreqSeparationDL != nil, ie.ScalingFactor != nil, ie.Dummy8 != nil, ie.ScellWithoutSSB != nil, ie.Csi_RS_MeasSCellWithoutSSB != nil, ie.Dummy1 != nil, ie.Type1_3_CSS != nil, ie.Pdcch_MonitoringAnyOccasions != nil, ie.Dummy2 != nil, ie.Ue_SpecificUL_DL_Assignment != nil, ie.SearchSpaceSharingCA_DL != nil, ie.TimeDurationForQCL != nil, ie.Pdsch_ProcessingType1_DifferentTB_PerSlot != nil, ie.Dummy3 != nil, ie.Dummy4 != nil, ie.Dummy5 != nil, ie.Dummy6 != nil, ie.Dummy7 != nil}); err != nil {
		return err
	}

	// 2. featureSetListPerDownlinkCC: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(featureSetDownlinkFeatureSetListPerDownlinkCCConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.FeatureSetListPerDownlinkCC))); err != nil {
			return err
		}
		for i := range ie.FeatureSetListPerDownlinkCC {
			if err := ie.FeatureSetListPerDownlinkCC[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. intraBandFreqSeparationDL: ref
	{
		if ie.IntraBandFreqSeparationDL != nil {
			if err := ie.IntraBandFreqSeparationDL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. scalingFactor: enumerated
	{
		if ie.ScalingFactor != nil {
			if err := e.EncodeEnumerated(*ie.ScalingFactor, featureSetDownlinkScalingFactorConstraints); err != nil {
				return err
			}
		}
	}

	// 5. dummy8: enumerated
	{
		if ie.Dummy8 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy8, featureSetDownlinkDummy8Constraints); err != nil {
				return err
			}
		}
	}

	// 6. scellWithoutSSB: enumerated
	{
		if ie.ScellWithoutSSB != nil {
			if err := e.EncodeEnumerated(*ie.ScellWithoutSSB, featureSetDownlinkScellWithoutSSBConstraints); err != nil {
				return err
			}
		}
	}

	// 7. csi-RS-MeasSCellWithoutSSB: enumerated
	{
		if ie.Csi_RS_MeasSCellWithoutSSB != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RS_MeasSCellWithoutSSB, featureSetDownlinkCsiRSMeasSCellWithoutSSBConstraints); err != nil {
				return err
			}
		}
	}

	// 8. dummy1: enumerated
	{
		if ie.Dummy1 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy1, featureSetDownlinkDummy1Constraints); err != nil {
				return err
			}
		}
	}

	// 9. type1-3-CSS: enumerated
	{
		if ie.Type1_3_CSS != nil {
			if err := e.EncodeEnumerated(*ie.Type1_3_CSS, featureSetDownlinkType13CSSConstraints); err != nil {
				return err
			}
		}
	}

	// 10. pdcch-MonitoringAnyOccasions: enumerated
	{
		if ie.Pdcch_MonitoringAnyOccasions != nil {
			if err := e.EncodeEnumerated(*ie.Pdcch_MonitoringAnyOccasions, featureSetDownlinkPdcchMonitoringAnyOccasionsConstraints); err != nil {
				return err
			}
		}
	}

	// 11. dummy2: enumerated
	{
		if ie.Dummy2 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy2, featureSetDownlinkDummy2Constraints); err != nil {
				return err
			}
		}
	}

	// 12. ue-SpecificUL-DL-Assignment: enumerated
	{
		if ie.Ue_SpecificUL_DL_Assignment != nil {
			if err := e.EncodeEnumerated(*ie.Ue_SpecificUL_DL_Assignment, featureSetDownlinkUeSpecificULDLAssignmentConstraints); err != nil {
				return err
			}
		}
	}

	// 13. searchSpaceSharingCA-DL: enumerated
	{
		if ie.SearchSpaceSharingCA_DL != nil {
			if err := e.EncodeEnumerated(*ie.SearchSpaceSharingCA_DL, featureSetDownlinkSearchSpaceSharingCADLConstraints); err != nil {
				return err
			}
		}
	}

	// 14. timeDurationForQCL: sequence
	{
		if ie.TimeDurationForQCL != nil {
			c := ie.TimeDurationForQCL
			featureSetDownlinkTimeDurationForQCLSeq := e.NewSequenceEncoder(featureSetDownlinkTimeDurationForQCLConstraints)
			if err := featureSetDownlinkTimeDurationForQCLSeq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
				return err
			}
			if c.Scs_60kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz), featureSetDownlinkTimeDurationForQCLScs60kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz), featureSetDownlinkTimeDurationForQCLScs120kHzConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 15. pdsch-ProcessingType1-DifferentTB-PerSlot: sequence
	{
		if ie.Pdsch_ProcessingType1_DifferentTB_PerSlot != nil {
			c := ie.Pdsch_ProcessingType1_DifferentTB_PerSlot
			featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq := e.NewSequenceEncoder(featureSetDownlinkPdschProcessingType1DifferentTBPerSlotConstraints)
			if err := featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz), featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs15kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz), featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs30kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz), featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs60kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz), featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs120kHzConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 16. dummy3: ref
	{
		if ie.Dummy3 != nil {
			if err := ie.Dummy3.Encode(e); err != nil {
				return err
			}
		}
	}

	// 17. dummy4: sequence-of
	{
		if ie.Dummy4 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkDummy4Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Dummy4))); err != nil {
				return err
			}
			for i := range ie.Dummy4 {
				if err := ie.Dummy4[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 18. dummy5: sequence-of
	{
		if ie.Dummy5 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkDummy5Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Dummy5))); err != nil {
				return err
			}
			for i := range ie.Dummy5 {
				if err := ie.Dummy5[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 19. dummy6: sequence-of
	{
		if ie.Dummy6 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkDummy6Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Dummy6))); err != nil {
				return err
			}
			for i := range ie.Dummy6 {
				if err := ie.Dummy6[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 20. dummy7: sequence-of
	{
		if ie.Dummy7 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkDummy7Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Dummy7))); err != nil {
				return err
			}
			for i := range ie.Dummy7 {
				if err := ie.Dummy7[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSetListPerDownlinkCC: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(featureSetDownlinkFeatureSetListPerDownlinkCCConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.FeatureSetListPerDownlinkCC = make([]FeatureSetDownlinkPerCC_Id, n)
		for i := int64(0); i < n; i++ {
			if err := ie.FeatureSetListPerDownlinkCC[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. intraBandFreqSeparationDL: ref
	{
		if seq.IsComponentPresent(1) {
			ie.IntraBandFreqSeparationDL = new(FreqSeparationClass)
			if err := ie.IntraBandFreqSeparationDL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. scalingFactor: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkScalingFactorConstraints)
			if err != nil {
				return err
			}
			ie.ScalingFactor = &idx
		}
	}

	// 5. dummy8: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkDummy8Constraints)
			if err != nil {
				return err
			}
			ie.Dummy8 = &idx
		}
	}

	// 6. scellWithoutSSB: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkScellWithoutSSBConstraints)
			if err != nil {
				return err
			}
			ie.ScellWithoutSSB = &idx
		}
	}

	// 7. csi-RS-MeasSCellWithoutSSB: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkCsiRSMeasSCellWithoutSSBConstraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_MeasSCellWithoutSSB = &idx
		}
	}

	// 8. dummy1: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkDummy1Constraints)
			if err != nil {
				return err
			}
			ie.Dummy1 = &idx
		}
	}

	// 9. type1-3-CSS: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkType13CSSConstraints)
			if err != nil {
				return err
			}
			ie.Type1_3_CSS = &idx
		}
	}

	// 10. pdcch-MonitoringAnyOccasions: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPdcchMonitoringAnyOccasionsConstraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringAnyOccasions = &idx
		}
	}

	// 11. dummy2: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkDummy2Constraints)
			if err != nil {
				return err
			}
			ie.Dummy2 = &idx
		}
	}

	// 12. ue-SpecificUL-DL-Assignment: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkUeSpecificULDLAssignmentConstraints)
			if err != nil {
				return err
			}
			ie.Ue_SpecificUL_DL_Assignment = &idx
		}
	}

	// 13. searchSpaceSharingCA-DL: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkSearchSpaceSharingCADLConstraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSharingCA_DL = &idx
		}
	}

	// 14. timeDurationForQCL: sequence
	{
		if seq.IsComponentPresent(12) {
			ie.TimeDurationForQCL = &struct {
				Scs_60kHz  *int64
				Scs_120kHz *int64
			}{}
			c := ie.TimeDurationForQCL
			featureSetDownlinkTimeDurationForQCLSeq := d.NewSequenceDecoder(featureSetDownlinkTimeDurationForQCLConstraints)
			if err := featureSetDownlinkTimeDurationForQCLSeq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkTimeDurationForQCLSeq.IsComponentPresent(0) {
				c.Scs_60kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkTimeDurationForQCLScs60kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz) = v
			}
			if featureSetDownlinkTimeDurationForQCLSeq.IsComponentPresent(1) {
				c.Scs_120kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkTimeDurationForQCLScs120kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz) = v
			}
		}
	}

	// 15. pdsch-ProcessingType1-DifferentTB-PerSlot: sequence
	{
		if seq.IsComponentPresent(13) {
			ie.Pdsch_ProcessingType1_DifferentTB_PerSlot = &struct {
				Scs_15kHz  *int64
				Scs_30kHz  *int64
				Scs_60kHz  *int64
				Scs_120kHz *int64
			}{}
			c := ie.Pdsch_ProcessingType1_DifferentTB_PerSlot
			featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq := d.NewSequenceDecoder(featureSetDownlinkPdschProcessingType1DifferentTBPerSlotConstraints)
			if err := featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(0) {
				c.Scs_15kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs15kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz) = v
			}
			if featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(1) {
				c.Scs_30kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs30kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz) = v
			}
			if featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(2) {
				c.Scs_60kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs60kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz) = v
			}
			if featureSetDownlinkPdschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(3) {
				c.Scs_120kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkPdschProcessingType1DifferentTBPerSlotScs120kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz) = v
			}
		}
	}

	// 16. dummy3: ref
	{
		if seq.IsComponentPresent(14) {
			ie.Dummy3 = new(DummyA)
			if err := ie.Dummy3.Decode(d); err != nil {
				return err
			}
		}
	}

	// 17. dummy4: sequence-of
	{
		if seq.IsComponentPresent(15) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkDummy4Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dummy4 = make([]DummyB, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dummy4[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 18. dummy5: sequence-of
	{
		if seq.IsComponentPresent(16) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkDummy5Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dummy5 = make([]DummyC, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dummy5[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 19. dummy6: sequence-of
	{
		if seq.IsComponentPresent(17) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkDummy6Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dummy6 = make([]DummyD, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dummy6[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 20. dummy7: sequence-of
	{
		if seq.IsComponentPresent(18) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkDummy7Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dummy7 = make([]DummyE, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dummy7[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
