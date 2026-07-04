// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxSwitchingBandPairList-r18 (line 5823).
// UplinkTxSwitchingBandPairList-r18::=      SEQUENCE (SIZE (1.. maxULTxSwitchingBandPairs)) OF UplinkTxSwitchingBandPairConfig-r18

var uplinkTxSwitchingBandPairListR18SizeConstraints = per.SizeRange(1, common.MaxULTxSwitchingBandPairs)

type UplinkTxSwitchingBandPairList_r18 []UplinkTxSwitchingBandPairConfig_r18

func (ie *UplinkTxSwitchingBandPairList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uplinkTxSwitchingBandPairListR18SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UplinkTxSwitchingBandPairList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uplinkTxSwitchingBandPairListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UplinkTxSwitchingBandPairList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
