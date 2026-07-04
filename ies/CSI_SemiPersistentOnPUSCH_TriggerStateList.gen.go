// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-SemiPersistentOnPUSCH-TriggerStateList (line 7597).
// CSI-SemiPersistentOnPUSCH-TriggerStateList ::= SEQUENCE(SIZE (1..maxNrOfSemiPersistentPUSCH-Triggers)) OF CSI-SemiPersistentOnPUSCH-TriggerState

var cSISemiPersistentOnPUSCHTriggerStateListSizeConstraints = per.SizeRange(1, common.MaxNrOfSemiPersistentPUSCH_Triggers)

type CSI_SemiPersistentOnPUSCH_TriggerStateList []CSI_SemiPersistentOnPUSCH_TriggerState

func (ie *CSI_SemiPersistentOnPUSCH_TriggerStateList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cSISemiPersistentOnPUSCHTriggerStateListSizeConstraints)
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

func (ie *CSI_SemiPersistentOnPUSCH_TriggerStateList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cSISemiPersistentOnPUSCHTriggerStateListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CSI_SemiPersistentOnPUSCH_TriggerStateList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
