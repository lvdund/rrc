package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17 struct {
	ParamCombination_r17                   int64                                                                        `lb:1,ub:8,madatory`
	ValueOfN_r17                           *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17 `optional`
	NumberOfPMI_SubbandsPerCQI_Subband_r17 *int64                                                                       `lb:1,ub:2,optional`
	TypeII_PortSelectionRI_Restriction_r17 aper.BitString                                                               `lb:4,ub:4,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ValueOfN_r17 != nil, ie.NumberOfPMI_SubbandsPerCQI_Subband_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.ParamCombination_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger ParamCombination_r17", err)
	}
	if ie.ValueOfN_r17 != nil {
		if err = ie.ValueOfN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ValueOfN_r17", err)
		}
	}
	if ie.NumberOfPMI_SubbandsPerCQI_Subband_r17 != nil {
		if err = w.WriteInteger(*ie.NumberOfPMI_SubbandsPerCQI_Subband_r17, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode NumberOfPMI_SubbandsPerCQI_Subband_r17", err)
		}
	}
	if err = w.WriteBitString(ie.TypeII_PortSelectionRI_Restriction_r17.Bytes, uint(ie.TypeII_PortSelectionRI_Restriction_r17.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteBitString TypeII_PortSelectionRI_Restriction_r17", err)
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17) Decode(r *aper.AperReader) error {
	var err error
	var ValueOfN_r17Present bool
	if ValueOfN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NumberOfPMI_SubbandsPerCQI_Subband_r17Present bool
	if NumberOfPMI_SubbandsPerCQI_Subband_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_ParamCombination_r17 int64
	if tmp_int_ParamCombination_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger ParamCombination_r17", err)
	}
	ie.ParamCombination_r17 = tmp_int_ParamCombination_r17
	if ValueOfN_r17Present {
		ie.ValueOfN_r17 = new(CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17)
		if err = ie.ValueOfN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ValueOfN_r17", err)
		}
	}
	if NumberOfPMI_SubbandsPerCQI_Subband_r17Present {
		var tmp_int_NumberOfPMI_SubbandsPerCQI_Subband_r17 int64
		if tmp_int_NumberOfPMI_SubbandsPerCQI_Subband_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode NumberOfPMI_SubbandsPerCQI_Subband_r17", err)
		}
		ie.NumberOfPMI_SubbandsPerCQI_Subband_r17 = &tmp_int_NumberOfPMI_SubbandsPerCQI_Subband_r17
	}
	var tmp_bs_TypeII_PortSelectionRI_Restriction_r17 []byte
	var n_TypeII_PortSelectionRI_Restriction_r17 uint
	if tmp_bs_TypeII_PortSelectionRI_Restriction_r17, n_TypeII_PortSelectionRI_Restriction_r17, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadBitString TypeII_PortSelectionRI_Restriction_r17", err)
	}
	ie.TypeII_PortSelectionRI_Restriction_r17 = aper.BitString{
		Bytes:   tmp_bs_TypeII_PortSelectionRI_Restriction_r17,
		NumBits: uint64(n_TypeII_PortSelectionRI_Restriction_r17),
	}
	return nil
}
