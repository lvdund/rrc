// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB26-r19 (line 4687).

var sIB26R19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "od-SIB1-ConfigList-r19"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB26R19OdSIB1ConfigListR19Constraints = per.SizeRange(1, common.MaxNrofOD_SIB1_r19)

var sIB26R19LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB26_r19 struct {
	Od_SIB1_ConfigList_r19   []OD_SIB1_Config_r19
	LateNonCriticalExtension []byte
}

func (ie *SIB26_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB26R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. od-SIB1-ConfigList-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sIB26R19OdSIB1ConfigListR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Od_SIB1_ConfigList_r19))); err != nil {
			return err
		}
		for i := range ie.Od_SIB1_ConfigList_r19 {
			if err := ie.Od_SIB1_ConfigList_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB26R19LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB26_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB26R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. od-SIB1-ConfigList-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sIB26R19OdSIB1ConfigListR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Od_SIB1_ConfigList_r19 = make([]OD_SIB1_Config_r19, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Od_SIB1_ConfigList_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB26R19LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
