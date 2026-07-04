// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-CSI-ResourceConfigToReleaseList-r18 (line 8749).
// LTM-CSI-ResourceConfigToReleaseList-r18 ::= SEQUENCE (SIZE (1..maxNrofLTM-CSI-ResourceConfigurations-r18)) OF LTM-CSI-ResourceConfigId-r18

var lTMCSIResourceConfigToReleaseListR18SizeConstraints = per.SizeRange(1, common.MaxNrofLTM_CSI_ResourceConfigurations_r18)

type LTM_CSI_ResourceConfigToReleaseList_r18 []LTM_CSI_ResourceConfigId_r18

func (ie *LTM_CSI_ResourceConfigToReleaseList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lTMCSIResourceConfigToReleaseListR18SizeConstraints)
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

func (ie *LTM_CSI_ResourceConfigToReleaseList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lTMCSIResourceConfigToReleaseListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LTM_CSI_ResourceConfigToReleaseList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
