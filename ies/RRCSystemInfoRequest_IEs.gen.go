// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RRCSystemInfoRequest-IEs (line 1818).

var rRCSystemInfoRequestIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requested-SI-List"},
		{Name: "spare"},
	},
}

var rRCSystemInfoRequestIEsRequestedSIListConstraints = per.FixedSize(common.MaxSI_Message)

var rRCSystemInfoRequestIEsSpareConstraints = per.FixedSize(12)

type RRCSystemInfoRequest_IEs struct {
	Requested_SI_List per.BitString
	Spare             per.BitString
}

func (ie *RRCSystemInfoRequest_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSystemInfoRequestIEsConstraints)
	_ = seq

	// 1. requested-SI-List: bit-string
	{
		if err := e.EncodeBitString(ie.Requested_SI_List, rRCSystemInfoRequestIEsRequestedSIListConstraints); err != nil {
			return err
		}
	}

	// 2. spare: bit-string
	{
		if err := e.EncodeBitString(ie.Spare, rRCSystemInfoRequestIEsSpareConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCSystemInfoRequest_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSystemInfoRequestIEsConstraints)
	_ = seq

	// 1. requested-SI-List: bit-string
	{
		v0, err := d.DecodeBitString(rRCSystemInfoRequestIEsRequestedSIListConstraints)
		if err != nil {
			return err
		}
		ie.Requested_SI_List = v0
	}

	// 2. spare: bit-string
	{
		v1, err := d.DecodeBitString(rRCSystemInfoRequestIEsSpareConstraints)
		if err != nil {
			return err
		}
		ie.Spare = v1
	}

	return nil
}
