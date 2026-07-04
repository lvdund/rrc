// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-v1900-IEs (line 1050).

var rRCReconfigurationV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "n3c-ExtIndirectPathAddChange-r19", Optional: true},
		{Name: "otherConfig-v1900", Optional: true},
		{Name: "onDemandPosSIB-RequestCtrlParam-r19", Optional: true},
		{Name: "retainLoggedMeasurements-r19", Optional: true},
		{Name: "ltm-ConfigNRDC-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfiguration_v1900_IEsN3cExtIndirectPathAddChangeR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1900_IEs_N3c_ExtIndirectPathAddChange_r19_Release = 0
	RRCReconfiguration_v1900_IEs_N3c_ExtIndirectPathAddChange_r19_Setup   = 1
)

const (
	RRCReconfiguration_v1900_IEs_OnDemandPosSIB_RequestCtrlParam_r19_Enabled = 0
)

var rRCReconfigurationV1900IEsOnDemandPosSIBRequestCtrlParamR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfiguration_v1900_IEs_OnDemandPosSIB_RequestCtrlParam_r19_Enabled},
}

const (
	RRCReconfiguration_v1900_IEs_RetainLoggedMeasurements_r19_True = 0
)

var rRCReconfigurationV1900IEsRetainLoggedMeasurementsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfiguration_v1900_IEs_RetainLoggedMeasurements_r19_True},
}

var rRCReconfiguration_v1900_IEsLtmConfigNRDCR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1900_IEs_Ltm_ConfigNRDC_r19_Release = 0
	RRCReconfiguration_v1900_IEs_Ltm_ConfigNRDC_r19_Setup   = 1
)

type RRCReconfiguration_v1900_IEs struct {
	N3c_ExtIndirectPathAddChange_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *N3C_ExtIndirectPathAddChange_r19
	}
	OtherConfig_v1900                   *OtherConfig_v1900
	OnDemandPosSIB_RequestCtrlParam_r19 *int64
	RetainLoggedMeasurements_r19        *int64
	Ltm_ConfigNRDC_r19                  *struct {
		Choice  int
		Release *struct{}
		Setup   *LTM_ConfigNRDC_r19
	}
	NonCriticalExtension *struct{}
}

func (ie *RRCReconfiguration_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.N3c_ExtIndirectPathAddChange_r19 != nil, ie.OtherConfig_v1900 != nil, ie.OnDemandPosSIB_RequestCtrlParam_r19 != nil, ie.RetainLoggedMeasurements_r19 != nil, ie.Ltm_ConfigNRDC_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. n3c-ExtIndirectPathAddChange-r19: choice
	{
		if ie.N3c_ExtIndirectPathAddChange_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1900_IEsN3cExtIndirectPathAddChangeR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.N3c_ExtIndirectPathAddChange_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.N3c_ExtIndirectPathAddChange_r19).Choice {
			case RRCReconfiguration_v1900_IEs_N3c_ExtIndirectPathAddChange_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1900_IEs_N3c_ExtIndirectPathAddChange_r19_Setup:
				if err := (*ie.N3c_ExtIndirectPathAddChange_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.N3c_ExtIndirectPathAddChange_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. otherConfig-v1900: ref
	{
		if ie.OtherConfig_v1900 != nil {
			if err := ie.OtherConfig_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. onDemandPosSIB-RequestCtrlParam-r19: enumerated
	{
		if ie.OnDemandPosSIB_RequestCtrlParam_r19 != nil {
			if err := e.EncodeEnumerated(*ie.OnDemandPosSIB_RequestCtrlParam_r19, rRCReconfigurationV1900IEsOnDemandPosSIBRequestCtrlParamR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. retainLoggedMeasurements-r19: enumerated
	{
		if ie.RetainLoggedMeasurements_r19 != nil {
			if err := e.EncodeEnumerated(*ie.RetainLoggedMeasurements_r19, rRCReconfigurationV1900IEsRetainLoggedMeasurementsR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. ltm-ConfigNRDC-r19: choice
	{
		if ie.Ltm_ConfigNRDC_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1900_IEsLtmConfigNRDCR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ltm_ConfigNRDC_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ltm_ConfigNRDC_r19).Choice {
			case RRCReconfiguration_v1900_IEs_Ltm_ConfigNRDC_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1900_IEs_Ltm_ConfigNRDC_r19_Setup:
				if err := (*ie.Ltm_ConfigNRDC_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ltm_ConfigNRDC_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *RRCReconfiguration_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. n3c-ExtIndirectPathAddChange-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.N3c_ExtIndirectPathAddChange_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *N3C_ExtIndirectPathAddChange_r19
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1900_IEsN3cExtIndirectPathAddChangeR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.N3c_ExtIndirectPathAddChange_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1900_IEs_N3c_ExtIndirectPathAddChange_r19_Release:
				(*ie.N3c_ExtIndirectPathAddChange_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1900_IEs_N3c_ExtIndirectPathAddChange_r19_Setup:
				(*ie.N3c_ExtIndirectPathAddChange_r19).Setup = new(N3C_ExtIndirectPathAddChange_r19)
				if err := (*ie.N3c_ExtIndirectPathAddChange_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. otherConfig-v1900: ref
	{
		if seq.IsComponentPresent(1) {
			ie.OtherConfig_v1900 = new(OtherConfig_v1900)
			if err := ie.OtherConfig_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. onDemandPosSIB-RequestCtrlParam-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationV1900IEsOnDemandPosSIBRequestCtrlParamR19Constraints)
			if err != nil {
				return err
			}
			ie.OnDemandPosSIB_RequestCtrlParam_r19 = &idx
		}
	}

	// 5. retainLoggedMeasurements-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationV1900IEsRetainLoggedMeasurementsR19Constraints)
			if err != nil {
				return err
			}
			ie.RetainLoggedMeasurements_r19 = &idx
		}
	}

	// 6. ltm-ConfigNRDC-r19: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Ltm_ConfigNRDC_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTM_ConfigNRDC_r19
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1900_IEsLtmConfigNRDCR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ltm_ConfigNRDC_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1900_IEs_Ltm_ConfigNRDC_r19_Release:
				(*ie.Ltm_ConfigNRDC_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1900_IEs_Ltm_ConfigNRDC_r19_Setup:
				(*ie.Ltm_ConfigNRDC_r19).Setup = new(LTM_ConfigNRDC_r19)
				if err := (*ie.Ltm_ConfigNRDC_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
