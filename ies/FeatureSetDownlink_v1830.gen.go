// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetDownlink-v1830 (line 19747).

var featureSetDownlinkV1830Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-RACH-AffectedBands-TargetBandList-r18", Optional: true},
		{Name: "pdcch-RACH-SwitchingTime-TargetBandList-r18", Optional: true},
		{Name: "pdcch-RACH-PrepTime-TargetBandList-r18", Optional: true},
	},
}

var featureSetDownlinkV1830PdcchRACHAffectedBandsTargetBandListR18Constraints = per.SizeRange(1, common.MaxBandsMRDC)

var featureSetDownlinkV1830PdcchRACHSwitchingTimeTargetBandListR18Constraints = per.SizeRange(1, common.MaxBandsMRDC)

var featureSetDownlinkV1830PdcchRACHPrepTimeTargetBandListR18Constraints = per.SizeRange(1, common.MaxBandsMRDC)

const (
	FeatureSetDownlink_v1830_Pdcch_RACH_AffectedBands_TargetBandList_r18_Elem_NoInterruption = 0
	FeatureSetDownlink_v1830_Pdcch_RACH_AffectedBands_TargetBandList_r18_Elem_Interruption   = 1
)

var featureSetDownlinkV1830PdcchRACHAffectedBandsTargetBandListR18ElemConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1830_Pdcch_RACH_AffectedBands_TargetBandList_r18_Elem_NoInterruption, FeatureSetDownlink_v1830_Pdcch_RACH_AffectedBands_TargetBandList_r18_Elem_Interruption},
}

const (
	FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms0          = 0
	FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms0dot25     = 1
	FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms0dot5      = 2
	FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms1          = 3
	FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms2          = 4
	FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_NotSupported = 5
)

var featureSetDownlinkV1830PdcchRACHSwitchingTimeTargetBandListR18ElemConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms0, FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms0dot25, FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms0dot5, FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms1, FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_Ms2, FeatureSetDownlink_v1830_Pdcch_RACH_SwitchingTime_TargetBandList_r18_Elem_NotSupported},
}

const (
	FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms1          = 0
	FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms3          = 1
	FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms5          = 2
	FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms10         = 3
	FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_NotSupported = 4
)

var featureSetDownlinkV1830PdcchRACHPrepTimeTargetBandListR18ElemConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms1, FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms3, FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms5, FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_Ms10, FeatureSetDownlink_v1830_Pdcch_RACH_PrepTime_TargetBandList_r18_Elem_NotSupported},
}

type FeatureSetDownlink_v1830 struct {
	Pdcch_RACH_AffectedBands_TargetBandList_r18 []int64
	Pdcch_RACH_SwitchingTime_TargetBandList_r18 []int64
	Pdcch_RACH_PrepTime_TargetBandList_r18      []int64
}

func (ie *FeatureSetDownlink_v1830) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1830Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_RACH_AffectedBands_TargetBandList_r18 != nil, ie.Pdcch_RACH_SwitchingTime_TargetBandList_r18 != nil, ie.Pdcch_RACH_PrepTime_TargetBandList_r18 != nil}); err != nil {
		return err
	}

	// 2. pdcch-RACH-AffectedBands-TargetBandList-r18: sequence-of
	{
		if ie.Pdcch_RACH_AffectedBands_TargetBandList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkV1830PdcchRACHAffectedBandsTargetBandListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_RACH_AffectedBands_TargetBandList_r18))); err != nil {
				return err
			}
			for i := range ie.Pdcch_RACH_AffectedBands_TargetBandList_r18 {
				if err := e.EncodeEnumerated(ie.Pdcch_RACH_AffectedBands_TargetBandList_r18[i], featureSetDownlinkV1830PdcchRACHAffectedBandsTargetBandListR18ElemConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. pdcch-RACH-SwitchingTime-TargetBandList-r18: sequence-of
	{
		if ie.Pdcch_RACH_SwitchingTime_TargetBandList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkV1830PdcchRACHSwitchingTimeTargetBandListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_RACH_SwitchingTime_TargetBandList_r18))); err != nil {
				return err
			}
			for i := range ie.Pdcch_RACH_SwitchingTime_TargetBandList_r18 {
				if err := e.EncodeEnumerated(ie.Pdcch_RACH_SwitchingTime_TargetBandList_r18[i], featureSetDownlinkV1830PdcchRACHSwitchingTimeTargetBandListR18ElemConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. pdcch-RACH-PrepTime-TargetBandList-r18: sequence-of
	{
		if ie.Pdcch_RACH_PrepTime_TargetBandList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkV1830PdcchRACHPrepTimeTargetBandListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_RACH_PrepTime_TargetBandList_r18))); err != nil {
				return err
			}
			for i := range ie.Pdcch_RACH_PrepTime_TargetBandList_r18 {
				if err := e.EncodeEnumerated(ie.Pdcch_RACH_PrepTime_TargetBandList_r18[i], featureSetDownlinkV1830PdcchRACHPrepTimeTargetBandListR18ElemConstraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1830) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1830Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcch-RACH-AffectedBands-TargetBandList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkV1830PdcchRACHAffectedBandsTargetBandListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_RACH_AffectedBands_TargetBandList_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeEnumerated(featureSetDownlinkV1830PdcchRACHAffectedBandsTargetBandListR18ElemConstraints)
				if err != nil {
					return err
				}
				ie.Pdcch_RACH_AffectedBands_TargetBandList_r18[i] = v
			}
		}
	}

	// 3. pdcch-RACH-SwitchingTime-TargetBandList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkV1830PdcchRACHSwitchingTimeTargetBandListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_RACH_SwitchingTime_TargetBandList_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeEnumerated(featureSetDownlinkV1830PdcchRACHSwitchingTimeTargetBandListR18ElemConstraints)
				if err != nil {
					return err
				}
				ie.Pdcch_RACH_SwitchingTime_TargetBandList_r18[i] = v
			}
		}
	}

	// 4. pdcch-RACH-PrepTime-TargetBandList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkV1830PdcchRACHPrepTimeTargetBandListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_RACH_PrepTime_TargetBandList_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeEnumerated(featureSetDownlinkV1830PdcchRACHPrepTimeTargetBandListR18ElemConstraints)
				if err != nil {
					return err
				}
				ie.Pdcch_RACH_PrepTime_TargetBandList_r18[i] = v
			}
		}
	}

	return nil
}
