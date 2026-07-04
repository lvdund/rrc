// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-FormatConfigExt-r17 (line 12056).

var pUCCHFormatConfigExtR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxCodeRateLP-r17", Optional: true},
	},
}

type PUCCH_FormatConfigExt_r17 struct {
	MaxCodeRateLP_r17 *PUCCH_MaxCodeRate
}

func (ie *PUCCH_FormatConfigExt_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHFormatConfigExtR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxCodeRateLP_r17 != nil}); err != nil {
		return err
	}

	// 3. maxCodeRateLP-r17: ref
	{
		if ie.MaxCodeRateLP_r17 != nil {
			if err := ie.MaxCodeRateLP_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_FormatConfigExt_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHFormatConfigExtR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maxCodeRateLP-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MaxCodeRateLP_r17 = new(PUCCH_MaxCodeRate)
			if err := ie.MaxCodeRateLP_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
