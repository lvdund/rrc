// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasConfig-r16 (line 27518).

var sLMeasConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MeasObjectToRemoveList-r16", Optional: true},
		{Name: "sl-MeasObjectToAddModList-r16", Optional: true},
		{Name: "sl-ReportConfigToRemoveList-r16", Optional: true},
		{Name: "sl-ReportConfigToAddModList-r16", Optional: true},
		{Name: "sl-MeasIdToRemoveList-r16", Optional: true},
		{Name: "sl-MeasIdToAddModList-r16", Optional: true},
		{Name: "sl-QuantityConfig-r16", Optional: true},
	},
}

type SL_MeasConfig_r16 struct {
	Sl_MeasObjectToRemoveList_r16   *SL_MeasObjectToRemoveList_r16
	Sl_MeasObjectToAddModList_r16   *SL_MeasObjectList_r16
	Sl_ReportConfigToRemoveList_r16 *SL_ReportConfigToRemoveList_r16
	Sl_ReportConfigToAddModList_r16 *SL_ReportConfigList_r16
	Sl_MeasIdToRemoveList_r16       *SL_MeasIdToRemoveList_r16
	Sl_MeasIdToAddModList_r16       *SL_MeasIdList_r16
	Sl_QuantityConfig_r16           *SL_QuantityConfig_r16
}

func (ie *SL_MeasConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMeasConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_MeasObjectToRemoveList_r16 != nil, ie.Sl_MeasObjectToAddModList_r16 != nil, ie.Sl_ReportConfigToRemoveList_r16 != nil, ie.Sl_ReportConfigToAddModList_r16 != nil, ie.Sl_MeasIdToRemoveList_r16 != nil, ie.Sl_MeasIdToAddModList_r16 != nil, ie.Sl_QuantityConfig_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-MeasObjectToRemoveList-r16: ref
	{
		if ie.Sl_MeasObjectToRemoveList_r16 != nil {
			if err := ie.Sl_MeasObjectToRemoveList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-MeasObjectToAddModList-r16: ref
	{
		if ie.Sl_MeasObjectToAddModList_r16 != nil {
			if err := ie.Sl_MeasObjectToAddModList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-ReportConfigToRemoveList-r16: ref
	{
		if ie.Sl_ReportConfigToRemoveList_r16 != nil {
			if err := ie.Sl_ReportConfigToRemoveList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-ReportConfigToAddModList-r16: ref
	{
		if ie.Sl_ReportConfigToAddModList_r16 != nil {
			if err := ie.Sl_ReportConfigToAddModList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. sl-MeasIdToRemoveList-r16: ref
	{
		if ie.Sl_MeasIdToRemoveList_r16 != nil {
			if err := ie.Sl_MeasIdToRemoveList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. sl-MeasIdToAddModList-r16: ref
	{
		if ie.Sl_MeasIdToAddModList_r16 != nil {
			if err := ie.Sl_MeasIdToAddModList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. sl-QuantityConfig-r16: ref
	{
		if ie.Sl_QuantityConfig_r16 != nil {
			if err := ie.Sl_QuantityConfig_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_MeasConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMeasConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-MeasObjectToRemoveList-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_MeasObjectToRemoveList_r16 = new(SL_MeasObjectToRemoveList_r16)
			if err := ie.Sl_MeasObjectToRemoveList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-MeasObjectToAddModList-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_MeasObjectToAddModList_r16 = new(SL_MeasObjectList_r16)
			if err := ie.Sl_MeasObjectToAddModList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-ReportConfigToRemoveList-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_ReportConfigToRemoveList_r16 = new(SL_ReportConfigToRemoveList_r16)
			if err := ie.Sl_ReportConfigToRemoveList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-ReportConfigToAddModList-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_ReportConfigToAddModList_r16 = new(SL_ReportConfigList_r16)
			if err := ie.Sl_ReportConfigToAddModList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. sl-MeasIdToRemoveList-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sl_MeasIdToRemoveList_r16 = new(SL_MeasIdToRemoveList_r16)
			if err := ie.Sl_MeasIdToRemoveList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. sl-MeasIdToAddModList-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Sl_MeasIdToAddModList_r16 = new(SL_MeasIdList_r16)
			if err := ie.Sl_MeasIdToAddModList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. sl-QuantityConfig-r16: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Sl_QuantityConfig_r16 = new(SL_QuantityConfig_r16)
			if err := ie.Sl_QuantityConfig_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
