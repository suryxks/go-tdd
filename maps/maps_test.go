package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrorNotFound
		assertError(t, err, want)
	})

}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add("test", "this is just a test")
		want := "this is just a test"
		assertError(t, err, nil)
		asserDefinition(t, dictionary, "test", want)

	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add("test", "this is just a test")
		want := "this is just a test"
		assertError(t, err, ErrWordExists)
		asserDefinition(t, dictionary, "test", want)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		asserDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new  word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExists)
	})
}
func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrorNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {

		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {

		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func asserDefinition(t testing.TB, dictionary Dictionary, word, defintion string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, defintion)
}
