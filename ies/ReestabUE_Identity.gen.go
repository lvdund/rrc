// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReestabUE-Identity (line 952).

var reestabUEIdentityConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "c-RNTI"},
		{Name: "physCellId"},
		{Name: "shortMAC-I"},
	},
}

type ReestabUE_Identity struct {
	C_RNTI     RNTI_Value
	PhysCellId PhysCellId
	ShortMAC_I ShortMAC_I
}

func (ie *ReestabUE_Identity) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reestabUEIdentityConstraints)
	_ = seq

	// 1. c-RNTI: ref
	{
		if err := ie.C_RNTI.Encode(e); err != nil {
			return err
		}
	}

	// 2. physCellId: ref
	{
		if err := ie.PhysCellId.Encode(e); err != nil {
			return err
		}
	}

	// 3. shortMAC-I: ref
	{
		if err := ie.ShortMAC_I.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReestabUE_Identity) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reestabUEIdentityConstraints)
	_ = seq

	// 1. c-RNTI: ref
	{
		if err := ie.C_RNTI.Decode(d); err != nil {
			return err
		}
	}

	// 2. physCellId: ref
	{
		if err := ie.PhysCellId.Decode(d); err != nil {
			return err
		}
	}

	// 3. shortMAC-I: ref
	{
		if err := ie.ShortMAC_I.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
