// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForNCSG-NR-r17 (line 10515).

var needForNCSGNRR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandNR-r17"},
		{Name: "gapIndication-r17"},
	},
}

const (
	NeedForNCSG_NR_r17_GapIndication_r17_Gap          = 0
	NeedForNCSG_NR_r17_GapIndication_r17_Ncsg         = 1
	NeedForNCSG_NR_r17_GapIndication_r17_Nogap_Noncsg = 2
)

var needForNCSGNRR17GapIndicationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NeedForNCSG_NR_r17_GapIndication_r17_Gap, NeedForNCSG_NR_r17_GapIndication_r17_Ncsg, NeedForNCSG_NR_r17_GapIndication_r17_Nogap_Noncsg},
}

type NeedForNCSG_NR_r17 struct {
	BandNR_r17        FreqBandIndicatorNR
	GapIndication_r17 int64
}

func (ie *NeedForNCSG_NR_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForNCSGNRR17Constraints)
	_ = seq

	// 1. bandNR-r17: ref
	{
		if err := ie.BandNR_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. gapIndication-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.GapIndication_r17, needForNCSGNRR17GapIndicationR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeedForNCSG_NR_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForNCSGNRR17Constraints)
	_ = seq

	// 1. bandNR-r17: ref
	{
		if err := ie.BandNR_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. gapIndication-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(needForNCSGNRR17GapIndicationR17Constraints)
		if err != nil {
			return err
		}
		ie.GapIndication_r17 = v1
	}

	return nil
}
