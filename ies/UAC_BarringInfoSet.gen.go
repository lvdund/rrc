// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UAC-BarringInfoSet (line 16182).

var uACBarringInfoSetConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uac-BarringFactor"},
		{Name: "uac-BarringTime"},
		{Name: "uac-BarringForAccessIdentity"},
	},
}

const (
	UAC_BarringInfoSet_Uac_BarringFactor_P00 = 0
	UAC_BarringInfoSet_Uac_BarringFactor_P05 = 1
	UAC_BarringInfoSet_Uac_BarringFactor_P10 = 2
	UAC_BarringInfoSet_Uac_BarringFactor_P15 = 3
	UAC_BarringInfoSet_Uac_BarringFactor_P20 = 4
	UAC_BarringInfoSet_Uac_BarringFactor_P25 = 5
	UAC_BarringInfoSet_Uac_BarringFactor_P30 = 6
	UAC_BarringInfoSet_Uac_BarringFactor_P40 = 7
	UAC_BarringInfoSet_Uac_BarringFactor_P50 = 8
	UAC_BarringInfoSet_Uac_BarringFactor_P60 = 9
	UAC_BarringInfoSet_Uac_BarringFactor_P70 = 10
	UAC_BarringInfoSet_Uac_BarringFactor_P75 = 11
	UAC_BarringInfoSet_Uac_BarringFactor_P80 = 12
	UAC_BarringInfoSet_Uac_BarringFactor_P85 = 13
	UAC_BarringInfoSet_Uac_BarringFactor_P90 = 14
	UAC_BarringInfoSet_Uac_BarringFactor_P95 = 15
)

var uACBarringInfoSetUacBarringFactorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UAC_BarringInfoSet_Uac_BarringFactor_P00, UAC_BarringInfoSet_Uac_BarringFactor_P05, UAC_BarringInfoSet_Uac_BarringFactor_P10, UAC_BarringInfoSet_Uac_BarringFactor_P15, UAC_BarringInfoSet_Uac_BarringFactor_P20, UAC_BarringInfoSet_Uac_BarringFactor_P25, UAC_BarringInfoSet_Uac_BarringFactor_P30, UAC_BarringInfoSet_Uac_BarringFactor_P40, UAC_BarringInfoSet_Uac_BarringFactor_P50, UAC_BarringInfoSet_Uac_BarringFactor_P60, UAC_BarringInfoSet_Uac_BarringFactor_P70, UAC_BarringInfoSet_Uac_BarringFactor_P75, UAC_BarringInfoSet_Uac_BarringFactor_P80, UAC_BarringInfoSet_Uac_BarringFactor_P85, UAC_BarringInfoSet_Uac_BarringFactor_P90, UAC_BarringInfoSet_Uac_BarringFactor_P95},
}

const (
	UAC_BarringInfoSet_Uac_BarringTime_S4   = 0
	UAC_BarringInfoSet_Uac_BarringTime_S8   = 1
	UAC_BarringInfoSet_Uac_BarringTime_S16  = 2
	UAC_BarringInfoSet_Uac_BarringTime_S32  = 3
	UAC_BarringInfoSet_Uac_BarringTime_S64  = 4
	UAC_BarringInfoSet_Uac_BarringTime_S128 = 5
	UAC_BarringInfoSet_Uac_BarringTime_S256 = 6
	UAC_BarringInfoSet_Uac_BarringTime_S512 = 7
)

var uACBarringInfoSetUacBarringTimeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UAC_BarringInfoSet_Uac_BarringTime_S4, UAC_BarringInfoSet_Uac_BarringTime_S8, UAC_BarringInfoSet_Uac_BarringTime_S16, UAC_BarringInfoSet_Uac_BarringTime_S32, UAC_BarringInfoSet_Uac_BarringTime_S64, UAC_BarringInfoSet_Uac_BarringTime_S128, UAC_BarringInfoSet_Uac_BarringTime_S256, UAC_BarringInfoSet_Uac_BarringTime_S512},
}

var uACBarringInfoSetUacBarringForAccessIdentityConstraints = per.FixedSize(7)

type UAC_BarringInfoSet struct {
	Uac_BarringFactor            int64
	Uac_BarringTime              int64
	Uac_BarringForAccessIdentity per.BitString
}

func (ie *UAC_BarringInfoSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uACBarringInfoSetConstraints)
	_ = seq

	// 1. uac-BarringFactor: enumerated
	{
		if err := e.EncodeEnumerated(ie.Uac_BarringFactor, uACBarringInfoSetUacBarringFactorConstraints); err != nil {
			return err
		}
	}

	// 2. uac-BarringTime: enumerated
	{
		if err := e.EncodeEnumerated(ie.Uac_BarringTime, uACBarringInfoSetUacBarringTimeConstraints); err != nil {
			return err
		}
	}

	// 3. uac-BarringForAccessIdentity: bit-string
	{
		if err := e.EncodeBitString(ie.Uac_BarringForAccessIdentity, uACBarringInfoSetUacBarringForAccessIdentityConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UAC_BarringInfoSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uACBarringInfoSetConstraints)
	_ = seq

	// 1. uac-BarringFactor: enumerated
	{
		v0, err := d.DecodeEnumerated(uACBarringInfoSetUacBarringFactorConstraints)
		if err != nil {
			return err
		}
		ie.Uac_BarringFactor = v0
	}

	// 2. uac-BarringTime: enumerated
	{
		v1, err := d.DecodeEnumerated(uACBarringInfoSetUacBarringTimeConstraints)
		if err != nil {
			return err
		}
		ie.Uac_BarringTime = v1
	}

	// 3. uac-BarringForAccessIdentity: bit-string
	{
		v2, err := d.DecodeBitString(uACBarringInfoSetUacBarringForAccessIdentityConstraints)
		if err != nil {
			return err
		}
		ie.Uac_BarringForAccessIdentity = v2
	}

	return nil
}
