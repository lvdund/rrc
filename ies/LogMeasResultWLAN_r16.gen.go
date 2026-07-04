// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LogMeasResultWLAN-r16 (line 26263).

var logMeasResultWLANR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "wlan-Identifiers-r16"},
		{Name: "rssiWLAN-r16", Optional: true},
		{Name: "rtt-WLAN-r16", Optional: true},
	},
}

type LogMeasResultWLAN_r16 struct {
	Wlan_Identifiers_r16 WLAN_Identifiers_r16
	RssiWLAN_r16         *WLAN_RSSI_Range_r16
	Rtt_WLAN_r16         *WLAN_RTT_r16
}

func (ie *LogMeasResultWLAN_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(logMeasResultWLANR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RssiWLAN_r16 != nil, ie.Rtt_WLAN_r16 != nil}); err != nil {
		return err
	}

	// 3. wlan-Identifiers-r16: ref
	{
		if err := ie.Wlan_Identifiers_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. rssiWLAN-r16: ref
	{
		if ie.RssiWLAN_r16 != nil {
			if err := ie.RssiWLAN_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. rtt-WLAN-r16: ref
	{
		if ie.Rtt_WLAN_r16 != nil {
			if err := ie.Rtt_WLAN_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LogMeasResultWLAN_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(logMeasResultWLANR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. wlan-Identifiers-r16: ref
	{
		if err := ie.Wlan_Identifiers_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. rssiWLAN-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.RssiWLAN_r16 = new(WLAN_RSSI_Range_r16)
			if err := ie.RssiWLAN_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. rtt-WLAN-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Rtt_WLAN_r16 = new(WLAN_RTT_r16)
			if err := ie.Rtt_WLAN_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
