package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_PUSCH_Resource_r16 struct {
	MsgA_MCS_r16                         int64                                                       `lb:0,ub:15,madatory`
	NrofSlotsMsgA_PUSCH_r16              int64                                                       `lb:1,ub:4,madatory`
	NrofMsgA_PO_PerSlot_r16              MsgA_PUSCH_Resource_r16_nrofMsgA_PO_PerSlot_r16             `madatory`
	MsgA_PUSCH_TimeDomainOffset_r16      int64                                                       `lb:1,ub:32,madatory`
	MsgA_PUSCH_TimeDomainAllocation_r16  *int64                                                      `lb:1,ub:maxNrofUL_Allocations,optional`
	StartSymbolAndLengthMsgA_PO_r16      *int64                                                      `lb:0,ub:127,optional`
	MappingTypeMsgA_PUSCH_r16            *MsgA_PUSCH_Resource_r16_mappingTypeMsgA_PUSCH_r16          `optional`
	GuardPeriodMsgA_PUSCH_r16            *int64                                                      `lb:0,ub:3,optional`
	GuardBandMsgA_PUSCH_r16              int64                                                       `lb:0,ub:1,madatory`
	FrequencyStartMsgA_PUSCH_r16         int64                                                       `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
	NrofPRBs_PerMsgA_PO_r16              int64                                                       `lb:1,ub:32,madatory`
	NrofMsgA_PO_FDM_r16                  MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16                 `madatory`
	MsgA_IntraSlotFrequencyHopping_r16   *MsgA_PUSCH_Resource_r16_msgA_IntraSlotFrequencyHopping_r16 `optional`
	MsgA_HoppingBits_r16                 *aper.BitString                                             `lb:2,ub:2,optional`
	MsgA_DMRS_Config_r16                 MsgA_DMRS_Config_r16                                        `madatory`
	NrofDMRS_Sequences_r16               int64                                                       `lb:1,ub:2,madatory`
	MsgA_Alpha_r16                       *MsgA_PUSCH_Resource_r16_msgA_Alpha_r16                     `optional`
	InterlaceIndexFirstPO_MsgA_PUSCH_r16 *int64                                                      `lb:1,ub:10,optional`
	NrofInterlacesPerMsgA_PO_r16         *int64                                                      `lb:1,ub:10,optional`
}

func (ie *MsgA_PUSCH_Resource_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MsgA_PUSCH_TimeDomainAllocation_r16 != nil, ie.StartSymbolAndLengthMsgA_PO_r16 != nil, ie.MappingTypeMsgA_PUSCH_r16 != nil, ie.GuardPeriodMsgA_PUSCH_r16 != nil, ie.MsgA_IntraSlotFrequencyHopping_r16 != nil, ie.MsgA_HoppingBits_r16 != nil, ie.MsgA_Alpha_r16 != nil, ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16 != nil, ie.NrofInterlacesPerMsgA_PO_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.MsgA_MCS_r16, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger MsgA_MCS_r16", err)
	}
	if err = w.WriteInteger(ie.NrofSlotsMsgA_PUSCH_r16, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger NrofSlotsMsgA_PUSCH_r16", err)
	}
	if err = ie.NrofMsgA_PO_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode NrofMsgA_PO_PerSlot_r16", err)
	}
	if err = w.WriteInteger(ie.MsgA_PUSCH_TimeDomainOffset_r16, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger MsgA_PUSCH_TimeDomainOffset_r16", err)
	}
	if ie.MsgA_PUSCH_TimeDomainAllocation_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_PUSCH_TimeDomainAllocation_r16, &aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
			return utils.WrapError("Encode MsgA_PUSCH_TimeDomainAllocation_r16", err)
		}
	}
	if ie.StartSymbolAndLengthMsgA_PO_r16 != nil {
		if err = w.WriteInteger(*ie.StartSymbolAndLengthMsgA_PO_r16, &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Encode StartSymbolAndLengthMsgA_PO_r16", err)
		}
	}
	if ie.MappingTypeMsgA_PUSCH_r16 != nil {
		if err = ie.MappingTypeMsgA_PUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MappingTypeMsgA_PUSCH_r16", err)
		}
	}
	if ie.GuardPeriodMsgA_PUSCH_r16 != nil {
		if err = w.WriteInteger(*ie.GuardPeriodMsgA_PUSCH_r16, &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode GuardPeriodMsgA_PUSCH_r16", err)
		}
	}
	if err = w.WriteInteger(ie.GuardBandMsgA_PUSCH_r16, &aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteInteger GuardBandMsgA_PUSCH_r16", err)
	}
	if err = w.WriteInteger(ie.FrequencyStartMsgA_PUSCH_r16, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("WriteInteger FrequencyStartMsgA_PUSCH_r16", err)
	}
	if err = w.WriteInteger(ie.NrofPRBs_PerMsgA_PO_r16, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger NrofPRBs_PerMsgA_PO_r16", err)
	}
	if err = ie.NrofMsgA_PO_FDM_r16.Encode(w); err != nil {
		return utils.WrapError("Encode NrofMsgA_PO_FDM_r16", err)
	}
	if ie.MsgA_IntraSlotFrequencyHopping_r16 != nil {
		if err = ie.MsgA_IntraSlotFrequencyHopping_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_IntraSlotFrequencyHopping_r16", err)
		}
	}
	if ie.MsgA_HoppingBits_r16 != nil {
		if err = w.WriteBitString(ie.MsgA_HoppingBits_r16.Bytes, uint(ie.MsgA_HoppingBits_r16.NumBits), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode MsgA_HoppingBits_r16", err)
		}
	}
	if err = ie.MsgA_DMRS_Config_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MsgA_DMRS_Config_r16", err)
	}
	if err = w.WriteInteger(ie.NrofDMRS_Sequences_r16, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger NrofDMRS_Sequences_r16", err)
	}
	if ie.MsgA_Alpha_r16 != nil {
		if err = ie.MsgA_Alpha_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_Alpha_r16", err)
		}
	}
	if ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16 != nil {
		if err = w.WriteInteger(*ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16, &aper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode InterlaceIndexFirstPO_MsgA_PUSCH_r16", err)
		}
	}
	if ie.NrofInterlacesPerMsgA_PO_r16 != nil {
		if err = w.WriteInteger(*ie.NrofInterlacesPerMsgA_PO_r16, &aper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode NrofInterlacesPerMsgA_PO_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_PUSCH_Resource_r16) Decode(r *aper.AperReader) error {
	var err error
	var MsgA_PUSCH_TimeDomainAllocation_r16Present bool
	if MsgA_PUSCH_TimeDomainAllocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var StartSymbolAndLengthMsgA_PO_r16Present bool
	if StartSymbolAndLengthMsgA_PO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MappingTypeMsgA_PUSCH_r16Present bool
	if MappingTypeMsgA_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GuardPeriodMsgA_PUSCH_r16Present bool
	if GuardPeriodMsgA_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_IntraSlotFrequencyHopping_r16Present bool
	if MsgA_IntraSlotFrequencyHopping_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_HoppingBits_r16Present bool
	if MsgA_HoppingBits_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_Alpha_r16Present bool
	if MsgA_Alpha_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterlaceIndexFirstPO_MsgA_PUSCH_r16Present bool
	if InterlaceIndexFirstPO_MsgA_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofInterlacesPerMsgA_PO_r16Present bool
	if NrofInterlacesPerMsgA_PO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_MsgA_MCS_r16 int64
	if tmp_int_MsgA_MCS_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger MsgA_MCS_r16", err)
	}
	ie.MsgA_MCS_r16 = tmp_int_MsgA_MCS_r16
	var tmp_int_NrofSlotsMsgA_PUSCH_r16 int64
	if tmp_int_NrofSlotsMsgA_PUSCH_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger NrofSlotsMsgA_PUSCH_r16", err)
	}
	ie.NrofSlotsMsgA_PUSCH_r16 = tmp_int_NrofSlotsMsgA_PUSCH_r16
	if err = ie.NrofMsgA_PO_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode NrofMsgA_PO_PerSlot_r16", err)
	}
	var tmp_int_MsgA_PUSCH_TimeDomainOffset_r16 int64
	if tmp_int_MsgA_PUSCH_TimeDomainOffset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger MsgA_PUSCH_TimeDomainOffset_r16", err)
	}
	ie.MsgA_PUSCH_TimeDomainOffset_r16 = tmp_int_MsgA_PUSCH_TimeDomainOffset_r16
	if MsgA_PUSCH_TimeDomainAllocation_r16Present {
		var tmp_int_MsgA_PUSCH_TimeDomainAllocation_r16 int64
		if tmp_int_MsgA_PUSCH_TimeDomainAllocation_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
			return utils.WrapError("Decode MsgA_PUSCH_TimeDomainAllocation_r16", err)
		}
		ie.MsgA_PUSCH_TimeDomainAllocation_r16 = &tmp_int_MsgA_PUSCH_TimeDomainAllocation_r16
	}
	if StartSymbolAndLengthMsgA_PO_r16Present {
		var tmp_int_StartSymbolAndLengthMsgA_PO_r16 int64
		if tmp_int_StartSymbolAndLengthMsgA_PO_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode StartSymbolAndLengthMsgA_PO_r16", err)
		}
		ie.StartSymbolAndLengthMsgA_PO_r16 = &tmp_int_StartSymbolAndLengthMsgA_PO_r16
	}
	if MappingTypeMsgA_PUSCH_r16Present {
		ie.MappingTypeMsgA_PUSCH_r16 = new(MsgA_PUSCH_Resource_r16_mappingTypeMsgA_PUSCH_r16)
		if err = ie.MappingTypeMsgA_PUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MappingTypeMsgA_PUSCH_r16", err)
		}
	}
	if GuardPeriodMsgA_PUSCH_r16Present {
		var tmp_int_GuardPeriodMsgA_PUSCH_r16 int64
		if tmp_int_GuardPeriodMsgA_PUSCH_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode GuardPeriodMsgA_PUSCH_r16", err)
		}
		ie.GuardPeriodMsgA_PUSCH_r16 = &tmp_int_GuardPeriodMsgA_PUSCH_r16
	}
	var tmp_int_GuardBandMsgA_PUSCH_r16 int64
	if tmp_int_GuardBandMsgA_PUSCH_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadInteger GuardBandMsgA_PUSCH_r16", err)
	}
	ie.GuardBandMsgA_PUSCH_r16 = tmp_int_GuardBandMsgA_PUSCH_r16
	var tmp_int_FrequencyStartMsgA_PUSCH_r16 int64
	if tmp_int_FrequencyStartMsgA_PUSCH_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("ReadInteger FrequencyStartMsgA_PUSCH_r16", err)
	}
	ie.FrequencyStartMsgA_PUSCH_r16 = tmp_int_FrequencyStartMsgA_PUSCH_r16
	var tmp_int_NrofPRBs_PerMsgA_PO_r16 int64
	if tmp_int_NrofPRBs_PerMsgA_PO_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger NrofPRBs_PerMsgA_PO_r16", err)
	}
	ie.NrofPRBs_PerMsgA_PO_r16 = tmp_int_NrofPRBs_PerMsgA_PO_r16
	if err = ie.NrofMsgA_PO_FDM_r16.Decode(r); err != nil {
		return utils.WrapError("Decode NrofMsgA_PO_FDM_r16", err)
	}
	if MsgA_IntraSlotFrequencyHopping_r16Present {
		ie.MsgA_IntraSlotFrequencyHopping_r16 = new(MsgA_PUSCH_Resource_r16_msgA_IntraSlotFrequencyHopping_r16)
		if err = ie.MsgA_IntraSlotFrequencyHopping_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_IntraSlotFrequencyHopping_r16", err)
		}
	}
	if MsgA_HoppingBits_r16Present {
		var tmp_bs_MsgA_HoppingBits_r16 []byte
		var n_MsgA_HoppingBits_r16 uint
		if tmp_bs_MsgA_HoppingBits_r16, n_MsgA_HoppingBits_r16, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode MsgA_HoppingBits_r16", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_MsgA_HoppingBits_r16,
			NumBits: uint64(n_MsgA_HoppingBits_r16),
		}
		ie.MsgA_HoppingBits_r16 = &tmp_bitstring
	}
	if err = ie.MsgA_DMRS_Config_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MsgA_DMRS_Config_r16", err)
	}
	var tmp_int_NrofDMRS_Sequences_r16 int64
	if tmp_int_NrofDMRS_Sequences_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger NrofDMRS_Sequences_r16", err)
	}
	ie.NrofDMRS_Sequences_r16 = tmp_int_NrofDMRS_Sequences_r16
	if MsgA_Alpha_r16Present {
		ie.MsgA_Alpha_r16 = new(MsgA_PUSCH_Resource_r16_msgA_Alpha_r16)
		if err = ie.MsgA_Alpha_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_Alpha_r16", err)
		}
	}
	if InterlaceIndexFirstPO_MsgA_PUSCH_r16Present {
		var tmp_int_InterlaceIndexFirstPO_MsgA_PUSCH_r16 int64
		if tmp_int_InterlaceIndexFirstPO_MsgA_PUSCH_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode InterlaceIndexFirstPO_MsgA_PUSCH_r16", err)
		}
		ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16 = &tmp_int_InterlaceIndexFirstPO_MsgA_PUSCH_r16
	}
	if NrofInterlacesPerMsgA_PO_r16Present {
		var tmp_int_NrofInterlacesPerMsgA_PO_r16 int64
		if tmp_int_NrofInterlacesPerMsgA_PO_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode NrofInterlacesPerMsgA_PO_r16", err)
		}
		ie.NrofInterlacesPerMsgA_PO_r16 = &tmp_int_NrofInterlacesPerMsgA_PO_r16
	}
	return nil
}
