// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParametersSidelinkEUTRA-NR-v1710 (line 17170).

var bandParametersSidelinkEUTRANRV1710Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eutra"},
		{Name: "nr"},
	},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Eutra = 0
	BandParametersSidelinkEUTRA_NR_v1710_Nr    = 1
)

var bandParametersSidelinkEUTRANRV1710NrConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TransmissionMode2-PartialSensing-r17", Optional: true},
		{Name: "rx-sidelinkPSFCH-r17", Optional: true},
		{Name: "tx-IUC-Scheme1-Mode2Sidelink-r17", Optional: true},
		{Name: "tx-IUC-Scheme2-Mode2Sidelink-r17", Optional: true},
	},
}

var bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-TxProcessModeTwoSidelink-r17"},
		{Name: "scs-CP-PatternTxSidelinkModeTwo-r17", Optional: true},
		{Name: "extendedCP-Mode2PartialSensing-r17", Optional: true},
		{Name: "dl-openLoopPC-Sidelink-r17", Optional: true},
	},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N8  = 0
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N16 = 1
)

var bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17HarqTxProcessModeTwoSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N8, BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N16},
}

var bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r17"},
		{Name: "fr2-r17"},
	},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17 = 0
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17 = 1
)

var bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r17", Optional: true},
		{Name: "scs-30kHz-r17", Optional: true},
		{Name: "scs-60kHz-r17", Optional: true},
	},
}

var bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r17", Optional: true},
		{Name: "scs-120kHz-r17", Optional: true},
	},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_ExtendedCP_Mode2PartialSensing_r17_Supported = 0
)

var bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ExtendedCPMode2PartialSensingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_ExtendedCP_Mode2PartialSensing_r17_Supported},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Dl_OpenLoopPC_Sidelink_r17_Supported = 0
)

var bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17DlOpenLoopPCSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Dl_OpenLoopPC_Sidelink_r17_Supported},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N5  = 0
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N15 = 1
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N25 = 2
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N32 = 3
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N35 = 4
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N45 = 5
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N50 = 6
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N64 = 7
)

var bandParametersSidelinkEUTRANRV1710NrRxSidelinkPSFCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N5, BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N15, BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N25, BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N32, BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N35, BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N45, BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N50, BandParametersSidelinkEUTRA_NR_v1710_Nr_Rx_SidelinkPSFCH_r17_N64},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme1_Mode2Sidelink_r17_Supported = 0
)

var bandParametersSidelinkEUTRANRV1710NrTxIUCScheme1Mode2SidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme1_Mode2Sidelink_r17_Supported},
}

const (
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme2_Mode2Sidelink_r17_N4  = 0
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme2_Mode2Sidelink_r17_N8  = 1
	BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme2_Mode2Sidelink_r17_N16 = 2
)

var bandParametersSidelinkEUTRANRV1710NrTxIUCScheme2Mode2SidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme2_Mode2Sidelink_r17_N4, BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme2_Mode2Sidelink_r17_N8, BandParametersSidelinkEUTRA_NR_v1710_Nr_Tx_IUC_Scheme2_Mode2Sidelink_r17_N16},
}

type BandParametersSidelinkEUTRA_NR_v1710 struct {
	Choice int
	Eutra  *struct{}
	Nr     *struct {
		Sl_TransmissionMode2_PartialSensing_r17 *struct {
			Harq_TxProcessModeTwoSidelink_r17   int64
			Scs_CP_PatternTxSidelinkModeTwo_r17 *struct {
				Choice  int
				Fr1_r17 *struct {
					Scs_15kHz_r17 *per.BitString
					Scs_30kHz_r17 *per.BitString
					Scs_60kHz_r17 *per.BitString
				}
				Fr2_r17 *struct {
					Scs_60kHz_r17  *per.BitString
					Scs_120kHz_r17 *per.BitString
				}
			}
			ExtendedCP_Mode2PartialSensing_r17 *int64
			Dl_OpenLoopPC_Sidelink_r17         *int64
		}
		Rx_SidelinkPSFCH_r17             *int64
		Tx_IUC_Scheme1_Mode2Sidelink_r17 *int64
		Tx_IUC_Scheme2_Mode2Sidelink_r17 *int64
	}
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(bandParametersSidelinkEUTRANRV1710Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_v1710_Eutra:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case BandParametersSidelinkEUTRA_NR_v1710_Nr:
		c := (*ie.Nr)
		bandParametersSidelinkEUTRANRV1710NrSeq := e.NewSequenceEncoder(bandParametersSidelinkEUTRANRV1710NrConstraints)
		if err := bandParametersSidelinkEUTRANRV1710NrSeq.EncodePreamble([]bool{c.Sl_TransmissionMode2_PartialSensing_r17 != nil, c.Rx_SidelinkPSFCH_r17 != nil, c.Tx_IUC_Scheme1_Mode2Sidelink_r17 != nil, c.Tx_IUC_Scheme2_Mode2Sidelink_r17 != nil}); err != nil {
			return err
		}
		if c.Sl_TransmissionMode2_PartialSensing_r17 != nil {
			c := (*c.Sl_TransmissionMode2_PartialSensing_r17)
			bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Seq := e.NewSequenceEncoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Constraints)
			if err := bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Seq.EncodePreamble([]bool{c.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil, c.ExtendedCP_Mode2PartialSensing_r17 != nil, c.Dl_OpenLoopPC_Sidelink_r17 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Harq_TxProcessModeTwoSidelink_r17, bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17HarqTxProcessModeTwoSidelinkR17Constraints); err != nil {
				return err
			}
			if c.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil {
				choiceEnc := e.NewChoiceEncoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice {
				case BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17:
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17)
					bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq := e.NewSequenceEncoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints)
					if err := bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.EncodePreamble([]bool{c.Scs_15kHz_r17 != nil, c.Scs_30kHz_r17 != nil, c.Scs_60kHz_r17 != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz_r17 != nil {
						if err := e.EncodeBitString((*c.Scs_15kHz_r17), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_30kHz_r17 != nil {
						if err := e.EncodeBitString((*c.Scs_30kHz_r17), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_60kHz_r17 != nil {
						if err := e.EncodeBitString((*c.Scs_60kHz_r17), per.FixedSize(16)); err != nil {
							return err
						}
					}
				case BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17:
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17)
					bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq := e.NewSequenceEncoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints)
					if err := bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.EncodePreamble([]bool{c.Scs_60kHz_r17 != nil, c.Scs_120kHz_r17 != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz_r17 != nil {
						if err := e.EncodeBitString((*c.Scs_60kHz_r17), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_120kHz_r17 != nil {
						if err := e.EncodeBitString((*c.Scs_120kHz_r17), per.FixedSize(16)); err != nil {
							return err
						}
					}
				}
			}
			if c.ExtendedCP_Mode2PartialSensing_r17 != nil {
				if err := e.EncodeEnumerated((*c.ExtendedCP_Mode2PartialSensing_r17), bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ExtendedCPMode2PartialSensingR17Constraints); err != nil {
					return err
				}
			}
			if c.Dl_OpenLoopPC_Sidelink_r17 != nil {
				if err := e.EncodeEnumerated((*c.Dl_OpenLoopPC_Sidelink_r17), bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17DlOpenLoopPCSidelinkR17Constraints); err != nil {
					return err
				}
			}
		}
		if c.Rx_SidelinkPSFCH_r17 != nil {
			if err := e.EncodeEnumerated((*c.Rx_SidelinkPSFCH_r17), bandParametersSidelinkEUTRANRV1710NrRxSidelinkPSFCHR17Constraints); err != nil {
				return err
			}
		}
		if c.Tx_IUC_Scheme1_Mode2Sidelink_r17 != nil {
			if err := e.EncodeEnumerated((*c.Tx_IUC_Scheme1_Mode2Sidelink_r17), bandParametersSidelinkEUTRANRV1710NrTxIUCScheme1Mode2SidelinkR17Constraints); err != nil {
				return err
			}
		}
		if c.Tx_IUC_Scheme2_Mode2Sidelink_r17 != nil {
			if err := e.EncodeEnumerated((*c.Tx_IUC_Scheme2_Mode2Sidelink_r17), bandParametersSidelinkEUTRANRV1710NrTxIUCScheme2Mode2SidelinkR17Constraints); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BandParametersSidelinkEUTRA-NR-v1710",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(bandParametersSidelinkEUTRANRV1710Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BandParametersSidelinkEUTRA-NR-v1710", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BandParametersSidelinkEUTRA_NR_v1710_Eutra:
		ie.Eutra = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case BandParametersSidelinkEUTRA_NR_v1710_Nr:
		ie.Nr = &struct {
			Sl_TransmissionMode2_PartialSensing_r17 *struct {
				Harq_TxProcessModeTwoSidelink_r17   int64
				Scs_CP_PatternTxSidelinkModeTwo_r17 *struct {
					Choice  int
					Fr1_r17 *struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}
					Fr2_r17 *struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}
				}
				ExtendedCP_Mode2PartialSensing_r17 *int64
				Dl_OpenLoopPC_Sidelink_r17         *int64
			}
			Rx_SidelinkPSFCH_r17             *int64
			Tx_IUC_Scheme1_Mode2Sidelink_r17 *int64
			Tx_IUC_Scheme2_Mode2Sidelink_r17 *int64
		}{}
		c := (*ie.Nr)
		bandParametersSidelinkEUTRANRV1710NrSeq := d.NewSequenceDecoder(bandParametersSidelinkEUTRANRV1710NrConstraints)
		if err := bandParametersSidelinkEUTRANRV1710NrSeq.DecodePreamble(); err != nil {
			return err
		}
		if bandParametersSidelinkEUTRANRV1710NrSeq.IsComponentPresent(0) {
			c.Sl_TransmissionMode2_PartialSensing_r17 = &struct {
				Harq_TxProcessModeTwoSidelink_r17   int64
				Scs_CP_PatternTxSidelinkModeTwo_r17 *struct {
					Choice  int
					Fr1_r17 *struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}
					Fr2_r17 *struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}
				}
				ExtendedCP_Mode2PartialSensing_r17 *int64
				Dl_OpenLoopPC_Sidelink_r17         *int64
			}{}
			c := (*c.Sl_TransmissionMode2_PartialSensing_r17)
			bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Seq := d.NewSequenceDecoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Constraints)
			if err := bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17HarqTxProcessModeTwoSidelinkR17Constraints)
				if err != nil {
					return err
				}
				c.Harq_TxProcessModeTwoSidelink_r17 = v
			}
			if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Seq.IsComponentPresent(1) {
				c.Scs_CP_PatternTxSidelinkModeTwo_r17 = &struct {
					Choice  int
					Fr1_r17 *struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}
					Fr2_r17 *struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}
				}{}
				choiceDec := d.NewChoiceDecoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice = int(index)
				switch index {
				case BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17:
					(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17 = &struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17)
					bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq := d.NewSequenceDecoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints)
					if err := bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r17) = v
					}
					if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r17) = v
					}
					if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r17) = v
					}
				case BandParametersSidelinkEUTRA_NR_v1710_Nr_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17:
					(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17 = &struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17)
					bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq := d.NewSequenceDecoder(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints)
					if err := bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r17) = v
					}
					if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r17) = v
					}
				}
			}
			if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Seq.IsComponentPresent(2) {
				c.ExtendedCP_Mode2PartialSensing_r17 = new(int64)
				v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17ExtendedCPMode2PartialSensingR17Constraints)
				if err != nil {
					return err
				}
				(*c.ExtendedCP_Mode2PartialSensing_r17) = v
			}
			if bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17Seq.IsComponentPresent(3) {
				c.Dl_OpenLoopPC_Sidelink_r17 = new(int64)
				v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1710NrSlTransmissionMode2PartialSensingR17DlOpenLoopPCSidelinkR17Constraints)
				if err != nil {
					return err
				}
				(*c.Dl_OpenLoopPC_Sidelink_r17) = v
			}
		}
		if bandParametersSidelinkEUTRANRV1710NrSeq.IsComponentPresent(1) {
			c.Rx_SidelinkPSFCH_r17 = new(int64)
			v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1710NrRxSidelinkPSFCHR17Constraints)
			if err != nil {
				return err
			}
			(*c.Rx_SidelinkPSFCH_r17) = v
		}
		if bandParametersSidelinkEUTRANRV1710NrSeq.IsComponentPresent(2) {
			c.Tx_IUC_Scheme1_Mode2Sidelink_r17 = new(int64)
			v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1710NrTxIUCScheme1Mode2SidelinkR17Constraints)
			if err != nil {
				return err
			}
			(*c.Tx_IUC_Scheme1_Mode2Sidelink_r17) = v
		}
		if bandParametersSidelinkEUTRANRV1710NrSeq.IsComponentPresent(3) {
			c.Tx_IUC_Scheme2_Mode2Sidelink_r17 = new(int64)
			v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1710NrTxIUCScheme2Mode2SidelinkR17Constraints)
			if err != nil {
				return err
			}
			(*c.Tx_IUC_Scheme2_Mode2Sidelink_r17) = v
		}
	default:
		return &per.DecodeError{TypeName: "BandParametersSidelinkEUTRA-NR-v1710", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
