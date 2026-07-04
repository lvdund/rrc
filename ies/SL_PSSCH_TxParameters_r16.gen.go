// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PSSCH-TxParameters-r16 (line 27732).

var sLPSSCHTxParametersR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MinMCS-PSSCH-r16"},
		{Name: "sl-MaxMCS-PSSCH-r16"},
		{Name: "sl-MinSubChannelNumPSSCH-r16"},
		{Name: "sl-MaxSubchannelNumPSSCH-r16"},
		{Name: "sl-MaxTxTransNumPSSCH-r16"},
		{Name: "sl-MaxTxPower-r16", Optional: true},
	},
}

var sLPSSCHTxParametersR16SlMinMCSPSSCHR16Constraints = per.Constrained(0, 27)

var sLPSSCHTxParametersR16SlMaxMCSPSSCHR16Constraints = per.Constrained(0, 31)

var sLPSSCHTxParametersR16SlMinSubChannelNumPSSCHR16Constraints = per.Constrained(1, 27)

var sLPSSCHTxParametersR16SlMaxSubchannelNumPSSCHR16Constraints = per.Constrained(1, 27)

var sLPSSCHTxParametersR16SlMaxTxTransNumPSSCHR16Constraints = per.Constrained(1, 32)

type SL_PSSCH_TxParameters_r16 struct {
	Sl_MinMCS_PSSCH_r16          int64
	Sl_MaxMCS_PSSCH_r16          int64
	Sl_MinSubChannelNumPSSCH_r16 int64
	Sl_MaxSubchannelNumPSSCH_r16 int64
	Sl_MaxTxTransNumPSSCH_r16    int64
	Sl_MaxTxPower_r16            *SL_TxPower_r16
}

func (ie *SL_PSSCH_TxParameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPSSCHTxParametersR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_MaxTxPower_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-MinMCS-PSSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MinMCS_PSSCH_r16, sLPSSCHTxParametersR16SlMinMCSPSSCHR16Constraints); err != nil {
			return err
		}
	}

	// 3. sl-MaxMCS-PSSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MaxMCS_PSSCH_r16, sLPSSCHTxParametersR16SlMaxMCSPSSCHR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-MinSubChannelNumPSSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MinSubChannelNumPSSCH_r16, sLPSSCHTxParametersR16SlMinSubChannelNumPSSCHR16Constraints); err != nil {
			return err
		}
	}

	// 5. sl-MaxSubchannelNumPSSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MaxSubchannelNumPSSCH_r16, sLPSSCHTxParametersR16SlMaxSubchannelNumPSSCHR16Constraints); err != nil {
			return err
		}
	}

	// 6. sl-MaxTxTransNumPSSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MaxTxTransNumPSSCH_r16, sLPSSCHTxParametersR16SlMaxTxTransNumPSSCHR16Constraints); err != nil {
			return err
		}
	}

	// 7. sl-MaxTxPower-r16: ref
	{
		if ie.Sl_MaxTxPower_r16 != nil {
			if err := ie.Sl_MaxTxPower_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PSSCH_TxParameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPSSCHTxParametersR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-MinMCS-PSSCH-r16: integer
	{
		v0, err := d.DecodeInteger(sLPSSCHTxParametersR16SlMinMCSPSSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MinMCS_PSSCH_r16 = v0
	}

	// 3. sl-MaxMCS-PSSCH-r16: integer
	{
		v1, err := d.DecodeInteger(sLPSSCHTxParametersR16SlMaxMCSPSSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MaxMCS_PSSCH_r16 = v1
	}

	// 4. sl-MinSubChannelNumPSSCH-r16: integer
	{
		v2, err := d.DecodeInteger(sLPSSCHTxParametersR16SlMinSubChannelNumPSSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MinSubChannelNumPSSCH_r16 = v2
	}

	// 5. sl-MaxSubchannelNumPSSCH-r16: integer
	{
		v3, err := d.DecodeInteger(sLPSSCHTxParametersR16SlMaxSubchannelNumPSSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MaxSubchannelNumPSSCH_r16 = v3
	}

	// 6. sl-MaxTxTransNumPSSCH-r16: integer
	{
		v4, err := d.DecodeInteger(sLPSSCHTxParametersR16SlMaxTxTransNumPSSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MaxTxTransNumPSSCH_r16 = v4
	}

	// 7. sl-MaxTxPower-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Sl_MaxTxPower_r16 = new(SL_TxPower_r16)
			if err := ie.Sl_MaxTxPower_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
