// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PRS-QoS-Info-r18 (line 2323).

var sLPRSQoSInfoR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-Priority-r18", Optional: true},
		{Name: "sl-PRS-DelayBudget-r18", Optional: true},
		{Name: "sl-PRS-Bandwidth-r18", Optional: true},
	},
}

var sLPRSQoSInfoR18SlPRSPriorityR18Constraints = per.Constrained(1, 8)

var sLPRSQoSInfoR18SlPRSDelayBudgetR18Constraints = per.Constrained(0, 1023)

const (
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz5    = 0
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz10   = 1
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz15   = 2
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz20   = 3
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz25   = 4
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz30   = 5
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz35   = 6
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz40   = 7
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz45   = 8
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz50   = 9
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz60   = 10
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz70   = 11
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz80   = 12
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz90   = 13
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz100  = 14
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz200  = 15
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz400  = 16
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare15 = 17
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare14 = 18
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare13 = 19
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare12 = 20
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare11 = 21
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare10 = 22
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare9  = 23
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare8  = 24
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare7  = 25
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare6  = 26
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare5  = 27
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare4  = 28
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare3  = 29
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare2  = 30
	SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare1  = 31
)

var sLPRSQoSInfoR18SlPRSBandwidthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz5, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz10, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz15, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz20, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz25, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz30, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz35, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz40, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz45, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz50, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz60, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz70, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz80, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz90, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz100, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz200, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Mhz400, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare15, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare14, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare13, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare12, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare11, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare10, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare9, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare8, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare7, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare6, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare5, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare4, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare3, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare2, SL_PRS_QoS_Info_r18_Sl_PRS_Bandwidth_r18_Spare1},
}

type SL_PRS_QoS_Info_r18 struct {
	Sl_PRS_Priority_r18    *int64
	Sl_PRS_DelayBudget_r18 *int64
	Sl_PRS_Bandwidth_r18   *int64
}

func (ie *SL_PRS_QoS_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPRSQoSInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_Priority_r18 != nil, ie.Sl_PRS_DelayBudget_r18 != nil, ie.Sl_PRS_Bandwidth_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-PRS-Priority-r18: integer
	{
		if ie.Sl_PRS_Priority_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_Priority_r18, sLPRSQoSInfoR18SlPRSPriorityR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-PRS-DelayBudget-r18: integer
	{
		if ie.Sl_PRS_DelayBudget_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_DelayBudget_r18, sLPRSQoSInfoR18SlPRSDelayBudgetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-PRS-Bandwidth-r18: enumerated
	{
		if ie.Sl_PRS_Bandwidth_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PRS_Bandwidth_r18, sLPRSQoSInfoR18SlPRSBandwidthR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PRS_QoS_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPRSQoSInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PRS-Priority-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLPRSQoSInfoR18SlPRSPriorityR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_Priority_r18 = &v
		}
	}

	// 4. sl-PRS-DelayBudget-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLPRSQoSInfoR18SlPRSDelayBudgetR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_DelayBudget_r18 = &v
		}
	}

	// 5. sl-PRS-Bandwidth-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLPRSQoSInfoR18SlPRSBandwidthR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_Bandwidth_r18 = &idx
		}
	}

	return nil
}
