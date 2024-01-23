package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("result if number  specified", func(t *testing.T) {
		result := Repeat("a", 3)
		expected := "aaa"
		assertCorrectMessage(t, result, expected)

	})
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}

}

func assertCorrectMessage(t testing.TB, got, want string) {

	t.Helper()
	if got != want {
		t.Errorf("Expected %q Recived %q", want, got)
	}
}
