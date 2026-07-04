// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TxPercentageDedicatedSL-PRS-RP-Config-r18 (line 27679).

var sLTxPercentageDedicatedSLPRSRPConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TxPercentageDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-Priority-DedicatedSL-PRS-RP", Optional: true},
	},
}

var sLTxPercentageDedicatedSLPRSRPConfigR18SlTxPercentageDedicatedSLPRSRPR18Constraints = per.Constrained(1, 8)

const (
	SL_TxPercentageDedicatedSL_PRS_RP_Config_r18_Sl_Priority_DedicatedSL_PRS_RP_P20 = 0
	SL_TxPercentageDedicatedSL_PRS_RP_Config_r18_Sl_Priority_DedicatedSL_PRS_RP_P35 = 1
	SL_TxPercentageDedicatedSL_PRS_RP_Config_r18_Sl_Priority_DedicatedSL_PRS_RP_P50 = 2
)

var sLTxPercentageDedicatedSLPRSRPConfigR18SlPriorityDedicatedSLPRSRPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxPercentageDedicatedSL_PRS_RP_Config_r18_Sl_Priority_DedicatedSL_PRS_RP_P20, SL_TxPercentageDedicatedSL_PRS_RP_Config_r18_Sl_Priority_DedicatedSL_PRS_RP_P35, SL_TxPercentageDedicatedSL_PRS_RP_Config_r18_Sl_Priority_DedicatedSL_PRS_RP_P50},
}

type SL_TxPercentageDedicatedSL_PRS_RP_Config_r18 struct {
	Sl_TxPercentageDedicatedSL_PRS_RP_r18 *int64
	Sl_Priority_DedicatedSL_PRS_RP        *int64
}

func (ie *SL_TxPercentageDedicatedSL_PRS_RP_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxPercentageDedicatedSLPRSRPConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TxPercentageDedicatedSL_PRS_RP_r18 != nil, ie.Sl_Priority_DedicatedSL_PRS_RP != nil}); err != nil {
		return err
	}

	// 2. sl-TxPercentageDedicatedSL-PRS-RP-r18: integer
	{
		if ie.Sl_TxPercentageDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_TxPercentageDedicatedSL_PRS_RP_r18, sLTxPercentageDedicatedSLPRSRPConfigR18SlTxPercentageDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-Priority-DedicatedSL-PRS-RP: enumerated
	{
		if ie.Sl_Priority_DedicatedSL_PRS_RP != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Priority_DedicatedSL_PRS_RP, sLTxPercentageDedicatedSLPRSRPConfigR18SlPriorityDedicatedSLPRSRPConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_TxPercentageDedicatedSL_PRS_RP_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxPercentageDedicatedSLPRSRPConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-TxPercentageDedicatedSL-PRS-RP-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLTxPercentageDedicatedSLPRSRPConfigR18SlTxPercentageDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TxPercentageDedicatedSL_PRS_RP_r18 = &v
		}
	}

	// 3. sl-Priority-DedicatedSL-PRS-RP: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLTxPercentageDedicatedSLPRSRPConfigR18SlPriorityDedicatedSLPRSRPConstraints)
			if err != nil {
				return err
			}
			ie.Sl_Priority_DedicatedSL_PRS_RP = &idx
		}
	}

	return nil
}
