// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupRequest-IEs (line 1789).

var rRCSetupRequestIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-Identity"},
		{Name: "establishmentCause"},
		{Name: "spare"},
	},
}

var rRCSetupRequestIEsSpareConstraints = per.FixedSize(1)

type RRCSetupRequest_IEs struct {
	Ue_Identity        InitialUE_Identity
	EstablishmentCause EstablishmentCause
	Spare              per.BitString
}

func (ie *RRCSetupRequest_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupRequestIEsConstraints)
	_ = seq

	// 1. ue-Identity: ref
	{
		if err := ie.Ue_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 2. establishmentCause: ref
	{
		if err := ie.EstablishmentCause.Encode(e); err != nil {
			return err
		}
	}

	// 3. spare: bit-string
	{
		if err := e.EncodeBitString(ie.Spare, rRCSetupRequestIEsSpareConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCSetupRequest_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupRequestIEsConstraints)
	_ = seq

	// 1. ue-Identity: ref
	{
		if err := ie.Ue_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 2. establishmentCause: ref
	{
		if err := ie.EstablishmentCause.Decode(d); err != nil {
			return err
		}
	}

	// 3. spare: bit-string
	{
		v2, err := d.DecodeBitString(rRCSetupRequestIEsSpareConstraints)
		if err != nil {
			return err
		}
		ie.Spare = v2
	}

	return nil
}
