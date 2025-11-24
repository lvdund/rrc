package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RadioLinkMonitoringConfig struct {
	FailureDetectionResourcesToAddModList  []RadioLinkMonitoringRS                                `lb:1,ub:maxNrofFailureDetectionResources,optional`
	FailureDetectionResourcesToReleaseList []RadioLinkMonitoringRS_Id                             `lb:1,ub:maxNrofFailureDetectionResources,optional`
	BeamFailureInstanceMaxCount            *RadioLinkMonitoringConfig_beamFailureInstanceMaxCount `optional`
	BeamFailureDetectionTimer              *RadioLinkMonitoringConfig_beamFailureDetectionTimer   `optional`
	Beamfailure_r17                        *BeamFailureDetection_r17                              `optional,ext-1`
}

func (ie *RadioLinkMonitoringConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Beamfailure_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.FailureDetectionResourcesToAddModList) > 0, len(ie.FailureDetectionResourcesToReleaseList) > 0, ie.BeamFailureInstanceMaxCount != nil, ie.BeamFailureDetectionTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.FailureDetectionResourcesToAddModList) > 0 {
		tmp_FailureDetectionResourcesToAddModList := utils.NewSequence[*RadioLinkMonitoringRS]([]*RadioLinkMonitoringRS{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		for _, i := range ie.FailureDetectionResourcesToAddModList {
			tmp_FailureDetectionResourcesToAddModList.Value = append(tmp_FailureDetectionResourcesToAddModList.Value, &i)
		}
		if err = tmp_FailureDetectionResourcesToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode FailureDetectionResourcesToAddModList", err)
		}
	}
	if len(ie.FailureDetectionResourcesToReleaseList) > 0 {
		tmp_FailureDetectionResourcesToReleaseList := utils.NewSequence[*RadioLinkMonitoringRS_Id]([]*RadioLinkMonitoringRS_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		for _, i := range ie.FailureDetectionResourcesToReleaseList {
			tmp_FailureDetectionResourcesToReleaseList.Value = append(tmp_FailureDetectionResourcesToReleaseList.Value, &i)
		}
		if err = tmp_FailureDetectionResourcesToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode FailureDetectionResourcesToReleaseList", err)
		}
	}
	if ie.BeamFailureInstanceMaxCount != nil {
		if err = ie.BeamFailureInstanceMaxCount.Encode(w); err != nil {
			return utils.WrapError("Encode BeamFailureInstanceMaxCount", err)
		}
	}
	if ie.BeamFailureDetectionTimer != nil {
		if err = ie.BeamFailureDetectionTimer.Encode(w); err != nil {
			return utils.WrapError("Encode BeamFailureDetectionTimer", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Beamfailure_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RadioLinkMonitoringConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Beamfailure_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Beamfailure_r17 optional
			if ie.Beamfailure_r17 != nil {
				if err = ie.Beamfailure_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Beamfailure_r17", err)
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

func (ie *RadioLinkMonitoringConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FailureDetectionResourcesToAddModListPresent bool
	if FailureDetectionResourcesToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FailureDetectionResourcesToReleaseListPresent bool
	if FailureDetectionResourcesToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamFailureInstanceMaxCountPresent bool
	if BeamFailureInstanceMaxCountPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamFailureDetectionTimerPresent bool
	if BeamFailureDetectionTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FailureDetectionResourcesToAddModListPresent {
		tmp_FailureDetectionResourcesToAddModList := utils.NewSequence[*RadioLinkMonitoringRS]([]*RadioLinkMonitoringRS{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		fn_FailureDetectionResourcesToAddModList := func() *RadioLinkMonitoringRS {
			return new(RadioLinkMonitoringRS)
		}
		if err = tmp_FailureDetectionResourcesToAddModList.Decode(r, fn_FailureDetectionResourcesToAddModList); err != nil {
			return utils.WrapError("Decode FailureDetectionResourcesToAddModList", err)
		}
		ie.FailureDetectionResourcesToAddModList = []RadioLinkMonitoringRS{}
		for _, i := range tmp_FailureDetectionResourcesToAddModList.Value {
			ie.FailureDetectionResourcesToAddModList = append(ie.FailureDetectionResourcesToAddModList, *i)
		}
	}
	if FailureDetectionResourcesToReleaseListPresent {
		tmp_FailureDetectionResourcesToReleaseList := utils.NewSequence[*RadioLinkMonitoringRS_Id]([]*RadioLinkMonitoringRS_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		fn_FailureDetectionResourcesToReleaseList := func() *RadioLinkMonitoringRS_Id {
			return new(RadioLinkMonitoringRS_Id)
		}
		if err = tmp_FailureDetectionResourcesToReleaseList.Decode(r, fn_FailureDetectionResourcesToReleaseList); err != nil {
			return utils.WrapError("Decode FailureDetectionResourcesToReleaseList", err)
		}
		ie.FailureDetectionResourcesToReleaseList = []RadioLinkMonitoringRS_Id{}
		for _, i := range tmp_FailureDetectionResourcesToReleaseList.Value {
			ie.FailureDetectionResourcesToReleaseList = append(ie.FailureDetectionResourcesToReleaseList, *i)
		}
	}
	if BeamFailureInstanceMaxCountPresent {
		ie.BeamFailureInstanceMaxCount = new(RadioLinkMonitoringConfig_beamFailureInstanceMaxCount)
		if err = ie.BeamFailureInstanceMaxCount.Decode(r); err != nil {
			return utils.WrapError("Decode BeamFailureInstanceMaxCount", err)
		}
	}
	if BeamFailureDetectionTimerPresent {
		ie.BeamFailureDetectionTimer = new(RadioLinkMonitoringConfig_beamFailureDetectionTimer)
		if err = ie.BeamFailureDetectionTimer.Decode(r); err != nil {
			return utils.WrapError("Decode BeamFailureDetectionTimer", err)
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

			Beamfailure_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Beamfailure_r17 optional
			if Beamfailure_r17Present {
				ie.Beamfailure_r17 = new(BeamFailureDetection_r17)
				if err = ie.Beamfailure_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Beamfailure_r17", err)
				}
			}
		}
	}
	return nil
}
