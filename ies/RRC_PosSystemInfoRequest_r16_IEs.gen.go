// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RRC-PosSystemInfoRequest-r16-IEs (line 1823).

var rRCPosSystemInfoRequestR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedPosSI-List"},
		{Name: "spare"},
	},
}

var rRCPosSystemInfoRequestR16IEsRequestedPosSIListConstraints = per.FixedSize(common.MaxSI_Message)

var rRCPosSystemInfoRequestR16IEsSpareConstraints = per.FixedSize(11)

type RRC_PosSystemInfoRequest_r16_IEs struct {
	RequestedPosSI_List per.BitString
	Spare               per.BitString
}

func (ie *RRC_PosSystemInfoRequest_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCPosSystemInfoRequestR16IEsConstraints)
	_ = seq

	// 1. requestedPosSI-List: bit-string
	{
		if err := e.EncodeBitString(ie.RequestedPosSI_List, rRCPosSystemInfoRequestR16IEsRequestedPosSIListConstraints); err != nil {
			return err
		}
	}

	// 2. spare: bit-string
	{
		if err := e.EncodeBitString(ie.Spare, rRCPosSystemInfoRequestR16IEsSpareConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRC_PosSystemInfoRequest_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCPosSystemInfoRequestR16IEsConstraints)
	_ = seq

	// 1. requestedPosSI-List: bit-string
	{
		v0, err := d.DecodeBitString(rRCPosSystemInfoRequestR16IEsRequestedPosSIListConstraints)
		if err != nil {
			return err
		}
		ie.RequestedPosSI_List = v0
	}

	// 2. spare: bit-string
	{
		v1, err := d.DecodeBitString(rRCPosSystemInfoRequestR16IEsSpareConstraints)
		if err != nil {
			return err
		}
		ie.Spare = v1
	}

	return nil
}
