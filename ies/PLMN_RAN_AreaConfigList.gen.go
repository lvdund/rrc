// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PLMN-RAN-AreaConfigList (line 1357).
// PLMN-RAN-AreaConfigList ::=         SEQUENCE (SIZE (1..maxPLMNIdentities)) OF PLMN-RAN-AreaConfig

var pLMNRANAreaConfigListSizeConstraints = per.SizeRange(1, common.MaxPLMNIdentities)

type PLMN_RAN_AreaConfigList []PLMN_RAN_AreaConfig

func (ie *PLMN_RAN_AreaConfigList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pLMNRANAreaConfigListSizeConstraints)
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

func (ie *PLMN_RAN_AreaConfigList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pLMNRANAreaConfigListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PLMN_RAN_AreaConfigList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
