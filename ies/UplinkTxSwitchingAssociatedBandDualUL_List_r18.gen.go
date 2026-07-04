// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxSwitchingAssociatedBandDualUL-List-r18 (line 5834).
// UplinkTxSwitchingAssociatedBandDualUL-List-r18::= SEQUENCE (SIZE (0..maxSimultaneousBands)) OF UplinkTxSwitchingAssociatedBandDualUL-r18

var uplinkTxSwitchingAssociatedBandDualULListR18SizeConstraints = per.SizeRange(0, common.MaxSimultaneousBands)

type UplinkTxSwitchingAssociatedBandDualUL_List_r18 []UplinkTxSwitchingAssociatedBandDualUL_r18

func (ie *UplinkTxSwitchingAssociatedBandDualUL_List_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uplinkTxSwitchingAssociatedBandDualULListR18SizeConstraints)
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

func (ie *UplinkTxSwitchingAssociatedBandDualUL_List_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uplinkTxSwitchingAssociatedBandDualULListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UplinkTxSwitchingAssociatedBandDualUL_List_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
