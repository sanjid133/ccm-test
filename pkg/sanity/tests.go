package sanity

import (
	. "github.com/onsi/ginkgo"
)

type test struct {
	text string
	body func(*SanityContext)
}

var tests []test

// DescribeSanity must be used instead of the usual Ginkgo Describe to
// register a test block. The difference is that the body function
// will be called multiple times with the right context (when
// setting up a Ginkgo suite or a testing.T test, with the right
// configuration).
func DescribeSanity(text string, body func(*SanityContext)) bool {
	tests = append(tests, test{text, body})
	return true
}

// registerTestsInGinkgo invokes the actual Gingko Describe
// for the tests registered earlier with DescribeSanity.
func registerTestsInGinkgo(sc *SanityContext) {
	for _, test := range tests {
		Describe(test.text, func() {
			BeforeEach(func() {
			//	sc.setup()
			})

			test.body(sc)

			AfterEach(func() {
			//	sc.teardown()
			})
		})
	}
}
