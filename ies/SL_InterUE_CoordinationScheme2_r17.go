package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_InterUE_CoordinationScheme2_r17 struct {
	Sl_IUC_Scheme2_r17                *SL_InterUE_CoordinationScheme2_r17_sl_IUC_Scheme2_r17                `optional`
	Sl_RB_SetPSFCH_r17                *uper.BitString                                                       `lb:10,ub:275,optional`
	Sl_TypeUE_A_r17                   *SL_InterUE_CoordinationScheme2_r17_sl_TypeUE_A_r17                   `optional`
	Sl_PSFCH_Occasion_r17             *int64                                                                `lb:0,ub:1,optional`
	Sl_SlotLevelResourceExclusion_r17 *SL_InterUE_CoordinationScheme2_r17_sl_SlotLevelResourceExclusion_r17 `optional`
	Sl_OptionForCondition2_A_1_r17    *int64                                                                `lb:0,ub:1,optional`
	Sl_IndicationUE_B_r17             *SL_InterUE_CoordinationScheme2_r17_sl_IndicationUE_B_r17             `optional`
	Sl_DeltaRSRP_Thresh_v1720         *int64                                                                `lb:-30,ub:30,optional,ext-1`
}

func (ie *SL_InterUE_CoordinationScheme2_r17) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Sl_DeltaRSRP_Thresh_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_IUC_Scheme2_r17 != nil, ie.Sl_RB_SetPSFCH_r17 != nil, ie.Sl_TypeUE_A_r17 != nil, ie.Sl_PSFCH_Occasion_r17 != nil, ie.Sl_SlotLevelResourceExclusion_r17 != nil, ie.Sl_OptionForCondition2_A_1_r17 != nil, ie.Sl_IndicationUE_B_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_IUC_Scheme2_r17 != nil {
		if err = ie.Sl_IUC_Scheme2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_IUC_Scheme2_r17", err)
		}
	}
	if ie.Sl_RB_SetPSFCH_r17 != nil {
		if err = w.WriteBitString(ie.Sl_RB_SetPSFCH_r17.Bytes, uint(ie.Sl_RB_SetPSFCH_r17.NumBits), &uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Encode Sl_RB_SetPSFCH_r17", err)
		}
	}
	if ie.Sl_TypeUE_A_r17 != nil {
		if err = ie.Sl_TypeUE_A_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TypeUE_A_r17", err)
		}
	}
	if ie.Sl_PSFCH_Occasion_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_PSFCH_Occasion_r17, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_Occasion_r17", err)
		}
	}
	if ie.Sl_SlotLevelResourceExclusion_r17 != nil {
		if err = ie.Sl_SlotLevelResourceExclusion_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SlotLevelResourceExclusion_r17", err)
		}
	}
	if ie.Sl_OptionForCondition2_A_1_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_OptionForCondition2_A_1_r17, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode Sl_OptionForCondition2_A_1_r17", err)
		}
	}
	if ie.Sl_IndicationUE_B_r17 != nil {
		if err = ie.Sl_IndicationUE_B_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_IndicationUE_B_r17", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_DeltaRSRP_Thresh_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_InterUE_CoordinationScheme2_r17", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_DeltaRSRP_Thresh_v1720 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_DeltaRSRP_Thresh_v1720 optional
			if ie.Sl_DeltaRSRP_Thresh_v1720 != nil {
				if err = extWriter.WriteInteger(*ie.Sl_DeltaRSRP_Thresh_v1720, &uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
					return utils.WrapError("Encode Sl_DeltaRSRP_Thresh_v1720", err)
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

func (ie *SL_InterUE_CoordinationScheme2_r17) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_IUC_Scheme2_r17Present bool
	if Sl_IUC_Scheme2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RB_SetPSFCH_r17Present bool
	if Sl_RB_SetPSFCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TypeUE_A_r17Present bool
	if Sl_TypeUE_A_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSFCH_Occasion_r17Present bool
	if Sl_PSFCH_Occasion_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SlotLevelResourceExclusion_r17Present bool
	if Sl_SlotLevelResourceExclusion_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_OptionForCondition2_A_1_r17Present bool
	if Sl_OptionForCondition2_A_1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_IndicationUE_B_r17Present bool
	if Sl_IndicationUE_B_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_IUC_Scheme2_r17Present {
		ie.Sl_IUC_Scheme2_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_IUC_Scheme2_r17)
		if err = ie.Sl_IUC_Scheme2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_IUC_Scheme2_r17", err)
		}
	}
	if Sl_RB_SetPSFCH_r17Present {
		var tmp_bs_Sl_RB_SetPSFCH_r17 []byte
		var n_Sl_RB_SetPSFCH_r17 uint
		if tmp_bs_Sl_RB_SetPSFCH_r17, n_Sl_RB_SetPSFCH_r17, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Decode Sl_RB_SetPSFCH_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Sl_RB_SetPSFCH_r17,
			NumBits: uint64(n_Sl_RB_SetPSFCH_r17),
		}
		ie.Sl_RB_SetPSFCH_r17 = &tmp_bitstring
	}
	if Sl_TypeUE_A_r17Present {
		ie.Sl_TypeUE_A_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_TypeUE_A_r17)
		if err = ie.Sl_TypeUE_A_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TypeUE_A_r17", err)
		}
	}
	if Sl_PSFCH_Occasion_r17Present {
		var tmp_int_Sl_PSFCH_Occasion_r17 int64
		if tmp_int_Sl_PSFCH_Occasion_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_Occasion_r17", err)
		}
		ie.Sl_PSFCH_Occasion_r17 = &tmp_int_Sl_PSFCH_Occasion_r17
	}
	if Sl_SlotLevelResourceExclusion_r17Present {
		ie.Sl_SlotLevelResourceExclusion_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_SlotLevelResourceExclusion_r17)
		if err = ie.Sl_SlotLevelResourceExclusion_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SlotLevelResourceExclusion_r17", err)
		}
	}
	if Sl_OptionForCondition2_A_1_r17Present {
		var tmp_int_Sl_OptionForCondition2_A_1_r17 int64
		if tmp_int_Sl_OptionForCondition2_A_1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Sl_OptionForCondition2_A_1_r17", err)
		}
		ie.Sl_OptionForCondition2_A_1_r17 = &tmp_int_Sl_OptionForCondition2_A_1_r17
	}
	if Sl_IndicationUE_B_r17Present {
		ie.Sl_IndicationUE_B_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_IndicationUE_B_r17)
		if err = ie.Sl_IndicationUE_B_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_IndicationUE_B_r17", err)
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

			Sl_DeltaRSRP_Thresh_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_DeltaRSRP_Thresh_v1720 optional
			if Sl_DeltaRSRP_Thresh_v1720Present {
				var tmp_int_Sl_DeltaRSRP_Thresh_v1720 int64
				if tmp_int_Sl_DeltaRSRP_Thresh_v1720, err = extReader.ReadInteger(&uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
					return utils.WrapError("Decode Sl_DeltaRSRP_Thresh_v1720", err)
				}
				ie.Sl_DeltaRSRP_Thresh_v1720 = &tmp_int_Sl_DeltaRSRP_Thresh_v1720
			}
		}
	}
	return nil
}
