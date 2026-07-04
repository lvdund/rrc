// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-SRS (line 12306).

var pUCCHSRSConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resource"},
		{Name: "uplinkBWP"},
	},
}

type PUCCH_SRS struct {
	Resource  SRS_ResourceId
	UplinkBWP BWP_Id
}

func (ie *PUCCH_SRS) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHSRSConstraints)
	_ = seq

	// 1. resource: ref
	{
		if err := ie.Resource.Encode(e); err != nil {
			return err
		}
	}

	// 2. uplinkBWP: ref
	{
		if err := ie.UplinkBWP.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_SRS) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHSRSConstraints)
	_ = seq

	// 1. resource: ref
	{
		if err := ie.Resource.Decode(d); err != nil {
			return err
		}
	}

	// 2. uplinkBWP: ref
	{
		if err := ie.UplinkBWP.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
