package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkCommon struct {
	GenericParameters                     BWP                                           `madatory`
	Rach_ConfigCommon                     *RACH_ConfigCommon                            `optional,setuprelease`
	Pusch_ConfigCommon                    *PUSCH_ConfigCommon                           `optional,setuprelease`
	Pucch_ConfigCommon                    *PUCCH_ConfigCommon                           `optional,setuprelease`
	Rach_ConfigCommonIAB_r16              *RACH_ConfigCommon                            `optional,ext-1,setuprelease`
	UseInterlacePUCCH_PUSCH_r16           *BWP_UplinkCommon_useInterlacePUCCH_PUSCH_r16 `optional,ext-1`
	MsgA_ConfigCommon_r16                 *MsgA_ConfigCommon_r16                        `optional,ext-1,setuprelease`
	EnableRA_PrioritizationForSlicing_r17 *bool                                         `optional,ext-2`
	AdditionalRACH_ConfigList_r17         *AdditionalRACH_ConfigList_r17                `optional,ext-2,setuprelease`
	Rsrp_ThresholdMsg3_r17                *RSRP_Range                                   `optional,ext-2`
	NumberOfMsg3_RepetitionsList_r17      []NumberOfMsg3_Repetitions_r17                `lb:4,ub:4,optional,ext-2`
	Mcs_Msg3_Repetitions_r17              []int64                                       `lb:8,ub:8,e_lb:0,e_ub:0,optional,ext-2`
}

func (ie *BWP_UplinkCommon) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Rach_ConfigCommonIAB_r16 != nil || ie.UseInterlacePUCCH_PUSCH_r16 != nil || ie.MsgA_ConfigCommon_r16 != nil || ie.EnableRA_PrioritizationForSlicing_r17 != nil || ie.AdditionalRACH_ConfigList_r17 != nil || ie.Rsrp_ThresholdMsg3_r17 != nil || len(ie.NumberOfMsg3_RepetitionsList_r17) > 0 || len(ie.Mcs_Msg3_Repetitions_r17) > 0
	preambleBits := []bool{hasExtensions, ie.Rach_ConfigCommon != nil, ie.Pusch_ConfigCommon != nil, ie.Pucch_ConfigCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.GenericParameters.Encode(w); err != nil {
		return utils.WrapError("Encode GenericParameters", err)
	}
	if ie.Rach_ConfigCommon != nil {
		tmp_Rach_ConfigCommon := utils.SetupRelease[*RACH_ConfigCommon]{
			Setup: ie.Rach_ConfigCommon,
		}
		if err = tmp_Rach_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Rach_ConfigCommon", err)
		}
	}
	if ie.Pusch_ConfigCommon != nil {
		tmp_Pusch_ConfigCommon := utils.SetupRelease[*PUSCH_ConfigCommon]{
			Setup: ie.Pusch_ConfigCommon,
		}
		if err = tmp_Pusch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_ConfigCommon", err)
		}
	}
	if ie.Pucch_ConfigCommon != nil {
		tmp_Pucch_ConfigCommon := utils.SetupRelease[*PUCCH_ConfigCommon]{
			Setup: ie.Pucch_ConfigCommon,
		}
		if err = tmp_Pucch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_ConfigCommon", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Rach_ConfigCommonIAB_r16 != nil || ie.UseInterlacePUCCH_PUSCH_r16 != nil || ie.MsgA_ConfigCommon_r16 != nil, ie.EnableRA_PrioritizationForSlicing_r17 != nil || ie.AdditionalRACH_ConfigList_r17 != nil || ie.Rsrp_ThresholdMsg3_r17 != nil || len(ie.NumberOfMsg3_RepetitionsList_r17) > 0 || len(ie.Mcs_Msg3_Repetitions_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BWP_UplinkCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Rach_ConfigCommonIAB_r16 != nil, ie.UseInterlacePUCCH_PUSCH_r16 != nil, ie.MsgA_ConfigCommon_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Rach_ConfigCommonIAB_r16 optional
			if ie.Rach_ConfigCommonIAB_r16 != nil {
				tmp_Rach_ConfigCommonIAB_r16 := utils.SetupRelease[*RACH_ConfigCommon]{
					Setup: ie.Rach_ConfigCommonIAB_r16,
				}
				if err = tmp_Rach_ConfigCommonIAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rach_ConfigCommonIAB_r16", err)
				}
			}
			// encode UseInterlacePUCCH_PUSCH_r16 optional
			if ie.UseInterlacePUCCH_PUSCH_r16 != nil {
				if err = ie.UseInterlacePUCCH_PUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UseInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// encode MsgA_ConfigCommon_r16 optional
			if ie.MsgA_ConfigCommon_r16 != nil {
				tmp_MsgA_ConfigCommon_r16 := utils.SetupRelease[*MsgA_ConfigCommon_r16]{
					Setup: ie.MsgA_ConfigCommon_r16,
				}
				if err = tmp_MsgA_ConfigCommon_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MsgA_ConfigCommon_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.EnableRA_PrioritizationForSlicing_r17 != nil, ie.AdditionalRACH_ConfigList_r17 != nil, ie.Rsrp_ThresholdMsg3_r17 != nil, len(ie.NumberOfMsg3_RepetitionsList_r17) > 0, len(ie.Mcs_Msg3_Repetitions_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnableRA_PrioritizationForSlicing_r17 optional
			if ie.EnableRA_PrioritizationForSlicing_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.EnableRA_PrioritizationForSlicing_r17); err != nil {
					return utils.WrapError("Encode EnableRA_PrioritizationForSlicing_r17", err)
				}
			}
			// encode AdditionalRACH_ConfigList_r17 optional
			if ie.AdditionalRACH_ConfigList_r17 != nil {
				tmp_AdditionalRACH_ConfigList_r17 := utils.SetupRelease[*AdditionalRACH_ConfigList_r17]{
					Setup: ie.AdditionalRACH_ConfigList_r17,
				}
				if err = tmp_AdditionalRACH_ConfigList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AdditionalRACH_ConfigList_r17", err)
				}
			}
			// encode Rsrp_ThresholdMsg3_r17 optional
			if ie.Rsrp_ThresholdMsg3_r17 != nil {
				if err = ie.Rsrp_ThresholdMsg3_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rsrp_ThresholdMsg3_r17", err)
				}
			}
			// encode NumberOfMsg3_RepetitionsList_r17 optional
			if len(ie.NumberOfMsg3_RepetitionsList_r17) > 0 {
				tmp_NumberOfMsg3_RepetitionsList_r17 := utils.NewSequence[*NumberOfMsg3_Repetitions_r17]([]*NumberOfMsg3_Repetitions_r17{}, aper.Constraint{Lb: 4, Ub: 4}, false)
				for _, i := range ie.NumberOfMsg3_RepetitionsList_r17 {
					tmp_NumberOfMsg3_RepetitionsList_r17.Value = append(tmp_NumberOfMsg3_RepetitionsList_r17.Value, &i)
				}
				if err = tmp_NumberOfMsg3_RepetitionsList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NumberOfMsg3_RepetitionsList_r17", err)
				}
			}
			// encode Mcs_Msg3_Repetitions_r17 optional
			if len(ie.Mcs_Msg3_Repetitions_r17) > 0 {
				tmp_Mcs_Msg3_Repetitions_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 8, Ub: 8}, false)
				for _, i := range ie.Mcs_Msg3_Repetitions_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 0}, false)
					tmp_Mcs_Msg3_Repetitions_r17.Value = append(tmp_Mcs_Msg3_Repetitions_r17.Value, &tmp_ie)
				}
				if err = tmp_Mcs_Msg3_Repetitions_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mcs_Msg3_Repetitions_r17", err)
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

func (ie *BWP_UplinkCommon) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Rach_ConfigCommonPresent bool
	if Rach_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_ConfigCommonPresent bool
	if Pusch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_ConfigCommonPresent bool
	if Pucch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.GenericParameters.Decode(r); err != nil {
		return utils.WrapError("Decode GenericParameters", err)
	}
	if Rach_ConfigCommonPresent {
		tmp_Rach_ConfigCommon := utils.SetupRelease[*RACH_ConfigCommon]{}
		if err = tmp_Rach_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Rach_ConfigCommon", err)
		}
		ie.Rach_ConfigCommon = tmp_Rach_ConfigCommon.Setup
	}
	if Pusch_ConfigCommonPresent {
		tmp_Pusch_ConfigCommon := utils.SetupRelease[*PUSCH_ConfigCommon]{}
		if err = tmp_Pusch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_ConfigCommon", err)
		}
		ie.Pusch_ConfigCommon = tmp_Pusch_ConfigCommon.Setup
	}
	if Pucch_ConfigCommonPresent {
		tmp_Pucch_ConfigCommon := utils.SetupRelease[*PUCCH_ConfigCommon]{}
		if err = tmp_Pucch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_ConfigCommon", err)
		}
		ie.Pucch_ConfigCommon = tmp_Pucch_ConfigCommon.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Rach_ConfigCommonIAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UseInterlacePUCCH_PUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_ConfigCommon_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Rach_ConfigCommonIAB_r16 optional
			if Rach_ConfigCommonIAB_r16Present {
				tmp_Rach_ConfigCommonIAB_r16 := utils.SetupRelease[*RACH_ConfigCommon]{}
				if err = tmp_Rach_ConfigCommonIAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rach_ConfigCommonIAB_r16", err)
				}
				ie.Rach_ConfigCommonIAB_r16 = tmp_Rach_ConfigCommonIAB_r16.Setup
			}
			// decode UseInterlacePUCCH_PUSCH_r16 optional
			if UseInterlacePUCCH_PUSCH_r16Present {
				ie.UseInterlacePUCCH_PUSCH_r16 = new(BWP_UplinkCommon_useInterlacePUCCH_PUSCH_r16)
				if err = ie.UseInterlacePUCCH_PUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UseInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// decode MsgA_ConfigCommon_r16 optional
			if MsgA_ConfigCommon_r16Present {
				tmp_MsgA_ConfigCommon_r16 := utils.SetupRelease[*MsgA_ConfigCommon_r16]{}
				if err = tmp_MsgA_ConfigCommon_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MsgA_ConfigCommon_r16", err)
				}
				ie.MsgA_ConfigCommon_r16 = tmp_MsgA_ConfigCommon_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			EnableRA_PrioritizationForSlicing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AdditionalRACH_ConfigList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rsrp_ThresholdMsg3_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NumberOfMsg3_RepetitionsList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mcs_Msg3_Repetitions_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnableRA_PrioritizationForSlicing_r17 optional
			if EnableRA_PrioritizationForSlicing_r17Present {
				var tmp_bool_EnableRA_PrioritizationForSlicing_r17 bool
				if tmp_bool_EnableRA_PrioritizationForSlicing_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode EnableRA_PrioritizationForSlicing_r17", err)
				}
				ie.EnableRA_PrioritizationForSlicing_r17 = &tmp_bool_EnableRA_PrioritizationForSlicing_r17
			}
			// decode AdditionalRACH_ConfigList_r17 optional
			if AdditionalRACH_ConfigList_r17Present {
				tmp_AdditionalRACH_ConfigList_r17 := utils.SetupRelease[*AdditionalRACH_ConfigList_r17]{}
				if err = tmp_AdditionalRACH_ConfigList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AdditionalRACH_ConfigList_r17", err)
				}
				ie.AdditionalRACH_ConfigList_r17 = tmp_AdditionalRACH_ConfigList_r17.Setup
			}
			// decode Rsrp_ThresholdMsg3_r17 optional
			if Rsrp_ThresholdMsg3_r17Present {
				ie.Rsrp_ThresholdMsg3_r17 = new(RSRP_Range)
				if err = ie.Rsrp_ThresholdMsg3_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rsrp_ThresholdMsg3_r17", err)
				}
			}
			// decode NumberOfMsg3_RepetitionsList_r17 optional
			if NumberOfMsg3_RepetitionsList_r17Present {
				tmp_NumberOfMsg3_RepetitionsList_r17 := utils.NewSequence[*NumberOfMsg3_Repetitions_r17]([]*NumberOfMsg3_Repetitions_r17{}, aper.Constraint{Lb: 4, Ub: 4}, false)
				fn_NumberOfMsg3_RepetitionsList_r17 := func() *NumberOfMsg3_Repetitions_r17 {
					return new(NumberOfMsg3_Repetitions_r17)
				}
				if err = tmp_NumberOfMsg3_RepetitionsList_r17.Decode(extReader, fn_NumberOfMsg3_RepetitionsList_r17); err != nil {
					return utils.WrapError("Decode NumberOfMsg3_RepetitionsList_r17", err)
				}
				ie.NumberOfMsg3_RepetitionsList_r17 = []NumberOfMsg3_Repetitions_r17{}
				for _, i := range tmp_NumberOfMsg3_RepetitionsList_r17.Value {
					ie.NumberOfMsg3_RepetitionsList_r17 = append(ie.NumberOfMsg3_RepetitionsList_r17, *i)
				}
			}
			// decode Mcs_Msg3_Repetitions_r17 optional
			if Mcs_Msg3_Repetitions_r17Present {
				tmp_Mcs_Msg3_Repetitions_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 8, Ub: 8}, false)
				fn_Mcs_Msg3_Repetitions_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 0}, false)
					return &ie
				}
				if err = tmp_Mcs_Msg3_Repetitions_r17.Decode(extReader, fn_Mcs_Msg3_Repetitions_r17); err != nil {
					return utils.WrapError("Decode Mcs_Msg3_Repetitions_r17", err)
				}
				ie.Mcs_Msg3_Repetitions_r17 = []int64{}
				for _, i := range tmp_Mcs_Msg3_Repetitions_r17.Value {
					ie.Mcs_Msg3_Repetitions_r17 = append(ie.Mcs_Msg3_Repetitions_r17, int64(i.Value))
				}
			}
		}
	}
	return nil
}
