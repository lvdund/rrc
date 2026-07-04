// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB19-r17 (line 4508).

var sIB19R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-Config-r17", Optional: true},
		{Name: "t-Service-r17", Optional: true},
		{Name: "referenceLocation-r17", Optional: true},
		{Name: "distanceThresh-r17", Optional: true},
		{Name: "ntn-NeighCellConfigList-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var sIB19R17TServiceR17Constraints = per.Constrained(0, 549755813887)

var sIB19R17DistanceThreshR17Constraints = per.Constrained(0, 65525)

var sIB19R17LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB19_r17 struct {
	Ntn_Config_r17                   *NTN_Config_r17
	T_Service_r17                    *int64
	ReferenceLocation_r17            *ReferenceLocation_r17
	DistanceThresh_r17               *int64
	Ntn_NeighCellConfigList_r17      *NTN_NeighCellConfigList_r17
	LateNonCriticalExtension         []byte
	Ntn_NeighCellConfigListExt_v1720 *NTN_NeighCellConfigList_r17
	MovingReferenceLocation_r18      *ReferenceLocation_r17
	Ntn_CovEnh_r18                   *NTN_CovEnh_r18
	SatSwitchWithReSync_r18          *SatSwitchWithReSync_r18
	RefLocList_r19                   *RefLocList_r19
}

func (ie *SIB19_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB19R17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ntn_NeighCellConfigListExt_v1720 != nil
	hasExtGroup1 := ie.MovingReferenceLocation_r18 != nil || ie.Ntn_CovEnh_r18 != nil || ie.SatSwitchWithReSync_r18 != nil
	hasExtGroup2 := ie.RefLocList_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ntn_Config_r17 != nil, ie.T_Service_r17 != nil, ie.ReferenceLocation_r17 != nil, ie.DistanceThresh_r17 != nil, ie.Ntn_NeighCellConfigList_r17 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. ntn-Config-r17: ref
	{
		if ie.Ntn_Config_r17 != nil {
			if err := ie.Ntn_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. t-Service-r17: integer
	{
		if ie.T_Service_r17 != nil {
			if err := e.EncodeInteger(*ie.T_Service_r17, sIB19R17TServiceR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. referenceLocation-r17: ref
	{
		if ie.ReferenceLocation_r17 != nil {
			if err := ie.ReferenceLocation_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. distanceThresh-r17: integer
	{
		if ie.DistanceThresh_r17 != nil {
			if err := e.EncodeInteger(*ie.DistanceThresh_r17, sIB19R17DistanceThreshR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ntn-NeighCellConfigList-r17: ref
	{
		if ie.Ntn_NeighCellConfigList_r17 != nil {
			if err := ie.Ntn_NeighCellConfigList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB19R17LateNonCriticalExtensionConstraints); err != nil {
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
					{Name: "ntn-NeighCellConfigListExt-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ntn_NeighCellConfigListExt_v1720 != nil}); err != nil {
				return err
			}

			if ie.Ntn_NeighCellConfigListExt_v1720 != nil {
				if err := ie.Ntn_NeighCellConfigListExt_v1720.Encode(ex); err != nil {
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
					{Name: "movingReferenceLocation-r18", Optional: true},
					{Name: "ntn-CovEnh-r18", Optional: true},
					{Name: "satSwitchWithReSync-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MovingReferenceLocation_r18 != nil, ie.Ntn_CovEnh_r18 != nil, ie.SatSwitchWithReSync_r18 != nil}); err != nil {
				return err
			}

			if ie.MovingReferenceLocation_r18 != nil {
				if err := ie.MovingReferenceLocation_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ntn_CovEnh_r18 != nil {
				if err := ie.Ntn_CovEnh_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SatSwitchWithReSync_r18 != nil {
				if err := ie.SatSwitchWithReSync_r18.Encode(ex); err != nil {
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
					{Name: "refLocList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RefLocList_r19 != nil}); err != nil {
				return err
			}

			if ie.RefLocList_r19 != nil {
				if err := ie.RefLocList_r19.Encode(ex); err != nil {
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

func (ie *SIB19_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB19R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ntn-Config-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ntn_Config_r17 = new(NTN_Config_r17)
			if err := ie.Ntn_Config_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. t-Service-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sIB19R17TServiceR17Constraints)
			if err != nil {
				return err
			}
			ie.T_Service_r17 = &v
		}
	}

	// 5. referenceLocation-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ReferenceLocation_r17 = new(ReferenceLocation_r17)
			if err := ie.ReferenceLocation_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. distanceThresh-r17: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sIB19R17DistanceThreshR17Constraints)
			if err != nil {
				return err
			}
			ie.DistanceThresh_r17 = &v
		}
	}

	// 7. ntn-NeighCellConfigList-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ntn_NeighCellConfigList_r17 = new(NTN_NeighCellConfigList_r17)
			if err := ie.Ntn_NeighCellConfigList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(sIB19R17LateNonCriticalExtensionConstraints)
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
				{Name: "ntn-NeighCellConfigListExt-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ntn_NeighCellConfigListExt_v1720 = new(NTN_NeighCellConfigList_r17)
			if err := ie.Ntn_NeighCellConfigListExt_v1720.Decode(dx); err != nil {
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
				{Name: "movingReferenceLocation-r18", Optional: true},
				{Name: "ntn-CovEnh-r18", Optional: true},
				{Name: "satSwitchWithReSync-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MovingReferenceLocation_r18 = new(ReferenceLocation_r17)
			if err := ie.MovingReferenceLocation_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ntn_CovEnh_r18 = new(NTN_CovEnh_r18)
			if err := ie.Ntn_CovEnh_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SatSwitchWithReSync_r18 = new(SatSwitchWithReSync_r18)
			if err := ie.SatSwitchWithReSync_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "refLocList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.RefLocList_r19 = new(RefLocList_r19)
			if err := ie.RefLocList_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
