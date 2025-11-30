package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DCP_Config_r16 struct {
	Ps_RNTI_r16                     RNTI_Value                                      `madatory`
	Ps_Offset_r16                   int64                                           `lb:1,ub:120,madatory`
	SizeDCI_2_6_r16                 int64                                           `lb:1,ub:maxDCI_2_6_Size_r16,madatory`
	Ps_PositionDCI_2_6_r16          int64                                           `lb:0,ub:maxDCI_2_6_Size_1_r16,madatory`
	Ps_WakeUp_r16                   *DCP_Config_r16_ps_WakeUp_r16                   `optional`
	Ps_TransmitPeriodicL1_RSRP_r16  *DCP_Config_r16_ps_TransmitPeriodicL1_RSRP_r16  `optional`
	Ps_TransmitOtherPeriodicCSI_r16 *DCP_Config_r16_ps_TransmitOtherPeriodicCSI_r16 `optional`
}

func (ie *DCP_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ps_WakeUp_r16 != nil, ie.Ps_TransmitPeriodicL1_RSRP_r16 != nil, ie.Ps_TransmitOtherPeriodicCSI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ps_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ps_RNTI_r16", err)
	}
	if err = w.WriteInteger(ie.Ps_Offset_r16, &aper.Constraint{Lb: 1, Ub: 120}, false); err != nil {
		return utils.WrapError("WriteInteger Ps_Offset_r16", err)
	}
	if err = w.WriteInteger(ie.SizeDCI_2_6_r16, &aper.Constraint{Lb: 1, Ub: maxDCI_2_6_Size_r16}, false); err != nil {
		return utils.WrapError("WriteInteger SizeDCI_2_6_r16", err)
	}
	if err = w.WriteInteger(ie.Ps_PositionDCI_2_6_r16, &aper.Constraint{Lb: 0, Ub: maxDCI_2_6_Size_1_r16}, false); err != nil {
		return utils.WrapError("WriteInteger Ps_PositionDCI_2_6_r16", err)
	}
	if ie.Ps_WakeUp_r16 != nil {
		if err = ie.Ps_WakeUp_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ps_WakeUp_r16", err)
		}
	}
	if ie.Ps_TransmitPeriodicL1_RSRP_r16 != nil {
		if err = ie.Ps_TransmitPeriodicL1_RSRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ps_TransmitPeriodicL1_RSRP_r16", err)
		}
	}
	if ie.Ps_TransmitOtherPeriodicCSI_r16 != nil {
		if err = ie.Ps_TransmitOtherPeriodicCSI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ps_TransmitOtherPeriodicCSI_r16", err)
		}
	}
	return nil
}

func (ie *DCP_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var Ps_WakeUp_r16Present bool
	if Ps_WakeUp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ps_TransmitPeriodicL1_RSRP_r16Present bool
	if Ps_TransmitPeriodicL1_RSRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ps_TransmitOtherPeriodicCSI_r16Present bool
	if Ps_TransmitOtherPeriodicCSI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ps_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ps_RNTI_r16", err)
	}
	var tmp_int_Ps_Offset_r16 int64
	if tmp_int_Ps_Offset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 120}, false); err != nil {
		return utils.WrapError("ReadInteger Ps_Offset_r16", err)
	}
	ie.Ps_Offset_r16 = tmp_int_Ps_Offset_r16
	var tmp_int_SizeDCI_2_6_r16 int64
	if tmp_int_SizeDCI_2_6_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxDCI_2_6_Size_r16}, false); err != nil {
		return utils.WrapError("ReadInteger SizeDCI_2_6_r16", err)
	}
	ie.SizeDCI_2_6_r16 = tmp_int_SizeDCI_2_6_r16
	var tmp_int_Ps_PositionDCI_2_6_r16 int64
	if tmp_int_Ps_PositionDCI_2_6_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxDCI_2_6_Size_1_r16}, false); err != nil {
		return utils.WrapError("ReadInteger Ps_PositionDCI_2_6_r16", err)
	}
	ie.Ps_PositionDCI_2_6_r16 = tmp_int_Ps_PositionDCI_2_6_r16
	if Ps_WakeUp_r16Present {
		ie.Ps_WakeUp_r16 = new(DCP_Config_r16_ps_WakeUp_r16)
		if err = ie.Ps_WakeUp_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ps_WakeUp_r16", err)
		}
	}
	if Ps_TransmitPeriodicL1_RSRP_r16Present {
		ie.Ps_TransmitPeriodicL1_RSRP_r16 = new(DCP_Config_r16_ps_TransmitPeriodicL1_RSRP_r16)
		if err = ie.Ps_TransmitPeriodicL1_RSRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ps_TransmitPeriodicL1_RSRP_r16", err)
		}
	}
	if Ps_TransmitOtherPeriodicCSI_r16Present {
		ie.Ps_TransmitOtherPeriodicCSI_r16 = new(DCP_Config_r16_ps_TransmitOtherPeriodicCSI_r16)
		if err = ie.Ps_TransmitOtherPeriodicCSI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ps_TransmitOtherPeriodicCSI_r16", err)
		}
	}
	return nil
}
