package ies

// MAC-ParametersFRX-Diff-r16 ::= SEQUENCE
// Extensible
type MacParametersfrxDiffR16 struct {
	DirectmcgScellactivationR16       *MacParametersfrxDiffR16DirectmcgScellactivationR16
	DirectmcgScellactivationresumeR16 *MacParametersfrxDiffR16DirectmcgScellactivationresumeR16
	DirectscgScellactivationR16       *MacParametersfrxDiffR16DirectscgScellactivationR16
	DirectscgScellactivationresumeR16 *MacParametersfrxDiffR16DirectscgScellactivationresumeR16
	DrxAdaptationR16                  *MacParametersfrxDiffR16DrxAdaptationR16
}
