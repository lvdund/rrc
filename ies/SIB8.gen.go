// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB8 (line 4213).

var sIB8Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "messageIdentifier"},
		{Name: "serialNumber"},
		{Name: "warningMessageSegmentType"},
		{Name: "warningMessageSegmentNumber"},
		{Name: "warningMessageSegment"},
		{Name: "dataCodingScheme", Optional: true},
		{Name: "warningAreaCoordinatesSegment", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB8MessageIdentifierConstraints = per.FixedSize(16)

var sIB8SerialNumberConstraints = per.FixedSize(16)

const (
	SIB8_WarningMessageSegmentType_NotLastSegment = 0
	SIB8_WarningMessageSegmentType_LastSegment    = 1
)

var sIB8WarningMessageSegmentTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB8_WarningMessageSegmentType_NotLastSegment, SIB8_WarningMessageSegmentType_LastSegment},
}

var sIB8WarningMessageSegmentNumberConstraints = per.Constrained(0, 63)

var sIB8WarningMessageSegmentConstraints = per.SizeConstraints{}

var sIB8DataCodingSchemeConstraints = per.FixedSize(1)

var sIB8WarningAreaCoordinatesSegmentConstraints = per.SizeConstraints{}

var sIB8LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB8 struct {
	MessageIdentifier             per.BitString
	SerialNumber                  per.BitString
	WarningMessageSegmentType     int64
	WarningMessageSegmentNumber   int64
	WarningMessageSegment         []byte
	DataCodingScheme              []byte
	WarningAreaCoordinatesSegment []byte
	LateNonCriticalExtension      []byte
}

func (ie *SIB8) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB8Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataCodingScheme != nil, ie.WarningAreaCoordinatesSegment != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. messageIdentifier: bit-string
	{
		if err := e.EncodeBitString(ie.MessageIdentifier, sIB8MessageIdentifierConstraints); err != nil {
			return err
		}
	}

	// 4. serialNumber: bit-string
	{
		if err := e.EncodeBitString(ie.SerialNumber, sIB8SerialNumberConstraints); err != nil {
			return err
		}
	}

	// 5. warningMessageSegmentType: enumerated
	{
		if err := e.EncodeEnumerated(ie.WarningMessageSegmentType, sIB8WarningMessageSegmentTypeConstraints); err != nil {
			return err
		}
	}

	// 6. warningMessageSegmentNumber: integer
	{
		if err := e.EncodeInteger(ie.WarningMessageSegmentNumber, sIB8WarningMessageSegmentNumberConstraints); err != nil {
			return err
		}
	}

	// 7. warningMessageSegment: octet-string
	{
		if err := e.EncodeOctetString(ie.WarningMessageSegment, sIB8WarningMessageSegmentConstraints); err != nil {
			return err
		}
	}

	// 8. dataCodingScheme: octet-string
	{
		if ie.DataCodingScheme != nil {
			if err := e.EncodeOctetString(ie.DataCodingScheme, sIB8DataCodingSchemeConstraints); err != nil {
				return err
			}
		}
	}

	// 9. warningAreaCoordinatesSegment: octet-string
	{
		if ie.WarningAreaCoordinatesSegment != nil {
			if err := e.EncodeOctetString(ie.WarningAreaCoordinatesSegment, sIB8WarningAreaCoordinatesSegmentConstraints); err != nil {
				return err
			}
		}
	}

	// 10. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB8LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB8) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB8Constraints)

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
		v0, err := d.DecodeBitString(sIB8MessageIdentifierConstraints)
		if err != nil {
			return err
		}
		ie.MessageIdentifier = v0
	}

	// 4. serialNumber: bit-string
	{
		v1, err := d.DecodeBitString(sIB8SerialNumberConstraints)
		if err != nil {
			return err
		}
		ie.SerialNumber = v1
	}

	// 5. warningMessageSegmentType: enumerated
	{
		v2, err := d.DecodeEnumerated(sIB8WarningMessageSegmentTypeConstraints)
		if err != nil {
			return err
		}
		ie.WarningMessageSegmentType = v2
	}

	// 6. warningMessageSegmentNumber: integer
	{
		v3, err := d.DecodeInteger(sIB8WarningMessageSegmentNumberConstraints)
		if err != nil {
			return err
		}
		ie.WarningMessageSegmentNumber = v3
	}

	// 7. warningMessageSegment: octet-string
	{
		v4, err := d.DecodeOctetString(sIB8WarningMessageSegmentConstraints)
		if err != nil {
			return err
		}
		ie.WarningMessageSegment = v4
	}

	// 8. dataCodingScheme: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(sIB8DataCodingSchemeConstraints)
			if err != nil {
				return err
			}
			ie.DataCodingScheme = v
		}
	}

	// 9. warningAreaCoordinatesSegment: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(sIB8WarningAreaCoordinatesSegmentConstraints)
			if err != nil {
				return err
			}
			ie.WarningAreaCoordinatesSegment = v
		}
	}

	// 10. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeOctetString(sIB8LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
