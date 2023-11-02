package application

import (
	"errors"
	"learn-to-code/internal/domain/quiz/participant"
)

type ParticipantApplicationService struct {
	participantRepository participant.Repository
}

func NewPartcipantApplicationService(participantRepository participant.Repository) *ParticipantApplicationService {
	return &ParticipantApplicationService{
		participantRepository: participantRepository,
	}
}

func (as *ParticipantApplicationService) GetStartedQuizCount(participantID string) (int, error) {
	p, err := as.participantRepository.FindByID(participantID)
	if errors.Is(err, participant.ErrParticipantNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return p.GetStartedQuizCount(), nil
}

func (as *ParticipantApplicationService) StartQuiz(participantID string, quizID string) error {
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

func (as *ParticipantApplicationService) createParticipantIfNotExists(participantID string) (participant.Participant, error) {
	p, err := as.participantRepository.FindByID(participantID)
	if err != nil && !errors.Is(err, participant.ErrParticipantNotFound) {
		return participant.Participant{}, err
	}

	if err != nil && errors.Is(err, participant.ErrParticipantNotFound) {
		p, err = participant.NewWithID(participantID)
		if err != nil {
			return participant.Participant{}, err
		}
	}
	return p, nil
}

func (as *ParticipantApplicationService) GetFinishedQuizCount(participantID string) (int, error) {
	p, err := as.participantRepository.FindByID(participantID)
	if errors.Is(err, participant.ErrParticipantNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return p.GetFinishedQuizCount(), nil
}

func (as *ParticipantApplicationService) FinishQuiz(participantID string, quizID string) error {
	p, err := as.participantRepository.FindByID(participantID)
	if err != nil && !errors.Is(err, participant.ErrParticipantNotFound) {
		return err
	}

	err = p.FinishQuiz(quizID)
	if err != nil {
		return err
	}

	appendEventErr := as.participantRepository.AppendEvents(participantID, p.GetNewEventsAndUpdatePersistedVersion())

	return appendEventErr
}
