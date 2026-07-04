// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxSwitchingAdditionalPeriodDualUL-r18 (line 17007).

var uplinkTxSwitchingAdditionalPeriodDualULR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxSwitchingBetweenBandPairs-r18"},
		{Name: "switchingAdditionalPeriodDualUL-r18"},
	},
}

const (
	UplinkTxSwitchingAdditionalPeriodDualUL_r18_SwitchingAdditionalPeriodDualUL_r18_N35us  = 0
	UplinkTxSwitchingAdditionalPeriodDualUL_r18_SwitchingAdditionalPeriodDualUL_r18_N140us = 1
	UplinkTxSwitchingAdditionalPeriodDualUL_r18_SwitchingAdditionalPeriodDualUL_r18_N210us = 2
)

var uplinkTxSwitchingAdditionalPeriodDualULR18SwitchingAdditionalPeriodDualULR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxSwitchingAdditionalPeriodDualUL_r18_SwitchingAdditionalPeriodDualUL_r18_N35us, UplinkTxSwitchingAdditionalPeriodDualUL_r18_SwitchingAdditionalPeriodDualUL_r18_N140us, UplinkTxSwitchingAdditionalPeriodDualUL_r18_SwitchingAdditionalPeriodDualUL_r18_N210us},
}

var uplinkTxSwitchingAdditionalPeriodDualULR18UplinkTxSwitchingBetweenBandPairsR18AnotherBandPairOrBandR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "bandPairIndex2-r18"},
		{Name: "bandIndex-r18"},
	},
}

const (
	UplinkTxSwitchingAdditionalPeriodDualUL_r18_UplinkTxSwitchingBetweenBandPairs_r18_AnotherBandPairOrBand_r18_BandPairIndex2_r18 = 0
	UplinkTxSwitchingAdditionalPeriodDualUL_r18_UplinkTxSwitchingBetweenBandPairs_r18_AnotherBandPairOrBand_r18_BandIndex_r18      = 1
)

type UplinkTxSwitchingAdditionalPeriodDualUL_r18 struct {
	UplinkTxSwitchingBetweenBandPairs_r18 struct {
		BandPairIndex1_r18        int64
		AnotherBandPairOrBand_r18 struct {
			Choice             int
			BandPairIndex2_r18 *int64
			BandIndex_r18      *int64
		}
	}
	SwitchingAdditionalPeriodDualUL_r18 int64
}

func (ie *UplinkTxSwitchingAdditionalPeriodDualUL_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxSwitchingAdditionalPeriodDualULR18Constraints)
	_ = seq

	// 1. uplinkTxSwitchingBetweenBandPairs-r18: sequence
	{
		{
			c := &ie.UplinkTxSwitchingBetweenBandPairs_r18
			if err := e.EncodeInteger(c.BandPairIndex1_r18, per.Constrained(1, common.MaxULTxSwitchingBandPairs)); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(uplinkTxSwitchingAdditionalPeriodDualULR18UplinkTxSwitchingBetweenBandPairsR18AnotherBandPairOrBandR18Constraints)
				if err := choiceEnc.EncodeChoice(int64(c.AnotherBandPairOrBand_r18.Choice), false, nil); err != nil {
					return err
				}
				switch c.AnotherBandPairOrBand_r18.Choice {
				case UplinkTxSwitchingAdditionalPeriodDualUL_r18_UplinkTxSwitchingBetweenBandPairs_r18_AnotherBandPairOrBand_r18_BandPairIndex2_r18:
					if err := e.EncodeInteger((*c.AnotherBandPairOrBand_r18.BandPairIndex2_r18), per.Constrained(1, common.MaxULTxSwitchingBandPairs)); err != nil {
						return err
					}
				case UplinkTxSwitchingAdditionalPeriodDualUL_r18_UplinkTxSwitchingBetweenBandPairs_r18_AnotherBandPairOrBand_r18_BandIndex_r18:
					if err := e.EncodeInteger((*c.AnotherBandPairOrBand_r18.BandIndex_r18), per.Constrained(1, common.MaxSimultaneousBands)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 2. switchingAdditionalPeriodDualUL-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.SwitchingAdditionalPeriodDualUL_r18, uplinkTxSwitchingAdditionalPeriodDualULR18SwitchingAdditionalPeriodDualULR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkTxSwitchingAdditionalPeriodDualUL_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxSwitchingAdditionalPeriodDualULR18Constraints)
	_ = seq

	// 1. uplinkTxSwitchingBetweenBandPairs-r18: sequence
	{
		{
			c := &ie.UplinkTxSwitchingBetweenBandPairs_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxULTxSwitchingBandPairs))
				if err != nil {
					return err
				}
				c.BandPairIndex1_r18 = v
			}
			{
				choiceDec := d.NewChoiceDecoder(uplinkTxSwitchingAdditionalPeriodDualULR18UplinkTxSwitchingBetweenBandPairsR18AnotherBandPairOrBandR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.AnotherBandPairOrBand_r18.Choice = int(index)
				switch index {
				case UplinkTxSwitchingAdditionalPeriodDualUL_r18_UplinkTxSwitchingBetweenBandPairs_r18_AnotherBandPairOrBand_r18_BandPairIndex2_r18:
					c.AnotherBandPairOrBand_r18.BandPairIndex2_r18 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, common.MaxULTxSwitchingBandPairs))
					if err != nil {
						return err
					}
					(*c.AnotherBandPairOrBand_r18.BandPairIndex2_r18) = v
				case UplinkTxSwitchingAdditionalPeriodDualUL_r18_UplinkTxSwitchingBetweenBandPairs_r18_AnotherBandPairOrBand_r18_BandIndex_r18:
					c.AnotherBandPairOrBand_r18.BandIndex_r18 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, common.MaxSimultaneousBands))
					if err != nil {
						return err
					}
					(*c.AnotherBandPairOrBand_r18.BandIndex_r18) = v
				}
			}
		}
	}

	// 2. switchingAdditionalPeriodDualUL-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(uplinkTxSwitchingAdditionalPeriodDualULR18SwitchingAdditionalPeriodDualULR18Constraints)
		if err != nil {
			return err
		}
		ie.SwitchingAdditionalPeriodDualUL_r18 = v1
	}

	return nil
}
