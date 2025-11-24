package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type2 struct {
	SubType           *CodebookConfig_codebookType_type2_subType          `lb:16,ub:16,optional`
	PhaseAlphabetSize CodebookConfig_codebookType_type2_phaseAlphabetSize `madatory`
	SubbandAmplitude  bool                                                `madatory`
	NumberOfBeams     CodebookConfig_codebookType_type2_numberOfBeams     `madatory`
}

func (ie *CodebookConfig_codebookType_type2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SubType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SubType != nil {
		if err = ie.SubType.Encode(w); err != nil {
			return utils.WrapError("Encode SubType", err)
		}
	}
	if err = ie.PhaseAlphabetSize.Encode(w); err != nil {
		return utils.WrapError("Encode PhaseAlphabetSize", err)
	}
	if err = w.WriteBoolean(ie.SubbandAmplitude); err != nil {
		return utils.WrapError("WriteBoolean SubbandAmplitude", err)
	}
	if err = ie.NumberOfBeams.Encode(w); err != nil {
		return utils.WrapError("Encode NumberOfBeams", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2) Decode(r *uper.UperReader) error {
	var err error
	var SubTypePresent bool
	if SubTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SubTypePresent {
		ie.SubType = new(CodebookConfig_codebookType_type2_subType)
		if err = ie.SubType.Decode(r); err != nil {
			return utils.WrapError("Decode SubType", err)
		}
	}
	if err = ie.PhaseAlphabetSize.Decode(r); err != nil {
		return utils.WrapError("Decode PhaseAlphabetSize", err)
	}
	var tmp_bool_SubbandAmplitude bool
	if tmp_bool_SubbandAmplitude, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean SubbandAmplitude", err)
	}
	ie.SubbandAmplitude = tmp_bool_SubbandAmplitude
	if err = ie.NumberOfBeams.Decode(r); err != nil {
		return utils.WrapError("Decode NumberOfBeams", err)
	}
	return nil
}
