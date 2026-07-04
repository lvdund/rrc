// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PathlossReferenceRS-r16 (line 15395).

var pathlossReferenceRSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PathlossReferenceRS-Id-r16"},
		{Name: "pathlossReferenceRS-r16"},
	},
}

type PathlossReferenceRS_r16 struct {
	Srs_PathlossReferenceRS_Id_r16 SRS_PathlossReferenceRS_Id_r16
	PathlossReferenceRS_r16        PathlossReferenceRS_Config
}

func (ie *PathlossReferenceRS_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pathlossReferenceRSR16Constraints)
	_ = seq

	// 1. srs-PathlossReferenceRS-Id-r16: ref
	{
		if err := ie.Srs_PathlossReferenceRS_Id_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. pathlossReferenceRS-r16: ref
	{
		if err := ie.PathlossReferenceRS_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PathlossReferenceRS_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pathlossReferenceRSR16Constraints)
	_ = seq

	// 1. srs-PathlossReferenceRS-Id-r16: ref
	{
		if err := ie.Srs_PathlossReferenceRS_Id_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. pathlossReferenceRS-r16: ref
	{
		if err := ie.PathlossReferenceRS_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
