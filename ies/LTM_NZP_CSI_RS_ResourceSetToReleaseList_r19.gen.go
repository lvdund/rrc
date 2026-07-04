// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-NZP-CSI-RS-ResourceSetToReleaseList-r19 (line 8669).
// LTM-NZP-CSI-RS-ResourceSetToReleaseList-r19 ::= SEQUENCE (SIZE (1..maxNrofNZP-CSI-RS-ResourceSets)) OF NZP-CSI-RS-ResourceSetId

var lTMNZPCSIRSResourceSetToReleaseListR19SizeConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourceSets)

type LTM_NZP_CSI_RS_ResourceSetToReleaseList_r19 []NZP_CSI_RS_ResourceSetId

func (ie *LTM_NZP_CSI_RS_ResourceSetToReleaseList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lTMNZPCSIRSResourceSetToReleaseListR19SizeConstraints)
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

func (ie *LTM_NZP_CSI_RS_ResourceSetToReleaseList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lTMNZPCSIRSResourceSetToReleaseListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LTM_NZP_CSI_RS_ResourceSetToReleaseList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
