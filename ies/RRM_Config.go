package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRM_Config struct {
	Ue_InactiveTime               *RRM_Config_ue_InactiveTime      `optional`
	CandidateCellInfoList         *MeasResultList2NR               `optional`
	CandidateCellInfoListSN_EUTRA *MeasResultServFreqListEUTRA_SCG `optional,ext-1`
}

func (ie *RRM_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.CandidateCellInfoListSN_EUTRA != nil
	preambleBits := []bool{hasExtensions, ie.Ue_InactiveTime != nil, ie.CandidateCellInfoList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ue_InactiveTime != nil {
		if err = ie.Ue_InactiveTime.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_InactiveTime", err)
		}
	}
	if ie.CandidateCellInfoList != nil {
		if err = ie.CandidateCellInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateCellInfoList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.CandidateCellInfoListSN_EUTRA != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RRM_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.CandidateCellInfoListSN_EUTRA != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CandidateCellInfoListSN_EUTRA optional
			if ie.CandidateCellInfoListSN_EUTRA != nil {
				if err = ie.CandidateCellInfoListSN_EUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CandidateCellInfoListSN_EUTRA", err)
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

func (ie *RRM_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_InactiveTimePresent bool
	if Ue_InactiveTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateCellInfoListPresent bool
	if CandidateCellInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ue_InactiveTimePresent {
		ie.Ue_InactiveTime = new(RRM_Config_ue_InactiveTime)
		if err = ie.Ue_InactiveTime.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_InactiveTime", err)
		}
	}
	if CandidateCellInfoListPresent {
		ie.CandidateCellInfoList = new(MeasResultList2NR)
		if err = ie.CandidateCellInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode CandidateCellInfoList", err)
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

			CandidateCellInfoListSN_EUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CandidateCellInfoListSN_EUTRA optional
			if CandidateCellInfoListSN_EUTRAPresent {
				ie.CandidateCellInfoListSN_EUTRA = new(MeasResultServFreqListEUTRA_SCG)
				if err = ie.CandidateCellInfoListSN_EUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode CandidateCellInfoListSN_EUTRA", err)
				}
			}
		}
	}
	return nil
}
