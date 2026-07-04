// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasIdToAddMod (line 9328).

var measIdToAddModConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measId"},
		{Name: "measObjectId"},
		{Name: "reportConfigId"},
	},
}

type MeasIdToAddMod struct {
	MeasId         MeasId
	MeasObjectId   MeasObjectId
	ReportConfigId ReportConfigId
}

func (ie *MeasIdToAddMod) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measIdToAddModConstraints)
	_ = seq

	// 1. measId: ref
	{
		if err := ie.MeasId.Encode(e); err != nil {
			return err
		}
	}

	// 2. measObjectId: ref
	{
		if err := ie.MeasObjectId.Encode(e); err != nil {
			return err
		}
	}

	// 3. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasIdToAddMod) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measIdToAddModConstraints)
	_ = seq

	// 1. measId: ref
	{
		if err := ie.MeasId.Decode(d); err != nil {
			return err
		}
	}

	// 2. measObjectId: ref
	{
		if err := ie.MeasObjectId.Decode(d); err != nil {
			return err
		}
	}

	// 3. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
