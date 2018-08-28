Feature: Plugin cruds
  In order to create some plugins
  As a developer
  I want to use jarvis API

  Scenario: simple plugin
    Given I create a 'plugin' with body '{"name":"demo plugin"}'
    When I read last created 'plugin'
    Then the 'plugin' name must be 'demo plugin'

  Scenario Outline: much more complex stuff
    Given I create a 'plugin' with body <body>
    When I read last created 'plugin'
    Then the 'plugin' name must be <name>

    Examples:
      | body | name |
      | '{"name":"demo001 plugin"}' | 'demo001 plugin' |
      | '{"name":"demo002 plugin"}' | 'demo002 plugin' |
      | '{"name":"demo003 plugin"}' | 'demo003 plugin' |
      | '{"name":"demo004 plugin"}' | 'demo004 plugin' |
      | '{"name":"demo005 plugin"}' | 'demo005 plugin' |
      | '{"name":"demo006 plugin"}' | 'demo006 plugin' |
      | '{"name":"demo007 plugin"}' | 'demo007 plugin' |
      | '{"name":"demo008 plugin"}' | 'demo008 plugin' |
      | '{"name":"demo009 plugin"}' | 'demo009 plugin' |

  Scenario: simple plugin search
    Given I create a 'plugin' with body '{"name":"searchable plugin"}'
    When I search a 'plugin' with name 'searchable plugin'
    Then the 'plugin' name must be 'searchable plugin'
