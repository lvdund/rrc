// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-AM-RLC (line 14066).

var uLAMRLCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sn-FieldLength", Optional: true},
		{Name: "t-PollRetransmit"},
		{Name: "pollPDU"},
		{Name: "pollByte"},
		{Name: "maxRetxThreshold"},
	},
}

const (
	UL_AM_RLC_MaxRetxThreshold_T1  = 0
	UL_AM_RLC_MaxRetxThreshold_T2  = 1
	UL_AM_RLC_MaxRetxThreshold_T3  = 2
	UL_AM_RLC_MaxRetxThreshold_T4  = 3
	UL_AM_RLC_MaxRetxThreshold_T6  = 4
	UL_AM_RLC_MaxRetxThreshold_T8  = 5
	UL_AM_RLC_MaxRetxThreshold_T16 = 6
	UL_AM_RLC_MaxRetxThreshold_T32 = 7
)

var uLAMRLCMaxRetxThresholdConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UL_AM_RLC_MaxRetxThreshold_T1, UL_AM_RLC_MaxRetxThreshold_T2, UL_AM_RLC_MaxRetxThreshold_T3, UL_AM_RLC_MaxRetxThreshold_T4, UL_AM_RLC_MaxRetxThreshold_T6, UL_AM_RLC_MaxRetxThreshold_T8, UL_AM_RLC_MaxRetxThreshold_T16, UL_AM_RLC_MaxRetxThreshold_T32},
}

type UL_AM_RLC struct {
	Sn_FieldLength   *SN_FieldLengthAM
	T_PollRetransmit T_PollRetransmit
	PollPDU          PollPDU
	PollByte         PollByte
	MaxRetxThreshold int64
}

func (ie *UL_AM_RLC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLAMRLCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sn_FieldLength != nil}); err != nil {
		return err
	}

	// 2. sn-FieldLength: ref
	{
		if ie.Sn_FieldLength != nil {
			if err := ie.Sn_FieldLength.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. t-PollRetransmit: ref
	{
		if err := ie.T_PollRetransmit.Encode(e); err != nil {
			return err
		}
	}

	// 4. pollPDU: ref
	{
		if err := ie.PollPDU.Encode(e); err != nil {
			return err
		}
	}

	// 5. pollByte: ref
	{
		if err := ie.PollByte.Encode(e); err != nil {
			return err
		}
	}

	// 6. maxRetxThreshold: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxRetxThreshold, uLAMRLCMaxRetxThresholdConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UL_AM_RLC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLAMRLCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sn-FieldLength: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sn_FieldLength = new(SN_FieldLengthAM)
			if err := ie.Sn_FieldLength.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. t-PollRetransmit: ref
	{
		if err := ie.T_PollRetransmit.Decode(d); err != nil {
			return err
		}
	}

	// 4. pollPDU: ref
	{
		if err := ie.PollPDU.Decode(d); err != nil {
			return err
		}
	}

	// 5. pollByte: ref
	{
		if err := ie.PollByte.Decode(d); err != nil {
			return err
		}
	}

	// 6. maxRetxThreshold: enumerated
	{
		v4, err := d.DecodeEnumerated(uLAMRLCMaxRetxThresholdConstraints)
		if err != nil {
			return err
		}
		ie.MaxRetxThreshold = v4
	}

	return nil
}
