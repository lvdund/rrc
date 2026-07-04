// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationList-UplinkTxSwitch-v1830 (line 16642).
// BandCombinationList-UplinkTxSwitch-v1830 ::= SEQUENCE (SIZE (1..maxBandComb)) OF BandCombination-UplinkTxSwitch-v1830

var bandCombinationListUplinkTxSwitchV1830SizeConstraints = per.SizeRange(1, common.MaxBandComb)

type BandCombinationList_UplinkTxSwitch_v1830 []BandCombination_UplinkTxSwitch_v1830

func (ie *BandCombinationList_UplinkTxSwitch_v1830) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationListUplinkTxSwitchV1830SizeConstraints)
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

func (ie *BandCombinationList_UplinkTxSwitch_v1830) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationListUplinkTxSwitchV1830SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationList_UplinkTxSwitch_v1830, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
