package mapper

import (
	"learn-to-code/internal/domain/quiz/participant/projection/quizattemptdetail"
	responseobject "learn-to-code/internal/interfaces/lambda/participant/quiz/responseobject"
)

type QuizAttemptDetailMapper struct {
}

func NewQuizAttemptDetailMapper() *QuizAttemptDetailMapper {
	return &QuizAttemptDetailMapper{}
}

func (cm *QuizAttemptDetailMapper) EntityToResponseObject(qad quizattemptdetail.QuizAttemptDetail) responseobject.QuizAttemptDetail {

	return responseobject.QuizAttemptDetail{
		QuestionsWithAnswer: qad.QuestionsWithAnswer,
		AttemptStatus:       string(qad.AttemptStatus),
		AttemptID:           qad.AttemptID,
		AttemptResult:       mapAttemptResult(qad.AttemptResult),
	}
}

func mapAttemptResult(domainObject quizattemptdetail.AttemptResult) responseobject.AttemptResult {
	return responseobject.AttemptResult{
		Pass:                                    domainObject.Pass,
		QuestionCorrectRatio:                    domainObject.QuestionCorrectRatio,
		TimeTakenMins:                           domainObject.TimeTakenMins,
		ComparedToTimeAveragePercentage:         domainObject.ComparedToTimeAveragePercentage,
		ComparedToCorrectRatioLastTryPercentage: domainObject.ComparedToCorrectRatioLastTryPercentage,
	}
}
