// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultUTRA-FDD-r16 (line 9840).

var measResultUTRAFDDR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r16"},
		{Name: "measResult-r16"},
	},
}

var measResultUTRAFDDR16MeasResultR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "utra-FDD-RSCP-r16", Optional: true},
		{Name: "utra-FDD-EcN0-r16", Optional: true},
	},
}

type MeasResultUTRA_FDD_r16 struct {
	PhysCellId_r16 PhysCellIdUTRA_FDD_r16
	MeasResult_r16 struct {
		Utra_FDD_RSCP_r16 *int64
		Utra_FDD_EcN0_r16 *int64
	}
}

func (ie *MeasResultUTRA_FDD_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultUTRAFDDR16Constraints)
	_ = seq

	// 1. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. measResult-r16: sequence
	{
		{
			c := &ie.MeasResult_r16
			measResultUTRAFDDR16MeasResultR16Seq := e.NewSequenceEncoder(measResultUTRAFDDR16MeasResultR16Constraints)
			if err := measResultUTRAFDDR16MeasResultR16Seq.EncodePreamble([]bool{c.Utra_FDD_RSCP_r16 != nil, c.Utra_FDD_EcN0_r16 != nil}); err != nil {
				return err
			}
			if c.Utra_FDD_RSCP_r16 != nil {
				if err := e.EncodeInteger((*c.Utra_FDD_RSCP_r16), per.Constrained(-5, 91)); err != nil {
					return err
				}
			}
			if c.Utra_FDD_EcN0_r16 != nil {
				if err := e.EncodeInteger((*c.Utra_FDD_EcN0_r16), per.Constrained(0, 49)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MeasResultUTRA_FDD_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultUTRAFDDR16Constraints)
	_ = seq

	// 1. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. measResult-r16: sequence
	{
		{
			c := &ie.MeasResult_r16
			measResultUTRAFDDR16MeasResultR16Seq := d.NewSequenceDecoder(measResultUTRAFDDR16MeasResultR16Constraints)
			if err := measResultUTRAFDDR16MeasResultR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measResultUTRAFDDR16MeasResultR16Seq.IsComponentPresent(0) {
				c.Utra_FDD_RSCP_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(-5, 91))
				if err != nil {
					return err
				}
				(*c.Utra_FDD_RSCP_r16) = v
			}
			if measResultUTRAFDDR16MeasResultR16Seq.IsComponentPresent(1) {
				c.Utra_FDD_EcN0_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 49))
				if err != nil {
					return err
				}
				(*c.Utra_FDD_EcN0_r16) = v
			}
		}
	}

	return nil
}
