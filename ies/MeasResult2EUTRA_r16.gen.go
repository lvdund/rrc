// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResult2EUTRA-r16 (line 3430).

var measResult2EUTRAR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "measResultList-r16"},
	},
}

type MeasResult2EUTRA_r16 struct {
	CarrierFreq_r16    ARFCN_ValueEUTRA
	MeasResultList_r16 MeasResultListEUTRA
}

func (ie *MeasResult2EUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResult2EUTRAR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. measResultList-r16: ref
	{
		if err := ie.MeasResultList_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResult2EUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResult2EUTRAR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. measResultList-r16: ref
	{
		if err := ie.MeasResultList_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
