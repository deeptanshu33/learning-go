package dictionarygo

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "updated"
		d := Dictionary{word: definition}

		err := d.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, newDefinition)
	})

	t.Run("word does not exist", func(t *testing.T) {
		word := "test"
		// definition := "this is just a test"
		newDefinition := "updated"
		d := Dictionary{}

		err := d.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "test definition"
		d := Dictionary{word: definition}

		err := d.Delete(word)
		assertError(t, err, nil)

		_, err = d.Search(word)

		assertError(t, err, ErrNotFound)
	})

	t.Run("non existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{}
		err := d.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertError(t testing.TB, got, expected error) {
	if got != expected {
		t.Errorf("got error %q, want error %q", got, expected)
	}
}

func assertStrings(t testing.TB, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q want %q given, %q", got, expected, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, definition)
}
