// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB7 (line 4196).

var sIB7Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "messageIdentifier"},
		{Name: "serialNumber"},
		{Name: "warningMessageSegmentType"},
		{Name: "warningMessageSegmentNumber"},
		{Name: "warningMessageSegment"},
		{Name: "dataCodingScheme", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sIB7MessageIdentifierConstraints = per.FixedSize(16)

var sIB7SerialNumberConstraints = per.FixedSize(16)

const (
	SIB7_WarningMessageSegmentType_NotLastSegment = 0
	SIB7_WarningMessageSegmentType_LastSegment    = 1
)

var sIB7WarningMessageSegmentTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB7_WarningMessageSegmentType_NotLastSegment, SIB7_WarningMessageSegmentType_LastSegment},
}

var sIB7WarningMessageSegmentNumberConstraints = per.Constrained(0, 63)

var sIB7WarningMessageSegmentConstraints = per.SizeConstraints{}

var sIB7DataCodingSchemeConstraints = per.FixedSize(1)

var sIB7LateNonCriticalExtensionConstraints = per.SizeConstraints{}

var sIB7WarningAreaCoordinatesSegmentR19Constraints = per.SizeConstraints{}

type SIB7 struct {
	MessageIdentifier                 per.BitString
	SerialNumber                      per.BitString
	WarningMessageSegmentType         int64
	WarningMessageSegmentNumber       int64
	WarningMessageSegment             []byte
	DataCodingScheme                  []byte
	LateNonCriticalExtension          []byte
	WarningAreaCoordinatesSegment_r19 []byte
}

func (ie *SIB7) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB7Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.WarningAreaCoordinatesSegment_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataCodingScheme != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. messageIdentifier: bit-string
	{
		if err := e.EncodeBitString(ie.MessageIdentifier, sIB7MessageIdentifierConstraints); err != nil {
			return err
		}
	}

	// 4. serialNumber: bit-string
	{
		if err := e.EncodeBitString(ie.SerialNumber, sIB7SerialNumberConstraints); err != nil {
			return err
		}
	}

	// 5. warningMessageSegmentType: enumerated
	{
		if err := e.EncodeEnumerated(ie.WarningMessageSegmentType, sIB7WarningMessageSegmentTypeConstraints); err != nil {
			return err
		}
	}

	// 6. warningMessageSegmentNumber: integer
	{
		if err := e.EncodeInteger(ie.WarningMessageSegmentNumber, sIB7WarningMessageSegmentNumberConstraints); err != nil {
			return err
		}
	}

	// 7. warningMessageSegment: octet-string
	{
		if err := e.EncodeOctetString(ie.WarningMessageSegment, sIB7WarningMessageSegmentConstraints); err != nil {
			return err
		}
	}

	// 8. dataCodingScheme: octet-string
	{
		if ie.DataCodingScheme != nil {
			if err := e.EncodeOctetString(ie.DataCodingScheme, sIB7DataCodingSchemeConstraints); err != nil {
				return err
			}
		}
	}

	// 9. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB7LateNonCriticalExtensionConstraints); err != nil {
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
					{Name: "warningAreaCoordinatesSegment-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.WarningAreaCoordinatesSegment_r19 != nil}); err != nil {
				return err
			}

			if ie.WarningAreaCoordinatesSegment_r19 != nil {
				if err := ex.EncodeOctetString(ie.WarningAreaCoordinatesSegment_r19, sIB7WarningAreaCoordinatesSegmentR19Constraints); err != nil {
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

func (ie *SIB7) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB7Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. messageIdentifier: bit-string
	{
		v0, err := d.DecodeBitString(sIB7MessageIdentifierConstraints)
		if err != nil {
			return err
		}
		ie.MessageIdentifier = v0
	}

	// 4. serialNumber: bit-string
	{
		v1, err := d.DecodeBitString(sIB7SerialNumberConstraints)
		if err != nil {
			return err
		}
		ie.SerialNumber = v1
	}

	// 5. warningMessageSegmentType: enumerated
	{
		v2, err := d.DecodeEnumerated(sIB7WarningMessageSegmentTypeConstraints)
		if err != nil {
			return err
		}
		ie.WarningMessageSegmentType = v2
	}

	// 6. warningMessageSegmentNumber: integer
	{
		v3, err := d.DecodeInteger(sIB7WarningMessageSegmentNumberConstraints)
		if err != nil {
			return err
		}
		ie.WarningMessageSegmentNumber = v3
	}

	// 7. warningMessageSegment: octet-string
	{
		v4, err := d.DecodeOctetString(sIB7WarningMessageSegmentConstraints)
		if err != nil {
			return err
		}
		ie.WarningMessageSegment = v4
	}

	// 8. dataCodingScheme: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(sIB7DataCodingSchemeConstraints)
			if err != nil {
				return err
			}
			ie.DataCodingScheme = v
		}
	}

	// 9. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(sIB7LateNonCriticalExtensionConstraints)
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
				{Name: "warningAreaCoordinatesSegment-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeOctetString(sIB7WarningAreaCoordinatesSegmentR19Constraints)
			if err != nil {
				return err
			}
			ie.WarningAreaCoordinatesSegment_r19 = v
		}
	}

	return nil
}
