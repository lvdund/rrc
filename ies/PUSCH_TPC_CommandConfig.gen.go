// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-TPC-CommandConfig (line 12739).

var pUSCHTPCCommandConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tpc-Index", Optional: true},
		{Name: "tpc-IndexSUL", Optional: true},
		{Name: "targetCell", Optional: true},
	},
}

var pUSCHTPCCommandConfigTpcIndexConstraints = per.Constrained(1, 15)

var pUSCHTPCCommandConfigTpcIndexSULConstraints = per.Constrained(1, 15)

type PUSCH_TPC_CommandConfig struct {
	Tpc_Index    *int64
	Tpc_IndexSUL *int64
	TargetCell   *ServCellIndex
}

func (ie *PUSCH_TPC_CommandConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHTPCCommandConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tpc_Index != nil, ie.Tpc_IndexSUL != nil, ie.TargetCell != nil}); err != nil {
		return err
	}

	// 3. tpc-Index: integer
	{
		if ie.Tpc_Index != nil {
			if err := e.EncodeInteger(*ie.Tpc_Index, pUSCHTPCCommandConfigTpcIndexConstraints); err != nil {
				return err
			}
		}
	}

	// 4. tpc-IndexSUL: integer
	{
		if ie.Tpc_IndexSUL != nil {
			if err := e.EncodeInteger(*ie.Tpc_IndexSUL, pUSCHTPCCommandConfigTpcIndexSULConstraints); err != nil {
				return err
			}
		}
	}

	// 5. targetCell: ref
	{
		if ie.TargetCell != nil {
			if err := ie.TargetCell.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUSCH_TPC_CommandConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHTPCCommandConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tpc-Index: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUSCHTPCCommandConfigTpcIndexConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_Index = &v
		}
	}

	// 4. tpc-IndexSUL: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(pUSCHTPCCommandConfigTpcIndexSULConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_IndexSUL = &v
		}
	}

	// 5. targetCell: ref
	{
		if seq.IsComponentPresent(2) {
			ie.TargetCell = new(ServCellIndex)
			if err := ie.TargetCell.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
