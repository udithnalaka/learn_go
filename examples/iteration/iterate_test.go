package iteration

import "testing"

func TestIterate(t *testing.T) {

	t.Run("Add", func(t *testing.T) {
		actual := Repeat("a")
		expected := "aaaaa"

		if actual != expected {
			t.Errorf("expected %q but was %q", expected, actual)
		}
	})
}
