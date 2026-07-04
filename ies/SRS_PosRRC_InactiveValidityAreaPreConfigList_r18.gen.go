// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosRRC-InactiveValidityAreaPreConfigList-r18 (line 1478).
// SRS-PosRRC-InactiveValidityAreaPreConfigList-r18  ::= SEQUENCE (SIZE(1..maxNrOfVA-r18)) OF SRS-PosRRC-InactiveValidityAreaConfig-r18

var sRSPosRRCInactiveValidityAreaPreConfigListR18SizeConstraints = per.SizeRange(1, common.MaxNrOfVA_r18)

type SRS_PosRRC_InactiveValidityAreaPreConfigList_r18 []SRS_PosRRC_InactiveValidityAreaConfig_r18

func (ie *SRS_PosRRC_InactiveValidityAreaPreConfigList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sRSPosRRCInactiveValidityAreaPreConfigListR18SizeConstraints)
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

func (ie *SRS_PosRRC_InactiveValidityAreaPreConfigList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sRSPosRRCInactiveValidityAreaPreConfigListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SRS_PosRRC_InactiveValidityAreaPreConfigList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
