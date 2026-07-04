// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB13-r16 (line 4361).

var sIB13R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-V2X-ConfigCommon-r16"},
		{Name: "dummy"},
		{Name: "tdd-Config-r16"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB13R16SlV2XConfigCommonR16Constraints = per.SizeConstraints{}

var sIB13R16DummyConstraints = per.SizeConstraints{}

var sIB13R16TddConfigR16Constraints = per.SizeConstraints{}

var sIB13R16LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB13_r16 struct {
	Sl_V2X_ConfigCommon_r16  []byte
	Dummy                    []byte
	Tdd_Config_r16           []byte
	LateNonCriticalExtension []byte
}

func (ie *SIB13_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB13R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. sl-V2X-ConfigCommon-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.Sl_V2X_ConfigCommon_r16, sIB13R16SlV2XConfigCommonR16Constraints); err != nil {
			return err
		}
	}

	// 4. dummy: octet-string
	{
		if err := e.EncodeOctetString(ie.Dummy, sIB13R16DummyConstraints); err != nil {
			return err
		}
	}

	// 5. tdd-Config-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.Tdd_Config_r16, sIB13R16TddConfigR16Constraints); err != nil {
			return err
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB13R16LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB13_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB13R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-V2X-ConfigCommon-r16: octet-string
	{
		v0, err := d.DecodeOctetString(sIB13R16SlV2XConfigCommonR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_V2X_ConfigCommon_r16 = v0
	}

	// 4. dummy: octet-string
	{
		v1, err := d.DecodeOctetString(sIB13R16DummyConstraints)
		if err != nil {
			return err
		}
		ie.Dummy = v1
	}

	// 5. tdd-Config-r16: octet-string
	{
		v2, err := d.DecodeOctetString(sIB13R16TddConfigR16Constraints)
		if err != nil {
			return err
		}
		ie.Tdd_Config_r16 = v2
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(sIB13R16LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
