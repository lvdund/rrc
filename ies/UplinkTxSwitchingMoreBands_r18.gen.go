// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxSwitchingMoreBands-r18 (line 5816).

var uplinkTxSwitchingMoreBandsR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxSwitchingBandList-r18", Optional: true},
		{Name: "uplinkTxSwitchingBandPairList-r18", Optional: true},
		{Name: "uplinkTxSwitchingAssociatedBandDualUL-List-r18", Optional: true},
	},
}

var uplinkTxSwitchingMoreBandsR18UplinkTxSwitchingBandListR18Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type UplinkTxSwitchingMoreBands_r18 struct {
	UplinkTxSwitchingBandList_r18                  []FreqBandIndicatorNR
	UplinkTxSwitchingBandPairList_r18              *UplinkTxSwitchingBandPairList_r18
	UplinkTxSwitchingAssociatedBandDualUL_List_r18 *UplinkTxSwitchingAssociatedBandDualUL_List_r18
}

func (ie *UplinkTxSwitchingMoreBands_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxSwitchingMoreBandsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxSwitchingBandList_r18 != nil, ie.UplinkTxSwitchingBandPairList_r18 != nil, ie.UplinkTxSwitchingAssociatedBandDualUL_List_r18 != nil}); err != nil {
		return err
	}

	// 3. uplinkTxSwitchingBandList-r18: sequence-of
	{
		if ie.UplinkTxSwitchingBandList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(uplinkTxSwitchingMoreBandsR18UplinkTxSwitchingBandListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.UplinkTxSwitchingBandList_r18))); err != nil {
				return err
			}
			for i := range ie.UplinkTxSwitchingBandList_r18 {
				if err := ie.UplinkTxSwitchingBandList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. uplinkTxSwitchingBandPairList-r18: ref
	{
		if ie.UplinkTxSwitchingBandPairList_r18 != nil {
			if err := ie.UplinkTxSwitchingBandPairList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. uplinkTxSwitchingAssociatedBandDualUL-List-r18: ref
	{
		if ie.UplinkTxSwitchingAssociatedBandDualUL_List_r18 != nil {
			if err := ie.UplinkTxSwitchingAssociatedBandDualUL_List_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UplinkTxSwitchingMoreBands_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxSwitchingMoreBandsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. uplinkTxSwitchingBandList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(uplinkTxSwitchingMoreBandsR18UplinkTxSwitchingBandListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingBandList_r18 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkTxSwitchingBandList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. uplinkTxSwitchingBandPairList-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.UplinkTxSwitchingBandPairList_r18 = new(UplinkTxSwitchingBandPairList_r18)
			if err := ie.UplinkTxSwitchingBandPairList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. uplinkTxSwitchingAssociatedBandDualUL-List-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.UplinkTxSwitchingAssociatedBandDualUL_List_r18 = new(UplinkTxSwitchingAssociatedBandDualUL_List_r18)
			if err := ie.UplinkTxSwitchingAssociatedBandDualUL_List_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
