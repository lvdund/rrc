package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_Configuration_r16 struct {
	Ssb_Freq_r16             ARFCN_ValueNR                              `madatory`
	HalfFrameIndex_r16       SSB_Configuration_r16_halfFrameIndex_r16   `madatory`
	SsbSubcarrierSpacing_r16 SubcarrierSpacing                          `madatory`
	Ssb_Periodicity_r16      *SSB_Configuration_r16_ssb_Periodicity_r16 `optional`
	Sfn0_Offset_r16          *SSB_Configuration_r16_sfn0_Offset_r16     `lb:0,ub:1023,optional`
	Sfn_SSB_Offset_r16       int64                                      `lb:0,ub:15,madatory`
	Ss_PBCH_BlockPower_r16   *int64                                     `lb:-60,ub:50,optional`
}

func (ie *SSB_Configuration_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_Periodicity_r16 != nil, ie.Sfn0_Offset_r16 != nil, ie.Ss_PBCH_BlockPower_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ssb_Freq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_Freq_r16", err)
	}
	if err = ie.HalfFrameIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode HalfFrameIndex_r16", err)
	}
	if err = ie.SsbSubcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SsbSubcarrierSpacing_r16", err)
	}
	if ie.Ssb_Periodicity_r16 != nil {
		if err = ie.Ssb_Periodicity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_Periodicity_r16", err)
		}
	}
	if ie.Sfn0_Offset_r16 != nil {
		if err = ie.Sfn0_Offset_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn0_Offset_r16", err)
		}
	}
	if err = w.WriteInteger(ie.Sfn_SSB_Offset_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Sfn_SSB_Offset_r16", err)
	}
	if ie.Ss_PBCH_BlockPower_r16 != nil {
		if err = w.WriteInteger(*ie.Ss_PBCH_BlockPower_r16, &uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
			return utils.WrapError("Encode Ss_PBCH_BlockPower_r16", err)
		}
	}
	return nil
}

func (ie *SSB_Configuration_r16) Decode(r *uper.UperReader) error {
	var err error
	var Ssb_Periodicity_r16Present bool
	if Ssb_Periodicity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn0_Offset_r16Present bool
	if Sfn0_Offset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ss_PBCH_BlockPower_r16Present bool
	if Ss_PBCH_BlockPower_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ssb_Freq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_Freq_r16", err)
	}
	if err = ie.HalfFrameIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode HalfFrameIndex_r16", err)
	}
	if err = ie.SsbSubcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SsbSubcarrierSpacing_r16", err)
	}
	if Ssb_Periodicity_r16Present {
		ie.Ssb_Periodicity_r16 = new(SSB_Configuration_r16_ssb_Periodicity_r16)
		if err = ie.Ssb_Periodicity_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Periodicity_r16", err)
		}
	}
	if Sfn0_Offset_r16Present {
		ie.Sfn0_Offset_r16 = new(SSB_Configuration_r16_sfn0_Offset_r16)
		if err = ie.Sfn0_Offset_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn0_Offset_r16", err)
		}
	}
	var tmp_int_Sfn_SSB_Offset_r16 int64
	if tmp_int_Sfn_SSB_Offset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Sfn_SSB_Offset_r16", err)
	}
	ie.Sfn_SSB_Offset_r16 = tmp_int_Sfn_SSB_Offset_r16
	if Ss_PBCH_BlockPower_r16Present {
		var tmp_int_Ss_PBCH_BlockPower_r16 int64
		if tmp_int_Ss_PBCH_BlockPower_r16, err = r.ReadInteger(&uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
			return utils.WrapError("Decode Ss_PBCH_BlockPower_r16", err)
		}
		ie.Ss_PBCH_BlockPower_r16 = &tmp_int_Ss_PBCH_BlockPower_r16
	}
	return nil
}
