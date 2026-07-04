// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-UE-AssistanceInformationNR-r16 (line 2700).
// SL-UE-AssistanceInformationNR-r16 ::= SEQUENCE (SIZE (1..maxNrofTrafficPattern-r16)) OF SL-TrafficPatternInfo-r16

var sLUEAssistanceInformationNRR16SizeConstraints = per.SizeRange(1, common.MaxNrofTrafficPattern_r16)

type SL_UE_AssistanceInformationNR_r16 []SL_TrafficPatternInfo_r16

func (ie *SL_UE_AssistanceInformationNR_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLUEAssistanceInformationNRR16SizeConstraints)
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

func (ie *SL_UE_AssistanceInformationNR_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLUEAssistanceInformationNRR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_UE_AssistanceInformationNR_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
