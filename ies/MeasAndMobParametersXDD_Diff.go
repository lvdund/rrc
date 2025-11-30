package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersXDD_Diff struct {
	IntraAndInterF_MeasAndReport *MeasAndMobParametersXDD_Diff_intraAndInterF_MeasAndReport `optional`
	EventA_MeasAndReport         *MeasAndMobParametersXDD_Diff_eventA_MeasAndReport         `optional`
	HandoverInterF               *MeasAndMobParametersXDD_Diff_handoverInterF               `optional,ext-1`
	HandoverLTE_EPC              *MeasAndMobParametersXDD_Diff_handoverLTE_EPC              `optional,ext-1`
	HandoverLTE_5GC              *MeasAndMobParametersXDD_Diff_handoverLTE_5GC              `optional,ext-1`
	Sftd_MeasNR_Neigh            *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh            `optional,ext-2`
	Sftd_MeasNR_Neigh_DRX        *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX        `optional,ext-2`
	Dummy                        *MeasAndMobParametersXDD_Diff_dummy                        `optional,ext-3`
}

func (ie *MeasAndMobParametersXDD_Diff) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.HandoverInterF != nil || ie.HandoverLTE_EPC != nil || ie.HandoverLTE_5GC != nil || ie.Sftd_MeasNR_Neigh != nil || ie.Sftd_MeasNR_Neigh_DRX != nil || ie.Dummy != nil
	preambleBits := []bool{hasExtensions, ie.IntraAndInterF_MeasAndReport != nil, ie.EventA_MeasAndReport != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IntraAndInterF_MeasAndReport != nil {
		if err = ie.IntraAndInterF_MeasAndReport.Encode(w); err != nil {
			return utils.WrapError("Encode IntraAndInterF_MeasAndReport", err)
		}
	}
	if ie.EventA_MeasAndReport != nil {
		if err = ie.EventA_MeasAndReport.Encode(w); err != nil {
			return utils.WrapError("Encode EventA_MeasAndReport", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.HandoverInterF != nil || ie.HandoverLTE_EPC != nil || ie.HandoverLTE_5GC != nil, ie.Sftd_MeasNR_Neigh != nil || ie.Sftd_MeasNR_Neigh_DRX != nil, ie.Dummy != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasAndMobParametersXDD_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.HandoverInterF != nil, ie.HandoverLTE_EPC != nil, ie.HandoverLTE_5GC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode HandoverInterF optional
			if ie.HandoverInterF != nil {
				if err = ie.HandoverInterF.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverInterF", err)
				}
			}
			// encode HandoverLTE_EPC optional
			if ie.HandoverLTE_EPC != nil {
				if err = ie.HandoverLTE_EPC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverLTE_EPC", err)
				}
			}
			// encode HandoverLTE_5GC optional
			if ie.HandoverLTE_5GC != nil {
				if err = ie.HandoverLTE_5GC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverLTE_5GC", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Sftd_MeasNR_Neigh != nil, ie.Sftd_MeasNR_Neigh_DRX != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sftd_MeasNR_Neigh optional
			if ie.Sftd_MeasNR_Neigh != nil {
				if err = ie.Sftd_MeasNR_Neigh.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sftd_MeasNR_Neigh", err)
				}
			}
			// encode Sftd_MeasNR_Neigh_DRX optional
			if ie.Sftd_MeasNR_Neigh_DRX != nil {
				if err = ie.Sftd_MeasNR_Neigh_DRX.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sftd_MeasNR_Neigh_DRX", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.Dummy != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dummy optional
			if ie.Dummy != nil {
				if err = ie.Dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy", err)
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

func (ie *MeasAndMobParametersXDD_Diff) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraAndInterF_MeasAndReportPresent bool
	if IntraAndInterF_MeasAndReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var EventA_MeasAndReportPresent bool
	if EventA_MeasAndReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if IntraAndInterF_MeasAndReportPresent {
		ie.IntraAndInterF_MeasAndReport = new(MeasAndMobParametersXDD_Diff_intraAndInterF_MeasAndReport)
		if err = ie.IntraAndInterF_MeasAndReport.Decode(r); err != nil {
			return utils.WrapError("Decode IntraAndInterF_MeasAndReport", err)
		}
	}
	if EventA_MeasAndReportPresent {
		ie.EventA_MeasAndReport = new(MeasAndMobParametersXDD_Diff_eventA_MeasAndReport)
		if err = ie.EventA_MeasAndReport.Decode(r); err != nil {
			return utils.WrapError("Decode EventA_MeasAndReport", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			HandoverInterFPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverLTE_EPCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverLTE_5GCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode HandoverInterF optional
			if HandoverInterFPresent {
				ie.HandoverInterF = new(MeasAndMobParametersXDD_Diff_handoverInterF)
				if err = ie.HandoverInterF.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverInterF", err)
				}
			}
			// decode HandoverLTE_EPC optional
			if HandoverLTE_EPCPresent {
				ie.HandoverLTE_EPC = new(MeasAndMobParametersXDD_Diff_handoverLTE_EPC)
				if err = ie.HandoverLTE_EPC.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverLTE_EPC", err)
				}
			}
			// decode HandoverLTE_5GC optional
			if HandoverLTE_5GCPresent {
				ie.HandoverLTE_5GC = new(MeasAndMobParametersXDD_Diff_handoverLTE_5GC)
				if err = ie.HandoverLTE_5GC.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverLTE_5GC", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sftd_MeasNR_NeighPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sftd_MeasNR_Neigh_DRXPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sftd_MeasNR_Neigh optional
			if Sftd_MeasNR_NeighPresent {
				ie.Sftd_MeasNR_Neigh = new(MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh)
				if err = ie.Sftd_MeasNR_Neigh.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sftd_MeasNR_Neigh", err)
				}
			}
			// decode Sftd_MeasNR_Neigh_DRX optional
			if Sftd_MeasNR_Neigh_DRXPresent {
				ie.Sftd_MeasNR_Neigh_DRX = new(MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX)
				if err = ie.Sftd_MeasNR_Neigh_DRX.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sftd_MeasNR_Neigh_DRX", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			DummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dummy optional
			if DummyPresent {
				ie.Dummy = new(MeasAndMobParametersXDD_Diff_dummy)
				if err = ie.Dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dummy", err)
				}
			}
		}
	}
	return nil
}
