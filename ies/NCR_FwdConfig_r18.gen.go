// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-FwdConfig-r18 (line 10304).

var nCRFwdConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicFwdRsrcSetToAddModList-r18", Optional: true},
		{Name: "periodicFwdRsrcSetToReleaseList-r18", Optional: true},
		{Name: "aperiodicFwdConfig-r18", Optional: true},
		{Name: "semiPersistentFwdRsrcSetToAddModList-r18", Optional: true},
		{Name: "semiPersistentFwdRsrcSetToReleaseList-r18", Optional: true},
	},
}

var nCRFwdConfigR18PeriodicFwdRsrcSetToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofPeriodicFwdResourceSet_r18)

var nCRFwdConfigR18PeriodicFwdRsrcSetToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofPeriodicFwdResourceSet_r18)

var nCR_FwdConfig_r18AperiodicFwdConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	NCR_FwdConfig_r18_AperiodicFwdConfig_r18_Release = 0
	NCR_FwdConfig_r18_AperiodicFwdConfig_r18_Setup   = 1
)

var nCRFwdConfigR18SemiPersistentFwdRsrcSetToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofSemiPersistentFwdResourceSet_r18)

var nCRFwdConfigR18SemiPersistentFwdRsrcSetToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofSemiPersistentFwdResourceSet_r18)

type NCR_FwdConfig_r18 struct {
	PeriodicFwdRsrcSetToAddModList_r18  []NCR_PeriodicFwdResourceSet_r18
	PeriodicFwdRsrcSetToReleaseList_r18 []NCR_PeriodicFwdResourceSetId_r18
	AperiodicFwdConfig_r18              *struct {
		Choice  int
		Release *struct{}
		Setup   *NCR_AperiodicFwdConfig_r18
	}
	SemiPersistentFwdRsrcSetToAddModList_r18  []NCR_SemiPersistentFwdResourceSet_r18
	SemiPersistentFwdRsrcSetToReleaseList_r18 []NCR_SemiPersistentFwdResourceSetId_r18
}

func (ie *NCR_FwdConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nCRFwdConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PeriodicFwdRsrcSetToAddModList_r18 != nil, ie.PeriodicFwdRsrcSetToReleaseList_r18 != nil, ie.AperiodicFwdConfig_r18 != nil, ie.SemiPersistentFwdRsrcSetToAddModList_r18 != nil, ie.SemiPersistentFwdRsrcSetToReleaseList_r18 != nil}); err != nil {
		return err
	}

	// 3. periodicFwdRsrcSetToAddModList-r18: sequence-of
	{
		if ie.PeriodicFwdRsrcSetToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRFwdConfigR18PeriodicFwdRsrcSetToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PeriodicFwdRsrcSetToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.PeriodicFwdRsrcSetToAddModList_r18 {
				if err := ie.PeriodicFwdRsrcSetToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. periodicFwdRsrcSetToReleaseList-r18: sequence-of
	{
		if ie.PeriodicFwdRsrcSetToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRFwdConfigR18PeriodicFwdRsrcSetToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PeriodicFwdRsrcSetToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.PeriodicFwdRsrcSetToReleaseList_r18 {
				if err := ie.PeriodicFwdRsrcSetToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. aperiodicFwdConfig-r18: choice
	{
		if ie.AperiodicFwdConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(nCR_FwdConfig_r18AperiodicFwdConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.AperiodicFwdConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.AperiodicFwdConfig_r18).Choice {
			case NCR_FwdConfig_r18_AperiodicFwdConfig_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case NCR_FwdConfig_r18_AperiodicFwdConfig_r18_Setup:
				if err := (*ie.AperiodicFwdConfig_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.AperiodicFwdConfig_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. semiPersistentFwdRsrcSetToAddModList-r18: sequence-of
	{
		if ie.SemiPersistentFwdRsrcSetToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRFwdConfigR18SemiPersistentFwdRsrcSetToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SemiPersistentFwdRsrcSetToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.SemiPersistentFwdRsrcSetToAddModList_r18 {
				if err := ie.SemiPersistentFwdRsrcSetToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. semiPersistentFwdRsrcSetToReleaseList-r18: sequence-of
	{
		if ie.SemiPersistentFwdRsrcSetToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRFwdConfigR18SemiPersistentFwdRsrcSetToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SemiPersistentFwdRsrcSetToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.SemiPersistentFwdRsrcSetToReleaseList_r18 {
				if err := ie.SemiPersistentFwdRsrcSetToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *NCR_FwdConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nCRFwdConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. periodicFwdRsrcSetToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(nCRFwdConfigR18PeriodicFwdRsrcSetToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PeriodicFwdRsrcSetToAddModList_r18 = make([]NCR_PeriodicFwdResourceSet_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PeriodicFwdRsrcSetToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. periodicFwdRsrcSetToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(nCRFwdConfigR18PeriodicFwdRsrcSetToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PeriodicFwdRsrcSetToReleaseList_r18 = make([]NCR_PeriodicFwdResourceSetId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PeriodicFwdRsrcSetToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. aperiodicFwdConfig-r18: choice
	{
		if seq.IsComponentPresent(2) {
			ie.AperiodicFwdConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NCR_AperiodicFwdConfig_r18
			}{}
			choiceDec := d.NewChoiceDecoder(nCR_FwdConfig_r18AperiodicFwdConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AperiodicFwdConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case NCR_FwdConfig_r18_AperiodicFwdConfig_r18_Release:
				(*ie.AperiodicFwdConfig_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case NCR_FwdConfig_r18_AperiodicFwdConfig_r18_Setup:
				(*ie.AperiodicFwdConfig_r18).Setup = new(NCR_AperiodicFwdConfig_r18)
				if err := (*ie.AperiodicFwdConfig_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. semiPersistentFwdRsrcSetToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(nCRFwdConfigR18SemiPersistentFwdRsrcSetToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SemiPersistentFwdRsrcSetToAddModList_r18 = make([]NCR_SemiPersistentFwdResourceSet_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SemiPersistentFwdRsrcSetToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. semiPersistentFwdRsrcSetToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(nCRFwdConfigR18SemiPersistentFwdRsrcSetToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SemiPersistentFwdRsrcSetToReleaseList_r18 = make([]NCR_SemiPersistentFwdResourceSetId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SemiPersistentFwdRsrcSetToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
