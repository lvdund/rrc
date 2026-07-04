// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: N3C-RelayUE-Info-r18 (line 10269).

var n3CRelayUEInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "n3c-CellGlobalId-r18"},
		{Name: "n3c-C-RNTI-r18"},
	},
}

type N3C_RelayUE_Info_r18 struct {
	N3c_CellGlobalId_r18 struct {
		N3c_PLMN_Id_r18      PLMN_Identity
		N3c_CellIdentity_r18 CellIdentity
	}
	N3c_C_RNTI_r18 RNTI_Value
}

func (ie *N3C_RelayUE_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(n3CRelayUEInfoR18Constraints)
	_ = seq

	// 1. n3c-CellGlobalId-r18: sequence
	{
		{
			c := &ie.N3c_CellGlobalId_r18
			if err := c.N3c_PLMN_Id_r18.Encode(e); err != nil {
				return err
			}
			if err := c.N3c_CellIdentity_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. n3c-C-RNTI-r18: ref
	{
		if err := ie.N3c_C_RNTI_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *N3C_RelayUE_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(n3CRelayUEInfoR18Constraints)
	_ = seq

	// 1. n3c-CellGlobalId-r18: sequence
	{
		{
			c := &ie.N3c_CellGlobalId_r18
			{
				if err := c.N3c_PLMN_Id_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.N3c_CellIdentity_r18.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 2. n3c-C-RNTI-r18: ref
	{
		if err := ie.N3c_C_RNTI_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
