// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RxTxTimeDiff-r17 (line 14221).

var rxTxTimeDiffR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "result-k5-r17", Optional: true},
	},
}

var rxTxTimeDiffR17ResultK5R17Constraints = per.Constrained(0, 61565)

type RxTxTimeDiff_r17 struct {
	Result_K5_r17 *int64
}

func (ie *RxTxTimeDiff_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rxTxTimeDiffR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Result_K5_r17 != nil}); err != nil {
		return err
	}

	// 3. result-k5-r17: integer
	{
		if ie.Result_K5_r17 != nil {
			if err := e.EncodeInteger(*ie.Result_K5_r17, rxTxTimeDiffR17ResultK5R17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RxTxTimeDiff_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rxTxTimeDiffR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. result-k5-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(rxTxTimeDiffR17ResultK5R17Constraints)
			if err != nil {
				return err
			}
			ie.Result_K5_r17 = &v
		}
	}

	return nil
}
