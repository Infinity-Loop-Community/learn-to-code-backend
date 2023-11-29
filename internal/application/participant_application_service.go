package application

import (
	"errors"
	"learn-to-code/internal/domain/command"
	"learn-to-code/internal/domain/quiz/participant"
)

type ParticipantApplicationService struct {
	participantRepository  participant.Repository
	startQuizToEventMapper *command.ParticipantCommandApplier
}

func NewPartcipantApplicationService(participantRepository participant.Repository, participantCommandApplier *command.ParticipantCommandApplier) *ParticipantApplicationService {
	return &ParticipantApplicationService{
		participantRepository:  participantRepository,
		startQuizToEventMapper: participantCommandApplier,
	}
}

func (as *ParticipantApplicationService) GetStartedQuizCount(participantID string) (int, error) {
	p, err := as.participantRepository.FindOrCreateByID(participantID)
	if err != nil {
		return 0, err
	}

	return p.GetStartedQuizCount(), nil
}

func (as *ParticipantApplicationService) ProcessCommand(commandDomainObject command.Command, participantID string) error {

	p, err := as.participantRepository.FindOrCreateByID(participantID)
	if err != nil {
		return err
	}

	p, err = as.startQuizToEventMapper.ApplyCommand(commandDomainObject, p)
	if err != nil {
		return err
	}

	appendEventErr := as.participantRepository.AppendEvents(p.GetID(), p.GetNewEventsAndUpdatePersistedVersion())

	return appendEventErr
}

func (as *ParticipantApplicationService) GetFinishedQuizCount(participantID string) (int, error) {
	p, err := as.participantRepository.FindOrCreateByID(participantID)
	if errors.Is(err, participant.ErrParticipantNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return p.GetFinishedQuizCount(), nil
}

func (as *ParticipantApplicationService) FinishQuiz(participantID string, quizID string) error {
	p, err := as.participantRepository.FindOrCreateByID(participantID)
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
