// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-TimeDomainResourceAllocationList (line 12701).
// PUSCH-TimeDomainResourceAllocationList ::=  SEQUENCE (SIZE(1..maxNrofUL-Allocations)) OF PUSCH-TimeDomainResourceAllocation

var pUSCHTimeDomainResourceAllocationListSizeConstraints = per.SizeRange(1, common.MaxNrofUL_Allocations)

type PUSCH_TimeDomainResourceAllocationList []PUSCH_TimeDomainResourceAllocation

func (ie *PUSCH_TimeDomainResourceAllocationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pUSCHTimeDomainResourceAllocationListSizeConstraints)
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

func (ie *PUSCH_TimeDomainResourceAllocationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pUSCHTimeDomainResourceAllocationListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PUSCH_TimeDomainResourceAllocationList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
