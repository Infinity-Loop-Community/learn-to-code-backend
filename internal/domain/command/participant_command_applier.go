package command

import (
	"fmt"
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/infrastructure/inmemory"
)

type ParticipantCommandApplier struct {
	courseRepository course.Repository
}

func NewParticipantCommandApplier(courseRepository course.Repository) *ParticipantCommandApplier {
	return &ParticipantCommandApplier{
		courseRepository: courseRepository,
	}
}

func (m *ParticipantCommandApplier) ApplyCommand(c Command, p participant.Participant) (participant.Participant, error) {
	var courses map[string]course.Course = map[string]course.Course{}

	switch c.Type {
	case StartQuizCommandType:
		startQuiz, err := DecodeCommand(c.Data, &StartQuiz{})
		if err != nil {
			return participant.Participant{}, err
		}

		err = p.StartQuiz(startQuiz.QuizID, startQuiz.RequiredQuestionsAnswered)
		if err != nil {
			return participant.Participant{}, err
		}

	case SelectAnswerCommandType:
		selectAnswerData, err := DecodeCommand(c.Data, &SelectAnswer{})
		if err != nil {
			return participant.Participant{}, err
		}

		_, ok := courses[selectAnswerData.QuizID]
		if !ok {
			c, err := m.courseRepository.FindByID(inmemory.CourseIDFrontendDevelopment)
			if err != nil {
				return participant.Participant{}, err
			}
			courses[selectAnswerData.QuizID] = c
		}

		isAnswerCorrect := m.isAnswerCorrect(courses, selectAnswerData)

		selectQuizErr := p.SelectQuizAnswer(selectAnswerData.QuizID, selectAnswerData.QuestionID, selectAnswerData.AnswerID, isAnswerCorrect)
		if selectQuizErr != nil {
			return participant.Participant{}, err
		}

	case FinishQuizCommandType:
		finishQuizData, err := DecodeCommand(c.Data, &FinishQuiz{})
		if err != nil {
			return participant.Participant{}, err
		}

		err = p.FinishQuiz(finishQuizData.QuizID)
		if err != nil {
			return participant.Participant{}, err
		}

	default:
		return p, fmt.Errorf("unknown c type '%s'", c.Type)
	}

	return p, nil
}

func (m *ParticipantCommandApplier) isAnswerCorrect(courses map[string]course.Course, selectAnswerData *SelectAnswer) bool {
	var isAnswerCorrect bool
	for _, step := range courses[selectAnswerData.QuizID].Steps {
		for _, quiz := range step.Quizzes {
			if quiz.ID == selectAnswerData.QuizID {
				for _, question := range quiz.Questions {
					if question.ID == selectAnswerData.QuestionID {
						for _, answer := range question.Answers {
							if answer.ID == selectAnswerData.AnswerID {
								isAnswerCorrect = answer.IsCorrect
							}
						}
					}
				}
			}
		}
	}
	return isAnswerCorrect
}
