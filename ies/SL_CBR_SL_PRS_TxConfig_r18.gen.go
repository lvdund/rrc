// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CBR-SL-PRS-TxConfig-r18 (line 26934).

var sLCBRSLPRSTxConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-CR-Limit-r18", Optional: true},
		{Name: "sl-PRS-MaxTx-power-r18", Optional: true},
		{Name: "sl-PRS-MaxNumTransmissions-r18", Optional: true},
	},
}

var sLCBRSLPRSTxConfigR18SlPRSCRLimitR18Constraints = per.Constrained(0, 10000)

var sLCBRSLPRSTxConfigR18SlPRSMaxTxPowerR18Constraints = per.Constrained(-30, 33)

var sLCBRSLPRSTxConfigR18SlPRSMaxNumTransmissionsR18Constraints = per.Constrained(1, 32)

type SL_CBR_SL_PRS_TxConfig_r18 struct {
	Sl_PRS_CR_Limit_r18            *int64
	Sl_PRS_MaxTx_Power_r18         *int64
	Sl_PRS_MaxNumTransmissions_r18 *int64
}

func (ie *SL_CBR_SL_PRS_TxConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLCBRSLPRSTxConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_CR_Limit_r18 != nil, ie.Sl_PRS_MaxTx_Power_r18 != nil, ie.Sl_PRS_MaxNumTransmissions_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-CR-Limit-r18: integer
	{
		if ie.Sl_PRS_CR_Limit_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_CR_Limit_r18, sLCBRSLPRSTxConfigR18SlPRSCRLimitR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-PRS-MaxTx-power-r18: integer
	{
		if ie.Sl_PRS_MaxTx_Power_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_MaxTx_Power_r18, sLCBRSLPRSTxConfigR18SlPRSMaxTxPowerR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-PRS-MaxNumTransmissions-r18: integer
	{
		if ie.Sl_PRS_MaxNumTransmissions_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_MaxNumTransmissions_r18, sLCBRSLPRSTxConfigR18SlPRSMaxNumTransmissionsR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_CBR_SL_PRS_TxConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLCBRSLPRSTxConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-CR-Limit-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLCBRSLPRSTxConfigR18SlPRSCRLimitR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_CR_Limit_r18 = &v
		}
	}

	// 3. sl-PRS-MaxTx-power-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLCBRSLPRSTxConfigR18SlPRSMaxTxPowerR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_MaxTx_Power_r18 = &v
		}
	}

	// 4. sl-PRS-MaxNumTransmissions-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLCBRSLPRSTxConfigR18SlPRSMaxNumTransmissionsR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_MaxNumTransmissions_r18 = &v
		}
	}

	return nil
}
