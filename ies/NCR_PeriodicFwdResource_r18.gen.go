// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-PeriodicFwdResource-r18 (line 10393).

var nCRPeriodicFwdResourceR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicFwdRsrcId-r18"},
		{Name: "beamIndex-r18"},
		{Name: "periodicTimeRsrc-r18"},
	},
}

var nCRPeriodicFwdResourceR18BeamIndexR18Constraints = per.Constrained(0, 63)

type NCR_PeriodicFwdResource_r18 struct {
	PeriodicFwdRsrcId_r18 NCR_PeriodicFwdResourceId_r18
	BeamIndex_r18         int64
	PeriodicTimeRsrc_r18  struct {
		PeriodicityAndOffset_r18 NCR_PeriodicityAndOffset_r18
		SymbolOffset_r18         int64
		DurationInSymbols_r18    int64
	}
}

func (ie *NCR_PeriodicFwdResource_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nCRPeriodicFwdResourceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. periodicFwdRsrcId-r18: ref
	{
		if err := ie.PeriodicFwdRsrcId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. beamIndex-r18: integer
	{
		if err := e.EncodeInteger(ie.BeamIndex_r18, nCRPeriodicFwdResourceR18BeamIndexR18Constraints); err != nil {
			return err
		}
	}

	// 4. periodicTimeRsrc-r18: sequence
	{
		{
			c := &ie.PeriodicTimeRsrc_r18
			if err := c.PeriodicityAndOffset_r18.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SymbolOffset_r18, per.Constrained(0, common.MaxNrofSymbols_1)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.DurationInSymbols_r18, per.Constrained(1, 112)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NCR_PeriodicFwdResource_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nCRPeriodicFwdResourceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. periodicFwdRsrcId-r18: ref
	{
		if err := ie.PeriodicFwdRsrcId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. beamIndex-r18: integer
	{
		v1, err := d.DecodeInteger(nCRPeriodicFwdResourceR18BeamIndexR18Constraints)
		if err != nil {
			return err
		}
		ie.BeamIndex_r18 = v1
	}

	// 4. periodicTimeRsrc-r18: sequence
	{
		{
			c := &ie.PeriodicTimeRsrc_r18
			{
				if err := c.PeriodicityAndOffset_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofSymbols_1))
				if err != nil {
					return err
				}
				c.SymbolOffset_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 112))
				if err != nil {
					return err
				}
				c.DurationInSymbols_r18 = v
			}
		}
	}

	return nil
}
