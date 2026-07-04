// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MinMaxMCS-Config-r16 (line 28098).

var sLMinMaxMCSConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MCS-Table-r16"},
		{Name: "sl-MinMCS-PSSCH-r16"},
		{Name: "sl-MaxMCS-PSSCH-r16"},
	},
}

const (
	SL_MinMaxMCS_Config_r16_Sl_MCS_Table_r16_Qam64      = 0
	SL_MinMaxMCS_Config_r16_Sl_MCS_Table_r16_Qam256     = 1
	SL_MinMaxMCS_Config_r16_Sl_MCS_Table_r16_Qam64LowSE = 2
)

var sLMinMaxMCSConfigR16SlMCSTableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_MinMaxMCS_Config_r16_Sl_MCS_Table_r16_Qam64, SL_MinMaxMCS_Config_r16_Sl_MCS_Table_r16_Qam256, SL_MinMaxMCS_Config_r16_Sl_MCS_Table_r16_Qam64LowSE},
}

var sLMinMaxMCSConfigR16SlMinMCSPSSCHR16Constraints = per.Constrained(0, 27)

var sLMinMaxMCSConfigR16SlMaxMCSPSSCHR16Constraints = per.Constrained(0, 31)

type SL_MinMaxMCS_Config_r16 struct {
	Sl_MCS_Table_r16    int64
	Sl_MinMCS_PSSCH_r16 int64
	Sl_MaxMCS_PSSCH_r16 int64
}

func (ie *SL_MinMaxMCS_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMinMaxMCSConfigR16Constraints)
	_ = seq

	// 1. sl-MCS-Table-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_MCS_Table_r16, sLMinMaxMCSConfigR16SlMCSTableR16Constraints); err != nil {
			return err
		}
	}

	// 2. sl-MinMCS-PSSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MinMCS_PSSCH_r16, sLMinMaxMCSConfigR16SlMinMCSPSSCHR16Constraints); err != nil {
			return err
		}
	}

	// 3. sl-MaxMCS-PSSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MaxMCS_PSSCH_r16, sLMinMaxMCSConfigR16SlMaxMCSPSSCHR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_MinMaxMCS_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMinMaxMCSConfigR16Constraints)
	_ = seq

	// 1. sl-MCS-Table-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sLMinMaxMCSConfigR16SlMCSTableR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MCS_Table_r16 = v0
	}

	// 2. sl-MinMCS-PSSCH-r16: integer
	{
		v1, err := d.DecodeInteger(sLMinMaxMCSConfigR16SlMinMCSPSSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MinMCS_PSSCH_r16 = v1
	}

	// 3. sl-MaxMCS-PSSCH-r16: integer
	{
		v2, err := d.DecodeInteger(sLMinMaxMCSConfigR16SlMaxMCSPSSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MaxMCS_PSSCH_r16 = v2
	}

	return nil
}
