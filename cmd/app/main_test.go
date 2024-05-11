package main

import (
	"p2/constant"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestAlgo(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) TestTC1() {
	expect := 2
	actual := howManyChickenAreRescue(constant.TC1)
	t.Equal(expect, actual)
}
func (t *TestSuite) TestTC2() {
	expect := 4
	actual := howManyChickenAreRescue(constant.TC2)
	t.Equal(expect, actual)
}
func (t *TestSuite) TestTC3() {
	expect := 9
	actual := howManyChickenAreRescue(constant.TC3)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestTC4() {
	expect := 10
	actual := howManyChickenAreRescue(constant.TC4)
	t.Equal(expect, actual)
}
func (t *TestSuite) TestTC5() {
	expect := 0
	actual := howManyChickenAreRescue(constant.TC5)
	t.Equal(expect, actual)
}
