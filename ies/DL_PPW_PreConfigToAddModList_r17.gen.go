// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DL-PPW-PreConfigToAddModList-r17 (line 5301).
// DL-PPW-PreConfigToAddModList-r17 ::=    SEQUENCE (SIZE (1..maxNrofPPW-Config-r17)) OF DL-PPW-PreConfig-r17

var dLPPWPreConfigToAddModListR17SizeConstraints = per.SizeRange(1, common.MaxNrofPPW_Config_r17)

type DL_PPW_PreConfigToAddModList_r17 []DL_PPW_PreConfig_r17

func (ie *DL_PPW_PreConfigToAddModList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dLPPWPreConfigToAddModListR17SizeConstraints)
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

func (ie *DL_PPW_PreConfigToAddModList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dLPPWPreConfigToAddModListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DL_PPW_PreConfigToAddModList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
