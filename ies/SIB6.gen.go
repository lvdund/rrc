// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB6 (line 4182).

var sIB6Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "messageIdentifier"},
		{Name: "serialNumber"},
		{Name: "warningType"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sIB6MessageIdentifierConstraints = per.FixedSize(16)

var sIB6SerialNumberConstraints = per.FixedSize(16)

var sIB6WarningTypeConstraints = per.FixedSize(2)

var sIB6LateNonCriticalExtensionConstraints = per.SizeConstraints{}

var sIB6WarningAreaCoordinatesR19Constraints = per.SizeConstraints{}

type SIB6 struct {
	MessageIdentifier          per.BitString
	SerialNumber               per.BitString
	WarningType                []byte
	LateNonCriticalExtension   []byte
	WarningAreaCoordinates_r19 []byte
}

func (ie *SIB6) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB6Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.WarningAreaCoordinates_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. messageIdentifier: bit-string
	{
		if err := e.EncodeBitString(ie.MessageIdentifier, sIB6MessageIdentifierConstraints); err != nil {
			return err
		}
	}

	// 4. serialNumber: bit-string
	{
		if err := e.EncodeBitString(ie.SerialNumber, sIB6SerialNumberConstraints); err != nil {
			return err
		}
	}

	// 5. warningType: octet-string
	{
		if err := e.EncodeOctetString(ie.WarningType, sIB6WarningTypeConstraints); err != nil {
			return err
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB6LateNonCriticalExtensionConstraints); err != nil {
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
					{Name: "warningAreaCoordinates-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.WarningAreaCoordinates_r19 != nil}); err != nil {
				return err
			}

			if ie.WarningAreaCoordinates_r19 != nil {
				if err := ex.EncodeOctetString(ie.WarningAreaCoordinates_r19, sIB6WarningAreaCoordinatesR19Constraints); err != nil {
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

func (ie *SIB6) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB6Constraints)

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
		v0, err := d.DecodeBitString(sIB6MessageIdentifierConstraints)
		if err != nil {
			return err
		}
		ie.MessageIdentifier = v0
	}

	// 4. serialNumber: bit-string
	{
		v1, err := d.DecodeBitString(sIB6SerialNumberConstraints)
		if err != nil {
			return err
		}
		ie.SerialNumber = v1
	}

	// 5. warningType: octet-string
	{
		v2, err := d.DecodeOctetString(sIB6WarningTypeConstraints)
		if err != nil {
			return err
		}
		ie.WarningType = v2
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(sIB6LateNonCriticalExtensionConstraints)
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
				{Name: "warningAreaCoordinates-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeOctetString(sIB6WarningAreaCoordinatesR19Constraints)
			if err != nil {
				return err
			}
			ie.WarningAreaCoordinates_r19 = v
		}
	}

	return nil
}
