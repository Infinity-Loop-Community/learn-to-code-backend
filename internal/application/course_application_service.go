package application

import (
	"learn-to-code/internal/domain/quiz/course"
)

func NewCourseApplicationService(courseRepository course.Repository) *CourseApplicationService {
	return &CourseApplicationService{
		courseRepository: courseRepository,
	}
}

type CourseApplicationService struct {
	courseRepository course.Repository
}

func (as *CourseApplicationService) GetCourseByID(id string) (course.Course, error) {
	return as.courseRepository.FindByID(id)
}
