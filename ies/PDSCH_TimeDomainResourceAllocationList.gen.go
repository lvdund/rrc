// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDSCH-TimeDomainResourceAllocationList (line 11513).
// PDSCH-TimeDomainResourceAllocationList ::=  SEQUENCE (SIZE(1..maxNrofDL-Allocations)) OF PDSCH-TimeDomainResourceAllocation

var pDSCHTimeDomainResourceAllocationListSizeConstraints = per.SizeRange(1, common.MaxNrofDL_Allocations)

type PDSCH_TimeDomainResourceAllocationList []PDSCH_TimeDomainResourceAllocation

func (ie *PDSCH_TimeDomainResourceAllocationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDSCHTimeDomainResourceAllocationListSizeConstraints)
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

func (ie *PDSCH_TimeDomainResourceAllocationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDSCHTimeDomainResourceAllocationListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PDSCH_TimeDomainResourceAllocationList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
