package mymaps_test

import (
	"testing"

	"mymaps"

	iz "github.com/matryer/is"
)

func TestSearch(t *testing.T) {
	is := iz.New(t)

	t.Run("valid value to lookup returns", func(t *testing.T) {
		dictionary := mymaps.Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("test")
		is.NoErr(err) // no error should be returned
		want := "this is just a test"
		is.Equal(got, want) // value found in map

	})
	t.Run("invalid value to lookup returns", func(t *testing.T) {
		dictionary := mymaps.Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("taco")
		is.Equal(err.Error(), "unable to find value in map") // error code for no matching value is 10
		is.Equal(got, "")                                    // no value should be in string when error message returns
	})
}
