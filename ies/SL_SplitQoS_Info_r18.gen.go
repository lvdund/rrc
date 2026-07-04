// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SplitQoS-Info-r18 (line 2312).

var sLSplitQoSInfoR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-QoS-FlowIdentity-r18"},
		{Name: "sl-SplitPacketDelayBudget-r18", Optional: true},
	},
}

var sLSplitQoSInfoR18SlSplitPacketDelayBudgetR18Constraints = per.Constrained(0, 1023)

type SL_SplitQoS_Info_r18 struct {
	Sl_QoS_FlowIdentity_r18       SL_QoS_FlowIdentity_r16
	Sl_SplitPacketDelayBudget_r18 *int64
}

func (ie *SL_SplitQoS_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSplitQoSInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_SplitPacketDelayBudget_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-QoS-FlowIdentity-r18: ref
	{
		if err := ie.Sl_QoS_FlowIdentity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-SplitPacketDelayBudget-r18: integer
	{
		if ie.Sl_SplitPacketDelayBudget_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_SplitPacketDelayBudget_r18, sLSplitQoSInfoR18SlSplitPacketDelayBudgetR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_SplitQoS_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSplitQoSInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-QoS-FlowIdentity-r18: ref
	{
		if err := ie.Sl_QoS_FlowIdentity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-SplitPacketDelayBudget-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLSplitQoSInfoR18SlSplitPacketDelayBudgetR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SplitPacketDelayBudget_r18 = &v
		}
	}

	return nil
}
