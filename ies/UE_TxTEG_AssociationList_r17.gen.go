// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-TxTEG-AssociationList-r17 (line 3592).
// UE-TxTEG-AssociationList-r17 ::= SEQUENCE (SIZE (1..maxNrOfTxTEGReport-r17)) OF UE-TxTEG-Association-r17

var uETxTEGAssociationListR17SizeConstraints = per.SizeRange(1, common.MaxNrOfTxTEGReport_r17)

type UE_TxTEG_AssociationList_r17 []UE_TxTEG_Association_r17

func (ie *UE_TxTEG_AssociationList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uETxTEGAssociationListR17SizeConstraints)
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

func (ie *UE_TxTEG_AssociationList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uETxTEGAssociationListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UE_TxTEG_AssociationList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
