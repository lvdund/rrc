// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PropDelayDiffReportConfig-r17 (line 26491).

var propDelayDiffReportConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "threshPropDelayDiff-r17", Optional: true},
		{Name: "neighCellInfoList-r17", Optional: true},
	},
}

const (
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms0dot5 = 0
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms1     = 1
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms2     = 2
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms3     = 3
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms4     = 4
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms5     = 5
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms6     = 6
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms7     = 7
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms8     = 8
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms9     = 9
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms10    = 10
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare5  = 11
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare4  = 12
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare3  = 13
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare2  = 14
	PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare1  = 15
)

var propDelayDiffReportConfigR17ThreshPropDelayDiffR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms0dot5, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms1, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms2, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms3, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms4, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms5, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms6, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms7, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms8, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms9, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Ms10, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare5, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare4, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare3, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare2, PropDelayDiffReportConfig_r17_ThreshPropDelayDiff_r17_Spare1},
}

var propDelayDiffReportConfigR17NeighCellInfoListR17Constraints = per.SizeRange(1, common.MaxCellNTN_r17)

type PropDelayDiffReportConfig_r17 struct {
	ThreshPropDelayDiff_r17 *int64
	NeighCellInfoList_r17   []NeighbourCellInfo_r17
}

func (ie *PropDelayDiffReportConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(propDelayDiffReportConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ThreshPropDelayDiff_r17 != nil, ie.NeighCellInfoList_r17 != nil}); err != nil {
		return err
	}

	// 2. threshPropDelayDiff-r17: enumerated
	{
		if ie.ThreshPropDelayDiff_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ThreshPropDelayDiff_r17, propDelayDiffReportConfigR17ThreshPropDelayDiffR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. neighCellInfoList-r17: sequence-of
	{
		if ie.NeighCellInfoList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(propDelayDiffReportConfigR17NeighCellInfoListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NeighCellInfoList_r17))); err != nil {
				return err
			}
			for i := range ie.NeighCellInfoList_r17 {
				if err := ie.NeighCellInfoList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *PropDelayDiffReportConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(propDelayDiffReportConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. threshPropDelayDiff-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(propDelayDiffReportConfigR17ThreshPropDelayDiffR17Constraints)
			if err != nil {
				return err
			}
			ie.ThreshPropDelayDiff_r17 = &idx
		}
	}

	// 3. neighCellInfoList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(propDelayDiffReportConfigR17NeighCellInfoListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NeighCellInfoList_r17 = make([]NeighbourCellInfo_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.NeighCellInfoList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
