// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ConfigDedicatedEUTRA-Info-r16 (line 1114).

var sLConfigDedicatedEUTRAInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfigDedicatedEUTRA-r16", Optional: true},
		{Name: "sl-TimeOffsetEUTRA-List-r16", Optional: true},
	},
}

var sLConfigDedicatedEUTRAInfoR16SlConfigDedicatedEUTRAR16Constraints = per.SizeConstraints{}

var sLConfigDedicatedEUTRAInfoR16SlTimeOffsetEUTRAListR16Constraints = per.FixedSize(8)

type SL_ConfigDedicatedEUTRA_Info_r16 struct {
	Sl_ConfigDedicatedEUTRA_r16 []byte
	Sl_TimeOffsetEUTRA_List_r16 []SL_TimeOffsetEUTRA_r16
}

func (ie *SL_ConfigDedicatedEUTRA_Info_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfigDedicatedEUTRAInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfigDedicatedEUTRA_r16 != nil, ie.Sl_TimeOffsetEUTRA_List_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedEUTRA-r16: octet-string
	{
		if ie.Sl_ConfigDedicatedEUTRA_r16 != nil {
			if err := e.EncodeOctetString(ie.Sl_ConfigDedicatedEUTRA_r16, sLConfigDedicatedEUTRAInfoR16SlConfigDedicatedEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-TimeOffsetEUTRA-List-r16: sequence-of
	{
		if ie.Sl_TimeOffsetEUTRA_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigDedicatedEUTRAInfoR16SlTimeOffsetEUTRAListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_TimeOffsetEUTRA_List_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_TimeOffsetEUTRA_List_r16 {
				if err := ie.Sl_TimeOffsetEUTRA_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_ConfigDedicatedEUTRA_Info_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfigDedicatedEUTRAInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedEUTRA-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(sLConfigDedicatedEUTRAInfoR16SlConfigDedicatedEUTRAR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ConfigDedicatedEUTRA_r16 = v
		}
	}

	// 3. sl-TimeOffsetEUTRA-List-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLConfigDedicatedEUTRAInfoR16SlTimeOffsetEUTRAListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_TimeOffsetEUTRA_List_r16 = make([]SL_TimeOffsetEUTRA_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_TimeOffsetEUTRA_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
