// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-CSI-IM-ResourceSetToAddModList-r19 (line 8672).
// LTM-CSI-IM-ResourceSetToAddModList-r19 ::= SEQUENCE (SIZE (1..maxNrofCSI-IM-ResourceSets)) OF CSI-IM-ResourceSet

var lTMCSIIMResourceSetToAddModListR19SizeConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_ResourceSets)

type LTM_CSI_IM_ResourceSetToAddModList_r19 []CSI_IM_ResourceSet

func (ie *LTM_CSI_IM_ResourceSetToAddModList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lTMCSIIMResourceSetToAddModListR19SizeConstraints)
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

func (ie *LTM_CSI_IM_ResourceSetToAddModList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lTMCSIIMResourceSetToAddModListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LTM_CSI_IM_ResourceSetToAddModList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
