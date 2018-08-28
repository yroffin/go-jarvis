Feature: Device cruds
  In order to create some devices
  As a developer
  I want to use jarvis API

  Scenario: simple device
    Given I create a 'device' with body '{"name":"demo device"}'
    When I read last created 'device'
    Then the 'device' name must be 'demo device'

  Scenario Outline: much more complex stuff
    Given I create a 'device' with body <body>
    When I read last created 'device'
    Then the 'device' name must be <name>

    Examples:
      | body | name |
      | '{"name":"demo001 device"}' | 'demo001 device' |
      | '{"name":"demo002 device"}' | 'demo002 device' |
      | '{"name":"demo003 device"}' | 'demo003 device' |
      | '{"name":"demo004 device"}' | 'demo004 device' |
      | '{"name":"demo005 device"}' | 'demo005 device' |
      | '{"name":"demo006 device"}' | 'demo006 device' |
      | '{"name":"demo007 device"}' | 'demo007 device' |
      | '{"name":"demo008 device"}' | 'demo008 device' |
      | '{"name":"demo009 device"}' | 'demo009 device' |

  Scenario: simple device search
    Given I create a 'device' with body '{"name":"searchable device"}'
    When I search a 'device' with name 'searchable device'
    Then the 'device' name must be 'searchable device'
