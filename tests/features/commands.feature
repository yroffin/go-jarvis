Feature: Command cruds
  In order to create some commands
  As a developer
  I want to use jarvis API

  Scenario: simple command
    Given I create a 'command' with body '{"name":"demo command"}'
    When I read last created 'command'
    Then the 'command' name must be 'demo command'

  Scenario Outline: much more complex stuff
    Given I create a 'command' with body <body>
    When I read last created 'command'
    Then the 'command' name must be <name>

    Examples:
      | body | name |
      | '{"name":"demo001 command"}' | 'demo001 command' |
      | '{"name":"demo002 command"}' | 'demo002 command' |
      | '{"name":"demo003 command"}' | 'demo003 command' |
      | '{"name":"demo004 command"}' | 'demo004 command' |
      | '{"name":"demo005 command"}' | 'demo005 command' |
      | '{"name":"demo006 command"}' | 'demo006 command' |
      | '{"name":"demo007 command"}' | 'demo007 command' |
      | '{"name":"demo008 command"}' | 'demo008 command' |
      | '{"name":"demo009 command"}' | 'demo009 command' |

  Scenario: simple command search
    Given I create a 'command' with body '{"name":"searchable command"}'
    When I search a 'command' with name 'searchable command'
    Then the 'command' name must be 'searchable command'
