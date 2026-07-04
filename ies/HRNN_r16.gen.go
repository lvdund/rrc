// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: HRNN-r16 (line 4256).

var hRNNR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "hrnn-r16", Optional: true},
	},
}

var hRNNR16HrnnR16Constraints = per.SizeRange(1, common.MaxHRNN_Len_r16)

type HRNN_r16 struct {
	Hrnn_r16 []byte
}

func (ie *HRNN_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(hRNNR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Hrnn_r16 != nil}); err != nil {
		return err
	}

	// 2. hrnn-r16: octet-string
	{
		if ie.Hrnn_r16 != nil {
			if err := e.EncodeOctetString(ie.Hrnn_r16, hRNNR16HrnnR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *HRNN_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(hRNNR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. hrnn-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(hRNNR16HrnnR16Constraints)
			if err != nil {
				return err
			}
			ie.Hrnn_r16 = v
		}
	}

	return nil
}
