// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDD-UL-DL-SlotConfig-IAB-MT-r16 (line 16119).

var tDDULDLSlotConfigIABMTR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "slotIndex-r16"},
		{Name: "symbols-IAB-MT-r16"},
	},
}

var tDD_UL_DL_SlotConfig_IAB_MT_r16SymbolsIABMTR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "allDownlink-r16"},
		{Name: "allUplink-r16"},
		{Name: "explicit-r16"},
		{Name: "explicit-IAB-MT-r16"},
	},
}

const (
	TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_AllDownlink_r16     = 0
	TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_AllUplink_r16       = 1
	TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_Explicit_r16        = 2
	TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_Explicit_IAB_MT_r16 = 3
)

var tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofDownlinkSymbols-r16", Optional: true},
		{Name: "nrofUplinkSymbols-r16", Optional: true},
	},
}

var tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofDownlinkSymbols-r16", Optional: true},
		{Name: "nrofUplinkSymbols-r16", Optional: true},
	},
}

type TDD_UL_DL_SlotConfig_IAB_MT_r16 struct {
	SlotIndex_r16      TDD_UL_DL_SlotIndex
	Symbols_IAB_MT_r16 struct {
		Choice          int
		AllDownlink_r16 *struct{}
		AllUplink_r16   *struct{}
		Explicit_r16    *struct {
			NrofDownlinkSymbols_r16 *int64
			NrofUplinkSymbols_r16   *int64
		}
		Explicit_IAB_MT_r16 *struct {
			NrofDownlinkSymbols_r16 *int64
			NrofUplinkSymbols_r16   *int64
		}
	}
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDDULDLSlotConfigIABMTR16Constraints)
	_ = seq

	// 1. slotIndex-r16: ref
	{
		if err := ie.SlotIndex_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. symbols-IAB-MT-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(tDD_UL_DL_SlotConfig_IAB_MT_r16SymbolsIABMTR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Symbols_IAB_MT_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Symbols_IAB_MT_r16.Choice {
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_AllDownlink_r16:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_AllUplink_r16:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_Explicit_r16:
			c := (*ie.Symbols_IAB_MT_r16.Explicit_r16)
			tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Seq := e.NewSequenceEncoder(tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Constraints)
			if err := tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Seq.EncodePreamble([]bool{c.NrofDownlinkSymbols_r16 != nil, c.NrofUplinkSymbols_r16 != nil}); err != nil {
				return err
			}
			if c.NrofDownlinkSymbols_r16 != nil {
				if err := e.EncodeInteger((*c.NrofDownlinkSymbols_r16), per.Constrained(1, common.MaxNrofSymbols_1)); err != nil {
					return err
				}
			}
			if c.NrofUplinkSymbols_r16 != nil {
				if err := e.EncodeInteger((*c.NrofUplinkSymbols_r16), per.Constrained(1, common.MaxNrofSymbols_1)); err != nil {
					return err
				}
			}
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_Explicit_IAB_MT_r16:
			c := (*ie.Symbols_IAB_MT_r16.Explicit_IAB_MT_r16)
			tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Seq := e.NewSequenceEncoder(tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Constraints)
			if err := tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Seq.EncodePreamble([]bool{c.NrofDownlinkSymbols_r16 != nil, c.NrofUplinkSymbols_r16 != nil}); err != nil {
				return err
			}
			if c.NrofDownlinkSymbols_r16 != nil {
				if err := e.EncodeInteger((*c.NrofDownlinkSymbols_r16), per.Constrained(1, common.MaxNrofSymbols_1)); err != nil {
					return err
				}
			}
			if c.NrofUplinkSymbols_r16 != nil {
				if err := e.EncodeInteger((*c.NrofUplinkSymbols_r16), per.Constrained(1, common.MaxNrofSymbols_1)); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Symbols_IAB_MT_r16.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDDULDLSlotConfigIABMTR16Constraints)
	_ = seq

	// 1. slotIndex-r16: ref
	{
		if err := ie.SlotIndex_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. symbols-IAB-MT-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(tDD_UL_DL_SlotConfig_IAB_MT_r16SymbolsIABMTR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Symbols_IAB_MT_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_AllDownlink_r16:
			ie.Symbols_IAB_MT_r16.AllDownlink_r16 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_AllUplink_r16:
			ie.Symbols_IAB_MT_r16.AllUplink_r16 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_Explicit_r16:
			ie.Symbols_IAB_MT_r16.Explicit_r16 = &struct {
				NrofDownlinkSymbols_r16 *int64
				NrofUplinkSymbols_r16   *int64
			}{}
			c := (*ie.Symbols_IAB_MT_r16.Explicit_r16)
			tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Seq := d.NewSequenceDecoder(tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Constraints)
			if err := tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Seq.IsComponentPresent(0) {
				c.NrofDownlinkSymbols_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSymbols_1))
				if err != nil {
					return err
				}
				(*c.NrofDownlinkSymbols_r16) = v
			}
			if tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitR16Seq.IsComponentPresent(1) {
				c.NrofUplinkSymbols_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSymbols_1))
				if err != nil {
					return err
				}
				(*c.NrofUplinkSymbols_r16) = v
			}
		case TDD_UL_DL_SlotConfig_IAB_MT_r16_Symbols_IAB_MT_r16_Explicit_IAB_MT_r16:
			ie.Symbols_IAB_MT_r16.Explicit_IAB_MT_r16 = &struct {
				NrofDownlinkSymbols_r16 *int64
				NrofUplinkSymbols_r16   *int64
			}{}
			c := (*ie.Symbols_IAB_MT_r16.Explicit_IAB_MT_r16)
			tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Seq := d.NewSequenceDecoder(tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Constraints)
			if err := tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Seq.IsComponentPresent(0) {
				c.NrofDownlinkSymbols_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSymbols_1))
				if err != nil {
					return err
				}
				(*c.NrofDownlinkSymbols_r16) = v
			}
			if tDDULDLSlotConfigIABMTR16SymbolsIABMTR16ExplicitIABMTR16Seq.IsComponentPresent(1) {
				c.NrofUplinkSymbols_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSymbols_1))
				if err != nil {
					return err
				}
				(*c.NrofUplinkSymbols_r16) = v
			}
		}
	}

	return nil
}
