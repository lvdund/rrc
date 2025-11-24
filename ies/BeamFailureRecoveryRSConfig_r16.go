package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureRecoveryRSConfig_r16 struct {
	Rsrp_ThresholdBFR_r16     *RSRP_Range           `optional`
	CandidateBeamRS_List_r16  []CandidateBeamRS_r16 `lb:1,ub:maxNrofCandidateBeams_r16,optional`
	CandidateBeamRS_List2_r17 []CandidateBeamRS_r16 `lb:1,ub:maxNrofCandidateBeams_r16,optional,ext-1`
}

func (ie *BeamFailureRecoveryRSConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.CandidateBeamRS_List2_r17) > 0
	preambleBits := []bool{hasExtensions, ie.Rsrp_ThresholdBFR_r16 != nil, len(ie.CandidateBeamRS_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rsrp_ThresholdBFR_r16 != nil {
		if err = ie.Rsrp_ThresholdBFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_ThresholdBFR_r16", err)
		}
	}
	if len(ie.CandidateBeamRS_List_r16) > 0 {
		tmp_CandidateBeamRS_List_r16 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
		for _, i := range ie.CandidateBeamRS_List_r16 {
			tmp_CandidateBeamRS_List_r16.Value = append(tmp_CandidateBeamRS_List_r16.Value, &i)
		}
		if err = tmp_CandidateBeamRS_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateBeamRS_List_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.CandidateBeamRS_List2_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BeamFailureRecoveryRSConfig_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.CandidateBeamRS_List2_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CandidateBeamRS_List2_r17 optional
			if len(ie.CandidateBeamRS_List2_r17) > 0 {
				tmp_CandidateBeamRS_List2_r17 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
				for _, i := range ie.CandidateBeamRS_List2_r17 {
					tmp_CandidateBeamRS_List2_r17.Value = append(tmp_CandidateBeamRS_List2_r17.Value, &i)
				}
				if err = tmp_CandidateBeamRS_List2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CandidateBeamRS_List2_r17", err)
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

func (ie *BeamFailureRecoveryRSConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Rsrp_ThresholdBFR_r16Present bool
	if Rsrp_ThresholdBFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateBeamRS_List_r16Present bool
	if CandidateBeamRS_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rsrp_ThresholdBFR_r16Present {
		ie.Rsrp_ThresholdBFR_r16 = new(RSRP_Range)
		if err = ie.Rsrp_ThresholdBFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_ThresholdBFR_r16", err)
		}
	}
	if CandidateBeamRS_List_r16Present {
		tmp_CandidateBeamRS_List_r16 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
		fn_CandidateBeamRS_List_r16 := func() *CandidateBeamRS_r16 {
			return new(CandidateBeamRS_r16)
		}
		if err = tmp_CandidateBeamRS_List_r16.Decode(r, fn_CandidateBeamRS_List_r16); err != nil {
			return utils.WrapError("Decode CandidateBeamRS_List_r16", err)
		}
		ie.CandidateBeamRS_List_r16 = []CandidateBeamRS_r16{}
		for _, i := range tmp_CandidateBeamRS_List_r16.Value {
			ie.CandidateBeamRS_List_r16 = append(ie.CandidateBeamRS_List_r16, *i)
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

			CandidateBeamRS_List2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CandidateBeamRS_List2_r17 optional
			if CandidateBeamRS_List2_r17Present {
				tmp_CandidateBeamRS_List2_r17 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
				fn_CandidateBeamRS_List2_r17 := func() *CandidateBeamRS_r16 {
					return new(CandidateBeamRS_r16)
				}
				if err = tmp_CandidateBeamRS_List2_r17.Decode(extReader, fn_CandidateBeamRS_List2_r17); err != nil {
					return utils.WrapError("Decode CandidateBeamRS_List2_r17", err)
				}
				ie.CandidateBeamRS_List2_r17 = []CandidateBeamRS_r16{}
				for _, i := range tmp_CandidateBeamRS_List2_r17.Value {
					ie.CandidateBeamRS_List2_r17 = append(ie.CandidateBeamRS_List2_r17, *i)
				}
			}
		}
	}
	return nil
}
