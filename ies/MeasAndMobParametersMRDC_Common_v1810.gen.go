// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-Common-v1810 (line 21439).

var measAndMobParametersMRDCCommonV1810Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mn-ConfiguredMN-TriggerSCPAC-r18", Optional: true},
		{Name: "mn-ConfiguredSN-TriggerSCPAC-r18", Optional: true},
		{Name: "sn-ConfiguredSCPAC-r18", Optional: true},
		{Name: "mn-ConfiguredMN-TriggerSCPAC-afterSCG-release-r18", Optional: true},
		{Name: "mn-ConfiguredReferenceConfigSCPAC-r18", Optional: true},
		{Name: "sn-ConfiguredReferenceConfigSCPAC-r18", Optional: true},
		{Name: "condHandoverWithCandSCG-Addition-r18", Optional: true},
		{Name: "condHandoverWithCandSCG-FR1-FR2-Change-r18", Optional: true},
		{Name: "condHandoverWithCandSCG-FDD-TDD-Change-r18", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredMN_TriggerSCPAC_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810MnConfiguredMNTriggerSCPACR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredMN_TriggerSCPAC_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredSN_TriggerSCPAC_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810MnConfiguredSNTriggerSCPACR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredSN_TriggerSCPAC_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_Sn_ConfiguredSCPAC_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810SnConfiguredSCPACR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_Sn_ConfiguredSCPAC_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredMN_TriggerSCPAC_AfterSCG_Release_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810MnConfiguredMNTriggerSCPACAfterSCGReleaseR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredMN_TriggerSCPAC_AfterSCG_Release_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredReferenceConfigSCPAC_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810MnConfiguredReferenceConfigSCPACR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_Mn_ConfiguredReferenceConfigSCPAC_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_Sn_ConfiguredReferenceConfigSCPAC_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810SnConfiguredReferenceConfigSCPACR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_Sn_ConfiguredReferenceConfigSCPAC_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_CondHandoverWithCandSCG_Addition_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGAdditionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_CondHandoverWithCandSCG_Addition_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_CondHandoverWithCandSCG_FR1_FR2_Change_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGFR1FR2ChangeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_CondHandoverWithCandSCG_FR1_FR2_Change_r18_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1810_CondHandoverWithCandSCG_FDD_TDD_Change_r18_Supported = 0
)

var measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGFDDTDDChangeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1810_CondHandoverWithCandSCG_FDD_TDD_Change_r18_Supported},
}

type MeasAndMobParametersMRDC_Common_v1810 struct {
	Mn_ConfiguredMN_TriggerSCPAC_r18                  *int64
	Mn_ConfiguredSN_TriggerSCPAC_r18                  *int64
	Sn_ConfiguredSCPAC_r18                            *int64
	Mn_ConfiguredMN_TriggerSCPAC_AfterSCG_Release_r18 *int64
	Mn_ConfiguredReferenceConfigSCPAC_r18             *int64
	Sn_ConfiguredReferenceConfigSCPAC_r18             *int64
	CondHandoverWithCandSCG_Addition_r18              *int64
	CondHandoverWithCandSCG_FR1_FR2_Change_r18        *int64
	CondHandoverWithCandSCG_FDD_TDD_Change_r18        *int64
}

func (ie *MeasAndMobParametersMRDC_Common_v1810) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1810Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mn_ConfiguredMN_TriggerSCPAC_r18 != nil, ie.Mn_ConfiguredSN_TriggerSCPAC_r18 != nil, ie.Sn_ConfiguredSCPAC_r18 != nil, ie.Mn_ConfiguredMN_TriggerSCPAC_AfterSCG_Release_r18 != nil, ie.Mn_ConfiguredReferenceConfigSCPAC_r18 != nil, ie.Sn_ConfiguredReferenceConfigSCPAC_r18 != nil, ie.CondHandoverWithCandSCG_Addition_r18 != nil, ie.CondHandoverWithCandSCG_FR1_FR2_Change_r18 != nil, ie.CondHandoverWithCandSCG_FDD_TDD_Change_r18 != nil}); err != nil {
		return err
	}

	// 2. mn-ConfiguredMN-TriggerSCPAC-r18: enumerated
	{
		if ie.Mn_ConfiguredMN_TriggerSCPAC_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Mn_ConfiguredMN_TriggerSCPAC_r18, measAndMobParametersMRDCCommonV1810MnConfiguredMNTriggerSCPACR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. mn-ConfiguredSN-TriggerSCPAC-r18: enumerated
	{
		if ie.Mn_ConfiguredSN_TriggerSCPAC_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Mn_ConfiguredSN_TriggerSCPAC_r18, measAndMobParametersMRDCCommonV1810MnConfiguredSNTriggerSCPACR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sn-ConfiguredSCPAC-r18: enumerated
	{
		if ie.Sn_ConfiguredSCPAC_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sn_ConfiguredSCPAC_r18, measAndMobParametersMRDCCommonV1810SnConfiguredSCPACR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. mn-ConfiguredMN-TriggerSCPAC-afterSCG-release-r18: enumerated
	{
		if ie.Mn_ConfiguredMN_TriggerSCPAC_AfterSCG_Release_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Mn_ConfiguredMN_TriggerSCPAC_AfterSCG_Release_r18, measAndMobParametersMRDCCommonV1810MnConfiguredMNTriggerSCPACAfterSCGReleaseR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. mn-ConfiguredReferenceConfigSCPAC-r18: enumerated
	{
		if ie.Mn_ConfiguredReferenceConfigSCPAC_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Mn_ConfiguredReferenceConfigSCPAC_r18, measAndMobParametersMRDCCommonV1810MnConfiguredReferenceConfigSCPACR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sn-ConfiguredReferenceConfigSCPAC-r18: enumerated
	{
		if ie.Sn_ConfiguredReferenceConfigSCPAC_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sn_ConfiguredReferenceConfigSCPAC_r18, measAndMobParametersMRDCCommonV1810SnConfiguredReferenceConfigSCPACR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. condHandoverWithCandSCG-Addition-r18: enumerated
	{
		if ie.CondHandoverWithCandSCG_Addition_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CondHandoverWithCandSCG_Addition_r18, measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGAdditionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. condHandoverWithCandSCG-FR1-FR2-Change-r18: enumerated
	{
		if ie.CondHandoverWithCandSCG_FR1_FR2_Change_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CondHandoverWithCandSCG_FR1_FR2_Change_r18, measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGFR1FR2ChangeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. condHandoverWithCandSCG-FDD-TDD-Change-r18: enumerated
	{
		if ie.CondHandoverWithCandSCG_FDD_TDD_Change_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CondHandoverWithCandSCG_FDD_TDD_Change_r18, measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGFDDTDDChangeR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1810) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1810Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mn-ConfiguredMN-TriggerSCPAC-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810MnConfiguredMNTriggerSCPACR18Constraints)
			if err != nil {
				return err
			}
			ie.Mn_ConfiguredMN_TriggerSCPAC_r18 = &idx
		}
	}

	// 3. mn-ConfiguredSN-TriggerSCPAC-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810MnConfiguredSNTriggerSCPACR18Constraints)
			if err != nil {
				return err
			}
			ie.Mn_ConfiguredSN_TriggerSCPAC_r18 = &idx
		}
	}

	// 4. sn-ConfiguredSCPAC-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810SnConfiguredSCPACR18Constraints)
			if err != nil {
				return err
			}
			ie.Sn_ConfiguredSCPAC_r18 = &idx
		}
	}

	// 5. mn-ConfiguredMN-TriggerSCPAC-afterSCG-release-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810MnConfiguredMNTriggerSCPACAfterSCGReleaseR18Constraints)
			if err != nil {
				return err
			}
			ie.Mn_ConfiguredMN_TriggerSCPAC_AfterSCG_Release_r18 = &idx
		}
	}

	// 6. mn-ConfiguredReferenceConfigSCPAC-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810MnConfiguredReferenceConfigSCPACR18Constraints)
			if err != nil {
				return err
			}
			ie.Mn_ConfiguredReferenceConfigSCPAC_r18 = &idx
		}
	}

	// 7. sn-ConfiguredReferenceConfigSCPAC-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810SnConfiguredReferenceConfigSCPACR18Constraints)
			if err != nil {
				return err
			}
			ie.Sn_ConfiguredReferenceConfigSCPAC_r18 = &idx
		}
	}

	// 8. condHandoverWithCandSCG-Addition-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGAdditionR18Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverWithCandSCG_Addition_r18 = &idx
		}
	}

	// 9. condHandoverWithCandSCG-FR1-FR2-Change-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGFR1FR2ChangeR18Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverWithCandSCG_FR1_FR2_Change_r18 = &idx
		}
	}

	// 10. condHandoverWithCandSCG-FDD-TDD-Change-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1810CondHandoverWithCandSCGFDDTDDChangeR18Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverWithCandSCG_FDD_TDD_Change_r18 = &idx
		}
	}

	return nil
}
