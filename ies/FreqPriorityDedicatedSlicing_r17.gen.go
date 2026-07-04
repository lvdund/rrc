// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqPriorityDedicatedSlicing-r17 (line 8358).

var freqPriorityDedicatedSlicingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-ExplicitCarrierFreq-r17"},
		{Name: "sliceInfoListDedicated-r17", Optional: true},
	},
}

type FreqPriorityDedicatedSlicing_r17 struct {
	Dl_ExplicitCarrierFreq_r17 ARFCN_ValueNR
	SliceInfoListDedicated_r17 *SliceInfoListDedicated_r17
}

func (ie *FreqPriorityDedicatedSlicing_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqPriorityDedicatedSlicingR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SliceInfoListDedicated_r17 != nil}); err != nil {
		return err
	}

	// 2. dl-ExplicitCarrierFreq-r17: ref
	{
		if err := ie.Dl_ExplicitCarrierFreq_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. sliceInfoListDedicated-r17: ref
	{
		if ie.SliceInfoListDedicated_r17 != nil {
			if err := ie.SliceInfoListDedicated_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FreqPriorityDedicatedSlicing_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqPriorityDedicatedSlicingR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-ExplicitCarrierFreq-r17: ref
	{
		if err := ie.Dl_ExplicitCarrierFreq_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. sliceInfoListDedicated-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SliceInfoListDedicated_r17 = new(SliceInfoListDedicated_r17)
			if err := ie.SliceInfoListDedicated_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
