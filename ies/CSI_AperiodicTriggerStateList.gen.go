// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-AperiodicTriggerStateList (line 6844).
// CSI-AperiodicTriggerStateList ::=   SEQUENCE (SIZE (1..maxNrOfCSI-AperiodicTriggers)) OF CSI-AperiodicTriggerState

var cSIAperiodicTriggerStateListSizeConstraints = per.SizeRange(1, common.MaxNrOfCSI_AperiodicTriggers)

type CSI_AperiodicTriggerStateList []CSI_AperiodicTriggerState

func (ie *CSI_AperiodicTriggerStateList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cSIAperiodicTriggerStateListSizeConstraints)
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

func (ie *CSI_AperiodicTriggerStateList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cSIAperiodicTriggerStateListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CSI_AperiodicTriggerStateList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
