package command

import (
	"fmt"
	"learn-to-code/internal/domain/command/data"
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/infrastructure/inmemory"

	"github.com/mitchellh/mapstructure"
)

type ParticipantCommandApplier struct {
	courseRepository course.Repository
}

func NewParticipantCommandApplier(courseRepository course.Repository) *ParticipantCommandApplier {
	return &ParticipantCommandApplier{
		courseRepository: courseRepository,
	}
}

func (m *ParticipantCommandApplier) ApplyCommand(command Command, p participant.Participant) (participant.Participant, error) {
	var err error

	var courses map[string]course.Course = map[string]course.Course{}

	switch command.Type {
	case data.StartQuizCommandType:
		startQuiz := data.StartQuiz{}
		err = m.decode(command.Data, &startQuiz)
		if err != nil {
			return participant.Participant{}, err
		}

		err = p.StartQuiz(startQuiz.QuizID, startQuiz.RequiredQuestionsAnswered)

	case data.SelectAnswerCommandType:
		selectAnswerData := data.SelectAnswer{}

		err := m.decode(command.Data, &selectAnswerData)
		if err != nil {
			return participant.Participant{}, err
		}

		_, ok := courses[selectAnswerData.QuizID]
		if !ok {
			c, err := m.courseRepository.FindByID(inmemory.CourseID)
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

	case data.FinishQuizCommandType:
		finishQuizData := data.FinishQuiz{}
		err = m.decode(command.Data, &finishQuizData)
		if err != nil {
			return participant.Participant{}, err
		}

		err = p.FinishQuiz(finishQuizData.QuizID)

	default:
		return p, fmt.Errorf("unknown command type '%s'", command.Type)
	}

	return p, err
}

func (m *ParticipantCommandApplier) isAnswerCorrect(courses map[string]course.Course, selectAnswerData data.SelectAnswer) bool {
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

func (m *ParticipantCommandApplier) decode(data any, result any) error {
	config := &mapstructure.DecoderConfig{
		ErrorUnset: true,
		Result:     result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(data)

	return err
}
