package types

const (
	StatusAccepted int64 = iota
	StatusWaitingForJudge
	StatusRunning
	StatusCompiling
	StatusCompileError
	StatusCompileTimeout
	StatusWrongAnswer
	StatusTimeLimitExceed
	StatusMemoryLimitExceed
	StatusOutputLimitExceed
	StatusSystemError
	StatusUnknownError
	StatusPresentationError
	StatusRuntimeError
	StatusJudgeError
	StatusExhaustedMatch

	NormalStatus = StatusCompileTimeout
)
