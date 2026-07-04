// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MasterKeyUpdate (line 1087).

var masterKeyUpdateConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "keySetChangeIndicator"},
		{Name: "nextHopChainingCount"},
		{Name: "nas-Container", Optional: true},
	},
}

var masterKeyUpdateNasContainerConstraints = per.SizeConstraints{}

type MasterKeyUpdate struct {
	KeySetChangeIndicator bool
	NextHopChainingCount  NextHopChainingCount
	Nas_Container         []byte
}

func (ie *MasterKeyUpdate) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(masterKeyUpdateConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nas_Container != nil}); err != nil {
		return err
	}

	// 3. keySetChangeIndicator: boolean
	{
		if err := e.EncodeBoolean(ie.KeySetChangeIndicator); err != nil {
			return err
		}
	}

	// 4. nextHopChainingCount: ref
	{
		if err := ie.NextHopChainingCount.Encode(e); err != nil {
			return err
		}
	}

	// 5. nas-Container: octet-string
	{
		if ie.Nas_Container != nil {
			if err := e.EncodeOctetString(ie.Nas_Container, masterKeyUpdateNasContainerConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MasterKeyUpdate) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(masterKeyUpdateConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. keySetChangeIndicator: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.KeySetChangeIndicator = v0
	}

	// 4. nextHopChainingCount: ref
	{
		if err := ie.NextHopChainingCount.Decode(d); err != nil {
			return err
		}
	}

	// 5. nas-Container: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(masterKeyUpdateNasContainerConstraints)
			if err != nil {
				return err
			}
			ie.Nas_Container = v
		}
	}

	return nil
}
