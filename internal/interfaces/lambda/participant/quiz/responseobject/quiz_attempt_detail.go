package responseobject

type AttemptResult struct {
	Pass                                    bool    `json:"pass"`
	QuestionCorrectRatio                    float64 `json:"questionCorrectRatio"`
	TimeTakenMins                           int     `json:"timeTakenMins"`
	ComparedToTimeAveragePercentage         int     `json:"comparedToTimeAveragePercentage"`
	ComparedToCorrectRatioLastTryPercentage int     `json:"comparedToCorrectRatioLastTryPercentage"`
}

type QuizAttemptDetail struct {
	QuestionsWithAnswer map[string]string `json:"questionsWithAnswer"`
	AttemptStatus       string            `json:"attemptStatus"`
	AttemptID           int               `json:"attemptId"`
	AttemptResult       AttemptResult     `json:"attemptResult"`
}
