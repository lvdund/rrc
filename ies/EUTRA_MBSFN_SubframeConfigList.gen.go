// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-MBSFN-SubframeConfigList (line 26139).
// EUTRA-MBSFN-SubframeConfigList ::= SEQUENCE (SIZE (1..maxMBSFN-Allocations)) OF EUTRA-MBSFN-SubframeConfig

var eUTRAMBSFNSubframeConfigListSizeConstraints = per.SizeRange(1, common.MaxMBSFN_Allocations)

type EUTRA_MBSFN_SubframeConfigList []EUTRA_MBSFN_SubframeConfig

func (ie *EUTRA_MBSFN_SubframeConfigList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRAMBSFNSubframeConfigListSizeConstraints)
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

func (ie *EUTRA_MBSFN_SubframeConfigList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRAMBSFNSubframeConfigListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_MBSFN_SubframeConfigList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
