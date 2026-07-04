// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IDC-FDM-Assistance-r18 (line 2715).

var iDCFDMAssistanceR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "affectedCarrierFreqRangeList-r18", Optional: true},
		{Name: "affectedCarrierFreqRangeCombList-r18", Optional: true},
	},
}

type IDC_FDM_Assistance_r18 struct {
	AffectedCarrierFreqRangeList_r18     *AffectedCarrierFreqRangeList_r18
	AffectedCarrierFreqRangeCombList_r18 *AffectedCarrierFreqRangeCombList_r18
}

func (ie *IDC_FDM_Assistance_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iDCFDMAssistanceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AffectedCarrierFreqRangeList_r18 != nil, ie.AffectedCarrierFreqRangeCombList_r18 != nil}); err != nil {
		return err
	}

	// 3. affectedCarrierFreqRangeList-r18: ref
	{
		if ie.AffectedCarrierFreqRangeList_r18 != nil {
			if err := ie.AffectedCarrierFreqRangeList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. affectedCarrierFreqRangeCombList-r18: ref
	{
		if ie.AffectedCarrierFreqRangeCombList_r18 != nil {
			if err := ie.AffectedCarrierFreqRangeCombList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IDC_FDM_Assistance_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iDCFDMAssistanceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. affectedCarrierFreqRangeList-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AffectedCarrierFreqRangeList_r18 = new(AffectedCarrierFreqRangeList_r18)
			if err := ie.AffectedCarrierFreqRangeList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. affectedCarrierFreqRangeCombList-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.AffectedCarrierFreqRangeCombList_r18 = new(AffectedCarrierFreqRangeCombList_r18)
			if err := ie.AffectedCarrierFreqRangeCombList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
