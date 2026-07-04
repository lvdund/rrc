// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-AperiodicFwdConfig-r18 (line 10280).

var nCRAperiodicFwdConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aperiodicFwdTimeRsrcToAddModList-r18", Optional: true},
		{Name: "aperiodicFwdTimeRsrcToReleaseList-r18", Optional: true},
		{Name: "referenceSCS-r18", Optional: true},
		{Name: "aperiodicBeamFieldWidth-r18", Optional: true},
		{Name: "numberOfFields-r18", Optional: true},
	},
}

var nCRAperiodicFwdConfigR18AperiodicFwdTimeRsrcToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofAperiodicFwdTimeResource_r18)

var nCRAperiodicFwdConfigR18AperiodicFwdTimeRsrcToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofAperiodicFwdTimeResource_r18)

var nCRAperiodicFwdConfigR18AperiodicBeamFieldWidthR18Constraints = per.Constrained(1, 6)

var nCRAperiodicFwdConfigR18NumberOfFieldsR18Constraints = per.Constrained(1, 32)

type NCR_AperiodicFwdConfig_r18 struct {
	AperiodicFwdTimeRsrcToAddModList_r18  []NCR_AperiodicFwdTimeResource_r18
	AperiodicFwdTimeRsrcToReleaseList_r18 []NCR_AperiodicFwdTimeResourceId_r18
	ReferenceSCS_r18                      *SubcarrierSpacing
	AperiodicBeamFieldWidth_r18           *int64
	NumberOfFields_r18                    *int64
}

func (ie *NCR_AperiodicFwdConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nCRAperiodicFwdConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AperiodicFwdTimeRsrcToAddModList_r18 != nil, ie.AperiodicFwdTimeRsrcToReleaseList_r18 != nil, ie.ReferenceSCS_r18 != nil, ie.AperiodicBeamFieldWidth_r18 != nil, ie.NumberOfFields_r18 != nil}); err != nil {
		return err
	}

	// 3. aperiodicFwdTimeRsrcToAddModList-r18: sequence-of
	{
		if ie.AperiodicFwdTimeRsrcToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRAperiodicFwdConfigR18AperiodicFwdTimeRsrcToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AperiodicFwdTimeRsrcToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.AperiodicFwdTimeRsrcToAddModList_r18 {
				if err := ie.AperiodicFwdTimeRsrcToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. aperiodicFwdTimeRsrcToReleaseList-r18: sequence-of
	{
		if ie.AperiodicFwdTimeRsrcToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRAperiodicFwdConfigR18AperiodicFwdTimeRsrcToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AperiodicFwdTimeRsrcToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.AperiodicFwdTimeRsrcToReleaseList_r18 {
				if err := ie.AperiodicFwdTimeRsrcToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. referenceSCS-r18: ref
	{
		if ie.ReferenceSCS_r18 != nil {
			if err := ie.ReferenceSCS_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. aperiodicBeamFieldWidth-r18: integer
	{
		if ie.AperiodicBeamFieldWidth_r18 != nil {
			if err := e.EncodeInteger(*ie.AperiodicBeamFieldWidth_r18, nCRAperiodicFwdConfigR18AperiodicBeamFieldWidthR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. numberOfFields-r18: integer
	{
		if ie.NumberOfFields_r18 != nil {
			if err := e.EncodeInteger(*ie.NumberOfFields_r18, nCRAperiodicFwdConfigR18NumberOfFieldsR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NCR_AperiodicFwdConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nCRAperiodicFwdConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. aperiodicFwdTimeRsrcToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(nCRAperiodicFwdConfigR18AperiodicFwdTimeRsrcToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AperiodicFwdTimeRsrcToAddModList_r18 = make([]NCR_AperiodicFwdTimeResource_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AperiodicFwdTimeRsrcToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. aperiodicFwdTimeRsrcToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(nCRAperiodicFwdConfigR18AperiodicFwdTimeRsrcToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AperiodicFwdTimeRsrcToReleaseList_r18 = make([]NCR_AperiodicFwdTimeResourceId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AperiodicFwdTimeRsrcToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. referenceSCS-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ReferenceSCS_r18 = new(SubcarrierSpacing)
			if err := ie.ReferenceSCS_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. aperiodicBeamFieldWidth-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(nCRAperiodicFwdConfigR18AperiodicBeamFieldWidthR18Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicBeamFieldWidth_r18 = &v
		}
	}

	// 7. numberOfFields-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(nCRAperiodicFwdConfigR18NumberOfFieldsR18Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfFields_r18 = &v
		}
	}

	return nil
}
