// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PRS-UE-AssistanceInformationNR-r18 (line 2791).
// SL-PRS-UE-AssistanceInformationNR-r18 ::= SEQUENCE (SIZE (1..maxNrofSL-PRS-TxConfig-r18)) OF SL-PRS-TxInfo-r18

var sLPRSUEAssistanceInformationNRR18SizeConstraints = per.SizeRange(1, common.MaxNrofSL_PRS_TxConfig_r18)

type SL_PRS_UE_AssistanceInformationNR_r18 []SL_PRS_TxInfo_r18

func (ie *SL_PRS_UE_AssistanceInformationNR_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLPRSUEAssistanceInformationNRR18SizeConstraints)
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

func (ie *SL_PRS_UE_AssistanceInformationNR_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLPRSUEAssistanceInformationNRR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_PRS_UE_AssistanceInformationNR_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
