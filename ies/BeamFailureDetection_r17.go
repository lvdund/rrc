package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureDetection_r17 struct {
	FailureDetectionSet1_r17 *BeamFailureDetectionSet_r17 `optional`
	FailureDetectionSet2_r17 *BeamFailureDetectionSet_r17 `optional`
	AdditionalPCI_r17        *AdditionalPCIIndex_r17      `optional`
}

func (ie *BeamFailureDetection_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FailureDetectionSet1_r17 != nil, ie.FailureDetectionSet2_r17 != nil, ie.AdditionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FailureDetectionSet1_r17 != nil {
		if err = ie.FailureDetectionSet1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode FailureDetectionSet1_r17", err)
		}
	}
	if ie.FailureDetectionSet2_r17 != nil {
		if err = ie.FailureDetectionSet2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode FailureDetectionSet2_r17", err)
		}
	}
	if ie.AdditionalPCI_r17 != nil {
		if err = ie.AdditionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPCI_r17", err)
		}
	}
	return nil
}

func (ie *BeamFailureDetection_r17) Decode(r *aper.AperReader) error {
	var err error
	var FailureDetectionSet1_r17Present bool
	if FailureDetectionSet1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FailureDetectionSet2_r17Present bool
	if FailureDetectionSet2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalPCI_r17Present bool
	if AdditionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if FailureDetectionSet1_r17Present {
		ie.FailureDetectionSet1_r17 = new(BeamFailureDetectionSet_r17)
		if err = ie.FailureDetectionSet1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode FailureDetectionSet1_r17", err)
		}
	}
	if FailureDetectionSet2_r17Present {
		ie.FailureDetectionSet2_r17 = new(BeamFailureDetectionSet_r17)
		if err = ie.FailureDetectionSet2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode FailureDetectionSet2_r17", err)
		}
	}
	if AdditionalPCI_r17Present {
		ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.AdditionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalPCI_r17", err)
		}
	}
	return nil
}
