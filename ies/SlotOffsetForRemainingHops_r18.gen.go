// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SlotOffsetForRemainingHops-r18 (line 15728).

var slotOffsetForRemainingHopsR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "slotOffsetRemainingHops-r18"},
	},
}

var slotOffsetForRemainingHops_r18SlotOffsetRemainingHopsR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "aperiodic-r18"},
		{Name: "semi-persistent-r18"},
		{Name: "periodic-r18"},
	},
}

const (
	SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Aperiodic_r18       = 0
	SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Semi_Persistent_r18 = 1
	SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Periodic_r18        = 2
)

var slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "slotOffset-r18", Optional: true},
		{Name: "startPosition-r18", Optional: true},
	},
}

var slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-sp-r18", Optional: true},
		{Name: "periodicityAndOffset-sp-Ext-r18", Optional: true},
		{Name: "startPosition-r18", Optional: true},
	},
}

var slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-p-r18", Optional: true},
		{Name: "periodicityAndOffset-p-Ext-r18", Optional: true},
		{Name: "startPosition-r18", Optional: true},
	},
}

type SlotOffsetForRemainingHops_r18 struct {
	SlotOffsetRemainingHops_r18 struct {
		Choice        int
		Aperiodic_r18 *struct {
			SlotOffset_r18    *int64
			StartPosition_r18 *int64
		}
		Semi_Persistent_r18 *struct {
			PeriodicityAndOffset_Sp_r18     *SRS_PeriodicityAndOffset_r16
			PeriodicityAndOffset_Sp_Ext_r18 *SRS_PeriodicityAndOffsetExt_r16
			StartPosition_r18               *int64
		}
		Periodic_r18 *struct {
			PeriodicityAndOffset_P_r18     *SRS_PeriodicityAndOffset_r16
			PeriodicityAndOffset_P_Ext_r18 *SRS_PeriodicityAndOffsetExt_r16
			StartPosition_r18              *int64
		}
	}
}

func (ie *SlotOffsetForRemainingHops_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(slotOffsetForRemainingHopsR18Constraints)
	_ = seq

	// 1. slotOffsetRemainingHops-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(slotOffsetForRemainingHops_r18SlotOffsetRemainingHopsR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.SlotOffsetRemainingHops_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.SlotOffsetRemainingHops_r18.Choice {
		case SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Aperiodic_r18:
			c := (*ie.SlotOffsetRemainingHops_r18.Aperiodic_r18)
			slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq := e.NewSequenceEncoder(slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Constraints)
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq.EncodePreamble([]bool{c.SlotOffset_r18 != nil, c.StartPosition_r18 != nil}); err != nil {
				return err
			}
			if c.SlotOffset_r18 != nil {
				if err := e.EncodeInteger((*c.SlotOffset_r18), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.StartPosition_r18 != nil {
				if err := e.EncodeInteger((*c.StartPosition_r18), per.Constrained(0, 13)); err != nil {
					return err
				}
			}
		case SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Semi_Persistent_r18:
			c := (*ie.SlotOffsetRemainingHops_r18.Semi_Persistent_r18)
			slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq := e.NewSequenceEncoder(slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Constraints)
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq.EncodePreamble([]bool{c.PeriodicityAndOffset_Sp_r18 != nil, c.PeriodicityAndOffset_Sp_Ext_r18 != nil, c.StartPosition_r18 != nil}); err != nil {
				return err
			}
			if c.PeriodicityAndOffset_Sp_r18 != nil {
				if err := c.PeriodicityAndOffset_Sp_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.PeriodicityAndOffset_Sp_Ext_r18 != nil {
				if err := c.PeriodicityAndOffset_Sp_Ext_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.StartPosition_r18 != nil {
				if err := e.EncodeInteger((*c.StartPosition_r18), per.Constrained(0, 13)); err != nil {
					return err
				}
			}
		case SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Periodic_r18:
			c := (*ie.SlotOffsetRemainingHops_r18.Periodic_r18)
			slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq := e.NewSequenceEncoder(slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Constraints)
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq.EncodePreamble([]bool{c.PeriodicityAndOffset_P_r18 != nil, c.PeriodicityAndOffset_P_Ext_r18 != nil, c.StartPosition_r18 != nil}); err != nil {
				return err
			}
			if c.PeriodicityAndOffset_P_r18 != nil {
				if err := c.PeriodicityAndOffset_P_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.PeriodicityAndOffset_P_Ext_r18 != nil {
				if err := c.PeriodicityAndOffset_P_Ext_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.StartPosition_r18 != nil {
				if err := e.EncodeInteger((*c.StartPosition_r18), per.Constrained(0, 13)); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.SlotOffsetRemainingHops_r18.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SlotOffsetForRemainingHops_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(slotOffsetForRemainingHopsR18Constraints)
	_ = seq

	// 1. slotOffsetRemainingHops-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(slotOffsetForRemainingHops_r18SlotOffsetRemainingHopsR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.SlotOffsetRemainingHops_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Aperiodic_r18:
			ie.SlotOffsetRemainingHops_r18.Aperiodic_r18 = &struct {
				SlotOffset_r18    *int64
				StartPosition_r18 *int64
			}{}
			c := (*ie.SlotOffsetRemainingHops_r18.Aperiodic_r18)
			slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq := d.NewSequenceDecoder(slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Constraints)
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq.IsComponentPresent(0) {
				c.SlotOffset_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.SlotOffset_r18) = v
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18AperiodicR18Seq.IsComponentPresent(1) {
				c.StartPosition_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 13))
				if err != nil {
					return err
				}
				(*c.StartPosition_r18) = v
			}
		case SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Semi_Persistent_r18:
			ie.SlotOffsetRemainingHops_r18.Semi_Persistent_r18 = &struct {
				PeriodicityAndOffset_Sp_r18     *SRS_PeriodicityAndOffset_r16
				PeriodicityAndOffset_Sp_Ext_r18 *SRS_PeriodicityAndOffsetExt_r16
				StartPosition_r18               *int64
			}{}
			c := (*ie.SlotOffsetRemainingHops_r18.Semi_Persistent_r18)
			slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq := d.NewSequenceDecoder(slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Constraints)
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq.IsComponentPresent(0) {
				c.PeriodicityAndOffset_Sp_r18 = new(SRS_PeriodicityAndOffset_r16)
				if err := (*c.PeriodicityAndOffset_Sp_r18).Decode(d); err != nil {
					return err
				}
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq.IsComponentPresent(1) {
				c.PeriodicityAndOffset_Sp_Ext_r18 = new(SRS_PeriodicityAndOffsetExt_r16)
				if err := (*c.PeriodicityAndOffset_Sp_Ext_r18).Decode(d); err != nil {
					return err
				}
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18SemiPersistentR18Seq.IsComponentPresent(2) {
				c.StartPosition_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 13))
				if err != nil {
					return err
				}
				(*c.StartPosition_r18) = v
			}
		case SlotOffsetForRemainingHops_r18_SlotOffsetRemainingHops_r18_Periodic_r18:
			ie.SlotOffsetRemainingHops_r18.Periodic_r18 = &struct {
				PeriodicityAndOffset_P_r18     *SRS_PeriodicityAndOffset_r16
				PeriodicityAndOffset_P_Ext_r18 *SRS_PeriodicityAndOffsetExt_r16
				StartPosition_r18              *int64
			}{}
			c := (*ie.SlotOffsetRemainingHops_r18.Periodic_r18)
			slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq := d.NewSequenceDecoder(slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Constraints)
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq.IsComponentPresent(0) {
				c.PeriodicityAndOffset_P_r18 = new(SRS_PeriodicityAndOffset_r16)
				if err := (*c.PeriodicityAndOffset_P_r18).Decode(d); err != nil {
					return err
				}
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq.IsComponentPresent(1) {
				c.PeriodicityAndOffset_P_Ext_r18 = new(SRS_PeriodicityAndOffsetExt_r16)
				if err := (*c.PeriodicityAndOffset_P_Ext_r18).Decode(d); err != nil {
					return err
				}
			}
			if slotOffsetForRemainingHopsR18SlotOffsetRemainingHopsR18PeriodicR18Seq.IsComponentPresent(2) {
				c.StartPosition_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 13))
				if err != nil {
					return err
				}
				(*c.StartPosition_r18) = v
			}
		}
	}

	return nil
}
