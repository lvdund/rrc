// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CBR-PSSCH-TxConfig-r16 (line 26915).

var sLCBRPSSCHTxConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-CR-Limit-r16", Optional: true},
		{Name: "sl-TxParameters-r16", Optional: true},
	},
}

var sLCBRPSSCHTxConfigR16SlCRLimitR16Constraints = per.Constrained(0, 10000)

type SL_CBR_PSSCH_TxConfig_r16 struct {
	Sl_CR_Limit_r16     *int64
	Sl_TxParameters_r16 *SL_PSSCH_TxParameters_r16
}

func (ie *SL_CBR_PSSCH_TxConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLCBRPSSCHTxConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_CR_Limit_r16 != nil, ie.Sl_TxParameters_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-CR-Limit-r16: integer
	{
		if ie.Sl_CR_Limit_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_CR_Limit_r16, sLCBRPSSCHTxConfigR16SlCRLimitR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-TxParameters-r16: ref
	{
		if ie.Sl_TxParameters_r16 != nil {
			if err := ie.Sl_TxParameters_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_CBR_PSSCH_TxConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLCBRPSSCHTxConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-CR-Limit-r16: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLCBRPSSCHTxConfigR16SlCRLimitR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CR_Limit_r16 = &v
		}
	}

	// 3. sl-TxParameters-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_TxParameters_r16 = new(SL_PSSCH_TxParameters_r16)
			if err := ie.Sl_TxParameters_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
