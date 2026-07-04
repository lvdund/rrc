// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: WLAN-Identifiers-r16 (line 26270).

var wLANIdentifiersR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssid-r16", Optional: true},
		{Name: "bssid-r16", Optional: true},
		{Name: "hessid-r16", Optional: true},
	},
}

var wLANIdentifiersR16SsidR16Constraints = per.SizeRange(1, 32)

var wLANIdentifiersR16BssidR16Constraints = per.FixedSize(6)

var wLANIdentifiersR16HessidR16Constraints = per.FixedSize(6)

type WLAN_Identifiers_r16 struct {
	Ssid_r16   []byte
	Bssid_r16  []byte
	Hessid_r16 []byte
}

func (ie *WLAN_Identifiers_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(wLANIdentifiersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssid_r16 != nil, ie.Bssid_r16 != nil, ie.Hessid_r16 != nil}); err != nil {
		return err
	}

	// 3. ssid-r16: octet-string
	{
		if ie.Ssid_r16 != nil {
			if err := e.EncodeOctetString(ie.Ssid_r16, wLANIdentifiersR16SsidR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. bssid-r16: octet-string
	{
		if ie.Bssid_r16 != nil {
			if err := e.EncodeOctetString(ie.Bssid_r16, wLANIdentifiersR16BssidR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. hessid-r16: octet-string
	{
		if ie.Hessid_r16 != nil {
			if err := e.EncodeOctetString(ie.Hessid_r16, wLANIdentifiersR16HessidR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *WLAN_Identifiers_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(wLANIdentifiersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ssid-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(wLANIdentifiersR16SsidR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssid_r16 = v
		}
	}

	// 4. bssid-r16: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(wLANIdentifiersR16BssidR16Constraints)
			if err != nil {
				return err
			}
			ie.Bssid_r16 = v
		}
	}

	// 5. hessid-r16: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(wLANIdentifiersR16HessidR16Constraints)
			if err != nil {
				return err
			}
			ie.Hessid_r16 = v
		}
	}

	return nil
}
