// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RF-Parameters-v16c0 (line 23662).

var rFParametersV16c0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandListNR-v16c0"},
	},
}

var rFParametersV16c0SupportedBandListNRV16c0Constraints = per.SizeRange(1, common.MaxBands)

type RF_Parameters_V16c0 struct {
	SupportedBandListNR_V16c0 []BandNR_V16c0
}

func (ie *RF_Parameters_V16c0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersV16c0Constraints)
	_ = seq

	// 1. supportedBandListNR-v16c0: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(rFParametersV16c0SupportedBandListNRV16c0Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.SupportedBandListNR_V16c0))); err != nil {
			return err
		}
		for i := range ie.SupportedBandListNR_V16c0 {
			if err := ie.SupportedBandListNR_V16c0[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RF_Parameters_V16c0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersV16c0Constraints)
	_ = seq

	// 1. supportedBandListNR-v16c0: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(rFParametersV16c0SupportedBandListNRV16c0Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SupportedBandListNR_V16c0 = make([]BandNR_V16c0, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SupportedBandListNR_V16c0[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
