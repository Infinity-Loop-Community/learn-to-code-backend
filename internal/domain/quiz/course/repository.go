package course

import "errors"

var ErrCourseNotFound = errors.New("participant not found")

type Repository interface {
	FindByID(id string) (Course, error)
}
