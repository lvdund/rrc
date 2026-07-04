// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-CapabilityRAT-ContainerList (line 25461).
// UE-CapabilityRAT-ContainerList ::=    SEQUENCE (SIZE (0..maxRAT-CapabilityContainers)) OF UE-CapabilityRAT-Container

var uECapabilityRATContainerListSizeConstraints = per.SizeRange(0, common.MaxRAT_CapabilityContainers)

type UE_CapabilityRAT_ContainerList []UE_CapabilityRAT_Container

func (ie *UE_CapabilityRAT_ContainerList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uECapabilityRATContainerListSizeConstraints)
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

func (ie *UE_CapabilityRAT_ContainerList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uECapabilityRATContainerListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UE_CapabilityRAT_ContainerList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
