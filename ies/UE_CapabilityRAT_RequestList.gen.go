// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-CapabilityRAT-RequestList (line 25471).
// UE-CapabilityRAT-RequestList ::=        SEQUENCE (SIZE (1..maxRAT-CapabilityContainers)) OF UE-CapabilityRAT-Request

var uECapabilityRATRequestListSizeConstraints = per.SizeRange(1, common.MaxRAT_CapabilityContainers)

type UE_CapabilityRAT_RequestList []UE_CapabilityRAT_Request

func (ie *UE_CapabilityRAT_RequestList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uECapabilityRATRequestListSizeConstraints)
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

func (ie *UE_CapabilityRAT_RequestList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uECapabilityRATRequestListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UE_CapabilityRAT_RequestList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
