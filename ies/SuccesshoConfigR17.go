package ies

import "rrc/utils"

// SuccessHO-Config-r17 ::= SEQUENCE
// Extensible
type SuccesshoConfigR17 struct {
	Thresholdpercentaget304R17    *SuccesshoConfigR17Thresholdpercentaget304R17
	Thresholdpercentaget310R17    *SuccesshoConfigR17Thresholdpercentaget310R17
	Thresholdpercentaget312R17    *SuccesshoConfigR17Thresholdpercentaget312R17
	SourcedapsFailurereportingR17 *utils.BOOLEAN
}
