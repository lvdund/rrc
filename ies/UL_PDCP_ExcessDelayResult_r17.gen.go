// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-PDCP-ExcessDelayResult-r17 (line 9882).

var uLPDCPExcessDelayResultR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-Id-r17"},
		{Name: "excessDelay-r17"},
	},
}

var uLPDCPExcessDelayResultR17ExcessDelayR17Constraints = per.Constrained(0, 31)

type UL_PDCP_ExcessDelayResult_r17 struct {
	Drb_Id_r17      DRB_Identity
	ExcessDelay_r17 int64
}

func (ie *UL_PDCP_ExcessDelayResult_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLPDCPExcessDelayResultR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. drb-Id-r17: ref
	{
		if err := ie.Drb_Id_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. excessDelay-r17: integer
	{
		if err := e.EncodeInteger(ie.ExcessDelay_r17, uLPDCPExcessDelayResultR17ExcessDelayR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UL_PDCP_ExcessDelayResult_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLPDCPExcessDelayResultR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. drb-Id-r17: ref
	{
		if err := ie.Drb_Id_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. excessDelay-r17: integer
	{
		v1, err := d.DecodeInteger(uLPDCPExcessDelayResultR17ExcessDelayR17Constraints)
		if err != nil {
			return err
		}
		ie.ExcessDelay_r17 = v1
	}

	return nil
}
