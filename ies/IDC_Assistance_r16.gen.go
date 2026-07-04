// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IDC-Assistance-r16 (line 2498).

var iDCAssistanceR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "affectedCarrierFreqList-r16", Optional: true},
		{Name: "affectedCarrierFreqCombList-r16", Optional: true},
	},
}

type IDC_Assistance_r16 struct {
	AffectedCarrierFreqList_r16     *AffectedCarrierFreqList_r16
	AffectedCarrierFreqCombList_r16 *AffectedCarrierFreqCombList_r16
}

func (ie *IDC_Assistance_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iDCAssistanceR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AffectedCarrierFreqList_r16 != nil, ie.AffectedCarrierFreqCombList_r16 != nil}); err != nil {
		return err
	}

	// 3. affectedCarrierFreqList-r16: ref
	{
		if ie.AffectedCarrierFreqList_r16 != nil {
			if err := ie.AffectedCarrierFreqList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. affectedCarrierFreqCombList-r16: ref
	{
		if ie.AffectedCarrierFreqCombList_r16 != nil {
			if err := ie.AffectedCarrierFreqCombList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IDC_Assistance_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iDCAssistanceR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. affectedCarrierFreqList-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AffectedCarrierFreqList_r16 = new(AffectedCarrierFreqList_r16)
			if err := ie.AffectedCarrierFreqList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. affectedCarrierFreqCombList-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.AffectedCarrierFreqCombList_r16 = new(AffectedCarrierFreqCombList_r16)
			if err := ie.AffectedCarrierFreqCombList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
