package inmemory

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/interfaces/lambda/course/responseobject"
	"learn-to-code/static"
)

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

// CourseRepository contains hardcoded data for now to validate the requirements and access patterns
type CourseRepository struct {
}

// frontend course
const CourseIDFrontendDevelopment = "ed86d338-84a0-4486-a314-b99b17175875"

const CourseStepIDEssentialsOfTheWeb = "c7486278-a50c-4629-89b9-cc1c74d7a538"
const QuizIDEssentialsOfTheWeb = "fcf7890f-9c72-46d3-931e-34494307be37"
const FirstQuestionID = "14c20d31-c7e1-416d-9c8e-1f2040141f0b"
const FirstAnswerID = "06a1956e-b659-493f-9533-b27733ddd7fe"
const FirstCorrectAnswerID = "48a293ee-7f43-4e3d-85d1-4737e6385c7c"

const CourseStepIDJavaScriptBasics = "123f372e-e176-41c8-ba8e-9fc406c9ad1e"
const QuizIDJavaScriptBasics = "e3ce1f8b-bb40-4bdd-b31b-33cbef24d267"

const CourseStepIDComputerScienceBasics = "6ed14552-fb4d-4e90-a300-c4a04b6197e4"
const QuizIDComputerScienceBasics = "addb1e53-eb07-44a3-881f-217eee3a926b"

const CourseStepIDJavaScriptAdvanced = "c8f398b8-a712-405e-ac95-f570ffe3a057"
const QuizIDJavaScriptAdvanced = "2819856b-ed81-4d48-92a3-acea534b3673"

const CourseStepIDGit = "99d51144-6a62-4c75-b19c-30508e2421c7"
const QuizIDGit = "9db0dd4a-7827-4306-92b5-02abf8706f4b"

const CourseStepIDTypeScript = "18d27e38-6336-4d00-929e-eceaebcd34a7"
const QuizIDTypeScript = "ad79ce21-c9bd-424a-9e4d-8136143a07b7"

const CourseStepIDNodeJS = "aa784d1b-6976-48fa-b9ba-f3e6bf6da585"
const QuizIDNodeJS = "b3e4b2a9-4759-46a1-9288-4a63a3a91e45"

const CourseStepIDWebpack = "36aac1a9-f786-49de-b910-629625c13355"
const QuizIDWebpack = "5bf61c9d-fa2a-4688-9930-781f5475f4c8"

const CourseStepIDTestingJest = "f7e4f315-d9e1-4d73-ae20-697369170c0e"
const QuizIDTestingJest = "f8668659-7e63-45dc-a87b-3c74872d1c74"

const CourseStepIDReact = "b38b83e4-136a-4e0f-a927-c96ab5104761"
const QuizIDReact = "4fb71fa7-33ad-4154-8e72-991033879c3e"

const CourseStepIDCIGithubActions = "f15e7c43-7b62-422e-bc5e-90c3163870f6"
const QuizIDCIGithubActions = "29308dff-0179-4f5b-b6ca-f29e11f17661"

func (q *CourseRepository) FindByID(courseID string) (course.Course, error) {

	if courseID == CourseIDFrontendDevelopment {
		return q.getFrontendCourse()
	}

	if courseID == CourseIDBackendDevelopment {
		return q.getBackendCourse()
	}

	return course.Course{}, course.ErrCourseNotFound
}

func (q *CourseRepository) getFrontendCourse() (course.Course, error) {
	quizEssentialsOfTheWeb, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDEssentialsOfTheWeb)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizJavaScriptBasics, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDJavaScriptBasics)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizComputerScience, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDComputerScienceBasics)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizAdvancedJavaScript, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDJavaScriptAdvanced)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizGit, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDGit)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizTypeScript, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDTypeScript)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizNodeJS, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDNodeJS)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizWebpack, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDWebpack)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizTestingJest, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDTestingJest)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizReact, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDReact)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizCIGithubActions, err := q.getQuiz(CourseIDFrontendDevelopment, QuizIDCIGithubActions)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	return course.Course{
		ID:   CourseIDFrontendDevelopment,
		Name: "Frontend Development",
		Steps: []course.Step{
			{
				ID:      CourseStepIDEssentialsOfTheWeb,
				Name:    "The essentials of the Web",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizEssentialsOfTheWeb)},
			},
			{
				ID:      CourseStepIDJavaScriptBasics,
				Name:    "JavaScript Basics",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizJavaScriptBasics)},
			},
			{
				ID:      CourseStepIDComputerScienceBasics,
				Name:    "Computer Science",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizComputerScience)},
			},
			{
				ID:      CourseStepIDJavaScriptAdvanced,
				Name:    "Advanced JavaScript",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizAdvancedJavaScript)},
			},
			{
				ID:      CourseStepIDGit,
				Name:    "Git",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizGit)},
			}, {
				ID:      CourseStepIDTypeScript,
				Name:    "TypeScript",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizTypeScript)},
			}, {
				ID:      CourseStepIDNodeJS,
				Name:    "NodeJS",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizNodeJS)},
			}, {
				ID:      CourseStepIDWebpack,
				Name:    "Webpack",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizWebpack)},
			}, {
				ID:      CourseStepIDTestingJest,
				Name:    "Testing with Jest",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizTestingJest)},
			}, {
				ID:      CourseStepIDReact,
				Name:    "React",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizReact)},
			}, {
				ID:      CourseStepIDCIGithubActions,
				Name:    "GitHub Actions",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizCIGithubActions)},
			},
		},
	}, nil
}

// backend course
const CourseIDBackendDevelopment = "dc64e1a8-2c83-4dca-b9fa-c5c99a11a773"

const CourseStepIDJavaBasics = "dbbed51d-845f-4084-a0c4-f5bea6502dd0"
const QuizIDJavaBasics = "ce2cae3e-5eb6-4293-afee-3e97bb5802ad"

const CourseStepIDCS = "859fe3e7-f05b-4e4d-afb3-f09f947cf244"
const QuizIDCS = "705d6ea3-4407-4818-acf6-4835168c648c"

const CourseStepIDGradle = "08d0fd0b-cb07-4e43-8b98-a6bed80ceecf"
const QuizIDGradle = "51c88734-2918-4928-83e8-0c80c7830cf8"

const CourseStepIDTesting = "247b2089-72dc-4d51-ad1d-dac3a70183f5"
const QuizIDTesting = "2dd320c6-e8b9-4d8d-8ebc-84afab4a5276"

const CourseStepIDTDD = "2656d99c-28f9-4b25-a842-7c4aa0c35329"
const QuizIDTDD = "fcb9218c-422e-42b5-94ab-abde486c3629"

const CourseStepIDSpringBoot = "adaa50f6-81a0-4c35-ada1-95b966e2fe61"
const QuizIDSpringBoot = "7577179a-4c59-4fcc-99a2-232132a65709"

const CourseStepIDDocker = "6ae1c0da-e46d-4c36-b5a1-2bafe9e020a7"
const QuizIDDocker = "492190b1-4123-4fcf-aeb8-dfaf2193309a"

const CourseStepIDGithubActions = "f4dfa8f8-1002-4cdb-91b9-7eda42edb04d"
const QuizIDGithubActions = "6c37bc46-05cf-4746-9c70-2a11545030fd"

const CourseStepIDAWS = "5d5b7116-9292-449f-a671-084d221b101c"
const QuizIDAWS = "c3b848ad-918d-4e41-b612-1a5920b9e499"

const CourseStepIDAdvancedJava = "98e54432-d7b6-43cc-b438-b0c57905a1c2"
const QuizIDAdvancedJava = "28d788d1-f8e8-4915-a48f-acd5eda4f1d5"

func (q *CourseRepository) getBackendCourse() (course.Course, error) {

	quizJavaBasics, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDJavaBasics)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizCS, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDCS)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizGit, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDGit)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizGradle, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDGradle)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizTesting, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDTesting)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizTDD, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDTDD)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizSpringBoot, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDSpringBoot)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizDocker, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDDocker)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizGithubActions, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDGithubActions)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizAWS, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDAWS)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizAdvancedJava, err := q.getQuiz(CourseIDBackendDevelopment, QuizIDAdvancedJava)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	return course.Course{
		ID:   CourseIDBackendDevelopment,
		Name: "Backend Development",
		Steps: []course.Step{
			{
				ID:      CourseStepIDJavaBasics,
				Name:    "Java Basics",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizJavaBasics)},
			}, {
				ID:      CourseStepIDCS,
				Name:    "Computer Science",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizCS)},
			}, {
				ID:      CourseStepIDGit,
				Name:    "Git",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizGit)},
			}, {
				ID:      CourseStepIDGradle,
				Name:    "Gradle",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizGradle)},
			}, {
				ID:      CourseStepIDTesting,
				Name:    "Testing",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizTesting)},
			}, {
				ID:      CourseStepIDTDD,
				Name:    "Test Driven Development",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizTDD)},
			}, {
				ID:      CourseStepIDSpringBoot,
				Name:    "Spring Boot",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizSpringBoot)},
			}, {
				ID:      CourseStepIDDocker,
				Name:    "Docker",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizDocker)},
			}, {
				ID:      CourseStepIDGithubActions,
				Name:    "Github Actions",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizGithubActions)},
			}, {
				ID:      CourseStepIDAWS,
				Name:    "AWS",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizAWS)},
			}, {
				ID:      CourseStepIDAdvancedJava,
				Name:    "Advanced Java",
				Quizzes: []course.StepQuiz{q.mapQuiz(quizAdvancedJava)},
			},
		},
	}, nil
}

func (q *CourseRepository) getQuiz(courseID string, quizID string) (responseobject.StepQuiz, error) {
	file, err := q.readQuizFromFile(courseID, quizID)
	if err != nil {
		return responseobject.StepQuiz{}, err
	}

	stepQuiz, err := q.parseJSON(file)
	if err != nil {
		return responseobject.StepQuiz{}, err
	}

	return stepQuiz, nil
}

func (q *CourseRepository) parseJSON(file []byte) (responseobject.StepQuiz, error) {
	stepQuiz := responseobject.StepQuiz{}
	err := json.Unmarshal(file, &stepQuiz)
	if err != nil {
		return responseobject.StepQuiz{}, err
	}

	return stepQuiz, nil
}

func (q *CourseRepository) readQuizFromFile(requestedCourseID string, requestedQuizID string) ([]byte, error) {
	file, err := fs.ReadFile(static.StaticJSONFiles, fmt.Sprintf("json/course/%s/quiz/%s.json", requestedCourseID, requestedQuizID))
	if err != nil {
		return nil, fmt.Errorf("Unknown course %s or quiz %s, err: %v", requestedCourseID, requestedQuizID, err)
	}

	return file, nil
}

func (q *CourseRepository) mapQuiz(quiz responseobject.StepQuiz) course.StepQuiz {
	return course.StepQuiz{
		ID:        quiz.ID,
		Questions: mapQuestions(quiz.Questions),
	}
}

func mapQuestions(responseObjectQuestions []responseobject.QuizQuestion) []course.QuizQuestion {
	questions := []course.QuizQuestion{}

	for _, responseObjectQuestion := range responseObjectQuestions {
		questions = append(questions, course.QuizQuestion{
			ID:          responseObjectQuestion.ID,
			Text:        responseObjectQuestion.Text,
			Difficulty:  responseObjectQuestion.Difficulty,
			Answers:     mapAnswers(responseObjectQuestion.Answers),
			Rating:      float64(responseObjectQuestion.Rating),
			RatingCount: responseObjectQuestion.RatingCount,
		})
	}

	return questions
}

func mapAnswers(responseObjectAnswers []responseobject.QuizAnswer) []course.QuizAnswer {
	answers := []course.QuizAnswer{}

	for _, responseObjectAnswer := range responseObjectAnswers {
		answers = append(answers, course.QuizAnswer{
			ID:          responseObjectAnswer.ID,
			Text:        responseObjectAnswer.Text,
			IsCorrect:   responseObjectAnswer.IsCorrect,
			Description: responseObjectAnswer.Description,
		})
	}

	return answers
}
