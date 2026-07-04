// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DRX-GC-Generic-r17 (line 27159).

var sLDRXGCGenericR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DRX-GC-HARQ-RTT-Timer1-r17", Optional: true},
		{Name: "sl-DRX-GC-HARQ-RTT-Timer2-r17", Optional: true},
		{Name: "sl-DRX-GC-RetransmissionTimer-r17"},
	},
}

const (
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl0    = 0
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl1    = 1
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl2    = 2
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl4    = 3
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare4 = 4
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare3 = 5
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare2 = 6
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare1 = 7
)

var sLDRXGCGenericR17SlDRXGCHARQRTTTimer1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl0, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl1, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl2, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Sl4, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare4, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare3, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare2, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer1_r17_Spare1},
}

const (
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl0    = 0
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl1    = 1
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl2    = 2
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl4    = 3
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare4 = 4
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare3 = 5
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare2 = 6
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare1 = 7
)

var sLDRXGCGenericR17SlDRXGCHARQRTTTimer2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl0, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl1, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl2, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Sl4, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare4, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare3, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare2, SL_DRX_GC_Generic_r17_Sl_DRX_GC_HARQ_RTT_Timer2_r17_Spare1},
}

const (
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl0     = 0
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl1     = 1
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl2     = 2
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl4     = 3
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl6     = 4
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl8     = 5
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl16    = 6
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl24    = 7
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl33    = 8
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl40    = 9
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl64    = 10
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl80    = 11
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl96    = 12
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl112   = 13
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl128   = 14
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl160   = 15
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl320   = 16
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare15 = 17
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare14 = 18
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare13 = 19
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare12 = 20
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare11 = 21
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare10 = 22
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare9  = 23
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare8  = 24
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare7  = 25
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare6  = 26
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare5  = 27
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare4  = 28
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare3  = 29
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare2  = 30
	SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare1  = 31
)

var sLDRXGCGenericR17SlDRXGCRetransmissionTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl0, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl1, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl2, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl4, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl6, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl8, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl16, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl24, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl33, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl40, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl64, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl80, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl96, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl112, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl128, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl160, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Sl320, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare15, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare14, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare13, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare12, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare11, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare10, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare9, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare8, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare7, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare6, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare5, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare4, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare3, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare2, SL_DRX_GC_Generic_r17_Sl_DRX_GC_RetransmissionTimer_r17_Spare1},
}

type SL_DRX_GC_Generic_r17 struct {
	Sl_DRX_GC_HARQ_RTT_Timer1_r17     *int64
	Sl_DRX_GC_HARQ_RTT_Timer2_r17     *int64
	Sl_DRX_GC_RetransmissionTimer_r17 int64
}

func (ie *SL_DRX_GC_Generic_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDRXGCGenericR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DRX_GC_HARQ_RTT_Timer1_r17 != nil, ie.Sl_DRX_GC_HARQ_RTT_Timer2_r17 != nil}); err != nil {
		return err
	}

	// 2. sl-DRX-GC-HARQ-RTT-Timer1-r17: enumerated
	{
		if ie.Sl_DRX_GC_HARQ_RTT_Timer1_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_DRX_GC_HARQ_RTT_Timer1_r17, sLDRXGCGenericR17SlDRXGCHARQRTTTimer1R17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-DRX-GC-HARQ-RTT-Timer2-r17: enumerated
	{
		if ie.Sl_DRX_GC_HARQ_RTT_Timer2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_DRX_GC_HARQ_RTT_Timer2_r17, sLDRXGCGenericR17SlDRXGCHARQRTTTimer2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-DRX-GC-RetransmissionTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_DRX_GC_RetransmissionTimer_r17, sLDRXGCGenericR17SlDRXGCRetransmissionTimerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_DRX_GC_Generic_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDRXGCGenericR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-DRX-GC-HARQ-RTT-Timer1-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLDRXGCGenericR17SlDRXGCHARQRTTTimer1R17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DRX_GC_HARQ_RTT_Timer1_r17 = &idx
		}
	}

	// 3. sl-DRX-GC-HARQ-RTT-Timer2-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLDRXGCGenericR17SlDRXGCHARQRTTTimer2R17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DRX_GC_HARQ_RTT_Timer2_r17 = &idx
		}
	}

	// 4. sl-DRX-GC-RetransmissionTimer-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(sLDRXGCGenericR17SlDRXGCRetransmissionTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_DRX_GC_RetransmissionTimer_r17 = v2
	}

	return nil
}
