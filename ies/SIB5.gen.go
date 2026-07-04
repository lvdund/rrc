// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB5 (line 4107).

var sIB5Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreqListEUTRA", Optional: true},
		{Name: "t-ReselectionEUTRA"},
		{Name: "t-ReselectionEUTRA-SF", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var sIB5LateNonCriticalExtensionConstraints = per.SizeConstraints{}

const (
	SIB5_Ext_IdleModeMeasVoiceFallback_r17_True = 0
)

var sIB5ExtIdleModeMeasVoiceFallbackR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB5_Ext_IdleModeMeasVoiceFallback_r17_True},
}

type SIB5 struct {
	CarrierFreqListEUTRA          *CarrierFreqListEUTRA
	T_ReselectionEUTRA            T_Reselection
	T_ReselectionEUTRA_SF         *SpeedStateScaleFactors
	LateNonCriticalExtension      []byte
	CarrierFreqListEUTRA_v1610    *CarrierFreqListEUTRA_v1610
	CarrierFreqListEUTRA_v1700    *CarrierFreqListEUTRA_v1700
	IdleModeMeasVoiceFallback_r17 *int64
	CarrierFreqListEUTRA_v1800    *CarrierFreqListEUTRA_v1800
}

func (ie *SIB5) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB5Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CarrierFreqListEUTRA_v1610 != nil
	hasExtGroup1 := ie.CarrierFreqListEUTRA_v1700 != nil || ie.IdleModeMeasVoiceFallback_r17 != nil
	hasExtGroup2 := ie.CarrierFreqListEUTRA_v1800 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CarrierFreqListEUTRA != nil, ie.T_ReselectionEUTRA_SF != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. carrierFreqListEUTRA: ref
	{
		if ie.CarrierFreqListEUTRA != nil {
			if err := ie.CarrierFreqListEUTRA.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. t-ReselectionEUTRA: ref
	{
		if err := ie.T_ReselectionEUTRA.Encode(e); err != nil {
			return err
		}
	}

	// 5. t-ReselectionEUTRA-SF: ref
	{
		if ie.T_ReselectionEUTRA_SF != nil {
			if err := ie.T_ReselectionEUTRA_SF.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB5LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "carrierFreqListEUTRA-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CarrierFreqListEUTRA_v1610 != nil}); err != nil {
				return err
			}

			if ie.CarrierFreqListEUTRA_v1610 != nil {
				if err := ie.CarrierFreqListEUTRA_v1610.Encode(ex); err != nil {
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
					{Name: "carrierFreqListEUTRA-v1700", Optional: true},
					{Name: "idleModeMeasVoiceFallback-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CarrierFreqListEUTRA_v1700 != nil, ie.IdleModeMeasVoiceFallback_r17 != nil}); err != nil {
				return err
			}

			if ie.CarrierFreqListEUTRA_v1700 != nil {
				if err := ie.CarrierFreqListEUTRA_v1700.Encode(ex); err != nil {
					return err
				}
			}

			if ie.IdleModeMeasVoiceFallback_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.IdleModeMeasVoiceFallback_r17, sIB5ExtIdleModeMeasVoiceFallbackR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "carrierFreqListEUTRA-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CarrierFreqListEUTRA_v1800 != nil}); err != nil {
				return err
			}

			if ie.CarrierFreqListEUTRA_v1800 != nil {
				if err := ie.CarrierFreqListEUTRA_v1800.Encode(ex); err != nil {
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

func (ie *SIB5) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB5Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreqListEUTRA: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CarrierFreqListEUTRA = new(CarrierFreqListEUTRA)
			if err := ie.CarrierFreqListEUTRA.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. t-ReselectionEUTRA: ref
	{
		if err := ie.T_ReselectionEUTRA.Decode(d); err != nil {
			return err
		}
	}

	// 5. t-ReselectionEUTRA-SF: ref
	{
		if seq.IsComponentPresent(2) {
			ie.T_ReselectionEUTRA_SF = new(SpeedStateScaleFactors)
			if err := ie.T_ReselectionEUTRA_SF.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(sIB5LateNonCriticalExtensionConstraints)
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
				{Name: "carrierFreqListEUTRA-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CarrierFreqListEUTRA_v1610 = new(CarrierFreqListEUTRA_v1610)
			if err := ie.CarrierFreqListEUTRA_v1610.Decode(dx); err != nil {
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
				{Name: "carrierFreqListEUTRA-v1700", Optional: true},
				{Name: "idleModeMeasVoiceFallback-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CarrierFreqListEUTRA_v1700 = new(CarrierFreqListEUTRA_v1700)
			if err := ie.CarrierFreqListEUTRA_v1700.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sIB5ExtIdleModeMeasVoiceFallbackR17Constraints)
			if err != nil {
				return err
			}
			ie.IdleModeMeasVoiceFallback_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "carrierFreqListEUTRA-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CarrierFreqListEUTRA_v1800 = new(CarrierFreqListEUTRA_v1800)
			if err := ie.CarrierFreqListEUTRA_v1800.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
