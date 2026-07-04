// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetup-IEs (line 1691).

var rRCSetupIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "radioBearerConfig"},
		{Name: "masterCellGroup"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCSetupIEsMasterCellGroupConstraints = per.SizeConstraints{}

var rRCSetupIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCSetup_IEs struct {
	RadioBearerConfig        RadioBearerConfig
	MasterCellGroup          []byte
	LateNonCriticalExtension []byte
	NonCriticalExtension     *RRCSetup_v1700_IEs
}

func (ie *RRCSetup_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. radioBearerConfig: ref
	{
		if err := ie.RadioBearerConfig.Encode(e); err != nil {
			return err
		}
	}

	// 3. masterCellGroup: octet-string
	{
		if err := e.EncodeOctetString(ie.MasterCellGroup, rRCSetupIEsMasterCellGroupConstraints); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCSetupIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCSetup_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. radioBearerConfig: ref
	{
		if err := ie.RadioBearerConfig.Decode(d); err != nil {
			return err
		}
	}

	// 3. masterCellGroup: octet-string
	{
		v1, err := d.DecodeOctetString(rRCSetupIEsMasterCellGroupConstraints)
		if err != nil {
			return err
		}
		ie.MasterCellGroup = v1
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(rRCSetupIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(RRCSetup_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
