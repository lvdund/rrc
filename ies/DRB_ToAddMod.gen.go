// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRB-ToAddMod (line 13153).

var dRBToAddModConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cnAssociation", Optional: true},
		{Name: "drb-Identity"},
		{Name: "reestablishPDCP", Optional: true},
		{Name: "recoverPDCP", Optional: true},
		{Name: "pdcp-Config", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var dRB_ToAddModCnAssociationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eps-BearerIdentity"},
		{Name: "sdap-Config"},
	},
}

const (
	DRB_ToAddMod_CnAssociation_Eps_BearerIdentity = 0
	DRB_ToAddMod_CnAssociation_Sdap_Config        = 1
)

const (
	DRB_ToAddMod_ReestablishPDCP_True = 0
)

var dRBToAddModReestablishPDCPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRB_ToAddMod_ReestablishPDCP_True},
}

const (
	DRB_ToAddMod_RecoverPDCP_True = 0
)

var dRBToAddModRecoverPDCPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRB_ToAddMod_RecoverPDCP_True},
}

const (
	DRB_ToAddMod_Ext_Daps_Config_r16_True = 0
)

var dRBToAddModExtDapsConfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRB_ToAddMod_Ext_Daps_Config_r16_True},
}

const (
	DRB_ToAddMod_Ext_N3c_BearerAssociated_r18_True = 0
)

var dRBToAddModExtN3cBearerAssociatedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRB_ToAddMod_Ext_N3c_BearerAssociated_r18_True},
}

type DRB_ToAddMod struct {
	CnAssociation *struct {
		Choice             int
		Eps_BearerIdentity *int64
		Sdap_Config        *SDAP_Config
	}
	Drb_Identity             DRB_Identity
	ReestablishPDCP          *int64
	RecoverPDCP              *int64
	Pdcp_Config              *PDCP_Config
	Daps_Config_r16          *int64
	N3c_BearerAssociated_r18 *int64
}

func (ie *DRB_ToAddMod) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBToAddModConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Daps_Config_r16 != nil
	hasExtGroup1 := ie.N3c_BearerAssociated_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CnAssociation != nil, ie.ReestablishPDCP != nil, ie.RecoverPDCP != nil, ie.Pdcp_Config != nil}); err != nil {
		return err
	}

	// 3. cnAssociation: choice
	{
		if ie.CnAssociation != nil {
			choiceEnc := e.NewChoiceEncoder(dRB_ToAddModCnAssociationConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CnAssociation).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CnAssociation).Choice {
			case DRB_ToAddMod_CnAssociation_Eps_BearerIdentity:
				if err := e.EncodeInteger((*(*ie.CnAssociation).Eps_BearerIdentity), per.Constrained(0, 15)); err != nil {
					return err
				}
			case DRB_ToAddMod_CnAssociation_Sdap_Config:
				if err := (*ie.CnAssociation).Sdap_Config.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.CnAssociation).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. drb-Identity: ref
	{
		if err := ie.Drb_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 5. reestablishPDCP: enumerated
	{
		if ie.ReestablishPDCP != nil {
			if err := e.EncodeEnumerated(*ie.ReestablishPDCP, dRBToAddModReestablishPDCPConstraints); err != nil {
				return err
			}
		}
	}

	// 6. recoverPDCP: enumerated
	{
		if ie.RecoverPDCP != nil {
			if err := e.EncodeEnumerated(*ie.RecoverPDCP, dRBToAddModRecoverPDCPConstraints); err != nil {
				return err
			}
		}
	}

	// 7. pdcp-Config: ref
	{
		if ie.Pdcp_Config != nil {
			if err := ie.Pdcp_Config.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "daps-Config-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Daps_Config_r16 != nil}); err != nil {
				return err
			}

			if ie.Daps_Config_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Daps_Config_r16, dRBToAddModExtDapsConfigR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "n3c-BearerAssociated-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.N3c_BearerAssociated_r18 != nil}); err != nil {
				return err
			}

			if ie.N3c_BearerAssociated_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.N3c_BearerAssociated_r18, dRBToAddModExtN3cBearerAssociatedR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRB_ToAddMod) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBToAddModConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cnAssociation: choice
	{
		if seq.IsComponentPresent(0) {
			ie.CnAssociation = &struct {
				Choice             int
				Eps_BearerIdentity *int64
				Sdap_Config        *SDAP_Config
			}{}
			choiceDec := d.NewChoiceDecoder(dRB_ToAddModCnAssociationConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CnAssociation).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case DRB_ToAddMod_CnAssociation_Eps_BearerIdentity:
				(*ie.CnAssociation).Eps_BearerIdentity = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				(*(*ie.CnAssociation).Eps_BearerIdentity) = v
			case DRB_ToAddMod_CnAssociation_Sdap_Config:
				(*ie.CnAssociation).Sdap_Config = new(SDAP_Config)
				if err := (*ie.CnAssociation).Sdap_Config.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. drb-Identity: ref
	{
		if err := ie.Drb_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 5. reestablishPDCP: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dRBToAddModReestablishPDCPConstraints)
			if err != nil {
				return err
			}
			ie.ReestablishPDCP = &idx
		}
	}

	// 6. recoverPDCP: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dRBToAddModRecoverPDCPConstraints)
			if err != nil {
				return err
			}
			ie.RecoverPDCP = &idx
		}
	}

	// 7. pdcp-Config: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Pdcp_Config = new(PDCP_Config)
			if err := ie.Pdcp_Config.Decode(d); err != nil {
				return err
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "daps-Config-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(dRBToAddModExtDapsConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.Daps_Config_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "n3c-BearerAssociated-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(dRBToAddModExtN3cBearerAssociatedR18Constraints)
			if err != nil {
				return err
			}
			ie.N3c_BearerAssociated_r18 = &v
		}
	}

	return nil
}
