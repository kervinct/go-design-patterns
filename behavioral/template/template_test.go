package behavioral

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplateExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{&TemplateImpl{}}
		res := s.ExecuteAlgorithm(s)

		expectedOrError(res, " world ", t)
	})

	t.Run("Using anonymous functions", func(t *testing.T) {
		m := new(AnonymousTemplate)

		res := m.ExecuteAlgorithm(func() string {
			return "world"
		})

		expectedOrError(res, " world ", t)
	})

	t.Run("Using anonymous functions adpated to an interface", func(t *testing.T) {
		messageRetriever := MessageRetrieverAdapter(func() string {
			return "world"
		})

		if messageRetriever == nil {
			t.Fatal("Can not continue with a nil MessageRetriever")
		}

		template := &TemplateImpl{}
		res := template.ExecuteAlgorithm(messageRetriever)

		expectedOrError(res, " world ", t)
	})
}

func expectedOrError(res string, expected string, t *testing.T) {
	if !strings.Contains(res, expected) {
		t.Errorf("Expceted string '%s' was not found on returned string", expected)
	}
}
