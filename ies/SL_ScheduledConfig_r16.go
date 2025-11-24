package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ScheduledConfig_r16 struct {
	Sl_RNTI_r16                      RNTI_Value                        `madatory`
	Mac_MainConfigSL_r16             *MAC_MainConfigSL_r16             `optional`
	Sl_CS_RNTI_r16                   *RNTI_Value                       `optional`
	Sl_PSFCH_ToPUCCH_r16             []int64                           `lb:1,ub:8,e_lb:0,e_ub:15,optional`
	Sl_ConfiguredGrantConfigList_r16 *SL_ConfiguredGrantConfigList_r16 `optional`
	Sl_DCI_ToSL_Trans_r16            []int64                           `lb:1,ub:8,e_lb:1,e_ub:32,optional,ext-1`
}

func (ie *SL_ScheduledConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.Sl_DCI_ToSL_Trans_r16) > 0
	preambleBits := []bool{hasExtensions, ie.Mac_MainConfigSL_r16 != nil, ie.Sl_CS_RNTI_r16 != nil, len(ie.Sl_PSFCH_ToPUCCH_r16) > 0, ie.Sl_ConfiguredGrantConfigList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RNTI_r16", err)
	}
	if ie.Mac_MainConfigSL_r16 != nil {
		if err = ie.Mac_MainConfigSL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_MainConfigSL_r16", err)
		}
	}
	if ie.Sl_CS_RNTI_r16 != nil {
		if err = ie.Sl_CS_RNTI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CS_RNTI_r16", err)
		}
	}
	if len(ie.Sl_PSFCH_ToPUCCH_r16) > 0 {
		tmp_Sl_PSFCH_ToPUCCH_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.Sl_PSFCH_ToPUCCH_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 15}, false)
			tmp_Sl_PSFCH_ToPUCCH_r16.Value = append(tmp_Sl_PSFCH_ToPUCCH_r16.Value, &tmp_ie)
		}
		if err = tmp_Sl_PSFCH_ToPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_ToPUCCH_r16", err)
		}
	}
	if ie.Sl_ConfiguredGrantConfigList_r16 != nil {
		if err = ie.Sl_ConfiguredGrantConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ConfiguredGrantConfigList_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.Sl_DCI_ToSL_Trans_r16) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_ScheduledConfig_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.Sl_DCI_ToSL_Trans_r16) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_DCI_ToSL_Trans_r16 optional
			if len(ie.Sl_DCI_ToSL_Trans_r16) > 0 {
				tmp_Sl_DCI_ToSL_Trans_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
				for _, i := range ie.Sl_DCI_ToSL_Trans_r16 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: 32}, false)
					tmp_Sl_DCI_ToSL_Trans_r16.Value = append(tmp_Sl_DCI_ToSL_Trans_r16.Value, &tmp_ie)
				}
				if err = tmp_Sl_DCI_ToSL_Trans_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_DCI_ToSL_Trans_r16", err)
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

func (ie *SL_ScheduledConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_MainConfigSL_r16Present bool
	if Mac_MainConfigSL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CS_RNTI_r16Present bool
	if Sl_CS_RNTI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSFCH_ToPUCCH_r16Present bool
	if Sl_PSFCH_ToPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ConfiguredGrantConfigList_r16Present bool
	if Sl_ConfiguredGrantConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RNTI_r16", err)
	}
	if Mac_MainConfigSL_r16Present {
		ie.Mac_MainConfigSL_r16 = new(MAC_MainConfigSL_r16)
		if err = ie.Mac_MainConfigSL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_MainConfigSL_r16", err)
		}
	}
	if Sl_CS_RNTI_r16Present {
		ie.Sl_CS_RNTI_r16 = new(RNTI_Value)
		if err = ie.Sl_CS_RNTI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CS_RNTI_r16", err)
		}
	}
	if Sl_PSFCH_ToPUCCH_r16Present {
		tmp_Sl_PSFCH_ToPUCCH_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_Sl_PSFCH_ToPUCCH_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 15}, false)
			return &ie
		}
		if err = tmp_Sl_PSFCH_ToPUCCH_r16.Decode(r, fn_Sl_PSFCH_ToPUCCH_r16); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_ToPUCCH_r16", err)
		}
		ie.Sl_PSFCH_ToPUCCH_r16 = []int64{}
		for _, i := range tmp_Sl_PSFCH_ToPUCCH_r16.Value {
			ie.Sl_PSFCH_ToPUCCH_r16 = append(ie.Sl_PSFCH_ToPUCCH_r16, int64(i.Value))
		}
	}
	if Sl_ConfiguredGrantConfigList_r16Present {
		ie.Sl_ConfiguredGrantConfigList_r16 = new(SL_ConfiguredGrantConfigList_r16)
		if err = ie.Sl_ConfiguredGrantConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ConfiguredGrantConfigList_r16", err)
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

			Sl_DCI_ToSL_Trans_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_DCI_ToSL_Trans_r16 optional
			if Sl_DCI_ToSL_Trans_r16Present {
				tmp_Sl_DCI_ToSL_Trans_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
				fn_Sl_DCI_ToSL_Trans_r16 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: 32}, false)
					return &ie
				}
				if err = tmp_Sl_DCI_ToSL_Trans_r16.Decode(extReader, fn_Sl_DCI_ToSL_Trans_r16); err != nil {
					return utils.WrapError("Decode Sl_DCI_ToSL_Trans_r16", err)
				}
				ie.Sl_DCI_ToSL_Trans_r16 = []int64{}
				for _, i := range tmp_Sl_DCI_ToSL_Trans_r16.Value {
					ie.Sl_DCI_ToSL_Trans_r16 = append(ie.Sl_DCI_ToSL_Trans_r16, int64(i.Value))
				}
			}
		}
	}
	return nil
}
