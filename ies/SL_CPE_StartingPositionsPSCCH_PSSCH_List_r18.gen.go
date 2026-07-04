// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CPE-StartingPositionsPSCCH-PSSCH-List-r18 (line 28004).
// SL-CPE-StartingPositionsPSCCH-PSSCH-List-r18 ::= SEQUENCE (SIZE (8)) OF SL-CPE-StartingPositionsPSCCH-PSSCH-r18

var sLCPEStartingPositionsPSCCHPSSCHListR18SizeConstraints = per.FixedSize(8)

type SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18 []SL_CPE_StartingPositionsPSCCH_PSSCH_r18

func (ie *SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLCPEStartingPositionsPSCCHPSSCHListR18SizeConstraints)
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

func (ie *SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLCPEStartingPositionsPSCCHPSSCHListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
