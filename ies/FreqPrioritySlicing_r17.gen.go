// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqPrioritySlicing-r17 (line 8376).

var freqPrioritySlicingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-ImplicitCarrierFreq-r17"},
		{Name: "sliceInfoList-r17", Optional: true},
	},
}

var freqPrioritySlicingR17DlImplicitCarrierFreqR17Constraints = per.Constrained(0, common.MaxFreq)

type FreqPrioritySlicing_r17 struct {
	Dl_ImplicitCarrierFreq_r17 int64
	SliceInfoList_r17          *SliceInfoList_r17
}

func (ie *FreqPrioritySlicing_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqPrioritySlicingR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SliceInfoList_r17 != nil}); err != nil {
		return err
	}

	// 2. dl-ImplicitCarrierFreq-r17: integer
	{
		if err := e.EncodeInteger(ie.Dl_ImplicitCarrierFreq_r17, freqPrioritySlicingR17DlImplicitCarrierFreqR17Constraints); err != nil {
			return err
		}
	}

	// 3. sliceInfoList-r17: ref
	{
		if ie.SliceInfoList_r17 != nil {
			if err := ie.SliceInfoList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FreqPrioritySlicing_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqPrioritySlicingR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-ImplicitCarrierFreq-r17: integer
	{
		v0, err := d.DecodeInteger(freqPrioritySlicingR17DlImplicitCarrierFreqR17Constraints)
		if err != nil {
			return err
		}
		ie.Dl_ImplicitCarrierFreq_r17 = v0
	}

	// 3. sliceInfoList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SliceInfoList_r17 = new(SliceInfoList_r17)
			if err := ie.SliceInfoList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
