// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: WLAN-RTT-r16 (line 26279).

var wLANRTTR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rttValue-r16"},
		{Name: "rttUnits-r16"},
		{Name: "rttAccuracy-r16", Optional: true},
	},
}

var wLANRTTR16RttValueR16Constraints = per.Constrained(0, 16777215)

const (
	WLAN_RTT_r16_RttUnits_r16_Microseconds          = 0
	WLAN_RTT_r16_RttUnits_r16_Hundredsofnanoseconds = 1
	WLAN_RTT_r16_RttUnits_r16_Tensofnanoseconds     = 2
	WLAN_RTT_r16_RttUnits_r16_Nanoseconds           = 3
	WLAN_RTT_r16_RttUnits_r16_Tenthsofnanoseconds   = 4
)

var wLANRTTR16RttUnitsR16Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{WLAN_RTT_r16_RttUnits_r16_Microseconds, WLAN_RTT_r16_RttUnits_r16_Hundredsofnanoseconds, WLAN_RTT_r16_RttUnits_r16_Tensofnanoseconds, WLAN_RTT_r16_RttUnits_r16_Nanoseconds, WLAN_RTT_r16_RttUnits_r16_Tenthsofnanoseconds},
}

var wLANRTTR16RttAccuracyR16Constraints = per.Constrained(0, 255)

type WLAN_RTT_r16 struct {
	RttValue_r16    int64
	RttUnits_r16    int64
	RttAccuracy_r16 *int64
}

func (ie *WLAN_RTT_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(wLANRTTR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RttAccuracy_r16 != nil}); err != nil {
		return err
	}

	// 3. rttValue-r16: integer
	{
		if err := e.EncodeInteger(ie.RttValue_r16, wLANRTTR16RttValueR16Constraints); err != nil {
			return err
		}
	}

	// 4. rttUnits-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.RttUnits_r16, wLANRTTR16RttUnitsR16Constraints); err != nil {
			return err
		}
	}

	// 5. rttAccuracy-r16: integer
	{
		if ie.RttAccuracy_r16 != nil {
			if err := e.EncodeInteger(*ie.RttAccuracy_r16, wLANRTTR16RttAccuracyR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *WLAN_RTT_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(wLANRTTR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rttValue-r16: integer
	{
		v0, err := d.DecodeInteger(wLANRTTR16RttValueR16Constraints)
		if err != nil {
			return err
		}
		ie.RttValue_r16 = v0
	}

	// 4. rttUnits-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(wLANRTTR16RttUnitsR16Constraints)
		if err != nil {
			return err
		}
		ie.RttUnits_r16 = v1
	}

	// 5. rttAccuracy-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(wLANRTTR16RttAccuracyR16Constraints)
			if err != nil {
				return err
			}
			ie.RttAccuracy_r16 = &v
		}
	}

	return nil
}
