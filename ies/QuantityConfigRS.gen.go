// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: QuantityConfigRS (line 12792).

var quantityConfigRSConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-FilterConfig"},
		{Name: "csi-RS-FilterConfig"},
	},
}

type QuantityConfigRS struct {
	Ssb_FilterConfig    FilterConfig
	Csi_RS_FilterConfig FilterConfig
}

func (ie *QuantityConfigRS) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(quantityConfigRSConstraints)
	_ = seq

	// 1. ssb-FilterConfig: ref
	{
		if err := ie.Ssb_FilterConfig.Encode(e); err != nil {
			return err
		}
	}

	// 2. csi-RS-FilterConfig: ref
	{
		if err := ie.Csi_RS_FilterConfig.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *QuantityConfigRS) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(quantityConfigRSConstraints)
	_ = seq

	// 1. ssb-FilterConfig: ref
	{
		if err := ie.Ssb_FilterConfig.Decode(d); err != nil {
			return err
		}
	}

	// 2. csi-RS-FilterConfig: ref
	{
		if err := ie.Csi_RS_FilterConfig.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
