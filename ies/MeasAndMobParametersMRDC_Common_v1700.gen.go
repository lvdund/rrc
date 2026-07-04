// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-Common-v1700 (line 21414).

var measAndMobParametersMRDCCommonV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "condPSCellChangeParameters-r17", Optional: true},
		{Name: "condHandoverWithSCG-ENDC-r17", Optional: true},
		{Name: "condHandoverWithSCG-NEDC-r17", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondHandoverWithSCG_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondHandoverWithSCGENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondHandoverWithSCG_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondHandoverWithSCG_NEDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondHandoverWithSCGNEDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondHandoverWithSCG_NEDC_r17_Supported},
}

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inter-SN-condPSCellChangeFDD-TDD-NRDC-r17", Optional: true},
		{Name: "inter-SN-condPSCellChangeFR1-FR2-NRDC-r17", Optional: true},
		{Name: "inter-SN-condPSCellChangeFDD-TDD-ENDC-r17", Optional: true},
		{Name: "inter-SN-condPSCellChangeFR1-FR2-ENDC-r17", Optional: true},
		{Name: "mn-InitiatedCondPSCellChange-FR1FDD-ENDC-r17", Optional: true},
		{Name: "mn-InitiatedCondPSCellChange-FR1TDD-ENDC-r17", Optional: true},
		{Name: "mn-InitiatedCondPSCellChange-FR2TDD-ENDC-r17", Optional: true},
		{Name: "sn-InitiatedCondPSCellChange-FR1FDD-ENDC-r17", Optional: true},
		{Name: "sn-InitiatedCondPSCellChange-FR1TDD-ENDC-r17", Optional: true},
		{Name: "sn-InitiatedCondPSCellChange-FR2TDD-ENDC-r17", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFDDTDDNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFR1FR2NRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFDDTDDENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFR1FR2ENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR1FDDENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR1TDDENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR2TDDENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR1FDDENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR1TDDENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17_Supported = 0
)

var measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR2TDDENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1700_CondPSCellChangeParameters_r17_Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17_Supported},
}

type MeasAndMobParametersMRDC_Common_v1700 struct {
	CondPSCellChangeParameters_r17 *struct {
		Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17    *int64
		Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17    *int64
		Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17    *int64
		Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17    *int64
		Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *int64
		Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *int64
		Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *int64
		Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *int64
		Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *int64
		Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *int64
	}
	CondHandoverWithSCG_ENDC_r17 *int64
	CondHandoverWithSCG_NEDC_r17 *int64
}

func (ie *MeasAndMobParametersMRDC_Common_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CondPSCellChangeParameters_r17 != nil, ie.CondHandoverWithSCG_ENDC_r17 != nil, ie.CondHandoverWithSCG_NEDC_r17 != nil}); err != nil {
		return err
	}

	// 2. condPSCellChangeParameters-r17: sequence
	{
		if ie.CondPSCellChangeParameters_r17 != nil {
			c := ie.CondPSCellChangeParameters_r17
			measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Constraints)
			if err := measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.EncodePreamble([]bool{c.Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17 != nil, c.Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17 != nil, c.Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17 != nil, c.Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17 != nil, c.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil, c.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil, c.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil, c.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil, c.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil, c.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil}); err != nil {
				return err
			}
			if c.Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFDDTDDNRDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFR1FR2NRDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFDDTDDENDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFR1FR2ENDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR1FDDENDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR1TDDENDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR2TDDENDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR1FDDENDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR1TDDENDCR17Constraints); err != nil {
					return err
				}
			}
			if c.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil {
				if err := e.EncodeEnumerated((*c.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17), measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR2TDDENDCR17Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. condHandoverWithSCG-ENDC-r17: enumerated
	{
		if ie.CondHandoverWithSCG_ENDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.CondHandoverWithSCG_ENDC_r17, measAndMobParametersMRDCCommonV1700CondHandoverWithSCGENDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. condHandoverWithSCG-NEDC-r17: enumerated
	{
		if ie.CondHandoverWithSCG_NEDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.CondHandoverWithSCG_NEDC_r17, measAndMobParametersMRDCCommonV1700CondHandoverWithSCGNEDCR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. condPSCellChangeParameters-r17: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.CondPSCellChangeParameters_r17 = &struct {
				Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17    *int64
				Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17    *int64
				Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17    *int64
				Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17    *int64
				Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *int64
				Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *int64
				Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *int64
				Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *int64
				Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *int64
				Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *int64
			}{}
			c := ie.CondPSCellChangeParameters_r17
			measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Constraints)
			if err := measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(0) {
				c.Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFDDTDDNRDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Inter_SN_CondPSCellChangeFDD_TDD_NRDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(1) {
				c.Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFR1FR2NRDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Inter_SN_CondPSCellChangeFR1_FR2_NRDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(2) {
				c.Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFDDTDDENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Inter_SN_CondPSCellChangeFDD_TDD_ENDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(3) {
				c.Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17InterSNCondPSCellChangeFR1FR2ENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Inter_SN_CondPSCellChangeFR1_FR2_ENDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(4) {
				c.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR1FDDENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(5) {
				c.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR1TDDENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(6) {
				c.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17MnInitiatedCondPSCellChangeFR2TDDENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(7) {
				c.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR1FDDENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(8) {
				c.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR1TDDENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17) = v
			}
			if measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17Seq.IsComponentPresent(9) {
				c.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondPSCellChangeParametersR17SnInitiatedCondPSCellChangeFR2TDDENDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17) = v
			}
		}
	}

	// 3. condHandoverWithSCG-ENDC-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondHandoverWithSCGENDCR17Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverWithSCG_ENDC_r17 = &idx
		}
	}

	// 4. condHandoverWithSCG-NEDC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1700CondHandoverWithSCGNEDCR17Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverWithSCG_NEDC_r17 = &idx
		}
	}

	return nil
}
