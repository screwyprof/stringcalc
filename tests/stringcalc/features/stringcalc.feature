Feature: Add strings
  In order to be happy
  As an awesome gopher
  I need to be able to Add strings

  Scenario: A few delimiters given with the input
    Given there is a {StringCalc}
    When I Add "//[*][%]\n1*2%3"
    Then the result should be 6