// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-CSI-IM-ResourceToReleaseList-r19 (line 8671).
// LTM-CSI-IM-ResourceToReleaseList-r19 ::= SEQUENCE (SIZE (1..maxNrofCSI-IM-Resources)) OF CSI-IM-ResourceId

var lTMCSIIMResourceToReleaseListR19SizeConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_Resources)

type LTM_CSI_IM_ResourceToReleaseList_r19 []CSI_IM_ResourceId

func (ie *LTM_CSI_IM_ResourceToReleaseList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lTMCSIIMResourceToReleaseListR19SizeConstraints)
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

func (ie *LTM_CSI_IM_ResourceToReleaseList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lTMCSIIMResourceToReleaseListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LTM_CSI_IM_ResourceToReleaseList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
