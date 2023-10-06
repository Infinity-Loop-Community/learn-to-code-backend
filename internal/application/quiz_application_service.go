package application

import (
	"errors"
	"hello-world/internal/domain/quiz/participant"
)

type QuizApplicationService struct {
	participantRepository participant.Repository
}

func NewQuizApplicationService(participantRepository participant.Repository) *QuizApplicationService {
	return &QuizApplicationService{
		participantRepository: participantRepository,
	}
}

func (as *QuizApplicationService) name() {

}

func (as *QuizApplicationService) GetStartedQuizCount(participantId string) (int, error) {
	p, err := as.participantRepository.FindById(participantId)
	if errors.Is(err, participant.ErrNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return p.GetStartedQuizCount(), nil
}

// StartQuiz is the first action of a participants in the quiz bounded context, hence if not created yet
// a first event for a participant is created, and with that the participant itself.
func (as *QuizApplicationService) StartQuiz(participantId string, quizId string) error {
	p, err := as.createParticipantIfNotExists(participantId)

	err = p.StartQuiz(quizId)
	if err != nil {
		return err
	}

	appendEventErr := as.participantRepository.AppendEvents(participantId, p.GetNewEventsAndUpdatePersistedVersion())

	return appendEventErr
}

func (as *QuizApplicationService) createParticipantIfNotExists(participantId string) (participant.Participant, error) {
	p, err := as.participantRepository.FindById(participantId)
	if err != nil && !errors.Is(err, participant.ErrNotFound) {
		return participant.Participant{}, err
	}

	if err != nil && errors.Is(err, participant.ErrNotFound) {
		p, err = participant.NewWithId(participantId)
		if err != nil {
			return participant.Participant{}, err
		}
	}
	return p, nil
}

func (as *QuizApplicationService) GetFinishedQuizCount(participantId string) (int, error) {
	p, err := as.participantRepository.FindById(participantId)
	if errors.Is(err, participant.ErrNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return p.GetFinishedQuizCount(), nil
}

func (as *QuizApplicationService) FinishQuiz(participantId string, quizId string) error {
	p, err := as.participantRepository.FindById(participantId)
	if err != nil && !errors.Is(err, participant.ErrNotFound) {
		return err
	}

	err = p.FinishQuiz(quizId)
	if err != nil {
		return err
	}

	appendEventErr := as.participantRepository.AppendEvents(participantId, p.GetNewEventsAndUpdatePersistedVersion())

	return appendEventErr
}
