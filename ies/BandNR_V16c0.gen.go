// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandNR-v16c0 (line 24662).

var bandNRV16c0Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-RepetitionTypeA-v16c0", Optional: true},
	},
}

const (
	BandNR_V16c0_Pusch_RepetitionTypeA_V16c0_Supported = 0
)

var bandNRV16c0PuschRepetitionTypeAV16c0Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_V16c0_Pusch_RepetitionTypeA_V16c0_Supported},
}

type BandNR_V16c0 struct {
	Pusch_RepetitionTypeA_V16c0 *int64
}

func (ie *BandNR_V16c0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandNRV16c0Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pusch_RepetitionTypeA_V16c0 != nil}); err != nil {
		return err
	}

	// 3. pusch-RepetitionTypeA-v16c0: enumerated
	{
		if ie.Pusch_RepetitionTypeA_V16c0 != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_RepetitionTypeA_V16c0, bandNRV16c0PuschRepetitionTypeAV16c0Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandNR_V16c0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandNRV16c0Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pusch-RepetitionTypeA-v16c0: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(bandNRV16c0PuschRepetitionTypeAV16c0Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_RepetitionTypeA_V16c0 = &idx
		}
	}

	return nil
}
