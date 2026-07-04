// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-RLC-ConfigMulticast-r18 (line 28600).

var mRBRLCConfigMulticastR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "logicalChannelIdentity-r18"},
		{Name: "sn-FieldLength-r18"},
		{Name: "t-Reassembly-r18", Optional: true},
	},
}

var mRB_RLC_ConfigMulticast_r18LogicalChannelIdentityR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "logicalChannelIdentitymulticast-r18"},
		{Name: "logicalChannelIdentityExt-r18"},
	},
}

const (
	MRB_RLC_ConfigMulticast_r18_LogicalChannelIdentity_r18_LogicalChannelIdentitymulticast_r18 = 0
	MRB_RLC_ConfigMulticast_r18_LogicalChannelIdentity_r18_LogicalChannelIdentityExt_r18       = 1
)

const (
	MRB_RLC_ConfigMulticast_r18_Sn_FieldLength_r18_Size6  = 0
	MRB_RLC_ConfigMulticast_r18_Sn_FieldLength_r18_Size12 = 1
)

var mRBRLCConfigMulticastR18SnFieldLengthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_RLC_ConfigMulticast_r18_Sn_FieldLength_r18_Size6, MRB_RLC_ConfigMulticast_r18_Sn_FieldLength_r18_Size12},
}

type MRB_RLC_ConfigMulticast_r18 struct {
	LogicalChannelIdentity_r18 struct {
		Choice                              int
		LogicalChannelIdentitymulticast_r18 *LogicalChannelIdentity
		LogicalChannelIdentityExt_r18       *LogicalChannelIdentityExt_r17
	}
	Sn_FieldLength_r18 int64
	T_Reassembly_r18   *T_Reassembly
}

func (ie *MRB_RLC_ConfigMulticast_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRBRLCConfigMulticastR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T_Reassembly_r18 != nil}); err != nil {
		return err
	}

	// 2. logicalChannelIdentity-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(mRB_RLC_ConfigMulticast_r18LogicalChannelIdentityR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.LogicalChannelIdentity_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.LogicalChannelIdentity_r18.Choice {
		case MRB_RLC_ConfigMulticast_r18_LogicalChannelIdentity_r18_LogicalChannelIdentitymulticast_r18:
			if err := ie.LogicalChannelIdentity_r18.LogicalChannelIdentitymulticast_r18.Encode(e); err != nil {
				return err
			}
		case MRB_RLC_ConfigMulticast_r18_LogicalChannelIdentity_r18_LogicalChannelIdentityExt_r18:
			if err := ie.LogicalChannelIdentity_r18.LogicalChannelIdentityExt_r18.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.LogicalChannelIdentity_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. sn-FieldLength-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sn_FieldLength_r18, mRBRLCConfigMulticastR18SnFieldLengthR18Constraints); err != nil {
			return err
		}
	}

	// 4. t-Reassembly-r18: ref
	{
		if ie.T_Reassembly_r18 != nil {
			if err := ie.T_Reassembly_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRB_RLC_ConfigMulticast_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRBRLCConfigMulticastR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. logicalChannelIdentity-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(mRB_RLC_ConfigMulticast_r18LogicalChannelIdentityR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.LogicalChannelIdentity_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MRB_RLC_ConfigMulticast_r18_LogicalChannelIdentity_r18_LogicalChannelIdentitymulticast_r18:
			ie.LogicalChannelIdentity_r18.LogicalChannelIdentitymulticast_r18 = new(LogicalChannelIdentity)
			if err := ie.LogicalChannelIdentity_r18.LogicalChannelIdentitymulticast_r18.Decode(d); err != nil {
				return err
			}
		case MRB_RLC_ConfigMulticast_r18_LogicalChannelIdentity_r18_LogicalChannelIdentityExt_r18:
			ie.LogicalChannelIdentity_r18.LogicalChannelIdentityExt_r18 = new(LogicalChannelIdentityExt_r17)
			if err := ie.LogicalChannelIdentity_r18.LogicalChannelIdentityExt_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sn-FieldLength-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(mRBRLCConfigMulticastR18SnFieldLengthR18Constraints)
		if err != nil {
			return err
		}
		ie.Sn_FieldLength_r18 = v1
	}

	// 4. t-Reassembly-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.T_Reassembly_r18 = new(T_Reassembly)
			if err := ie.T_Reassembly_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
