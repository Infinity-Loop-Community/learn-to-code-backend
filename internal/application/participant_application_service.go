package application

import (
	"fmt"
	"learn-to-code/internal/domain/command"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/projection"
	"learn-to-code/internal/domain/quiz/participant/projection/quizattemptdetail"
	"strconv"
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

	appendEventErr := as.participantRepository.StoreEvents(p.GetID(), p.GetNewEventsAndUpdatePersistedVersion())

	return appendEventErr
}

func (as *ParticipantApplicationService) GetQuizzes(participantID string) (projection.QuizOverview, error) {
	p, err := as.participantRepository.FindOrCreateByID(participantID)
	if err != nil {
		return projection.QuizOverview{}, err
	}

	quizOverview, err := projection.NewQuizOverview(p)

	return quizOverview, err
}

func (as *ParticipantApplicationService) GetQuizAttemptDetail(participantID string, quizID string, attemptID string) (quizattemptdetail.QuizAttemptDetail, error) {
	p, err := as.participantRepository.FindOrCreateByID(participantID)
	if err != nil {
		return quizattemptdetail.QuizAttemptDetail{}, err
	}

	var attemptIDNumber int
	if attemptID == "latest" {
		attemptIDNumber = p.GetQuizAttemptCount(quizID)
	} else {
		attemptIDInt64, err := strconv.ParseInt(attemptID, 10, 0)
		if err != nil {
			return quizattemptdetail.QuizAttemptDetail{}, err
		}
		attemptIDNumber = int(attemptIDInt64)
	}

	quizOverview, err := quizattemptdetail.NewQuizAttemptDetail(p, quizID, attemptIDNumber)

	return quizOverview, err
}

func (as *ParticipantApplicationService) GetLatestQuizAttemptDetail(participantID string, quizID string) (quizattemptdetail.QuizAttemptDetail, error) {
	p, err := as.participantRepository.FindOrCreateByID(participantID)
	if err != nil {
		return quizattemptdetail.QuizAttemptDetail{}, err
	}

	quizOverview, err := projection.NewQuizOverview(p)
	if err != nil {
		return quizattemptdetail.QuizAttemptDetail{}, err
	}

	quizAttemptCount := len(quizOverview.FinishedQuizzes[quizID])

	if quizAttemptCount == 0 {
		return quizattemptdetail.QuizAttemptDetail{}, fmt.Errorf("no quiz with id %s finished yet", quizID)
	}

	latestAttempt := quizOverview.FinishedQuizzes[quizID][quizAttemptCount-1]

	attemptDetail, err := quizattemptdetail.NewQuizAttemptDetail(p, quizID, latestAttempt.AttemptID)
	if err != nil {
		return quizattemptdetail.QuizAttemptDetail{}, err
	}

	return attemptDetail, nil
}
