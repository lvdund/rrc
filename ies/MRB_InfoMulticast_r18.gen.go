// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-InfoMulticast-r18 (line 28578).

var mRBInfoMulticastR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcp-Config-r18"},
		{Name: "rlc-Config-r18"},
	},
}

type MRB_InfoMulticast_r18 struct {
	Pdcp_Config_r18 MRB_PDCP_ConfigMulticast_r18
	Rlc_Config_r18  MRB_RLC_ConfigMulticast_r18
}

func (ie *MRB_InfoMulticast_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRBInfoMulticastR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. pdcp-Config-r18: ref
	{
		if err := ie.Pdcp_Config_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. rlc-Config-r18: ref
	{
		if err := ie.Rlc_Config_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MRB_InfoMulticast_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRBInfoMulticastR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. pdcp-Config-r18: ref
	{
		if err := ie.Pdcp_Config_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. rlc-Config-r18: ref
	{
		if err := ie.Rlc_Config_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
