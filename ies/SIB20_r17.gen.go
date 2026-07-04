// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB20-r17 (line 4557).

var sIB20R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mcch-Config-r17"},
		{Name: "cfr-ConfigMCCH-MTCH-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sIB20R17LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB20_r17 struct {
	Mcch_Config_r17                MCCH_Config_r17
	Cfr_ConfigMCCH_MTCH_r17        *CFR_ConfigMCCH_MTCH_r17
	LateNonCriticalExtension       []byte
	Cfr_ConfigMCCH_MTCH_RedCap_r18 *CFR_ConfigMCCH_MTCH_r17
	Mcch_ConfigRedCap_r18          *MCCH_Config_r17
}

func (ie *SIB20_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB20R17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Cfr_ConfigMCCH_MTCH_RedCap_r18 != nil || ie.Mcch_ConfigRedCap_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cfr_ConfigMCCH_MTCH_r17 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. mcch-Config-r17: ref
	{
		if err := ie.Mcch_Config_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. cfr-ConfigMCCH-MTCH-r17: ref
	{
		if ie.Cfr_ConfigMCCH_MTCH_r17 != nil {
			if err := ie.Cfr_ConfigMCCH_MTCH_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB20R17LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cfr-ConfigMCCH-MTCH-RedCap-r18", Optional: true},
					{Name: "mcch-ConfigRedCap-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cfr_ConfigMCCH_MTCH_RedCap_r18 != nil, ie.Mcch_ConfigRedCap_r18 != nil}); err != nil {
				return err
			}

			if ie.Cfr_ConfigMCCH_MTCH_RedCap_r18 != nil {
				if err := ie.Cfr_ConfigMCCH_MTCH_RedCap_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Mcch_ConfigRedCap_r18 != nil {
				if err := ie.Mcch_ConfigRedCap_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SIB20_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB20R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mcch-Config-r17: ref
	{
		if err := ie.Mcch_Config_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. cfr-ConfigMCCH-MTCH-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Cfr_ConfigMCCH_MTCH_r17 = new(CFR_ConfigMCCH_MTCH_r17)
			if err := ie.Cfr_ConfigMCCH_MTCH_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB20R17LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cfr-ConfigMCCH-MTCH-RedCap-r18", Optional: true},
				{Name: "mcch-ConfigRedCap-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cfr_ConfigMCCH_MTCH_RedCap_r18 = new(CFR_ConfigMCCH_MTCH_r17)
			if err := ie.Cfr_ConfigMCCH_MTCH_RedCap_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Mcch_ConfigRedCap_r18 = new(MCCH_Config_r17)
			if err := ie.Mcch_ConfigRedCap_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
