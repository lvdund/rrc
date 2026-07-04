// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDSCH-TimeDomainResourceAllocationList-r16 (line 11521).
// PDSCH-TimeDomainResourceAllocationList-r16 ::=  SEQUENCE (SIZE(1..maxNrofDL-Allocations)) OF PDSCH-TimeDomainResourceAllocation-r16

var pDSCHTimeDomainResourceAllocationListR16SizeConstraints = per.SizeRange(1, common.MaxNrofDL_Allocations)

type PDSCH_TimeDomainResourceAllocationList_r16 []PDSCH_TimeDomainResourceAllocation_r16

func (ie *PDSCH_TimeDomainResourceAllocationList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDSCHTimeDomainResourceAllocationListR16SizeConstraints)
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

func (ie *PDSCH_TimeDomainResourceAllocationList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDSCHTimeDomainResourceAllocationListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PDSCH_TimeDomainResourceAllocationList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
