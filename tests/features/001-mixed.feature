Feature: Mixed cruds
  In order to create some commands
  As a developer
  I want to use jarvis API

  Scenario: simple mixed cruds
    Given I erase all 'device' with 'name' starting with 'mixed'
    Given I erase all 'plugin' with 'name' starting with 'mixed'
    Given I erase all 'command' with 'name' starting with 'mixed'
    Given I create a 'device' with body '{"name":"mixed device"}'
    Given I create a 'plugin' with body '{"name":"mixed plugin"}'
    Given I create a 'command' with body '{"name":"mixed command"}'
    When I search a 'device' with name 'mixed device'
    Then the 'device' name must be 'mixed device'
    When I search a 'plugin' with name 'mixed plugin'
    Then the 'plugin' name must be 'mixed plugin'
    When I search a 'command' with name 'mixed command'
    Then the 'command' name must be 'mixed command'
    Then the 'command' search 'name' = 'mixed command' count is 1
