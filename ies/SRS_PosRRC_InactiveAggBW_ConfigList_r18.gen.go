// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosRRC-InactiveAggBW-ConfigList-r18 (line 1504).

var sRSPosRRCInactiveAggBWConfigListR18SizeConstraints = per.SizeRange(1, common.MaxNrOfLinkedSRS_PosResSetCombInactive_r18)

type SRS_PosRRC_InactiveAggBW_ConfigList_r18 []SRS_InactivePosResourceSetLinkedForAggBW_List_r18

func (ie *SRS_PosRRC_InactiveAggBW_ConfigList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sRSPosRRCInactiveAggBWConfigListR18SizeConstraints)
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

func (ie *SRS_PosRRC_InactiveAggBW_ConfigList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sRSPosRRCInactiveAggBWConfigListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SRS_PosRRC_InactiveAggBW_ConfigList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
