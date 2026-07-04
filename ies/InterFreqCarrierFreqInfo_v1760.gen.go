// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InterFreqCarrierFreqInfo-v1760 (line 4030).

var interFreqCarrierFreqInfoV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyBandList-v1760", Optional: true},
		{Name: "frequencyBandListSUL-v1760", Optional: true},
	},
}

type InterFreqCarrierFreqInfo_v1760 struct {
	FrequencyBandList_v1760    *MultiFrequencyBandListNR_SIB_v1760
	FrequencyBandListSUL_v1760 *MultiFrequencyBandListNR_SIB_v1760
}

func (ie *InterFreqCarrierFreqInfo_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1760Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyBandList_v1760 != nil, ie.FrequencyBandListSUL_v1760 != nil}); err != nil {
		return err
	}

	// 2. frequencyBandList-v1760: ref
	{
		if ie.FrequencyBandList_v1760 != nil {
			if err := ie.FrequencyBandList_v1760.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. frequencyBandListSUL-v1760: ref
	{
		if ie.FrequencyBandListSUL_v1760 != nil {
			if err := ie.FrequencyBandListSUL_v1760.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1760Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. frequencyBandList-v1760: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FrequencyBandList_v1760 = new(MultiFrequencyBandListNR_SIB_v1760)
			if err := ie.FrequencyBandList_v1760.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. frequencyBandListSUL-v1760: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FrequencyBandListSUL_v1760 = new(MultiFrequencyBandListNR_SIB_v1760)
			if err := ie.FrequencyBandListSUL_v1760.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
