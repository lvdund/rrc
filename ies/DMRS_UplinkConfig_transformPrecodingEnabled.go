package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkConfig_transformPrecodingEnabled struct {
	NPUSCH_Identity                   *int64                                                            `lb:0,ub:1007,optional`
	SequenceGroupHopping              *DMRS_UplinkConfig_transformPrecodingEnabled_sequenceGroupHopping `optional`
	SequenceHopping                   *DMRS_UplinkConfig_transformPrecodingEnabled_sequenceHopping      `optional`
	Dmrs_UplinkTransformPrecoding_r16 *DMRS_UplinkTransformPrecoding_r16                                `optional,setuprelease`
}

func (ie *DMRS_UplinkConfig_transformPrecodingEnabled) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NPUSCH_Identity != nil, ie.SequenceGroupHopping != nil, ie.SequenceHopping != nil, ie.Dmrs_UplinkTransformPrecoding_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NPUSCH_Identity != nil {
		if err = w.WriteInteger(*ie.NPUSCH_Identity, &aper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
			return utils.WrapError("Encode NPUSCH_Identity", err)
		}
	}
	if ie.SequenceGroupHopping != nil {
		if err = ie.SequenceGroupHopping.Encode(w); err != nil {
			return utils.WrapError("Encode SequenceGroupHopping", err)
		}
	}
	if ie.SequenceHopping != nil {
		if err = ie.SequenceHopping.Encode(w); err != nil {
			return utils.WrapError("Encode SequenceHopping", err)
		}
	}
	if ie.Dmrs_UplinkTransformPrecoding_r16 != nil {
		tmp_Dmrs_UplinkTransformPrecoding_r16 := utils.SetupRelease[*DMRS_UplinkTransformPrecoding_r16]{
			Setup: ie.Dmrs_UplinkTransformPrecoding_r16,
		}
		if err = tmp_Dmrs_UplinkTransformPrecoding_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_UplinkTransformPrecoding_r16", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkConfig_transformPrecodingEnabled) Decode(r *aper.AperReader) error {
	var err error
	var NPUSCH_IdentityPresent bool
	if NPUSCH_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SequenceGroupHoppingPresent bool
	if SequenceGroupHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SequenceHoppingPresent bool
	if SequenceHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_UplinkTransformPrecoding_r16Present bool
	if Dmrs_UplinkTransformPrecoding_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if NPUSCH_IdentityPresent {
		var tmp_int_NPUSCH_Identity int64
		if tmp_int_NPUSCH_Identity, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
			return utils.WrapError("Decode NPUSCH_Identity", err)
		}
		ie.NPUSCH_Identity = &tmp_int_NPUSCH_Identity
	}
	if SequenceGroupHoppingPresent {
		ie.SequenceGroupHopping = new(DMRS_UplinkConfig_transformPrecodingEnabled_sequenceGroupHopping)
		if err = ie.SequenceGroupHopping.Decode(r); err != nil {
			return utils.WrapError("Decode SequenceGroupHopping", err)
		}
	}
	if SequenceHoppingPresent {
		ie.SequenceHopping = new(DMRS_UplinkConfig_transformPrecodingEnabled_sequenceHopping)
		if err = ie.SequenceHopping.Decode(r); err != nil {
			return utils.WrapError("Decode SequenceHopping", err)
		}
	}
	if Dmrs_UplinkTransformPrecoding_r16Present {
		tmp_Dmrs_UplinkTransformPrecoding_r16 := utils.SetupRelease[*DMRS_UplinkTransformPrecoding_r16]{}
		if err = tmp_Dmrs_UplinkTransformPrecoding_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_UplinkTransformPrecoding_r16", err)
		}
		ie.Dmrs_UplinkTransformPrecoding_r16 = tmp_Dmrs_UplinkTransformPrecoding_r16.Setup
	}
	return nil
}
