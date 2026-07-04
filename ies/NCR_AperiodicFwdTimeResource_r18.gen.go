// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-AperiodicFwdTimeResource-r18 (line 10291).

var nCRAperiodicFwdTimeResourceR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aperiodicFwdTimeRsrcId-r18"},
		{Name: "slotOffsetAperiodic-r18"},
		{Name: "symbolOffset-r18"},
		{Name: "durationInSymbols-r18"},
	},
}

var nCRAperiodicFwdTimeResourceR18SlotOffsetAperiodicR18Constraints = per.Constrained(0, 14)

var nCRAperiodicFwdTimeResourceR18SymbolOffsetR18Constraints = per.Constrained(0, common.MaxNrofSymbols_1)

var nCRAperiodicFwdTimeResourceR18DurationInSymbolsR18Constraints = per.Constrained(1, 28)

type NCR_AperiodicFwdTimeResource_r18 struct {
	AperiodicFwdTimeRsrcId_r18 NCR_AperiodicFwdTimeResourceId_r18
	SlotOffsetAperiodic_r18    int64
	SymbolOffset_r18           int64
	DurationInSymbols_r18      int64
}

func (ie *NCR_AperiodicFwdTimeResource_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nCRAperiodicFwdTimeResourceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. aperiodicFwdTimeRsrcId-r18: ref
	{
		if err := ie.AperiodicFwdTimeRsrcId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. slotOffsetAperiodic-r18: integer
	{
		if err := e.EncodeInteger(ie.SlotOffsetAperiodic_r18, nCRAperiodicFwdTimeResourceR18SlotOffsetAperiodicR18Constraints); err != nil {
			return err
		}
	}

	// 4. symbolOffset-r18: integer
	{
		if err := e.EncodeInteger(ie.SymbolOffset_r18, nCRAperiodicFwdTimeResourceR18SymbolOffsetR18Constraints); err != nil {
			return err
		}
	}

	// 5. durationInSymbols-r18: integer
	{
		if err := e.EncodeInteger(ie.DurationInSymbols_r18, nCRAperiodicFwdTimeResourceR18DurationInSymbolsR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NCR_AperiodicFwdTimeResource_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nCRAperiodicFwdTimeResourceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. aperiodicFwdTimeRsrcId-r18: ref
	{
		if err := ie.AperiodicFwdTimeRsrcId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. slotOffsetAperiodic-r18: integer
	{
		v1, err := d.DecodeInteger(nCRAperiodicFwdTimeResourceR18SlotOffsetAperiodicR18Constraints)
		if err != nil {
			return err
		}
		ie.SlotOffsetAperiodic_r18 = v1
	}

	// 4. symbolOffset-r18: integer
	{
		v2, err := d.DecodeInteger(nCRAperiodicFwdTimeResourceR18SymbolOffsetR18Constraints)
		if err != nil {
			return err
		}
		ie.SymbolOffset_r18 = v2
	}

	// 5. durationInSymbols-r18: integer
	{
		v3, err := d.DecodeInteger(nCRAperiodicFwdTimeResourceR18DurationInSymbolsR18Constraints)
		if err != nil {
			return err
		}
		ie.DurationInSymbols_r18 = v3
	}

	return nil
}
