// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultCBR-DedicatedSL-PRS-r18 (line 10031).

var measResultCBRDedicatedSLPRSR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-ResourcePoolID-r18"},
		{Name: "sL-CBR-ResultsDedicatedSL-PRS-RP-r18"},
	},
}

type MeasResultCBR_DedicatedSL_PRS_r18 struct {
	Sl_PRS_ResourcePoolID_r18            SL_PRS_ResourcePoolID_r18
	SL_CBR_ResultsDedicatedSL_PRS_RP_r18 SL_CBR_r16
}

func (ie *MeasResultCBR_DedicatedSL_PRS_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultCBRDedicatedSLPRSR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-PRS-ResourcePoolID-r18: ref
	{
		if err := ie.Sl_PRS_ResourcePoolID_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. sL-CBR-ResultsDedicatedSL-PRS-RP-r18: ref
	{
		if err := ie.SL_CBR_ResultsDedicatedSL_PRS_RP_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultCBR_DedicatedSL_PRS_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultCBRDedicatedSLPRSR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-PRS-ResourcePoolID-r18: ref
	{
		if err := ie.Sl_PRS_ResourcePoolID_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. sL-CBR-ResultsDedicatedSL-PRS-RP-r18: ref
	{
		if err := ie.SL_CBR_ResultsDedicatedSL_PRS_RP_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
