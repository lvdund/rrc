// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-RLC-ConfigBroadcast-r17 (line 28546).

var mRBRLCConfigBroadcastR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "logicalChannelIdentity-r17"},
		{Name: "sn-FieldLength-r17", Optional: true},
		{Name: "t-Reassembly-r17", Optional: true},
	},
}

const (
	MRB_RLC_ConfigBroadcast_r17_Sn_FieldLength_r17_Size6 = 0
)

var mRBRLCConfigBroadcastR17SnFieldLengthR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_RLC_ConfigBroadcast_r17_Sn_FieldLength_r17_Size6},
}

type MRB_RLC_ConfigBroadcast_r17 struct {
	LogicalChannelIdentity_r17 LogicalChannelIdentity
	Sn_FieldLength_r17         *int64
	T_Reassembly_r17           *T_Reassembly
}

func (ie *MRB_RLC_ConfigBroadcast_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRBRLCConfigBroadcastR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sn_FieldLength_r17 != nil, ie.T_Reassembly_r17 != nil}); err != nil {
		return err
	}

	// 2. logicalChannelIdentity-r17: ref
	{
		if err := ie.LogicalChannelIdentity_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. sn-FieldLength-r17: enumerated
	{
		if ie.Sn_FieldLength_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sn_FieldLength_r17, mRBRLCConfigBroadcastR17SnFieldLengthR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. t-Reassembly-r17: ref
	{
		if ie.T_Reassembly_r17 != nil {
			if err := ie.T_Reassembly_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRB_RLC_ConfigBroadcast_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRBRLCConfigBroadcastR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. logicalChannelIdentity-r17: ref
	{
		if err := ie.LogicalChannelIdentity_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. sn-FieldLength-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mRBRLCConfigBroadcastR17SnFieldLengthR17Constraints)
			if err != nil {
				return err
			}
			ie.Sn_FieldLength_r17 = &idx
		}
	}

	// 4. t-Reassembly-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.T_Reassembly_r17 = new(T_Reassembly)
			if err := ie.T_Reassembly_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
