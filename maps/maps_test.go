package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("search exist word", func(t *testing.T) {
		word := "test"
		description := "this is just a test"
		dictionary := Dictionary{word: description}
		got, err := dictionary.Search(word)
		assertSuccessSearchResult(t, description, got, err)
	})

	t.Run("search not exist word", func(t *testing.T) {
		word := "test"
		description := "this is just a test"
		dictionary := Dictionary{word: description}
		_, err := dictionary.Search("other")
		assertFailureSearchResult(t, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add not exist word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		description := "this is just a test"
		errAdd := dictionary.Add(word, description)

		got, errSearch := dictionary.Search("test")

		assertNilError(t, errAdd)
		assertSuccessSearchResult(t, description, got, errSearch)
	})

	t.Run("add exist word", func(t *testing.T) {
		word := "test"
		description := "this is just a test"
		dictionary := Dictionary{word: description}
		errAdd := dictionary.Add(word, description)

		assertError(t, ErrAlreadyExistWord, errAdd)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update exist word", func(t *testing.T) {
		word := "test"
		newDescription := "this is just a test"
		dictionary := Dictionary{word: "old description"}

		errUpdate := dictionary.Update(word, newDescription)

		got, errSearch := dictionary.Search("test")

		assertNilError(t, errUpdate)
		assertSuccessSearchResult(t, newDescription, got, errSearch)
	})

	t.Run("update not exist word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		description := "this is just a test"
		errUpdate := dictionary.Update(word, description)
		assertError(t, ErrNotFoundWord, errUpdate)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete exist word", func(t *testing.T) {
		word := "test"
		description := "this is just a test"
		dictionary := Dictionary{word: description}

		errDelete := dictionary.Delete(word)
		_, errSearch := dictionary.Search(word)

		assertNilError(t, errDelete)
		assertError(t, ErrNotFoundWord, errSearch)
	})

	t.Run("delete not exist word", func(t *testing.T) {
		word := "test"
		description := "this is just a test"
		dictionary := Dictionary{word: description}
		errDelete := dictionary.Delete("other")
		assertError(t, ErrNotFoundWord, errDelete)
	})
}

func assertError(t testing.TB, want, got error) {
	t.Helper()
	if !errors.Is(want, got) {
		t.Errorf("want err %v, got err %v", want, got)
	}
}

func assertNilError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("want nil error, got %v", got)
	}
}

func assertSuccessSearchResult(t testing.TB, want, got string, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("want nil error but got %v", err)
	}
	if got != want {
		t.Errorf("want search result %q but got %q", want, got)
	}
}

func assertFailureSearchResult(t testing.TB, err error) {
	t.Helper()
	if !errors.Is(err, ErrNotFoundWord) {
		t.Errorf("want ErrNotFoundWord err but got %v", err)
	}
}
