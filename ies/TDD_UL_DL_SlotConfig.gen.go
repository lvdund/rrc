// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDD-UL-DL-SlotConfig (line 16107).

var tDDULDLSlotConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "slotIndex"},
		{Name: "symbols"},
	},
}

var tDD_UL_DL_SlotConfigSymbolsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "allDownlink"},
		{Name: "allUplink"},
		{Name: "explicit"},
	},
}

const (
	TDD_UL_DL_SlotConfig_Symbols_AllDownlink = 0
	TDD_UL_DL_SlotConfig_Symbols_AllUplink   = 1
	TDD_UL_DL_SlotConfig_Symbols_Explicit    = 2
)

var tDDULDLSlotConfigSymbolsExplicitConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofDownlinkSymbols", Optional: true},
		{Name: "nrofUplinkSymbols", Optional: true},
	},
}

type TDD_UL_DL_SlotConfig struct {
	SlotIndex TDD_UL_DL_SlotIndex
	Symbols   struct {
		Choice      int
		AllDownlink *struct{}
		AllUplink   *struct{}
		Explicit    *struct {
			NrofDownlinkSymbols *int64
			NrofUplinkSymbols   *int64
		}
	}
}

func (ie *TDD_UL_DL_SlotConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDDULDLSlotConfigConstraints)
	_ = seq

	// 1. slotIndex: ref
	{
		if err := ie.SlotIndex.Encode(e); err != nil {
			return err
		}
	}

	// 2. symbols: choice
	{
		choiceEnc := e.NewChoiceEncoder(tDD_UL_DL_SlotConfigSymbolsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Symbols.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Symbols.Choice {
		case TDD_UL_DL_SlotConfig_Symbols_AllDownlink:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_Symbols_AllUplink:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_Symbols_Explicit:
			c := (*ie.Symbols.Explicit)
			tDDULDLSlotConfigSymbolsExplicitSeq := e.NewSequenceEncoder(tDDULDLSlotConfigSymbolsExplicitConstraints)
			if err := tDDULDLSlotConfigSymbolsExplicitSeq.EncodePreamble([]bool{c.NrofDownlinkSymbols != nil, c.NrofUplinkSymbols != nil}); err != nil {
				return err
			}
			if c.NrofDownlinkSymbols != nil {
				if err := e.EncodeInteger((*c.NrofDownlinkSymbols), per.Constrained(1, common.MaxNrofSymbols_1)); err != nil {
					return err
				}
			}
			if c.NrofUplinkSymbols != nil {
				if err := e.EncodeInteger((*c.NrofUplinkSymbols), per.Constrained(1, common.MaxNrofSymbols_1)); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Symbols.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *TDD_UL_DL_SlotConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDDULDLSlotConfigConstraints)
	_ = seq

	// 1. slotIndex: ref
	{
		if err := ie.SlotIndex.Decode(d); err != nil {
			return err
		}
	}

	// 2. symbols: choice
	{
		choiceDec := d.NewChoiceDecoder(tDD_UL_DL_SlotConfigSymbolsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Symbols.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TDD_UL_DL_SlotConfig_Symbols_AllDownlink:
			ie.Symbols.AllDownlink = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_Symbols_AllUplink:
			ie.Symbols.AllUplink = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_Symbols_Explicit:
			ie.Symbols.Explicit = &struct {
				NrofDownlinkSymbols *int64
				NrofUplinkSymbols   *int64
			}{}
			c := (*ie.Symbols.Explicit)
			tDDULDLSlotConfigSymbolsExplicitSeq := d.NewSequenceDecoder(tDDULDLSlotConfigSymbolsExplicitConstraints)
			if err := tDDULDLSlotConfigSymbolsExplicitSeq.DecodePreamble(); err != nil {
				return err
			}
			if tDDULDLSlotConfigSymbolsExplicitSeq.IsComponentPresent(0) {
				c.NrofDownlinkSymbols = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSymbols_1))
				if err != nil {
					return err
				}
				(*c.NrofDownlinkSymbols) = v
			}
			if tDDULDLSlotConfigSymbolsExplicitSeq.IsComponentPresent(1) {
				c.NrofUplinkSymbols = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSymbols_1))
				if err != nil {
					return err
				}
				(*c.NrofUplinkSymbols) = v
			}
		}
	}

	return nil
}
