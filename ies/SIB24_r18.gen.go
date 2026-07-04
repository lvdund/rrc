// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB24-r18 (line 4660).

var sIB24R18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "multicastMCCH-Config-r18", Optional: true},
		{Name: "cfr-ConfigMCCH-MTCH-r18", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB24R18LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB24_r18 struct {
	MulticastMCCH_Config_r18 *MCCH_Config_r17
	Cfr_ConfigMCCH_MTCH_r18  *CFR_ConfigMCCH_MTCH_r17
	LateNonCriticalExtension []byte
}

func (ie *SIB24_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB24R18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MulticastMCCH_Config_r18 != nil, ie.Cfr_ConfigMCCH_MTCH_r18 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. multicastMCCH-Config-r18: ref
	{
		if ie.MulticastMCCH_Config_r18 != nil {
			if err := ie.MulticastMCCH_Config_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. cfr-ConfigMCCH-MTCH-r18: ref
	{
		if ie.Cfr_ConfigMCCH_MTCH_r18 != nil {
			if err := ie.Cfr_ConfigMCCH_MTCH_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB24R18LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB24_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB24R18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. multicastMCCH-Config-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MulticastMCCH_Config_r18 = new(MCCH_Config_r17)
			if err := ie.MulticastMCCH_Config_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. cfr-ConfigMCCH-MTCH-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Cfr_ConfigMCCH_MTCH_r18 = new(CFR_ConfigMCCH_MTCH_r17)
			if err := ie.Cfr_ConfigMCCH_MTCH_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB24R18LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
