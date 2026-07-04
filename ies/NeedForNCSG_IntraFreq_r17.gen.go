// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForNCSG-IntraFreq-r17 (line 10510).

var needForNCSGIntraFreqR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellId-r17"},
		{Name: "gapIndicationIntra-r17"},
	},
}

const (
	NeedForNCSG_IntraFreq_r17_GapIndicationIntra_r17_Gap          = 0
	NeedForNCSG_IntraFreq_r17_GapIndicationIntra_r17_Ncsg         = 1
	NeedForNCSG_IntraFreq_r17_GapIndicationIntra_r17_Nogap_Noncsg = 2
)

var needForNCSGIntraFreqR17GapIndicationIntraR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NeedForNCSG_IntraFreq_r17_GapIndicationIntra_r17_Gap, NeedForNCSG_IntraFreq_r17_GapIndicationIntra_r17_Ncsg, NeedForNCSG_IntraFreq_r17_GapIndicationIntra_r17_Nogap_Noncsg},
}

type NeedForNCSG_IntraFreq_r17 struct {
	ServCellId_r17         ServCellIndex
	GapIndicationIntra_r17 int64
}

func (ie *NeedForNCSG_IntraFreq_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForNCSGIntraFreqR17Constraints)
	_ = seq

	// 1. servCellId-r17: ref
	{
		if err := ie.ServCellId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. gapIndicationIntra-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.GapIndicationIntra_r17, needForNCSGIntraFreqR17GapIndicationIntraR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeedForNCSG_IntraFreq_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForNCSGIntraFreqR17Constraints)
	_ = seq

	// 1. servCellId-r17: ref
	{
		if err := ie.ServCellId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. gapIndicationIntra-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(needForNCSGIntraFreqR17GapIndicationIntraR17Constraints)
		if err != nil {
			return err
		}
		ie.GapIndicationIntra_r17 = v1
	}

	return nil
}
