// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParametersSidelinkDiscovery-r17 (line 17208).

var bandParametersSidelinkDiscoveryR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-CrossCarrierScheduling-r17", Optional: true},
		{Name: "sl-TransmissionMode2-PartialSensing-r17", Optional: true},
		{Name: "tx-IUC-Scheme1-Mode2Sidelink-r17", Optional: true},
	},
}

const (
	BandParametersSidelinkDiscovery_r17_Sl_CrossCarrierScheduling_r17_Supported = 0
)

var bandParametersSidelinkDiscoveryR17SlCrossCarrierSchedulingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkDiscovery_r17_Sl_CrossCarrierScheduling_r17_Supported},
}

const (
	BandParametersSidelinkDiscovery_r17_Tx_IUC_Scheme1_Mode2Sidelink_r17_Supported = 0
)

var bandParametersSidelinkDiscoveryR17TxIUCScheme1Mode2SidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkDiscovery_r17_Tx_IUC_Scheme1_Mode2Sidelink_r17_Supported},
}

var bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-TxProcessModeTwoSidelink-r17"},
		{Name: "scs-CP-PatternTxSidelinkModeTwo-r17", Optional: true},
		{Name: "extendedCP-Mode2PartialSensing-r17", Optional: true},
		{Name: "dl-openLoopPC-Sidelink-r17", Optional: true},
	},
}

const (
	BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N8  = 0
	BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N16 = 1
)

var bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17HarqTxProcessModeTwoSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N8, BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Harq_TxProcessModeTwoSidelink_r17_N16},
}

var bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r17"},
		{Name: "fr2-r17"},
	},
}

const (
	BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17 = 0
	BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17 = 1
)

var bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r17", Optional: true},
		{Name: "scs-30kHz-r17", Optional: true},
		{Name: "scs-60kHz-r17", Optional: true},
	},
}

var bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r17", Optional: true},
		{Name: "scs-120kHz-r17", Optional: true},
	},
}

const (
	BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_ExtendedCP_Mode2PartialSensing_r17_Supported = 0
)

var bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ExtendedCPMode2PartialSensingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_ExtendedCP_Mode2PartialSensing_r17_Supported},
}

const (
	BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Dl_OpenLoopPC_Sidelink_r17_Supported = 0
)

var bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17DlOpenLoopPCSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Dl_OpenLoopPC_Sidelink_r17_Supported},
}

type BandParametersSidelinkDiscovery_r17 struct {
	Sl_CrossCarrierScheduling_r17           *int64
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
	Tx_IUC_Scheme1_Mode2Sidelink_r17 *int64
}

func (ie *BandParametersSidelinkDiscovery_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersSidelinkDiscoveryR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_CrossCarrierScheduling_r17 != nil, ie.Sl_TransmissionMode2_PartialSensing_r17 != nil, ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 != nil}); err != nil {
		return err
	}

	// 2. sl-CrossCarrierScheduling-r17: enumerated
	{
		if ie.Sl_CrossCarrierScheduling_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_CrossCarrierScheduling_r17, bandParametersSidelinkDiscoveryR17SlCrossCarrierSchedulingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-TransmissionMode2-PartialSensing-r17: sequence
	{
		if ie.Sl_TransmissionMode2_PartialSensing_r17 != nil {
			c := ie.Sl_TransmissionMode2_PartialSensing_r17
			bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Seq := e.NewSequenceEncoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Constraints)
			if err := bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Seq.EncodePreamble([]bool{c.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil, c.ExtendedCP_Mode2PartialSensing_r17 != nil, c.Dl_OpenLoopPC_Sidelink_r17 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Harq_TxProcessModeTwoSidelink_r17, bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17HarqTxProcessModeTwoSidelinkR17Constraints); err != nil {
				return err
			}
			if c.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil {
				choiceEnc := e.NewChoiceEncoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice {
				case BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17:
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17)
					bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq := e.NewSequenceEncoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints)
					if err := bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.EncodePreamble([]bool{c.Scs_15kHz_r17 != nil, c.Scs_30kHz_r17 != nil, c.Scs_60kHz_r17 != nil}); err != nil {
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
				case BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17:
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17)
					bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq := e.NewSequenceEncoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints)
					if err := bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.EncodePreamble([]bool{c.Scs_60kHz_r17 != nil, c.Scs_120kHz_r17 != nil}); err != nil {
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
				if err := e.EncodeEnumerated((*c.ExtendedCP_Mode2PartialSensing_r17), bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ExtendedCPMode2PartialSensingR17Constraints); err != nil {
					return err
				}
			}
			if c.Dl_OpenLoopPC_Sidelink_r17 != nil {
				if err := e.EncodeEnumerated((*c.Dl_OpenLoopPC_Sidelink_r17), bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17DlOpenLoopPCSidelinkR17Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. tx-IUC-Scheme1-Mode2Sidelink-r17: enumerated
	{
		if ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Tx_IUC_Scheme1_Mode2Sidelink_r17, bandParametersSidelinkDiscoveryR17TxIUCScheme1Mode2SidelinkR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandParametersSidelinkDiscovery_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersSidelinkDiscoveryR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-CrossCarrierScheduling-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(bandParametersSidelinkDiscoveryR17SlCrossCarrierSchedulingR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CrossCarrierScheduling_r17 = &idx
		}
	}

	// 3. sl-TransmissionMode2-PartialSensing-r17: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_TransmissionMode2_PartialSensing_r17 = &struct {
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
			c := ie.Sl_TransmissionMode2_PartialSensing_r17
			bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Seq := d.NewSequenceDecoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Constraints)
			if err := bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17HarqTxProcessModeTwoSidelinkR17Constraints)
				if err != nil {
					return err
				}
				c.Harq_TxProcessModeTwoSidelink_r17 = v
			}
			if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Seq.IsComponentPresent(1) {
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
				choiceDec := d.NewChoiceDecoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice = int(index)
				switch index {
				case BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17:
					(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17 = &struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17)
					bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq := d.NewSequenceDecoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints)
					if err := bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r17) = v
					}
					if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r17) = v
					}
					if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r17) = v
					}
				case BandParametersSidelinkDiscovery_r17_Sl_TransmissionMode2_PartialSensing_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17:
					(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17 = &struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17)
					bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq := d.NewSequenceDecoder(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints)
					if err := bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r17) = v
					}
					if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r17 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r17) = v
					}
				}
			}
			if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Seq.IsComponentPresent(2) {
				c.ExtendedCP_Mode2PartialSensing_r17 = new(int64)
				v, err := d.DecodeEnumerated(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17ExtendedCPMode2PartialSensingR17Constraints)
				if err != nil {
					return err
				}
				(*c.ExtendedCP_Mode2PartialSensing_r17) = v
			}
			if bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17Seq.IsComponentPresent(3) {
				c.Dl_OpenLoopPC_Sidelink_r17 = new(int64)
				v, err := d.DecodeEnumerated(bandParametersSidelinkDiscoveryR17SlTransmissionMode2PartialSensingR17DlOpenLoopPCSidelinkR17Constraints)
				if err != nil {
					return err
				}
				(*c.Dl_OpenLoopPC_Sidelink_r17) = v
			}
		}
	}

	// 4. tx-IUC-Scheme1-Mode2Sidelink-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(bandParametersSidelinkDiscoveryR17TxIUCScheme1Mode2SidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.Tx_IUC_Scheme1_Mode2Sidelink_r17 = &idx
		}
	}

	return nil
}
