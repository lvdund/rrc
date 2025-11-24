package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MulticastConfig_r17 struct {
	Pdsch_HARQ_ACK_CodebookListMulticast_r17 *PDSCH_HARQ_ACK_CodebookList_r16                       `optional,setuprelease`
	Type1_Codebook_GenerationMode_r17        *MulticastConfig_r17_type1_Codebook_GenerationMode_r17 `optional`
}

func (ie *MulticastConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17 != nil, ie.Type1_Codebook_GenerationMode_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17 != nil {
		tmp_Pdsch_HARQ_ACK_CodebookListMulticast_r17 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{
			Setup: ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17,
		}
		if err = tmp_Pdsch_HARQ_ACK_CodebookListMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_HARQ_ACK_CodebookListMulticast_r17", err)
		}
	}
	if ie.Type1_Codebook_GenerationMode_r17 != nil {
		if err = ie.Type1_Codebook_GenerationMode_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1_Codebook_GenerationMode_r17", err)
		}
	}
	return nil
}

func (ie *MulticastConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var Pdsch_HARQ_ACK_CodebookListMulticast_r17Present bool
	if Pdsch_HARQ_ACK_CodebookListMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1_Codebook_GenerationMode_r17Present bool
	if Type1_Codebook_GenerationMode_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdsch_HARQ_ACK_CodebookListMulticast_r17Present {
		tmp_Pdsch_HARQ_ACK_CodebookListMulticast_r17 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{}
		if err = tmp_Pdsch_HARQ_ACK_CodebookListMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_HARQ_ACK_CodebookListMulticast_r17", err)
		}
		ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17 = tmp_Pdsch_HARQ_ACK_CodebookListMulticast_r17.Setup
	}
	if Type1_Codebook_GenerationMode_r17Present {
		ie.Type1_Codebook_GenerationMode_r17 = new(MulticastConfig_r17_type1_Codebook_GenerationMode_r17)
		if err = ie.Type1_Codebook_GenerationMode_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_Codebook_GenerationMode_r17", err)
		}
	}
	return nil
}
