// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-TimeDomainResourceAllocationList-r16 (line 12709).
// PUSCH-TimeDomainResourceAllocationList-r16 ::=  SEQUENCE (SIZE(1..maxNrofUL-Allocations-r16)) OF PUSCH-TimeDomainResourceAllocation-r16

var pUSCHTimeDomainResourceAllocationListR16SizeConstraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

type PUSCH_TimeDomainResourceAllocationList_r16 []PUSCH_TimeDomainResourceAllocation_r16

func (ie *PUSCH_TimeDomainResourceAllocationList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pUSCHTimeDomainResourceAllocationListR16SizeConstraints)
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

func (ie *PUSCH_TimeDomainResourceAllocationList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pUSCHTimeDomainResourceAllocationListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PUSCH_TimeDomainResourceAllocationList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
