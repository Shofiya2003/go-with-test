package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"chris"},
			[]string{"chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name1 string
				Name2 string
			}{"chris", "London"},
			[]string{"chris", "London"},
		},
		{
			"struct with no string fields",
			struct {
				Name string
				Age  int
			}{"chris", 33},
			[]string{"chris"},
		},
		{
			"nested fields",
			Person{
				"chris",
				Profile{22, "London"},
			},
			[]string{"chris", "London"},
		},
		{
			"pointers to the things",
			&Person{
				"chris",
				Profile{22, "London"},
			},
			[]string{"chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{34, "London"},
				{22, "Vadodara"},
			},
			[]string{"London", "Vadodara"},
		},
		{
			"arrays",
			[2]Profile{
				{34, "London"},
				{22, "Vadodara"},
			},
			[]string{"London", "Vadodara"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)

			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		testMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		Walk(testMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")

	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{33, "London"}
			ch <- Profile{22, "Vadodara"}
			close(ch)
		}()

		var got []string

		Walk(ch, func(input string) {
			got = append(got, input)
		})

		want := []string{"London", "Vadodara"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v ", got, want)
		}

	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "London"}, Profile{22, "Vadodara"}
		}

		var got []string
		want := []string{"London", "Vadodara"}

		Walk(aFunction, func(intput string) {
			got = append(got, intput)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it did not", haystack, needle)
	}
}
