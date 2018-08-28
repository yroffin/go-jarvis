Feature: Device cruds
  In order to create some devices
  As a developer
  I want to use jarvis API

  Scenario: simple device
    Given I erase all 'device' with 'name' starting with 'demo'
    Given I create a 'device' with body '{"name":"demo device"}' and store it to 'created'
    Then the store 'created'.'name' must be 'demo device'

  Scenario Outline: much more complex stuff
    Given I erase all 'device' with 'name' starting with <name>
    Given I create a 'device' with body <body> and store it to 'created'
    Then the store 'created'.'name' must be <name>

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
    Given I erase all 'device' with 'name' starting with 'searchable'
    Given I create a 'device' with body '{"name":"searchable device"}' and store it to 'created'
    When I search a 'device' with 'name' equals to 'searchable device' and store it to 'found'
    Then the store 'found'.'name' must be 'searchable device'
