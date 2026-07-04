// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReestablishmentRequest-IEs (line 946).

var rRCReestablishmentRequestIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-Identity"},
		{Name: "reestablishmentCause"},
		{Name: "spare"},
	},
}

var rRCReestablishmentRequestIEsSpareConstraints = per.FixedSize(1)

type RRCReestablishmentRequest_IEs struct {
	Ue_Identity          ReestabUE_Identity
	ReestablishmentCause ReestablishmentCause
	Spare                per.BitString
}

func (ie *RRCReestablishmentRequest_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestablishmentRequestIEsConstraints)
	_ = seq

	// 1. ue-Identity: ref
	{
		if err := ie.Ue_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 2. reestablishmentCause: ref
	{
		if err := ie.ReestablishmentCause.Encode(e); err != nil {
			return err
		}
	}

	// 3. spare: bit-string
	{
		if err := e.EncodeBitString(ie.Spare, rRCReestablishmentRequestIEsSpareConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCReestablishmentRequest_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestablishmentRequestIEsConstraints)
	_ = seq

	// 1. ue-Identity: ref
	{
		if err := ie.Ue_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 2. reestablishmentCause: ref
	{
		if err := ie.ReestablishmentCause.Decode(d); err != nil {
			return err
		}
	}

	// 3. spare: bit-string
	{
		v2, err := d.DecodeBitString(rRCReestablishmentRequestIEsSpareConstraints)
		if err != nil {
			return err
		}
		ie.Spare = v2
	}

	return nil
}
