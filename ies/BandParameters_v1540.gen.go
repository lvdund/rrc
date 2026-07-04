// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandParameters-v1540 (line 17041).

var bandParametersV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-CarrierSwitch", Optional: true},
		{Name: "srs-TxSwitch", Optional: true},
	},
}

var bandParameters_v1540SrsCarrierSwitchConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr"},
		{Name: "eutra"},
	},
}

const (
	BandParameters_v1540_Srs_CarrierSwitch_Nr    = 0
	BandParameters_v1540_Srs_CarrierSwitch_Eutra = 1
)

var bandParametersV1540SrsCarrierSwitchNrSrsSwitchingTimesListNRConstraints = per.SizeRange(1, common.MaxSimultaneousBands)

var bandParametersV1540SrsCarrierSwitchEutraSrsSwitchingTimesListEUTRAConstraints = per.SizeRange(1, common.MaxSimultaneousBands)

var bandParametersV1540SrsTxSwitchConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedSRS-TxPortSwitch"},
		{Name: "txSwitchImpactToRx", Optional: true},
		{Name: "txSwitchWithAnotherBand", Optional: true},
	},
}

const (
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r2         = 0
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r4         = 1
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T2r4         = 2
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r4_T2r4    = 3
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r1         = 4
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T2r2         = 5
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T4r4         = 6
	BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_NotSupported = 7
)

var bandParametersV1540SrsTxSwitchSupportedSRSTxPortSwitchConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r2, BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r4, BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T2r4, BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r4_T2r4, BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T1r1, BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T2r2, BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_T4r4, BandParameters_v1540_Srs_TxSwitch_SupportedSRS_TxPortSwitch_NotSupported},
}

type BandParameters_v1540 struct {
	Srs_CarrierSwitch *struct {
		Choice int
		Nr     *struct{ Srs_SwitchingTimesListNR []SRS_SwitchingTimeNR }
		Eutra  *struct{ Srs_SwitchingTimesListEUTRA []SRS_SwitchingTimeEUTRA }
	}
	Srs_TxSwitch *struct {
		SupportedSRS_TxPortSwitch int64
		TxSwitchImpactToRx        *int64
		TxSwitchWithAnotherBand   *int64
	}
}

func (ie *BandParameters_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_CarrierSwitch != nil, ie.Srs_TxSwitch != nil}); err != nil {
		return err
	}

	// 2. srs-CarrierSwitch: choice
	{
		if ie.Srs_CarrierSwitch != nil {
			choiceEnc := e.NewChoiceEncoder(bandParameters_v1540SrsCarrierSwitchConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_CarrierSwitch).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_CarrierSwitch).Choice {
			case BandParameters_v1540_Srs_CarrierSwitch_Nr:
				c := (*(*ie.Srs_CarrierSwitch).Nr)
				{
					seqOf := e.NewSequenceOfEncoder(bandParametersV1540SrsCarrierSwitchNrSrsSwitchingTimesListNRConstraints)
					if err := seqOf.EncodeLength(int64(len(c.Srs_SwitchingTimesListNR))); err != nil {
						return err
					}
					for i := range c.Srs_SwitchingTimesListNR {
						if err := c.Srs_SwitchingTimesListNR[i].Encode(e); err != nil {
							return err
						}
					}
				}
			case BandParameters_v1540_Srs_CarrierSwitch_Eutra:
				c := (*(*ie.Srs_CarrierSwitch).Eutra)
				{
					seqOf := e.NewSequenceOfEncoder(bandParametersV1540SrsCarrierSwitchEutraSrsSwitchingTimesListEUTRAConstraints)
					if err := seqOf.EncodeLength(int64(len(c.Srs_SwitchingTimesListEUTRA))); err != nil {
						return err
					}
					for i := range c.Srs_SwitchingTimesListEUTRA {
						if err := c.Srs_SwitchingTimesListEUTRA[i].Encode(e); err != nil {
							return err
						}
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_CarrierSwitch).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. srs-TxSwitch: sequence
	{
		if ie.Srs_TxSwitch != nil {
			c := ie.Srs_TxSwitch
			bandParametersV1540SrsTxSwitchSeq := e.NewSequenceEncoder(bandParametersV1540SrsTxSwitchConstraints)
			if err := bandParametersV1540SrsTxSwitchSeq.EncodePreamble([]bool{c.TxSwitchImpactToRx != nil, c.TxSwitchWithAnotherBand != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.SupportedSRS_TxPortSwitch, bandParametersV1540SrsTxSwitchSupportedSRSTxPortSwitchConstraints); err != nil {
				return err
			}
			if c.TxSwitchImpactToRx != nil {
				if err := e.EncodeInteger((*c.TxSwitchImpactToRx), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.TxSwitchWithAnotherBand != nil {
				if err := e.EncodeInteger((*c.TxSwitchWithAnotherBand), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-CarrierSwitch: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_CarrierSwitch = &struct {
				Choice int
				Nr     *struct{ Srs_SwitchingTimesListNR []SRS_SwitchingTimeNR }
				Eutra  *struct{ Srs_SwitchingTimesListEUTRA []SRS_SwitchingTimeEUTRA }
			}{}
			choiceDec := d.NewChoiceDecoder(bandParameters_v1540SrsCarrierSwitchConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_CarrierSwitch).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BandParameters_v1540_Srs_CarrierSwitch_Nr:
				(*ie.Srs_CarrierSwitch).Nr = &struct{ Srs_SwitchingTimesListNR []SRS_SwitchingTimeNR }{}
				c := (*(*ie.Srs_CarrierSwitch).Nr)
				{
					seqOf := d.NewSequenceOfDecoder(bandParametersV1540SrsCarrierSwitchNrSrsSwitchingTimesListNRConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Srs_SwitchingTimesListNR = make([]SRS_SwitchingTimeNR, n)
					for i := int64(0); i < n; i++ {
						if err := c.Srs_SwitchingTimesListNR[i].Decode(d); err != nil {
							return err
						}
					}
				}
			case BandParameters_v1540_Srs_CarrierSwitch_Eutra:
				(*ie.Srs_CarrierSwitch).Eutra = &struct{ Srs_SwitchingTimesListEUTRA []SRS_SwitchingTimeEUTRA }{}
				c := (*(*ie.Srs_CarrierSwitch).Eutra)
				{
					seqOf := d.NewSequenceOfDecoder(bandParametersV1540SrsCarrierSwitchEutraSrsSwitchingTimesListEUTRAConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Srs_SwitchingTimesListEUTRA = make([]SRS_SwitchingTimeEUTRA, n)
					for i := int64(0); i < n; i++ {
						if err := c.Srs_SwitchingTimesListEUTRA[i].Decode(d); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	// 3. srs-TxSwitch: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_TxSwitch = &struct {
				SupportedSRS_TxPortSwitch int64
				TxSwitchImpactToRx        *int64
				TxSwitchWithAnotherBand   *int64
			}{}
			c := ie.Srs_TxSwitch
			bandParametersV1540SrsTxSwitchSeq := d.NewSequenceDecoder(bandParametersV1540SrsTxSwitchConstraints)
			if err := bandParametersV1540SrsTxSwitchSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(bandParametersV1540SrsTxSwitchSupportedSRSTxPortSwitchConstraints)
				if err != nil {
					return err
				}
				c.SupportedSRS_TxPortSwitch = v
			}
			if bandParametersV1540SrsTxSwitchSeq.IsComponentPresent(1) {
				c.TxSwitchImpactToRx = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.TxSwitchImpactToRx) = v
			}
			if bandParametersV1540SrsTxSwitchSeq.IsComponentPresent(2) {
				c.TxSwitchWithAnotherBand = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.TxSwitchWithAnotherBand) = v
			}
		}
	}

	return nil
}
