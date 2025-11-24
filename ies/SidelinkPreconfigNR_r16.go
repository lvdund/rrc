package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkPreconfigNR_r16 struct {
	Sl_PreconfigFreqInfoList_r16                []SL_FreqConfigCommon_r16                            `lb:1,ub:maxNrofFreqSL_r16,optional`
	Sl_PreconfigNR_AnchorCarrierFreqList_r16    *SL_NR_AnchorCarrierFreqList_r16                     `optional`
	Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 *SL_EUTRA_AnchorCarrierFreqList_r16                  `optional`
	Sl_RadioBearerPreConfigList_r16             []SL_RadioBearerConfig_r16                           `lb:1,ub:maxNrofSLRB_r16,optional`
	Sl_RLC_BearerPreConfigList_r16              []SL_RLC_BearerConfig_r16                            `lb:1,ub:maxSL_LCID_r16,optional`
	Sl_MeasPreConfig_r16                        *SL_MeasConfigCommon_r16                             `optional`
	Sl_OffsetDFN_r16                            *int64                                               `lb:1,ub:1000,optional`
	T400_r16                                    *SidelinkPreconfigNR_r16_t400_r16                    `optional`
	Sl_MaxNumConsecutiveDTX_r16                 *SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16 `optional`
	Sl_SSB_PriorityNR_r16                       *int64                                               `lb:1,ub:8,optional`
	Sl_PreconfigGeneral_r16                     *SL_PreconfigGeneral_r16                             `optional`
	Sl_UE_SelectedPreConfig_r16                 *SL_UE_SelectedConfig_r16                            `optional`
	Sl_CSI_Acquisition_r16                      *SidelinkPreconfigNR_r16_sl_CSI_Acquisition_r16      `optional`
	Sl_RoHC_Profiles_r16                        *SL_RoHC_Profiles_r16                                `optional`
	Sl_MaxCID_r16                               int64                                                `lb:1,ub:16383,madatory`
	Sl_DRX_PreConfigGC_BC_r17                   *SL_DRX_ConfigGC_BC_r17                              `optional,ext-1`
	Sl_TxProfileList_r17                        *SL_TxProfileList_r17                                `optional,ext-1`
	Sl_PreconfigDiscConfig_r17                  *SL_RemoteUE_Config_r17                              `optional,ext-1`
}

func (ie *SidelinkPreconfigNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Sl_DRX_PreConfigGC_BC_r17 != nil || ie.Sl_TxProfileList_r17 != nil || ie.Sl_PreconfigDiscConfig_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.Sl_PreconfigFreqInfoList_r16) > 0, ie.Sl_PreconfigNR_AnchorCarrierFreqList_r16 != nil, ie.Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 != nil, len(ie.Sl_RadioBearerPreConfigList_r16) > 0, len(ie.Sl_RLC_BearerPreConfigList_r16) > 0, ie.Sl_MeasPreConfig_r16 != nil, ie.Sl_OffsetDFN_r16 != nil, ie.T400_r16 != nil, ie.Sl_MaxNumConsecutiveDTX_r16 != nil, ie.Sl_SSB_PriorityNR_r16 != nil, ie.Sl_PreconfigGeneral_r16 != nil, ie.Sl_UE_SelectedPreConfig_r16 != nil, ie.Sl_CSI_Acquisition_r16 != nil, ie.Sl_RoHC_Profiles_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_PreconfigFreqInfoList_r16) > 0 {
		tmp_Sl_PreconfigFreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.Sl_PreconfigFreqInfoList_r16 {
			tmp_Sl_PreconfigFreqInfoList_r16.Value = append(tmp_Sl_PreconfigFreqInfoList_r16.Value, &i)
		}
		if err = tmp_Sl_PreconfigFreqInfoList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PreconfigFreqInfoList_r16", err)
		}
	}
	if ie.Sl_PreconfigNR_AnchorCarrierFreqList_r16 != nil {
		if err = ie.Sl_PreconfigNR_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PreconfigNR_AnchorCarrierFreqList_r16", err)
		}
	}
	if ie.Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 != nil {
		if err = ie.Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if len(ie.Sl_RadioBearerPreConfigList_r16) > 0 {
		tmp_Sl_RadioBearerPreConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.Sl_RadioBearerPreConfigList_r16 {
			tmp_Sl_RadioBearerPreConfigList_r16.Value = append(tmp_Sl_RadioBearerPreConfigList_r16.Value, &i)
		}
		if err = tmp_Sl_RadioBearerPreConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RadioBearerPreConfigList_r16", err)
		}
	}
	if len(ie.Sl_RLC_BearerPreConfigList_r16) > 0 {
		tmp_Sl_RLC_BearerPreConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_BearerPreConfigList_r16 {
			tmp_Sl_RLC_BearerPreConfigList_r16.Value = append(tmp_Sl_RLC_BearerPreConfigList_r16.Value, &i)
		}
		if err = tmp_Sl_RLC_BearerPreConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_BearerPreConfigList_r16", err)
		}
	}
	if ie.Sl_MeasPreConfig_r16 != nil {
		if err = ie.Sl_MeasPreConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasPreConfig_r16", err)
		}
	}
	if ie.Sl_OffsetDFN_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_OffsetDFN_r16, &uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Encode Sl_OffsetDFN_r16", err)
		}
	}
	if ie.T400_r16 != nil {
		if err = ie.T400_r16.Encode(w); err != nil {
			return utils.WrapError("Encode T400_r16", err)
		}
	}
	if ie.Sl_MaxNumConsecutiveDTX_r16 != nil {
		if err = ie.Sl_MaxNumConsecutiveDTX_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if ie.Sl_SSB_PriorityNR_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_SSB_PriorityNR_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_SSB_PriorityNR_r16", err)
		}
	}
	if ie.Sl_PreconfigGeneral_r16 != nil {
		if err = ie.Sl_PreconfigGeneral_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PreconfigGeneral_r16", err)
		}
	}
	if ie.Sl_UE_SelectedPreConfig_r16 != nil {
		if err = ie.Sl_UE_SelectedPreConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_UE_SelectedPreConfig_r16", err)
		}
	}
	if ie.Sl_CSI_Acquisition_r16 != nil {
		if err = ie.Sl_CSI_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CSI_Acquisition_r16", err)
		}
	}
	if ie.Sl_RoHC_Profiles_r16 != nil {
		if err = ie.Sl_RoHC_Profiles_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RoHC_Profiles_r16", err)
		}
	}
	if err = w.WriteInteger(ie.Sl_MaxCID_r16, &uper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MaxCID_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_DRX_PreConfigGC_BC_r17 != nil || ie.Sl_TxProfileList_r17 != nil || ie.Sl_PreconfigDiscConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SidelinkPreconfigNR_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_DRX_PreConfigGC_BC_r17 != nil, ie.Sl_TxProfileList_r17 != nil, ie.Sl_PreconfigDiscConfig_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_DRX_PreConfigGC_BC_r17 optional
			if ie.Sl_DRX_PreConfigGC_BC_r17 != nil {
				if err = ie.Sl_DRX_PreConfigGC_BC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_DRX_PreConfigGC_BC_r17", err)
				}
			}
			// encode Sl_TxProfileList_r17 optional
			if ie.Sl_TxProfileList_r17 != nil {
				if err = ie.Sl_TxProfileList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_TxProfileList_r17", err)
				}
			}
			// encode Sl_PreconfigDiscConfig_r17 optional
			if ie.Sl_PreconfigDiscConfig_r17 != nil {
				if err = ie.Sl_PreconfigDiscConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_PreconfigDiscConfig_r17", err)
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

func (ie *SidelinkPreconfigNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PreconfigFreqInfoList_r16Present bool
	if Sl_PreconfigFreqInfoList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PreconfigNR_AnchorCarrierFreqList_r16Present bool
	if Sl_PreconfigNR_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16Present bool
	if Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RadioBearerPreConfigList_r16Present bool
	if Sl_RadioBearerPreConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_BearerPreConfigList_r16Present bool
	if Sl_RLC_BearerPreConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasPreConfig_r16Present bool
	if Sl_MeasPreConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_OffsetDFN_r16Present bool
	if Sl_OffsetDFN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T400_r16Present bool
	if T400_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MaxNumConsecutiveDTX_r16Present bool
	if Sl_MaxNumConsecutiveDTX_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SSB_PriorityNR_r16Present bool
	if Sl_SSB_PriorityNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PreconfigGeneral_r16Present bool
	if Sl_PreconfigGeneral_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_UE_SelectedPreConfig_r16Present bool
	if Sl_UE_SelectedPreConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CSI_Acquisition_r16Present bool
	if Sl_CSI_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RoHC_Profiles_r16Present bool
	if Sl_RoHC_Profiles_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PreconfigFreqInfoList_r16Present {
		tmp_Sl_PreconfigFreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_Sl_PreconfigFreqInfoList_r16 := func() *SL_FreqConfigCommon_r16 {
			return new(SL_FreqConfigCommon_r16)
		}
		if err = tmp_Sl_PreconfigFreqInfoList_r16.Decode(r, fn_Sl_PreconfigFreqInfoList_r16); err != nil {
			return utils.WrapError("Decode Sl_PreconfigFreqInfoList_r16", err)
		}
		ie.Sl_PreconfigFreqInfoList_r16 = []SL_FreqConfigCommon_r16{}
		for _, i := range tmp_Sl_PreconfigFreqInfoList_r16.Value {
			ie.Sl_PreconfigFreqInfoList_r16 = append(ie.Sl_PreconfigFreqInfoList_r16, *i)
		}
	}
	if Sl_PreconfigNR_AnchorCarrierFreqList_r16Present {
		ie.Sl_PreconfigNR_AnchorCarrierFreqList_r16 = new(SL_NR_AnchorCarrierFreqList_r16)
		if err = ie.Sl_PreconfigNR_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PreconfigNR_AnchorCarrierFreqList_r16", err)
		}
	}
	if Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16Present {
		ie.Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 = new(SL_EUTRA_AnchorCarrierFreqList_r16)
		if err = ie.Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PreconfigEUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if Sl_RadioBearerPreConfigList_r16Present {
		tmp_Sl_RadioBearerPreConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_Sl_RadioBearerPreConfigList_r16 := func() *SL_RadioBearerConfig_r16 {
			return new(SL_RadioBearerConfig_r16)
		}
		if err = tmp_Sl_RadioBearerPreConfigList_r16.Decode(r, fn_Sl_RadioBearerPreConfigList_r16); err != nil {
			return utils.WrapError("Decode Sl_RadioBearerPreConfigList_r16", err)
		}
		ie.Sl_RadioBearerPreConfigList_r16 = []SL_RadioBearerConfig_r16{}
		for _, i := range tmp_Sl_RadioBearerPreConfigList_r16.Value {
			ie.Sl_RadioBearerPreConfigList_r16 = append(ie.Sl_RadioBearerPreConfigList_r16, *i)
		}
	}
	if Sl_RLC_BearerPreConfigList_r16Present {
		tmp_Sl_RLC_BearerPreConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_BearerPreConfigList_r16 := func() *SL_RLC_BearerConfig_r16 {
			return new(SL_RLC_BearerConfig_r16)
		}
		if err = tmp_Sl_RLC_BearerPreConfigList_r16.Decode(r, fn_Sl_RLC_BearerPreConfigList_r16); err != nil {
			return utils.WrapError("Decode Sl_RLC_BearerPreConfigList_r16", err)
		}
		ie.Sl_RLC_BearerPreConfigList_r16 = []SL_RLC_BearerConfig_r16{}
		for _, i := range tmp_Sl_RLC_BearerPreConfigList_r16.Value {
			ie.Sl_RLC_BearerPreConfigList_r16 = append(ie.Sl_RLC_BearerPreConfigList_r16, *i)
		}
	}
	if Sl_MeasPreConfig_r16Present {
		ie.Sl_MeasPreConfig_r16 = new(SL_MeasConfigCommon_r16)
		if err = ie.Sl_MeasPreConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasPreConfig_r16", err)
		}
	}
	if Sl_OffsetDFN_r16Present {
		var tmp_int_Sl_OffsetDFN_r16 int64
		if tmp_int_Sl_OffsetDFN_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Decode Sl_OffsetDFN_r16", err)
		}
		ie.Sl_OffsetDFN_r16 = &tmp_int_Sl_OffsetDFN_r16
	}
	if T400_r16Present {
		ie.T400_r16 = new(SidelinkPreconfigNR_r16_t400_r16)
		if err = ie.T400_r16.Decode(r); err != nil {
			return utils.WrapError("Decode T400_r16", err)
		}
	}
	if Sl_MaxNumConsecutiveDTX_r16Present {
		ie.Sl_MaxNumConsecutiveDTX_r16 = new(SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16)
		if err = ie.Sl_MaxNumConsecutiveDTX_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if Sl_SSB_PriorityNR_r16Present {
		var tmp_int_Sl_SSB_PriorityNR_r16 int64
		if tmp_int_Sl_SSB_PriorityNR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_SSB_PriorityNR_r16", err)
		}
		ie.Sl_SSB_PriorityNR_r16 = &tmp_int_Sl_SSB_PriorityNR_r16
	}
	if Sl_PreconfigGeneral_r16Present {
		ie.Sl_PreconfigGeneral_r16 = new(SL_PreconfigGeneral_r16)
		if err = ie.Sl_PreconfigGeneral_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PreconfigGeneral_r16", err)
		}
	}
	if Sl_UE_SelectedPreConfig_r16Present {
		ie.Sl_UE_SelectedPreConfig_r16 = new(SL_UE_SelectedConfig_r16)
		if err = ie.Sl_UE_SelectedPreConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_UE_SelectedPreConfig_r16", err)
		}
	}
	if Sl_CSI_Acquisition_r16Present {
		ie.Sl_CSI_Acquisition_r16 = new(SidelinkPreconfigNR_r16_sl_CSI_Acquisition_r16)
		if err = ie.Sl_CSI_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CSI_Acquisition_r16", err)
		}
	}
	if Sl_RoHC_Profiles_r16Present {
		ie.Sl_RoHC_Profiles_r16 = new(SL_RoHC_Profiles_r16)
		if err = ie.Sl_RoHC_Profiles_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RoHC_Profiles_r16", err)
		}
	}
	var tmp_int_Sl_MaxCID_r16 int64
	if tmp_int_Sl_MaxCID_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MaxCID_r16", err)
	}
	ie.Sl_MaxCID_r16 = tmp_int_Sl_MaxCID_r16

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

			Sl_DRX_PreConfigGC_BC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_TxProfileList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_PreconfigDiscConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_DRX_PreConfigGC_BC_r17 optional
			if Sl_DRX_PreConfigGC_BC_r17Present {
				ie.Sl_DRX_PreConfigGC_BC_r17 = new(SL_DRX_ConfigGC_BC_r17)
				if err = ie.Sl_DRX_PreConfigGC_BC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_DRX_PreConfigGC_BC_r17", err)
				}
			}
			// decode Sl_TxProfileList_r17 optional
			if Sl_TxProfileList_r17Present {
				ie.Sl_TxProfileList_r17 = new(SL_TxProfileList_r17)
				if err = ie.Sl_TxProfileList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_TxProfileList_r17", err)
				}
			}
			// decode Sl_PreconfigDiscConfig_r17 optional
			if Sl_PreconfigDiscConfig_r17Present {
				ie.Sl_PreconfigDiscConfig_r17 = new(SL_RemoteUE_Config_r17)
				if err = ie.Sl_PreconfigDiscConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_PreconfigDiscConfig_r17", err)
				}
			}
		}
	}
	return nil
}
