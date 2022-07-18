package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type MyFirstSuite struct {
	suite.Suite
}

// Triggers once (before suite run)
func (s *MyFirstSuite) BeforeAll(t provider.T) {
}

// Triggers once (after all tests finished)
func (s *MyFirstSuite) AfterAll(t provider.T) {
}

// Triggers each time before test run
func (s *MyFirstSuite) BeforeEach(t provider.T) {
	t.Epic("My Epic")
	t.Feature("My Feature")
}

// Triggers each time after test finished
func (s *MyFirstSuite) AfterEach(t provider.T) {
}

func (s *MyFirstSuite) TestMyFirstTest(t provider.T) {
	test := "test"
	t.Parallel()
	t.WithNewStep("My First Step", func(sCtx provider.StepCtx) {
		sCtx.Require().NotNil(test)
		sCtx.Require().Equal(test, "test")
	}, allure.NewParameter("time", time.Now()))
}

func (s *MyFirstSuite) TestXSkip(t provider.T) {
	var test string
	t.XSkip()
	t.Require().Equal("test", test)
}

func (s *MyFirstSuite) TestMySecondTest(t provider.T) {
	t.Parallel()
	test := "test"
	for idx, text := range []string{"test0", "test1", "test2"} {
		t.Run(text, func(t provider.T) {
			testText := text
			data := fmt.Sprintf("%s%d", test, idx)
			t.Parallel()
			t.WithNewStep("My First Step", func(sCtx provider.StepCtx) {
				sCtx.Require().NotNil(testText)
				sCtx.Require().Equal(data, testText)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}

func TestSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(MyFirstSuite))
}
