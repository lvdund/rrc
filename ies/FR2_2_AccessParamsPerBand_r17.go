package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FR2_2_AccessParamsPerBand_r17 struct {
	Dl_FR2_2_SCS_120kHz_r17                   *FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_120kHz_r17                   `optional`
	Ul_FR2_2_SCS_120kHz_r17                   *FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_120kHz_r17                   `optional`
	InitialAccessSSB_120kHz_r17               *FR2_2_AccessParamsPerBand_r17_initialAccessSSB_120kHz_r17               `optional`
	WidebandPRACH_SCS_120kHz_r17              *FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_120kHz_r17              `optional`
	MultiRB_PUCCH_SCS_120kHz_r17              *FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_120kHz_r17              `optional`
	MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 *FR2_2_AccessParamsPerBand_r17_multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 `optional`
	MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 *FR2_2_AccessParamsPerBand_r17_multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 `optional`
	Dl_FR2_2_SCS_480kHz_r17                   *FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_480kHz_r17                   `optional`
	Ul_FR2_2_SCS_480kHz_r17                   *FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_480kHz_r17                   `optional`
	InitialAccessSSB_480kHz_r17               *FR2_2_AccessParamsPerBand_r17_initialAccessSSB_480kHz_r17               `optional`
	WidebandPRACH_SCS_480kHz_r17              *FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_480kHz_r17              `optional`
	MultiRB_PUCCH_SCS_480kHz_r17              *FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_480kHz_r17              `optional`
	EnhancedPDCCH_monitoringSCS_480kHz_r17    *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_480kHz_r17    `optional`
	Dl_FR2_2_SCS_960kHz_r17                   *FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_960kHz_r17                   `optional`
	Ul_FR2_2_SCS_960kHz_r17                   *FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_960kHz_r17                   `optional`
	MultiRB_PUCCH_SCS_960kHz_r17              *FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_960kHz_r17              `optional`
	EnhancedPDCCH_monitoringSCS_960kHz_r17    *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17    `optional`
	Type1_ChannelAccess_FR2_2_r17             *FR2_2_AccessParamsPerBand_r17_type1_ChannelAccess_FR2_2_r17             `optional`
	Type2_ChannelAccess_FR2_2_r17             *FR2_2_AccessParamsPerBand_r17_type2_ChannelAccess_FR2_2_r17             `optional`
	Reduced_BeamSwitchTiming_FR2_2_r17        *FR2_2_AccessParamsPerBand_r17_reduced_BeamSwitchTiming_FR2_2_r17        `optional`
	Support32_DL_HARQ_ProcessPerSCS_r17       *HARQ_ProcessPerSCS_r17                                                  `optional`
	Support32_UL_HARQ_ProcessPerSCS_r17       *HARQ_ProcessPerSCS_r17                                                  `optional`
	Modulation64_QAM_PUSCH_FR2_2_r17          *FR2_2_AccessParamsPerBand_r17_modulation64_QAM_PUSCH_FR2_2_r17          `optional,ext-1`
}

func (ie *FR2_2_AccessParamsPerBand_r17) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Modulation64_QAM_PUSCH_FR2_2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Dl_FR2_2_SCS_120kHz_r17 != nil, ie.Ul_FR2_2_SCS_120kHz_r17 != nil, ie.InitialAccessSSB_120kHz_r17 != nil, ie.WidebandPRACH_SCS_120kHz_r17 != nil, ie.MultiRB_PUCCH_SCS_120kHz_r17 != nil, ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil, ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil, ie.Dl_FR2_2_SCS_480kHz_r17 != nil, ie.Ul_FR2_2_SCS_480kHz_r17 != nil, ie.InitialAccessSSB_480kHz_r17 != nil, ie.WidebandPRACH_SCS_480kHz_r17 != nil, ie.MultiRB_PUCCH_SCS_480kHz_r17 != nil, ie.EnhancedPDCCH_monitoringSCS_480kHz_r17 != nil, ie.Dl_FR2_2_SCS_960kHz_r17 != nil, ie.Ul_FR2_2_SCS_960kHz_r17 != nil, ie.MultiRB_PUCCH_SCS_960kHz_r17 != nil, ie.EnhancedPDCCH_monitoringSCS_960kHz_r17 != nil, ie.Type1_ChannelAccess_FR2_2_r17 != nil, ie.Type2_ChannelAccess_FR2_2_r17 != nil, ie.Reduced_BeamSwitchTiming_FR2_2_r17 != nil, ie.Support32_DL_HARQ_ProcessPerSCS_r17 != nil, ie.Support32_UL_HARQ_ProcessPerSCS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dl_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.Dl_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.Ul_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.Ul_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.InitialAccessSSB_120kHz_r17 != nil {
		if err = ie.InitialAccessSSB_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InitialAccessSSB_120kHz_r17", err)
		}
	}
	if ie.WidebandPRACH_SCS_120kHz_r17 != nil {
		if err = ie.WidebandPRACH_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode WidebandPRACH_SCS_120kHz_r17", err)
		}
	}
	if ie.MultiRB_PUCCH_SCS_120kHz_r17 != nil {
		if err = ie.MultiRB_PUCCH_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MultiRB_PUCCH_SCS_120kHz_r17", err)
		}
	}
	if ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.Dl_FR2_2_SCS_480kHz_r17 != nil {
		if err = ie.Dl_FR2_2_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if ie.Ul_FR2_2_SCS_480kHz_r17 != nil {
		if err = ie.Ul_FR2_2_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if ie.InitialAccessSSB_480kHz_r17 != nil {
		if err = ie.InitialAccessSSB_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InitialAccessSSB_480kHz_r17", err)
		}
	}
	if ie.WidebandPRACH_SCS_480kHz_r17 != nil {
		if err = ie.WidebandPRACH_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode WidebandPRACH_SCS_480kHz_r17", err)
		}
	}
	if ie.MultiRB_PUCCH_SCS_480kHz_r17 != nil {
		if err = ie.MultiRB_PUCCH_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MultiRB_PUCCH_SCS_480kHz_r17", err)
		}
	}
	if ie.EnhancedPDCCH_monitoringSCS_480kHz_r17 != nil {
		if err = ie.EnhancedPDCCH_monitoringSCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EnhancedPDCCH_monitoringSCS_480kHz_r17", err)
		}
	}
	if ie.Dl_FR2_2_SCS_960kHz_r17 != nil {
		if err = ie.Dl_FR2_2_SCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if ie.Ul_FR2_2_SCS_960kHz_r17 != nil {
		if err = ie.Ul_FR2_2_SCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if ie.MultiRB_PUCCH_SCS_960kHz_r17 != nil {
		if err = ie.MultiRB_PUCCH_SCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MultiRB_PUCCH_SCS_960kHz_r17", err)
		}
	}
	if ie.EnhancedPDCCH_monitoringSCS_960kHz_r17 != nil {
		if err = ie.EnhancedPDCCH_monitoringSCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EnhancedPDCCH_monitoringSCS_960kHz_r17", err)
		}
	}
	if ie.Type1_ChannelAccess_FR2_2_r17 != nil {
		if err = ie.Type1_ChannelAccess_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1_ChannelAccess_FR2_2_r17", err)
		}
	}
	if ie.Type2_ChannelAccess_FR2_2_r17 != nil {
		if err = ie.Type2_ChannelAccess_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type2_ChannelAccess_FR2_2_r17", err)
		}
	}
	if ie.Reduced_BeamSwitchTiming_FR2_2_r17 != nil {
		if err = ie.Reduced_BeamSwitchTiming_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Reduced_BeamSwitchTiming_FR2_2_r17", err)
		}
	}
	if ie.Support32_DL_HARQ_ProcessPerSCS_r17 != nil {
		if err = ie.Support32_DL_HARQ_ProcessPerSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Support32_DL_HARQ_ProcessPerSCS_r17", err)
		}
	}
	if ie.Support32_UL_HARQ_ProcessPerSCS_r17 != nil {
		if err = ie.Support32_UL_HARQ_ProcessPerSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Support32_UL_HARQ_ProcessPerSCS_r17", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Modulation64_QAM_PUSCH_FR2_2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap FR2_2_AccessParamsPerBand_r17", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Modulation64_QAM_PUSCH_FR2_2_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Modulation64_QAM_PUSCH_FR2_2_r17 optional
			if ie.Modulation64_QAM_PUSCH_FR2_2_r17 != nil {
				if err = ie.Modulation64_QAM_PUSCH_FR2_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Modulation64_QAM_PUSCH_FR2_2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *FR2_2_AccessParamsPerBand_r17) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_FR2_2_SCS_120kHz_r17Present bool
	if Dl_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FR2_2_SCS_120kHz_r17Present bool
	if Ul_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InitialAccessSSB_120kHz_r17Present bool
	if InitialAccessSSB_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var WidebandPRACH_SCS_120kHz_r17Present bool
	if WidebandPRACH_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiRB_PUCCH_SCS_120kHz_r17Present bool
	if MultiRB_PUCCH_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present bool
	if MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present bool
	if MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_FR2_2_SCS_480kHz_r17Present bool
	if Dl_FR2_2_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FR2_2_SCS_480kHz_r17Present bool
	if Ul_FR2_2_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InitialAccessSSB_480kHz_r17Present bool
	if InitialAccessSSB_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var WidebandPRACH_SCS_480kHz_r17Present bool
	if WidebandPRACH_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiRB_PUCCH_SCS_480kHz_r17Present bool
	if MultiRB_PUCCH_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EnhancedPDCCH_monitoringSCS_480kHz_r17Present bool
	if EnhancedPDCCH_monitoringSCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_FR2_2_SCS_960kHz_r17Present bool
	if Dl_FR2_2_SCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FR2_2_SCS_960kHz_r17Present bool
	if Ul_FR2_2_SCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiRB_PUCCH_SCS_960kHz_r17Present bool
	if MultiRB_PUCCH_SCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EnhancedPDCCH_monitoringSCS_960kHz_r17Present bool
	if EnhancedPDCCH_monitoringSCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1_ChannelAccess_FR2_2_r17Present bool
	if Type1_ChannelAccess_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2_ChannelAccess_FR2_2_r17Present bool
	if Type2_ChannelAccess_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Reduced_BeamSwitchTiming_FR2_2_r17Present bool
	if Reduced_BeamSwitchTiming_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Support32_DL_HARQ_ProcessPerSCS_r17Present bool
	if Support32_DL_HARQ_ProcessPerSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Support32_UL_HARQ_ProcessPerSCS_r17Present bool
	if Support32_UL_HARQ_ProcessPerSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_FR2_2_SCS_120kHz_r17Present {
		ie.Dl_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_120kHz_r17)
		if err = ie.Dl_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if Ul_FR2_2_SCS_120kHz_r17Present {
		ie.Ul_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_120kHz_r17)
		if err = ie.Ul_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if InitialAccessSSB_120kHz_r17Present {
		ie.InitialAccessSSB_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_initialAccessSSB_120kHz_r17)
		if err = ie.InitialAccessSSB_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InitialAccessSSB_120kHz_r17", err)
		}
	}
	if WidebandPRACH_SCS_120kHz_r17Present {
		ie.WidebandPRACH_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_120kHz_r17)
		if err = ie.WidebandPRACH_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode WidebandPRACH_SCS_120kHz_r17", err)
		}
	}
	if MultiRB_PUCCH_SCS_120kHz_r17Present {
		ie.MultiRB_PUCCH_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_120kHz_r17)
		if err = ie.MultiRB_PUCCH_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MultiRB_PUCCH_SCS_120kHz_r17", err)
		}
	}
	if MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present {
		ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17)
		if err = ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present {
		ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17)
		if err = ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if Dl_FR2_2_SCS_480kHz_r17Present {
		ie.Dl_FR2_2_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_480kHz_r17)
		if err = ie.Dl_FR2_2_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if Ul_FR2_2_SCS_480kHz_r17Present {
		ie.Ul_FR2_2_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_480kHz_r17)
		if err = ie.Ul_FR2_2_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if InitialAccessSSB_480kHz_r17Present {
		ie.InitialAccessSSB_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_initialAccessSSB_480kHz_r17)
		if err = ie.InitialAccessSSB_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InitialAccessSSB_480kHz_r17", err)
		}
	}
	if WidebandPRACH_SCS_480kHz_r17Present {
		ie.WidebandPRACH_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_480kHz_r17)
		if err = ie.WidebandPRACH_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode WidebandPRACH_SCS_480kHz_r17", err)
		}
	}
	if MultiRB_PUCCH_SCS_480kHz_r17Present {
		ie.MultiRB_PUCCH_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_480kHz_r17)
		if err = ie.MultiRB_PUCCH_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MultiRB_PUCCH_SCS_480kHz_r17", err)
		}
	}
	if EnhancedPDCCH_monitoringSCS_480kHz_r17Present {
		ie.EnhancedPDCCH_monitoringSCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_480kHz_r17)
		if err = ie.EnhancedPDCCH_monitoringSCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EnhancedPDCCH_monitoringSCS_480kHz_r17", err)
		}
	}
	if Dl_FR2_2_SCS_960kHz_r17Present {
		ie.Dl_FR2_2_SCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_960kHz_r17)
		if err = ie.Dl_FR2_2_SCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if Ul_FR2_2_SCS_960kHz_r17Present {
		ie.Ul_FR2_2_SCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_960kHz_r17)
		if err = ie.Ul_FR2_2_SCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if MultiRB_PUCCH_SCS_960kHz_r17Present {
		ie.MultiRB_PUCCH_SCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_960kHz_r17)
		if err = ie.MultiRB_PUCCH_SCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MultiRB_PUCCH_SCS_960kHz_r17", err)
		}
	}
	if EnhancedPDCCH_monitoringSCS_960kHz_r17Present {
		ie.EnhancedPDCCH_monitoringSCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17)
		if err = ie.EnhancedPDCCH_monitoringSCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EnhancedPDCCH_monitoringSCS_960kHz_r17", err)
		}
	}
	if Type1_ChannelAccess_FR2_2_r17Present {
		ie.Type1_ChannelAccess_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_type1_ChannelAccess_FR2_2_r17)
		if err = ie.Type1_ChannelAccess_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_ChannelAccess_FR2_2_r17", err)
		}
	}
	if Type2_ChannelAccess_FR2_2_r17Present {
		ie.Type2_ChannelAccess_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_type2_ChannelAccess_FR2_2_r17)
		if err = ie.Type2_ChannelAccess_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_ChannelAccess_FR2_2_r17", err)
		}
	}
	if Reduced_BeamSwitchTiming_FR2_2_r17Present {
		ie.Reduced_BeamSwitchTiming_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_reduced_BeamSwitchTiming_FR2_2_r17)
		if err = ie.Reduced_BeamSwitchTiming_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Reduced_BeamSwitchTiming_FR2_2_r17", err)
		}
	}
	if Support32_DL_HARQ_ProcessPerSCS_r17Present {
		ie.Support32_DL_HARQ_ProcessPerSCS_r17 = new(HARQ_ProcessPerSCS_r17)
		if err = ie.Support32_DL_HARQ_ProcessPerSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Support32_DL_HARQ_ProcessPerSCS_r17", err)
		}
	}
	if Support32_UL_HARQ_ProcessPerSCS_r17Present {
		ie.Support32_UL_HARQ_ProcessPerSCS_r17 = new(HARQ_ProcessPerSCS_r17)
		if err = ie.Support32_UL_HARQ_ProcessPerSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Support32_UL_HARQ_ProcessPerSCS_r17", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Modulation64_QAM_PUSCH_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Modulation64_QAM_PUSCH_FR2_2_r17 optional
			if Modulation64_QAM_PUSCH_FR2_2_r17Present {
				ie.Modulation64_QAM_PUSCH_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_modulation64_QAM_PUSCH_FR2_2_r17)
				if err = ie.Modulation64_QAM_PUSCH_FR2_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Modulation64_QAM_PUSCH_FR2_2_r17", err)
				}
			}
		}
	}
	return nil
}
