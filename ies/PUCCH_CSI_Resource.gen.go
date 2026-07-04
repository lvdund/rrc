// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-CSI-Resource (line 12213).

var pUCCHCSIResourceConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkBandwidthPartId"},
		{Name: "pucch-Resource"},
	},
}

type PUCCH_CSI_Resource struct {
	UplinkBandwidthPartId BWP_Id
	Pucch_Resource        PUCCH_ResourceId
}

func (ie *PUCCH_CSI_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHCSIResourceConstraints)
	_ = seq

	// 1. uplinkBandwidthPartId: ref
	{
		if err := ie.UplinkBandwidthPartId.Encode(e); err != nil {
			return err
		}
	}

	// 2. pucch-Resource: ref
	{
		if err := ie.Pucch_Resource.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_CSI_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHCSIResourceConstraints)
	_ = seq

	// 1. uplinkBandwidthPartId: ref
	{
		if err := ie.UplinkBandwidthPartId.Decode(d); err != nil {
			return err
		}
	}

	// 2. pucch-Resource: ref
	{
		if err := ie.Pucch_Resource.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
