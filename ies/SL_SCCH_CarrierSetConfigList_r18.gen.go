// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-SCCH-CarrierSetConfigList-r18 (line 27028).
// SL-SCCH-CarrierSetConfigList-r18 ::= SEQUENCE (SIZE (1..maxNrofSL-CarrierSetConfig-r18)) OF SL-SCCH-CarrierSetConfig-r18

var sLSCCHCarrierSetConfigListR18SizeConstraints = per.SizeRange(1, common.MaxNrofSL_CarrierSetConfig_r18)

type SL_SCCH_CarrierSetConfigList_r18 []SL_SCCH_CarrierSetConfig_r18

func (ie *SL_SCCH_CarrierSetConfigList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLSCCHCarrierSetConfigListR18SizeConstraints)
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

func (ie *SL_SCCH_CarrierSetConfigList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLSCCHCarrierSetConfigListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_SCCH_CarrierSetConfigList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
