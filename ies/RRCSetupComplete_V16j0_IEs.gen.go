// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupComplete-v16j0-IEs (line 1766).

var rRCSetupCompleteV16j0IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCSetupCompleteV16j0IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCSetupComplete_V16j0_IEs struct {
	LateNonCriticalExtension []byte
	NonCriticalExtension     *RRCSetupComplete_V17b0_IEs
}

func (ie *RRCSetupComplete_V16j0_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteV16j0IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCSetupCompleteV16j0IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCSetupComplete_V16j0_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteV16j0IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(rRCSetupCompleteV16j0IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCSetupComplete_V17b0_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
