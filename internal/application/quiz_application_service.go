package application

import (
	"errors"
	"learn-to-code/internal/domain/quiz/participant"
)

type QuizApplicationService struct {
	participantRepository participant.Repository
}

func NewQuizApplicationService(participantRepository participant.Repository) *QuizApplicationService {
	return &QuizApplicationService{
		participantRepository: participantRepository,
	}
}

func (as *QuizApplicationService) GetStartedQuizCount(participantID string) (int, error) {
	p, err := as.participantRepository.FindByID(participantID)
	if errors.Is(err, participant.ErrNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return p.GetStartedQuizCount(), nil
}

func (as *QuizApplicationService) StartQuiz(participantID string, quizID string) error {
	p, err := as.createParticipantIfNotExists(participantID)
	if err != nil {
		return err
	}

	err = p.StartQuiz(quizID)
	if err != nil {
		return err
	}

	appendEventErr := as.participantRepository.AppendEvents(participantID, p.GetNewEventsAndUpdatePersistedVersion())

	return appendEventErr
}

func (as *QuizApplicationService) createParticipantIfNotExists(participantID string) (participant.Participant, error) {
	p, err := as.participantRepository.FindByID(participantID)
	if err != nil && !errors.Is(err, participant.ErrNotFound) {
		return participant.Participant{}, err
	}

	if err != nil && errors.Is(err, participant.ErrNotFound) {
		p, err = participant.NewWithID(participantID)
		if err != nil {
			return participant.Participant{}, err
		}
	}
	return p, nil
}

func (as *QuizApplicationService) GetFinishedQuizCount(participantID string) (int, error) {
	p, err := as.participantRepository.FindByID(participantID)
	if errors.Is(err, participant.ErrNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return p.GetFinishedQuizCount(), nil
}

func (as *QuizApplicationService) FinishQuiz(participantID string, quizID string) error {
	p, err := as.participantRepository.FindByID(participantID)
	if err != nil && !errors.Is(err, participant.ErrNotFound) {
		return err
	}

	err = p.FinishQuiz(quizID)
	if err != nil {
		return err
	}

	appendEventErr := as.participantRepository.AppendEvents(participantID, p.GetNewEventsAndUpdatePersistedVersion())

	return appendEventErr
}
