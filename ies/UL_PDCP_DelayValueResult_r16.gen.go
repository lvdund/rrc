// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-PDCP-DelayValueResult-r16 (line 9874).

var uLPDCPDelayValueResultR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-Id-r16"},
		{Name: "averageDelay-r16"},
	},
}

var uLPDCPDelayValueResultR16AverageDelayR16Constraints = per.Constrained(0, 10000)

type UL_PDCP_DelayValueResult_r16 struct {
	Drb_Id_r16       DRB_Identity
	AverageDelay_r16 int64
}

func (ie *UL_PDCP_DelayValueResult_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLPDCPDelayValueResultR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. drb-Id-r16: ref
	{
		if err := ie.Drb_Id_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. averageDelay-r16: integer
	{
		if err := e.EncodeInteger(ie.AverageDelay_r16, uLPDCPDelayValueResultR16AverageDelayR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UL_PDCP_DelayValueResult_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLPDCPDelayValueResultR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. drb-Id-r16: ref
	{
		if err := ie.Drb_Id_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. averageDelay-r16: integer
	{
		v1, err := d.DecodeInteger(uLPDCPDelayValueResultR16AverageDelayR16Constraints)
		if err != nil {
			return err
		}
		ie.AverageDelay_r16 = v1
	}

	return nil
}
