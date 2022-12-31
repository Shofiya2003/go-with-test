package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("search a word present in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing stuff"}
		got, err := dictionary.Search("test")
		want := "testing stuff"
		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("word not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("build")
		assertError(t, err, WordNotFoundError)
	})

}

func TestAddWord(t *testing.T) {

	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.AddWord("build", "building stuff")
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "build", "building stuff")
	})

	t.Run("word exists", func(t *testing.T) {
		word := "build"
		definition := "building stuff"
		dictionary := Dictionary{word: definition}
		err := dictionary.AddWord(word, definition)
		assertError(t, err, WordAlreadyExistsError)

	})

}

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		word := "test"
		definition := "testing stuff"
		dictionary := Dictionary{word: definition}
		newDefinition := "new def test"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("word does not exist", func(t *testing.T) {
		word := "test"
		definition := "testing stuff"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, WordNotFoundError)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "testing stuff"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != WordNotFoundError {
		t.Errorf("Expected %q to be deleted ", word)
	}
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q given %q want %q", got, "test", want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word string, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)
}
