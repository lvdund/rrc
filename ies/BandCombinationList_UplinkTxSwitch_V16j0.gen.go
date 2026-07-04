// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationList-UplinkTxSwitch-v16j0 (line 16620).
// BandCombinationList-UplinkTxSwitch-v16j0 ::= SEQUENCE (SIZE (1..maxBandComb)) OF BandCombination-UplinkTxSwitch-v16j0

var bandCombinationListUplinkTxSwitchV16j0SizeConstraints = per.SizeRange(1, common.MaxBandComb)

type BandCombinationList_UplinkTxSwitch_V16j0 []BandCombination_UplinkTxSwitch_V16j0

func (ie *BandCombinationList_UplinkTxSwitch_V16j0) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationListUplinkTxSwitchV16j0SizeConstraints)
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

func (ie *BandCombinationList_UplinkTxSwitch_V16j0) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationListUplinkTxSwitchV16j0SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationList_UplinkTxSwitch_V16j0, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
