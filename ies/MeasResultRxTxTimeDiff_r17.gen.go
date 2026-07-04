// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultRxTxTimeDiff-r17 (line 9983).

var measResultRxTxTimeDiffR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rxTxTimeDiff-ue-r17", Optional: true},
	},
}

type MeasResultRxTxTimeDiff_r17 struct {
	RxTxTimeDiff_Ue_r17 *RxTxTimeDiff_r17
}

func (ie *MeasResultRxTxTimeDiff_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultRxTxTimeDiffR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RxTxTimeDiff_Ue_r17 != nil}); err != nil {
		return err
	}

	// 3. rxTxTimeDiff-ue-r17: ref
	{
		if ie.RxTxTimeDiff_Ue_r17 != nil {
			if err := ie.RxTxTimeDiff_Ue_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultRxTxTimeDiff_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultRxTxTimeDiffR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rxTxTimeDiff-ue-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.RxTxTimeDiff_Ue_r17 = new(RxTxTimeDiff_r17)
			if err := ie.RxTxTimeDiff_Ue_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
