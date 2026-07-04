// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: Adapt-SSB-Config-r19 (line 5858).

var adaptSSBConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-BurstPeriodicityList-r19", Optional: true},
		{Name: "posInDCI-SSB-PeriodicityIndicationForScell-r19", Optional: true},
	},
}

var adaptSSBConfigR19SsbBurstPeriodicityListR19Constraints = per.SizeRange(1, 2)

var adaptSSBConfigR19PosInDCISSBPeriodicityIndicationForScellR19Constraints = per.Constrained(1, common.MaxDCI_2_9_Size_r18)

type Adapt_SSB_Config_r19 struct {
	Ssb_BurstPeriodicityList_r19                   []Adapt_SSB_BurstPeriodicity_r19
	PosInDCI_SSB_PeriodicityIndicationForScell_r19 *int64
}

func (ie *Adapt_SSB_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(adaptSSBConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_BurstPeriodicityList_r19 != nil, ie.PosInDCI_SSB_PeriodicityIndicationForScell_r19 != nil}); err != nil {
		return err
	}

	// 2. ssb-BurstPeriodicityList-r19: sequence-of
	{
		if ie.Ssb_BurstPeriodicityList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(adaptSSBConfigR19SsbBurstPeriodicityListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ssb_BurstPeriodicityList_r19))); err != nil {
				return err
			}
			for i := range ie.Ssb_BurstPeriodicityList_r19 {
				if err := ie.Ssb_BurstPeriodicityList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. posInDCI-SSB-PeriodicityIndicationForScell-r19: integer
	{
		if ie.PosInDCI_SSB_PeriodicityIndicationForScell_r19 != nil {
			if err := e.EncodeInteger(*ie.PosInDCI_SSB_PeriodicityIndicationForScell_r19, adaptSSBConfigR19PosInDCISSBPeriodicityIndicationForScellR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Adapt_SSB_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(adaptSSBConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-BurstPeriodicityList-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(adaptSSBConfigR19SsbBurstPeriodicityListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ssb_BurstPeriodicityList_r19 = make([]Adapt_SSB_BurstPeriodicity_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ssb_BurstPeriodicityList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. posInDCI-SSB-PeriodicityIndicationForScell-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(adaptSSBConfigR19PosInDCISSBPeriodicityIndicationForScellR19Constraints)
			if err != nil {
				return err
			}
			ie.PosInDCI_SSB_PeriodicityIndicationForScell_r19 = &v
		}
	}

	return nil
}
