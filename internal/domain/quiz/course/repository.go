package course

import "errors"

var ErrCourseNotFound = errors.New("course not found")

type Repository interface {
	FindByID(id string) (Course, error)
}
