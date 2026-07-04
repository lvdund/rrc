// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-BWP-ConfigCommon-r16 (line 26791).

var sLBWPConfigCommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-BWP-Generic-r16", Optional: true},
		{Name: "sl-BWP-PoolConfigCommon-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

type SL_BWP_ConfigCommon_r16 struct {
	Sl_BWP_Generic_r16              *SL_BWP_Generic_r16
	Sl_BWP_PoolConfigCommon_r16     *SL_BWP_PoolConfigCommon_r16
	Sl_BWP_PoolConfigCommonPS_r17   *SL_BWP_PoolConfigCommon_r16
	Sl_BWP_DiscPoolConfigCommon_r17 *SL_BWP_DiscPoolConfigCommon_r17
	Sl_BWP_PoolConfigCommonA2X_r18  *SL_BWP_PoolConfigCommon_r16
}

func (ie *SL_BWP_ConfigCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLBWPConfigCommonR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_BWP_PoolConfigCommonPS_r17 != nil || ie.Sl_BWP_DiscPoolConfigCommon_r17 != nil
	hasExtGroup1 := ie.Sl_BWP_PoolConfigCommonA2X_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_BWP_Generic_r16 != nil, ie.Sl_BWP_PoolConfigCommon_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-BWP-Generic-r16: ref
	{
		if ie.Sl_BWP_Generic_r16 != nil {
			if err := ie.Sl_BWP_Generic_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-BWP-PoolConfigCommon-r16: ref
	{
		if ie.Sl_BWP_PoolConfigCommon_r16 != nil {
			if err := ie.Sl_BWP_PoolConfigCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-BWP-PoolConfigCommonPS-r17", Optional: true},
					{Name: "sl-BWP-DiscPoolConfigCommon-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_BWP_PoolConfigCommonPS_r17 != nil, ie.Sl_BWP_DiscPoolConfigCommon_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_BWP_PoolConfigCommonPS_r17 != nil {
				if err := ie.Sl_BWP_PoolConfigCommonPS_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_BWP_DiscPoolConfigCommon_r17 != nil {
				if err := ie.Sl_BWP_DiscPoolConfigCommon_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-BWP-PoolConfigCommonA2X-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_BWP_PoolConfigCommonA2X_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_BWP_PoolConfigCommonA2X_r18 != nil {
				if err := ie.Sl_BWP_PoolConfigCommonA2X_r18.Encode(ex); err != nil {
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

func (ie *SL_BWP_ConfigCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLBWPConfigCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-BWP-Generic-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_BWP_Generic_r16 = new(SL_BWP_Generic_r16)
			if err := ie.Sl_BWP_Generic_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-BWP-PoolConfigCommon-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_BWP_PoolConfigCommon_r16 = new(SL_BWP_PoolConfigCommon_r16)
			if err := ie.Sl_BWP_PoolConfigCommon_r16.Decode(d); err != nil {
				return err
			}
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
				{Name: "sl-BWP-PoolConfigCommonPS-r17", Optional: true},
				{Name: "sl-BWP-DiscPoolConfigCommon-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_BWP_PoolConfigCommonPS_r17 = new(SL_BWP_PoolConfigCommon_r16)
			if err := ie.Sl_BWP_PoolConfigCommonPS_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_BWP_DiscPoolConfigCommon_r17 = new(SL_BWP_DiscPoolConfigCommon_r17)
			if err := ie.Sl_BWP_DiscPoolConfigCommon_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-BWP-PoolConfigCommonA2X-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_BWP_PoolConfigCommonA2X_r18 = new(SL_BWP_PoolConfigCommon_r16)
			if err := ie.Sl_BWP_PoolConfigCommonA2X_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
