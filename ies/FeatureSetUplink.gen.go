// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetUplink (line 20088).

var featureSetUplinkConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSetListPerUplinkCC"},
		{Name: "scalingFactor", Optional: true},
		{Name: "dummy3", Optional: true},
		{Name: "intraBandFreqSeparationUL", Optional: true},
		{Name: "searchSpaceSharingCA-UL", Optional: true},
		{Name: "dummy1", Optional: true},
		{Name: "supportedSRS-Resources", Optional: true},
		{Name: "twoPUCCH-Group", Optional: true},
		{Name: "dynamicSwitchSUL", Optional: true},
		{Name: "simultaneousTxSUL-NonSUL", Optional: true},
		{Name: "pusch-ProcessingType1-DifferentTB-PerSlot", Optional: true},
		{Name: "dummy2", Optional: true},
	},
}

var featureSetUplinkFeatureSetListPerUplinkCCConstraints = per.SizeRange(1, common.MaxNrofServingCells)

const (
	FeatureSetUplink_ScalingFactor_F0p4  = 0
	FeatureSetUplink_ScalingFactor_F0p75 = 1
	FeatureSetUplink_ScalingFactor_F0p8  = 2
)

var featureSetUplinkScalingFactorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_ScalingFactor_F0p4, FeatureSetUplink_ScalingFactor_F0p75, FeatureSetUplink_ScalingFactor_F0p8},
}

const (
	FeatureSetUplink_Dummy3_Supported = 0
)

var featureSetUplinkDummy3Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_Dummy3_Supported},
}

const (
	FeatureSetUplink_SearchSpaceSharingCA_UL_Supported = 0
)

var featureSetUplinkSearchSpaceSharingCAULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_SearchSpaceSharingCA_UL_Supported},
}

const (
	FeatureSetUplink_TwoPUCCH_Group_Supported = 0
)

var featureSetUplinkTwoPUCCHGroupConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_TwoPUCCH_Group_Supported},
}

const (
	FeatureSetUplink_DynamicSwitchSUL_Supported = 0
)

var featureSetUplinkDynamicSwitchSULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_DynamicSwitchSUL_Supported},
}

const (
	FeatureSetUplink_SimultaneousTxSUL_NonSUL_Supported = 0
)

var featureSetUplinkSimultaneousTxSULNonSULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_SimultaneousTxSUL_NonSUL_Supported},
}

var featureSetUplinkPuschProcessingType1DifferentTBPerSlotConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto2 = 0
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto4 = 1
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto7 = 2
)

var featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto2, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto4, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_15kHz_Upto7},
}

const (
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto2 = 0
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto4 = 1
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto7 = 2
)

var featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto2, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto4, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_30kHz_Upto7},
}

const (
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto2 = 0
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto4 = 1
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto7 = 2
)

var featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto2, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto4, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_60kHz_Upto7},
}

const (
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto2 = 0
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto4 = 1
	FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto7 = 2
)

var featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto2, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto4, FeatureSetUplink_Pusch_ProcessingType1_DifferentTB_PerSlot_Scs_120kHz_Upto7},
}

type FeatureSetUplink struct {
	FeatureSetListPerUplinkCC                 []FeatureSetUplinkPerCC_Id
	ScalingFactor                             *int64
	Dummy3                                    *int64
	IntraBandFreqSeparationUL                 *FreqSeparationClass
	SearchSpaceSharingCA_UL                   *int64
	Dummy1                                    *DummyI
	SupportedSRS_Resources                    *SRS_Resources
	TwoPUCCH_Group                            *int64
	DynamicSwitchSUL                          *int64
	SimultaneousTxSUL_NonSUL                  *int64
	Pusch_ProcessingType1_DifferentTB_PerSlot *struct {
		Scs_15kHz  *int64
		Scs_30kHz  *int64
		Scs_60kHz  *int64
		Scs_120kHz *int64
	}
	Dummy2 *DummyF
}

func (ie *FeatureSetUplink) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ScalingFactor != nil, ie.Dummy3 != nil, ie.IntraBandFreqSeparationUL != nil, ie.SearchSpaceSharingCA_UL != nil, ie.Dummy1 != nil, ie.SupportedSRS_Resources != nil, ie.TwoPUCCH_Group != nil, ie.DynamicSwitchSUL != nil, ie.SimultaneousTxSUL_NonSUL != nil, ie.Pusch_ProcessingType1_DifferentTB_PerSlot != nil, ie.Dummy2 != nil}); err != nil {
		return err
	}

	// 2. featureSetListPerUplinkCC: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(featureSetUplinkFeatureSetListPerUplinkCCConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.FeatureSetListPerUplinkCC))); err != nil {
			return err
		}
		for i := range ie.FeatureSetListPerUplinkCC {
			if err := ie.FeatureSetListPerUplinkCC[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. scalingFactor: enumerated
	{
		if ie.ScalingFactor != nil {
			if err := e.EncodeEnumerated(*ie.ScalingFactor, featureSetUplinkScalingFactorConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dummy3: enumerated
	{
		if ie.Dummy3 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy3, featureSetUplinkDummy3Constraints); err != nil {
				return err
			}
		}
	}

	// 5. intraBandFreqSeparationUL: ref
	{
		if ie.IntraBandFreqSeparationUL != nil {
			if err := ie.IntraBandFreqSeparationUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. searchSpaceSharingCA-UL: enumerated
	{
		if ie.SearchSpaceSharingCA_UL != nil {
			if err := e.EncodeEnumerated(*ie.SearchSpaceSharingCA_UL, featureSetUplinkSearchSpaceSharingCAULConstraints); err != nil {
				return err
			}
		}
	}

	// 7. dummy1: ref
	{
		if ie.Dummy1 != nil {
			if err := ie.Dummy1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. supportedSRS-Resources: ref
	{
		if ie.SupportedSRS_Resources != nil {
			if err := ie.SupportedSRS_Resources.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. twoPUCCH-Group: enumerated
	{
		if ie.TwoPUCCH_Group != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Group, featureSetUplinkTwoPUCCHGroupConstraints); err != nil {
				return err
			}
		}
	}

	// 10. dynamicSwitchSUL: enumerated
	{
		if ie.DynamicSwitchSUL != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSwitchSUL, featureSetUplinkDynamicSwitchSULConstraints); err != nil {
				return err
			}
		}
	}

	// 11. simultaneousTxSUL-NonSUL: enumerated
	{
		if ie.SimultaneousTxSUL_NonSUL != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousTxSUL_NonSUL, featureSetUplinkSimultaneousTxSULNonSULConstraints); err != nil {
				return err
			}
		}
	}

	// 12. pusch-ProcessingType1-DifferentTB-PerSlot: sequence
	{
		if ie.Pusch_ProcessingType1_DifferentTB_PerSlot != nil {
			c := ie.Pusch_ProcessingType1_DifferentTB_PerSlot
			featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq := e.NewSequenceEncoder(featureSetUplinkPuschProcessingType1DifferentTBPerSlotConstraints)
			if err := featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz), featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs15kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz), featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs30kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz), featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs60kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz), featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs120kHzConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 13. dummy2: ref
	{
		if ie.Dummy2 != nil {
			if err := ie.Dummy2.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSetListPerUplinkCC: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(featureSetUplinkFeatureSetListPerUplinkCCConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.FeatureSetListPerUplinkCC = make([]FeatureSetUplinkPerCC_Id, n)
		for i := int64(0); i < n; i++ {
			if err := ie.FeatureSetListPerUplinkCC[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. scalingFactor: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkScalingFactorConstraints)
			if err != nil {
				return err
			}
			ie.ScalingFactor = &idx
		}
	}

	// 4. dummy3: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkDummy3Constraints)
			if err != nil {
				return err
			}
			ie.Dummy3 = &idx
		}
	}

	// 5. intraBandFreqSeparationUL: ref
	{
		if seq.IsComponentPresent(3) {
			ie.IntraBandFreqSeparationUL = new(FreqSeparationClass)
			if err := ie.IntraBandFreqSeparationUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. searchSpaceSharingCA-UL: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetUplinkSearchSpaceSharingCAULConstraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSharingCA_UL = &idx
		}
	}

	// 7. dummy1: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Dummy1 = new(DummyI)
			if err := ie.Dummy1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. supportedSRS-Resources: ref
	{
		if seq.IsComponentPresent(6) {
			ie.SupportedSRS_Resources = new(SRS_Resources)
			if err := ie.SupportedSRS_Resources.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. twoPUCCH-Group: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetUplinkTwoPUCCHGroupConstraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Group = &idx
		}
	}

	// 10. dynamicSwitchSUL: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(featureSetUplinkDynamicSwitchSULConstraints)
			if err != nil {
				return err
			}
			ie.DynamicSwitchSUL = &idx
		}
	}

	// 11. simultaneousTxSUL-NonSUL: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(featureSetUplinkSimultaneousTxSULNonSULConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousTxSUL_NonSUL = &idx
		}
	}

	// 12. pusch-ProcessingType1-DifferentTB-PerSlot: sequence
	{
		if seq.IsComponentPresent(10) {
			ie.Pusch_ProcessingType1_DifferentTB_PerSlot = &struct {
				Scs_15kHz  *int64
				Scs_30kHz  *int64
				Scs_60kHz  *int64
				Scs_120kHz *int64
			}{}
			c := ie.Pusch_ProcessingType1_DifferentTB_PerSlot
			featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq := d.NewSequenceDecoder(featureSetUplinkPuschProcessingType1DifferentTBPerSlotConstraints)
			if err := featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(0) {
				c.Scs_15kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs15kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz) = v
			}
			if featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(1) {
				c.Scs_30kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs30kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz) = v
			}
			if featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(2) {
				c.Scs_60kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs60kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz) = v
			}
			if featureSetUplinkPuschProcessingType1DifferentTBPerSlotSeq.IsComponentPresent(3) {
				c.Scs_120kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPuschProcessingType1DifferentTBPerSlotScs120kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz) = v
			}
		}
	}

	// 13. dummy2: ref
	{
		if seq.IsComponentPresent(11) {
			ie.Dummy2 = new(DummyF)
			if err := ie.Dummy2.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
