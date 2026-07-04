// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB4 (line 3931).

var sIB4Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "interFreqCarrierFreqList"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
	},
}

var sIB4LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB4 struct {
	InterFreqCarrierFreqList       InterFreqCarrierFreqList
	LateNonCriticalExtension       []byte
	InterFreqCarrierFreqList_v1610 *InterFreqCarrierFreqList_v1610
	InterFreqCarrierFreqList_v1700 *InterFreqCarrierFreqList_v1700
	InterFreqCarrierFreqList_v1720 *InterFreqCarrierFreqList_v1720
	InterFreqCarrierFreqList_v1730 *InterFreqCarrierFreqList_v1730
	InterFreqCarrierFreqList_v1760 *InterFreqCarrierFreqList_v1760
	InterFreqCarrierFreqList_v1800 *InterFreqCarrierFreqList_v1800
	InterFreqCarrierFreqList_v1900 *InterFreqCarrierFreqList_v1900
}

func (ie *SIB4) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB4Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.InterFreqCarrierFreqList_v1610 != nil
	hasExtGroup1 := ie.InterFreqCarrierFreqList_v1700 != nil
	hasExtGroup2 := ie.InterFreqCarrierFreqList_v1720 != nil
	hasExtGroup3 := ie.InterFreqCarrierFreqList_v1730 != nil
	hasExtGroup4 := ie.InterFreqCarrierFreqList_v1760 != nil
	hasExtGroup5 := ie.InterFreqCarrierFreqList_v1800 != nil
	hasExtGroup6 := ie.InterFreqCarrierFreqList_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. interFreqCarrierFreqList: ref
	{
		if err := ie.InterFreqCarrierFreqList.Encode(e); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB4LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "interFreqCarrierFreqList-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFreqCarrierFreqList_v1610 != nil}); err != nil {
				return err
			}

			if ie.InterFreqCarrierFreqList_v1610 != nil {
				if err := ie.InterFreqCarrierFreqList_v1610.Encode(ex); err != nil {
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
					{Name: "interFreqCarrierFreqList-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFreqCarrierFreqList_v1700 != nil}); err != nil {
				return err
			}

			if ie.InterFreqCarrierFreqList_v1700 != nil {
				if err := ie.InterFreqCarrierFreqList_v1700.Encode(ex); err != nil {
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
					{Name: "interFreqCarrierFreqList-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFreqCarrierFreqList_v1720 != nil}); err != nil {
				return err
			}

			if ie.InterFreqCarrierFreqList_v1720 != nil {
				if err := ie.InterFreqCarrierFreqList_v1720.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "interFreqCarrierFreqList-v1730", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFreqCarrierFreqList_v1730 != nil}); err != nil {
				return err
			}

			if ie.InterFreqCarrierFreqList_v1730 != nil {
				if err := ie.InterFreqCarrierFreqList_v1730.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "interFreqCarrierFreqList-v1760", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFreqCarrierFreqList_v1760 != nil}); err != nil {
				return err
			}

			if ie.InterFreqCarrierFreqList_v1760 != nil {
				if err := ie.InterFreqCarrierFreqList_v1760.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "interFreqCarrierFreqList-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFreqCarrierFreqList_v1800 != nil}); err != nil {
				return err
			}

			if ie.InterFreqCarrierFreqList_v1800 != nil {
				if err := ie.InterFreqCarrierFreqList_v1800.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 6:
		if hasExtGroup6 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "interFreqCarrierFreqList-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFreqCarrierFreqList_v1900 != nil}); err != nil {
				return err
			}

			if ie.InterFreqCarrierFreqList_v1900 != nil {
				if err := ie.InterFreqCarrierFreqList_v1900.Encode(ex); err != nil {
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

func (ie *SIB4) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB4Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. interFreqCarrierFreqList: ref
	{
		if err := ie.InterFreqCarrierFreqList.Decode(d); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB4LateNonCriticalExtensionConstraints)
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
				{Name: "interFreqCarrierFreqList-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InterFreqCarrierFreqList_v1610 = new(InterFreqCarrierFreqList_v1610)
			if err := ie.InterFreqCarrierFreqList_v1610.Decode(dx); err != nil {
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
				{Name: "interFreqCarrierFreqList-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InterFreqCarrierFreqList_v1700 = new(InterFreqCarrierFreqList_v1700)
			if err := ie.InterFreqCarrierFreqList_v1700.Decode(dx); err != nil {
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
				{Name: "interFreqCarrierFreqList-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InterFreqCarrierFreqList_v1720 = new(InterFreqCarrierFreqList_v1720)
			if err := ie.InterFreqCarrierFreqList_v1720.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "interFreqCarrierFreqList-v1730", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InterFreqCarrierFreqList_v1730 = new(InterFreqCarrierFreqList_v1730)
			if err := ie.InterFreqCarrierFreqList_v1730.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "interFreqCarrierFreqList-v1760", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InterFreqCarrierFreqList_v1760 = new(InterFreqCarrierFreqList_v1760)
			if err := ie.InterFreqCarrierFreqList_v1760.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "interFreqCarrierFreqList-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InterFreqCarrierFreqList_v1800 = new(InterFreqCarrierFreqList_v1800)
			if err := ie.InterFreqCarrierFreqList_v1800.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "interFreqCarrierFreqList-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InterFreqCarrierFreqList_v1900 = new(InterFreqCarrierFreqList_v1900)
			if err := ie.InterFreqCarrierFreqList_v1900.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
