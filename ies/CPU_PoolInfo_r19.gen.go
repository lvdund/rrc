// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CPU-PoolInfo-r19 (line 18367).

var cPUPoolInfoR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumCPU-PerCC-r19"},
		{Name: "maxNumCPU-AllCC-r19"},
	},
}

var cPUPoolInfoR19MaxNumCPUPerCCR19Constraints = per.Constrained(1, 8)

var cPUPoolInfoR19MaxNumCPUAllCCR19Constraints = per.Constrained(1, 32)

type CPU_PoolInfo_r19 struct {
	MaxNumCPU_PerCC_r19 int64
	MaxNumCPU_AllCC_r19 int64
}

func (ie *CPU_PoolInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cPUPoolInfoR19Constraints)
	_ = seq

	// 1. maxNumCPU-PerCC-r19: integer
	{
		if err := e.EncodeInteger(ie.MaxNumCPU_PerCC_r19, cPUPoolInfoR19MaxNumCPUPerCCR19Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumCPU-AllCC-r19: integer
	{
		if err := e.EncodeInteger(ie.MaxNumCPU_AllCC_r19, cPUPoolInfoR19MaxNumCPUAllCCR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CPU_PoolInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cPUPoolInfoR19Constraints)
	_ = seq

	// 1. maxNumCPU-PerCC-r19: integer
	{
		v0, err := d.DecodeInteger(cPUPoolInfoR19MaxNumCPUPerCCR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumCPU_PerCC_r19 = v0
	}

	// 2. maxNumCPU-AllCC-r19: integer
	{
		v1, err := d.DecodeInteger(cPUPoolInfoR19MaxNumCPUAllCCR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumCPU_AllCC_r19 = v1
	}

	return nil
}
