// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ULTxSwitchingBandPair-r18 (line 16965).

var uLTxSwitchingBandPairR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandIndexUL1-r18"},
		{Name: "bandIndexUL2-r18"},
		{Name: "uplinkTxSwitchingOptionForBandPair-r18"},
		{Name: "uplinkTxSwitchingPeriodForBandPair-r18"},
		{Name: "uplinkTxSwitching-DL-Interruption-r18", Optional: true},
		{Name: "uplinkTxSwitchingPeriodUnaffectedBandDualUL-List-r18", Optional: true},
	},
}

var uLTxSwitchingBandPairR18BandIndexUL1R18Constraints = per.Constrained(1, common.MaxSimultaneousBands)

var uLTxSwitchingBandPairR18BandIndexUL2R18Constraints = per.Constrained(1, common.MaxSimultaneousBands)

const (
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingOptionForBandPair_r18_SwitchedUL = 0
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingOptionForBandPair_r18_DualUL     = 1
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingOptionForBandPair_r18_Both       = 2
)

var uLTxSwitchingBandPairR18UplinkTxSwitchingOptionForBandPairR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULTxSwitchingBandPair_r18_UplinkTxSwitchingOptionForBandPair_r18_SwitchedUL, ULTxSwitchingBandPair_r18_UplinkTxSwitchingOptionForBandPair_r18_DualUL, ULTxSwitchingBandPair_r18_UplinkTxSwitchingOptionForBandPair_r18_Both},
}

var uLTxSwitchingBandPairR18UplinkTxSwitchingDLInterruptionR18Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

var uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodUnaffectedBandDualULListR18Constraints = per.SizeRange(1, common.MaxSimultaneousBands_2_r18)

var uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "switchingPeriodFor2T-r18", Optional: true},
		{Name: "switchingPeriodFor1T-r18"},
	},
}

const (
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor2T_r18_N35us  = 0
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor2T_r18_N140us = 1
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor2T_r18_N210us = 2
)

var uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18SwitchingPeriodFor2TR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor2T_r18_N35us, ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor2T_r18_N140us, ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor2T_r18_N210us},
}

const (
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor1T_r18_N35us  = 0
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor1T_r18_N140us = 1
	ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor1T_r18_N210us = 2
)

var uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18SwitchingPeriodFor1TR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor1T_r18_N35us, ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor1T_r18_N140us, ULTxSwitchingBandPair_r18_UplinkTxSwitchingPeriodForBandPair_r18_SwitchingPeriodFor1T_r18_N210us},
}

type ULTxSwitchingBandPair_r18 struct {
	BandIndexUL1_r18                       int64
	BandIndexUL2_r18                       int64
	UplinkTxSwitchingOptionForBandPair_r18 int64
	UplinkTxSwitchingPeriodForBandPair_r18 struct {
		SwitchingPeriodFor2T_r18 *int64
		SwitchingPeriodFor1T_r18 int64
	}
	UplinkTxSwitching_DL_Interruption_r18                *per.BitString
	UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18 []SwitchingPeriodUnaffectedBandDualUL_r18
}

func (ie *ULTxSwitchingBandPair_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLTxSwitchingBandPairR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxSwitching_DL_Interruption_r18 != nil, ie.UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18 != nil}); err != nil {
		return err
	}

	// 2. bandIndexUL1-r18: integer
	{
		if err := e.EncodeInteger(ie.BandIndexUL1_r18, uLTxSwitchingBandPairR18BandIndexUL1R18Constraints); err != nil {
			return err
		}
	}

	// 3. bandIndexUL2-r18: integer
	{
		if err := e.EncodeInteger(ie.BandIndexUL2_r18, uLTxSwitchingBandPairR18BandIndexUL2R18Constraints); err != nil {
			return err
		}
	}

	// 4. uplinkTxSwitchingOptionForBandPair-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.UplinkTxSwitchingOptionForBandPair_r18, uLTxSwitchingBandPairR18UplinkTxSwitchingOptionForBandPairR18Constraints); err != nil {
			return err
		}
	}

	// 5. uplinkTxSwitchingPeriodForBandPair-r18: sequence
	{
		{
			c := &ie.UplinkTxSwitchingPeriodForBandPair_r18
			uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Seq := e.NewSequenceEncoder(uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Constraints)
			if err := uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Seq.EncodePreamble([]bool{c.SwitchingPeriodFor2T_r18 != nil}); err != nil {
				return err
			}
			if c.SwitchingPeriodFor2T_r18 != nil {
				if err := e.EncodeEnumerated((*c.SwitchingPeriodFor2T_r18), uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18SwitchingPeriodFor2TR18Constraints); err != nil {
					return err
				}
			}
			if err := e.EncodeEnumerated(c.SwitchingPeriodFor1T_r18, uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18SwitchingPeriodFor1TR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. uplinkTxSwitching-DL-Interruption-r18: bit-string
	{
		if ie.UplinkTxSwitching_DL_Interruption_r18 != nil {
			if err := e.EncodeBitString(*ie.UplinkTxSwitching_DL_Interruption_r18, uLTxSwitchingBandPairR18UplinkTxSwitchingDLInterruptionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. uplinkTxSwitchingPeriodUnaffectedBandDualUL-List-r18: sequence-of
	{
		if ie.UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodUnaffectedBandDualULListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18))); err != nil {
				return err
			}
			for i := range ie.UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18 {
				if err := ie.UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *ULTxSwitchingBandPair_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLTxSwitchingBandPairR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandIndexUL1-r18: integer
	{
		v0, err := d.DecodeInteger(uLTxSwitchingBandPairR18BandIndexUL1R18Constraints)
		if err != nil {
			return err
		}
		ie.BandIndexUL1_r18 = v0
	}

	// 3. bandIndexUL2-r18: integer
	{
		v1, err := d.DecodeInteger(uLTxSwitchingBandPairR18BandIndexUL2R18Constraints)
		if err != nil {
			return err
		}
		ie.BandIndexUL2_r18 = v1
	}

	// 4. uplinkTxSwitchingOptionForBandPair-r18: enumerated
	{
		v2, err := d.DecodeEnumerated(uLTxSwitchingBandPairR18UplinkTxSwitchingOptionForBandPairR18Constraints)
		if err != nil {
			return err
		}
		ie.UplinkTxSwitchingOptionForBandPair_r18 = v2
	}

	// 5. uplinkTxSwitchingPeriodForBandPair-r18: sequence
	{
		{
			c := &ie.UplinkTxSwitchingPeriodForBandPair_r18
			uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Seq := d.NewSequenceDecoder(uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Constraints)
			if err := uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18Seq.IsComponentPresent(0) {
				c.SwitchingPeriodFor2T_r18 = new(int64)
				v, err := d.DecodeEnumerated(uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18SwitchingPeriodFor2TR18Constraints)
				if err != nil {
					return err
				}
				(*c.SwitchingPeriodFor2T_r18) = v
			}
			{
				v, err := d.DecodeEnumerated(uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodForBandPairR18SwitchingPeriodFor1TR18Constraints)
				if err != nil {
					return err
				}
				c.SwitchingPeriodFor1T_r18 = v
			}
		}
	}

	// 6. uplinkTxSwitching-DL-Interruption-r18: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(uLTxSwitchingBandPairR18UplinkTxSwitchingDLInterruptionR18Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_DL_Interruption_r18 = &v
		}
	}

	// 7. uplinkTxSwitchingPeriodUnaffectedBandDualUL-List-r18: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(uLTxSwitchingBandPairR18UplinkTxSwitchingPeriodUnaffectedBandDualULListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18 = make([]SwitchingPeriodUnaffectedBandDualUL_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkTxSwitchingPeriodUnaffectedBandDualUL_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
