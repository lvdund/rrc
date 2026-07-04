// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PLMN-RAN-AreaCellList (line 1350).
// PLMN-RAN-AreaCellList ::=           SEQUENCE (SIZE (1.. maxPLMNIdentities)) OF PLMN-RAN-AreaCell

var pLMNRANAreaCellListSizeConstraints = per.SizeRange(1, common.MaxPLMNIdentities)

type PLMN_RAN_AreaCellList []PLMN_RAN_AreaCell

func (ie *PLMN_RAN_AreaCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pLMNRANAreaCellListSizeConstraints)
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

func (ie *PLMN_RAN_AreaCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pLMNRANAreaCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PLMN_RAN_AreaCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
