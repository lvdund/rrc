// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BCCH-Config (line 7853).

var bCCHConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "modificationPeriodCoeff"},
	},
}

const (
	BCCH_Config_ModificationPeriodCoeff_N2  = 0
	BCCH_Config_ModificationPeriodCoeff_N4  = 1
	BCCH_Config_ModificationPeriodCoeff_N8  = 2
	BCCH_Config_ModificationPeriodCoeff_N16 = 3
)

var bCCHConfigModificationPeriodCoeffConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BCCH_Config_ModificationPeriodCoeff_N2, BCCH_Config_ModificationPeriodCoeff_N4, BCCH_Config_ModificationPeriodCoeff_N8, BCCH_Config_ModificationPeriodCoeff_N16},
}

type BCCH_Config struct {
	ModificationPeriodCoeff int64
}

func (ie *BCCH_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bCCHConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. modificationPeriodCoeff: enumerated
	{
		if err := e.EncodeEnumerated(ie.ModificationPeriodCoeff, bCCHConfigModificationPeriodCoeffConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BCCH_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bCCHConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. modificationPeriodCoeff: enumerated
	{
		v0, err := d.DecodeEnumerated(bCCHConfigModificationPeriodCoeffConstraints)
		if err != nil {
			return err
		}
		ie.ModificationPeriodCoeff = v0
	}

	return nil
}
