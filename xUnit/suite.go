package xunit

type TestSuite struct {
	tests []Test
}

func NewTestSuite() *TestSuite {
	return &TestSuite{}
}

func (suite *TestSuite) Add(test Test) {
	suite.tests = append(suite.tests, test)
}

func (suite *TestSuite) Run(result *TestResult) {
	for _, test := range suite.tests {
		test.Run(test, result)
	}
}
