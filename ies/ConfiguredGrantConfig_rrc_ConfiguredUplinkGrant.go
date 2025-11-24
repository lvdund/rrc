package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant struct {
	TimeDomainOffset                   int64                                                                               `lb:0,ub:5119,madatory`
	TimeDomainAllocation               int64                                                                               `lb:0,ub:15,madatory`
	FrequencyDomainAllocation          uper.BitString                                                                      `lb:18,ub:18,madatory`
	AntennaPort                        int64                                                                               `lb:0,ub:31,madatory`
	Dmrs_SeqInitialization             *int64                                                                              `lb:0,ub:1,optional`
	PrecodingAndNumberOfLayers         int64                                                                               `lb:0,ub:63,madatory`
	Srs_ResourceIndicator              *int64                                                                              `lb:0,ub:15,optional`
	McsAndTBS                          int64                                                                               `lb:0,ub:31,madatory`
	FrequencyHoppingOffset             *int64                                                                              `lb:1,ub:maxNrofPhysicalResourceBlocks_1,optional`
	PathlossReferenceIndex             int64                                                                               `lb:0,ub:maxNrofPUSCH_PathlossReferenceRSs_1,madatory`
	Pusch_RepTypeIndicator_r16         *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_pusch_RepTypeIndicator_r16         `optional`
	FrequencyHoppingPUSCH_RepTypeB_r16 *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16 `optional`
	TimeReferenceSFN_r16               *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_timeReferenceSFN_r16               `optional`
	PathlossReferenceIndex2_r17        *int64                                                                              `lb:0,ub:maxNrofPUSCH_PathlossReferenceRSs_1,optional`
	Srs_ResourceIndicator2_r17         *int64                                                                              `lb:0,ub:15,optional`
	PrecodingAndNumberOfLayers2_r17    *int64                                                                              `lb:0,ub:63,optional`
	TimeDomainAllocation_v1710         *int64                                                                              `lb:16,ub:63,optional`
	TimeDomainOffset_r17               *int64                                                                              `lb:0,ub:40959,optional`
	Cg_SDT_Configuration_r17           *CG_SDT_Configuration_r17                                                           `optional`
}

func (ie *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dmrs_SeqInitialization != nil, ie.Srs_ResourceIndicator != nil, ie.FrequencyHoppingOffset != nil, ie.Pusch_RepTypeIndicator_r16 != nil, ie.FrequencyHoppingPUSCH_RepTypeB_r16 != nil, ie.TimeReferenceSFN_r16 != nil, ie.PathlossReferenceIndex2_r17 != nil, ie.Srs_ResourceIndicator2_r17 != nil, ie.PrecodingAndNumberOfLayers2_r17 != nil, ie.TimeDomainAllocation_v1710 != nil, ie.TimeDomainOffset_r17 != nil, ie.Cg_SDT_Configuration_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.TimeDomainOffset, &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDomainOffset", err)
	}
	if err = w.WriteInteger(ie.TimeDomainAllocation, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDomainAllocation", err)
	}
	if err = w.WriteBitString(ie.FrequencyDomainAllocation.Bytes, uint(ie.FrequencyDomainAllocation.NumBits), &uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
		return utils.WrapError("WriteBitString FrequencyDomainAllocation", err)
	}
	if err = w.WriteInteger(ie.AntennaPort, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger AntennaPort", err)
	}
	if ie.Dmrs_SeqInitialization != nil {
		if err = w.WriteInteger(*ie.Dmrs_SeqInitialization, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode Dmrs_SeqInitialization", err)
		}
	}
	if err = w.WriteInteger(ie.PrecodingAndNumberOfLayers, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger PrecodingAndNumberOfLayers", err)
	}
	if ie.Srs_ResourceIndicator != nil {
		if err = w.WriteInteger(*ie.Srs_ResourceIndicator, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Srs_ResourceIndicator", err)
		}
	}
	if err = w.WriteInteger(ie.McsAndTBS, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger McsAndTBS", err)
	}
	if ie.FrequencyHoppingOffset != nil {
		if err = w.WriteInteger(*ie.FrequencyHoppingOffset, &uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode FrequencyHoppingOffset", err)
		}
	}
	if err = w.WriteInteger(ie.PathlossReferenceIndex, &uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
		return utils.WrapError("WriteInteger PathlossReferenceIndex", err)
	}
	if ie.Pusch_RepTypeIndicator_r16 != nil {
		if err = ie.Pusch_RepTypeIndicator_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_RepTypeIndicator_r16", err)
		}
	}
	if ie.FrequencyHoppingPUSCH_RepTypeB_r16 != nil {
		if err = ie.FrequencyHoppingPUSCH_RepTypeB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyHoppingPUSCH_RepTypeB_r16", err)
		}
	}
	if ie.TimeReferenceSFN_r16 != nil {
		if err = ie.TimeReferenceSFN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TimeReferenceSFN_r16", err)
		}
	}
	if ie.PathlossReferenceIndex2_r17 != nil {
		if err = w.WriteInteger(*ie.PathlossReferenceIndex2_r17, &uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
			return utils.WrapError("Encode PathlossReferenceIndex2_r17", err)
		}
	}
	if ie.Srs_ResourceIndicator2_r17 != nil {
		if err = w.WriteInteger(*ie.Srs_ResourceIndicator2_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Srs_ResourceIndicator2_r17", err)
		}
	}
	if ie.PrecodingAndNumberOfLayers2_r17 != nil {
		if err = w.WriteInteger(*ie.PrecodingAndNumberOfLayers2_r17, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode PrecodingAndNumberOfLayers2_r17", err)
		}
	}
	if ie.TimeDomainAllocation_v1710 != nil {
		if err = w.WriteInteger(*ie.TimeDomainAllocation_v1710, &uper.Constraint{Lb: 16, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode TimeDomainAllocation_v1710", err)
		}
	}
	if ie.TimeDomainOffset_r17 != nil {
		if err = w.WriteInteger(*ie.TimeDomainOffset_r17, &uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Encode TimeDomainOffset_r17", err)
		}
	}
	if ie.Cg_SDT_Configuration_r17 != nil {
		if err = ie.Cg_SDT_Configuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_Configuration_r17", err)
		}
	}
	return nil
}

func (ie *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant) Decode(r *uper.UperReader) error {
	var err error
	var Dmrs_SeqInitializationPresent bool
	if Dmrs_SeqInitializationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ResourceIndicatorPresent bool
	if Srs_ResourceIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyHoppingOffsetPresent bool
	if FrequencyHoppingOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_RepTypeIndicator_r16Present bool
	if Pusch_RepTypeIndicator_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyHoppingPUSCH_RepTypeB_r16Present bool
	if FrequencyHoppingPUSCH_RepTypeB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeReferenceSFN_r16Present bool
	if TimeReferenceSFN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceIndex2_r17Present bool
	if PathlossReferenceIndex2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ResourceIndicator2_r17Present bool
	if Srs_ResourceIndicator2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PrecodingAndNumberOfLayers2_r17Present bool
	if PrecodingAndNumberOfLayers2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeDomainAllocation_v1710Present bool
	if TimeDomainAllocation_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeDomainOffset_r17Present bool
	if TimeDomainOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_Configuration_r17Present bool
	if Cg_SDT_Configuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_TimeDomainOffset int64
	if tmp_int_TimeDomainOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDomainOffset", err)
	}
	ie.TimeDomainOffset = tmp_int_TimeDomainOffset
	var tmp_int_TimeDomainAllocation int64
	if tmp_int_TimeDomainAllocation, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDomainAllocation", err)
	}
	ie.TimeDomainAllocation = tmp_int_TimeDomainAllocation
	var tmp_bs_FrequencyDomainAllocation []byte
	var n_FrequencyDomainAllocation uint
	if tmp_bs_FrequencyDomainAllocation, n_FrequencyDomainAllocation, err = r.ReadBitString(&uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
		return utils.WrapError("ReadBitString FrequencyDomainAllocation", err)
	}
	ie.FrequencyDomainAllocation = uper.BitString{
		Bytes:   tmp_bs_FrequencyDomainAllocation,
		NumBits: uint64(n_FrequencyDomainAllocation),
	}
	var tmp_int_AntennaPort int64
	if tmp_int_AntennaPort, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger AntennaPort", err)
	}
	ie.AntennaPort = tmp_int_AntennaPort
	if Dmrs_SeqInitializationPresent {
		var tmp_int_Dmrs_SeqInitialization int64
		if tmp_int_Dmrs_SeqInitialization, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Dmrs_SeqInitialization", err)
		}
		ie.Dmrs_SeqInitialization = &tmp_int_Dmrs_SeqInitialization
	}
	var tmp_int_PrecodingAndNumberOfLayers int64
	if tmp_int_PrecodingAndNumberOfLayers, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger PrecodingAndNumberOfLayers", err)
	}
	ie.PrecodingAndNumberOfLayers = tmp_int_PrecodingAndNumberOfLayers
	if Srs_ResourceIndicatorPresent {
		var tmp_int_Srs_ResourceIndicator int64
		if tmp_int_Srs_ResourceIndicator, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Srs_ResourceIndicator", err)
		}
		ie.Srs_ResourceIndicator = &tmp_int_Srs_ResourceIndicator
	}
	var tmp_int_McsAndTBS int64
	if tmp_int_McsAndTBS, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger McsAndTBS", err)
	}
	ie.McsAndTBS = tmp_int_McsAndTBS
	if FrequencyHoppingOffsetPresent {
		var tmp_int_FrequencyHoppingOffset int64
		if tmp_int_FrequencyHoppingOffset, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode FrequencyHoppingOffset", err)
		}
		ie.FrequencyHoppingOffset = &tmp_int_FrequencyHoppingOffset
	}
	var tmp_int_PathlossReferenceIndex int64
	if tmp_int_PathlossReferenceIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
		return utils.WrapError("ReadInteger PathlossReferenceIndex", err)
	}
	ie.PathlossReferenceIndex = tmp_int_PathlossReferenceIndex
	if Pusch_RepTypeIndicator_r16Present {
		ie.Pusch_RepTypeIndicator_r16 = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_pusch_RepTypeIndicator_r16)
		if err = ie.Pusch_RepTypeIndicator_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_RepTypeIndicator_r16", err)
		}
	}
	if FrequencyHoppingPUSCH_RepTypeB_r16Present {
		ie.FrequencyHoppingPUSCH_RepTypeB_r16 = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16)
		if err = ie.FrequencyHoppingPUSCH_RepTypeB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyHoppingPUSCH_RepTypeB_r16", err)
		}
	}
	if TimeReferenceSFN_r16Present {
		ie.TimeReferenceSFN_r16 = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_timeReferenceSFN_r16)
		if err = ie.TimeReferenceSFN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TimeReferenceSFN_r16", err)
		}
	}
	if PathlossReferenceIndex2_r17Present {
		var tmp_int_PathlossReferenceIndex2_r17 int64
		if tmp_int_PathlossReferenceIndex2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
			return utils.WrapError("Decode PathlossReferenceIndex2_r17", err)
		}
		ie.PathlossReferenceIndex2_r17 = &tmp_int_PathlossReferenceIndex2_r17
	}
	if Srs_ResourceIndicator2_r17Present {
		var tmp_int_Srs_ResourceIndicator2_r17 int64
		if tmp_int_Srs_ResourceIndicator2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Srs_ResourceIndicator2_r17", err)
		}
		ie.Srs_ResourceIndicator2_r17 = &tmp_int_Srs_ResourceIndicator2_r17
	}
	if PrecodingAndNumberOfLayers2_r17Present {
		var tmp_int_PrecodingAndNumberOfLayers2_r17 int64
		if tmp_int_PrecodingAndNumberOfLayers2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode PrecodingAndNumberOfLayers2_r17", err)
		}
		ie.PrecodingAndNumberOfLayers2_r17 = &tmp_int_PrecodingAndNumberOfLayers2_r17
	}
	if TimeDomainAllocation_v1710Present {
		var tmp_int_TimeDomainAllocation_v1710 int64
		if tmp_int_TimeDomainAllocation_v1710, err = r.ReadInteger(&uper.Constraint{Lb: 16, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode TimeDomainAllocation_v1710", err)
		}
		ie.TimeDomainAllocation_v1710 = &tmp_int_TimeDomainAllocation_v1710
	}
	if TimeDomainOffset_r17Present {
		var tmp_int_TimeDomainOffset_r17 int64
		if tmp_int_TimeDomainOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Decode TimeDomainOffset_r17", err)
		}
		ie.TimeDomainOffset_r17 = &tmp_int_TimeDomainOffset_r17
	}
	if Cg_SDT_Configuration_r17Present {
		ie.Cg_SDT_Configuration_r17 = new(CG_SDT_Configuration_r17)
		if err = ie.Cg_SDT_Configuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_Configuration_r17", err)
		}
	}
	return nil
}
