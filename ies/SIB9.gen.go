// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB9 (line 4228).

var sIB9Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "timeInfo", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sIB9LateNonCriticalExtensionConstraints = per.SizeConstraints{}

var sIB9EventIDTSSR18Constraints = per.Constrained(0, 63)

var sIB9TimeInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "timeInfoUTC"},
		{Name: "dayLightSavingTime", Optional: true},
		{Name: "leapSeconds", Optional: true},
		{Name: "localTimeOffset", Optional: true},
	},
}

type SIB9 struct {
	TimeInfo *struct {
		TimeInfoUTC        int64
		DayLightSavingTime *per.BitString
		LeapSeconds        *int64
		LocalTimeOffset    *int64
	}
	LateNonCriticalExtension []byte
	ReferenceTimeInfo_r16    *ReferenceTimeInfo_r16
	EventID_TSS_r18          *int64
}

func (ie *SIB9) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB9Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReferenceTimeInfo_r16 != nil
	hasExtGroup1 := ie.EventID_TSS_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TimeInfo != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. timeInfo: sequence
	{
		if ie.TimeInfo != nil {
			c := ie.TimeInfo
			sIB9TimeInfoSeq := e.NewSequenceEncoder(sIB9TimeInfoConstraints)
			if err := sIB9TimeInfoSeq.EncodePreamble([]bool{c.DayLightSavingTime != nil, c.LeapSeconds != nil, c.LocalTimeOffset != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.TimeInfoUTC, per.Constrained(0, 549755813887)); err != nil {
				return err
			}
			if c.DayLightSavingTime != nil {
				if err := e.EncodeBitString((*c.DayLightSavingTime), per.FixedSize(2)); err != nil {
					return err
				}
			}
			if c.LeapSeconds != nil {
				if err := e.EncodeInteger((*c.LeapSeconds), per.Constrained(-127, 128)); err != nil {
					return err
				}
			}
			if c.LocalTimeOffset != nil {
				if err := e.EncodeInteger((*c.LocalTimeOffset), per.Constrained(-63, 64)); err != nil {
					return err
				}
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB9LateNonCriticalExtensionConstraints); err != nil {
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
					{Name: "referenceTimeInfo-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReferenceTimeInfo_r16 != nil}); err != nil {
				return err
			}

			if ie.ReferenceTimeInfo_r16 != nil {
				if err := ie.ReferenceTimeInfo_r16.Encode(ex); err != nil {
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
					{Name: "eventID-TSS-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EventID_TSS_r18 != nil}); err != nil {
				return err
			}

			if ie.EventID_TSS_r18 != nil {
				if err := ex.EncodeInteger(*ie.EventID_TSS_r18, sIB9EventIDTSSR18Constraints); err != nil {
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

func (ie *SIB9) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB9Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. timeInfo: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.TimeInfo = &struct {
				TimeInfoUTC        int64
				DayLightSavingTime *per.BitString
				LeapSeconds        *int64
				LocalTimeOffset    *int64
			}{}
			c := ie.TimeInfo
			sIB9TimeInfoSeq := d.NewSequenceDecoder(sIB9TimeInfoConstraints)
			if err := sIB9TimeInfoSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 549755813887))
				if err != nil {
					return err
				}
				c.TimeInfoUTC = v
			}
			if sIB9TimeInfoSeq.IsComponentPresent(1) {
				c.DayLightSavingTime = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				(*c.DayLightSavingTime) = v
			}
			if sIB9TimeInfoSeq.IsComponentPresent(2) {
				c.LeapSeconds = new(int64)
				v, err := d.DecodeInteger(per.Constrained(-127, 128))
				if err != nil {
					return err
				}
				(*c.LeapSeconds) = v
			}
			if sIB9TimeInfoSeq.IsComponentPresent(3) {
				c.LocalTimeOffset = new(int64)
				v, err := d.DecodeInteger(per.Constrained(-63, 64))
				if err != nil {
					return err
				}
				(*c.LocalTimeOffset) = v
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB9LateNonCriticalExtensionConstraints)
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
				{Name: "referenceTimeInfo-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ReferenceTimeInfo_r16 = new(ReferenceTimeInfo_r16)
			if err := ie.ReferenceTimeInfo_r16.Decode(dx); err != nil {
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
				{Name: "eventID-TSS-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sIB9EventIDTSSR18Constraints)
			if err != nil {
				return err
			}
			ie.EventID_TSS_r18 = &v
		}
	}

	return nil
}
