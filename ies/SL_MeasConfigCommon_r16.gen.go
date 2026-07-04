// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasConfigCommon-r16 (line 27501).

var sLMeasConfigCommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MeasObjectListCommon-r16", Optional: true},
		{Name: "sl-ReportConfigListCommon-r16", Optional: true},
		{Name: "sl-MeasIdListCommon-r16", Optional: true},
		{Name: "sl-QuantityConfigCommon-r16", Optional: true},
	},
}

type SL_MeasConfigCommon_r16 struct {
	Sl_MeasObjectListCommon_r16   *SL_MeasObjectList_r16
	Sl_ReportConfigListCommon_r16 *SL_ReportConfigList_r16
	Sl_MeasIdListCommon_r16       *SL_MeasIdList_r16
	Sl_QuantityConfigCommon_r16   *SL_QuantityConfig_r16
}

func (ie *SL_MeasConfigCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMeasConfigCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_MeasObjectListCommon_r16 != nil, ie.Sl_ReportConfigListCommon_r16 != nil, ie.Sl_MeasIdListCommon_r16 != nil, ie.Sl_QuantityConfigCommon_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-MeasObjectListCommon-r16: ref
	{
		if ie.Sl_MeasObjectListCommon_r16 != nil {
			if err := ie.Sl_MeasObjectListCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-ReportConfigListCommon-r16: ref
	{
		if ie.Sl_ReportConfigListCommon_r16 != nil {
			if err := ie.Sl_ReportConfigListCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-MeasIdListCommon-r16: ref
	{
		if ie.Sl_MeasIdListCommon_r16 != nil {
			if err := ie.Sl_MeasIdListCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-QuantityConfigCommon-r16: ref
	{
		if ie.Sl_QuantityConfigCommon_r16 != nil {
			if err := ie.Sl_QuantityConfigCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_MeasConfigCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMeasConfigCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-MeasObjectListCommon-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_MeasObjectListCommon_r16 = new(SL_MeasObjectList_r16)
			if err := ie.Sl_MeasObjectListCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-ReportConfigListCommon-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_ReportConfigListCommon_r16 = new(SL_ReportConfigList_r16)
			if err := ie.Sl_ReportConfigListCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-MeasIdListCommon-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_MeasIdListCommon_r16 = new(SL_MeasIdList_r16)
			if err := ie.Sl_MeasIdListCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-QuantityConfigCommon-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_QuantityConfigCommon_r16 = new(SL_QuantityConfig_r16)
			if err := ie.Sl_QuantityConfigCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
