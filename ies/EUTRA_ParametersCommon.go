package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_ParametersCommon struct {
	Mfbi_EUTRA                *EUTRA_ParametersCommon_mfbi_EUTRA         `optional`
	ModifiedMPR_BehaviorEUTRA *uper.BitString                            `lb:32,ub:32,optional`
	MultiNS_Pmax_EUTRA        *EUTRA_ParametersCommon_multiNS_Pmax_EUTRA `optional`
	Rs_SINR_MeasEUTRA         *EUTRA_ParametersCommon_rs_SINR_MeasEUTRA  `optional`
	Ne_DC                     *EUTRA_ParametersCommon_ne_DC              `optional,ext-1`
	Nr_HO_ToEN_DC_r16         *EUTRA_ParametersCommon_nr_HO_ToEN_DC_r16  `optional,ext-2`
}

func (ie *EUTRA_ParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Ne_DC != nil || ie.Nr_HO_ToEN_DC_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Mfbi_EUTRA != nil, ie.ModifiedMPR_BehaviorEUTRA != nil, ie.MultiNS_Pmax_EUTRA != nil, ie.Rs_SINR_MeasEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mfbi_EUTRA != nil {
		if err = ie.Mfbi_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode Mfbi_EUTRA", err)
		}
	}
	if ie.ModifiedMPR_BehaviorEUTRA != nil {
		if err = w.WriteBitString(ie.ModifiedMPR_BehaviorEUTRA.Bytes, uint(ie.ModifiedMPR_BehaviorEUTRA.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode ModifiedMPR_BehaviorEUTRA", err)
		}
	}
	if ie.MultiNS_Pmax_EUTRA != nil {
		if err = ie.MultiNS_Pmax_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode MultiNS_Pmax_EUTRA", err)
		}
	}
	if ie.Rs_SINR_MeasEUTRA != nil {
		if err = ie.Rs_SINR_MeasEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode Rs_SINR_MeasEUTRA", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Ne_DC != nil, ie.Nr_HO_ToEN_DC_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap EUTRA_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ne_DC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ne_DC optional
			if ie.Ne_DC != nil {
				if err = ie.Ne_DC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ne_DC", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Nr_HO_ToEN_DC_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Nr_HO_ToEN_DC_r16 optional
			if ie.Nr_HO_ToEN_DC_r16 != nil {
				if err = ie.Nr_HO_ToEN_DC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_HO_ToEN_DC_r16", err)
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

func (ie *EUTRA_ParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Mfbi_EUTRAPresent bool
	if Mfbi_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ModifiedMPR_BehaviorEUTRAPresent bool
	if ModifiedMPR_BehaviorEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiNS_Pmax_EUTRAPresent bool
	if MultiNS_Pmax_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rs_SINR_MeasEUTRAPresent bool
	if Rs_SINR_MeasEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mfbi_EUTRAPresent {
		ie.Mfbi_EUTRA = new(EUTRA_ParametersCommon_mfbi_EUTRA)
		if err = ie.Mfbi_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode Mfbi_EUTRA", err)
		}
	}
	if ModifiedMPR_BehaviorEUTRAPresent {
		var tmp_bs_ModifiedMPR_BehaviorEUTRA []byte
		var n_ModifiedMPR_BehaviorEUTRA uint
		if tmp_bs_ModifiedMPR_BehaviorEUTRA, n_ModifiedMPR_BehaviorEUTRA, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode ModifiedMPR_BehaviorEUTRA", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_ModifiedMPR_BehaviorEUTRA,
			NumBits: uint64(n_ModifiedMPR_BehaviorEUTRA),
		}
		ie.ModifiedMPR_BehaviorEUTRA = &tmp_bitstring
	}
	if MultiNS_Pmax_EUTRAPresent {
		ie.MultiNS_Pmax_EUTRA = new(EUTRA_ParametersCommon_multiNS_Pmax_EUTRA)
		if err = ie.MultiNS_Pmax_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode MultiNS_Pmax_EUTRA", err)
		}
	}
	if Rs_SINR_MeasEUTRAPresent {
		ie.Rs_SINR_MeasEUTRA = new(EUTRA_ParametersCommon_rs_SINR_MeasEUTRA)
		if err = ie.Rs_SINR_MeasEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode Rs_SINR_MeasEUTRA", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			Ne_DCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ne_DC optional
			if Ne_DCPresent {
				ie.Ne_DC = new(EUTRA_ParametersCommon_ne_DC)
				if err = ie.Ne_DC.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ne_DC", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Nr_HO_ToEN_DC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Nr_HO_ToEN_DC_r16 optional
			if Nr_HO_ToEN_DC_r16Present {
				ie.Nr_HO_ToEN_DC_r16 = new(EUTRA_ParametersCommon_nr_HO_ToEN_DC_r16)
				if err = ie.Nr_HO_ToEN_DC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_HO_ToEN_DC_r16", err)
				}
			}
		}
	}
	return nil
}
