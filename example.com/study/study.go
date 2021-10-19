package study

// study.go

import (
	"github.com/bigwhite/functrace"
	"github.com/pkg/errors"
)

var _ MStudy = (*study)(nil)

type MStudy interface {
	Listen(msg string) string
	Speak(msg string) string
	Read(msg string) string
	Write(msg string) string
}

type study struct {
	Name string
}

func (s *study) Listen(msg string) string {
	defer functrace.Trace()()
	return s.Name + " 听 " + msg
}

func (s *study) Speak(msg string) string {
	defer functrace.Trace()()
	return s.Name + " 说 " + msg
}

func (s *study) Read(msg string) string {
	defer functrace.Trace()()
	return s.Name + " 读 " + msg
}

func (s *study) Write(msg string) string {
	defer functrace.Trace()()
	return s.Name + " 写 " + msg
}

func New(name string) (MStudy, error) {
	defer functrace.Trace()()
	if name == "" {
		return nil, errors.New("name required")
	}

	return &study{
		Name: name,
	}, nil
}
