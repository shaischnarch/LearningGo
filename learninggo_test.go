package learninggo_test

import (
	"testing"

	learninggo "github.com/shaischnarch/LearningGo"
)

func TestAdd(t *testing.T) {
	exepected := 5
	actual := learninggo.Add(3, 2)
	if exepected != actual {
		t.Errorf("Expected value: %d, Got value: %d", exepected, actual)
	}
}

func TestAddFloat(t *testing.T) {
	exepected := 5.0
	actual := learninggo.Add(3.0, 2)
	if exepected != actual {
		t.Errorf("Expected value: %f, Got value: %f", exepected, actual)
	}
}