Feature: Mixed cruds
  In order to create some commands
  As a developer
  I want to use jarvis API

  Scenario: simple mixed cruds
    Given I erase all 'device' with 'name' starting with 'mixed'
    Given I erase all 'plugin' with 'name' starting with 'mixed'
    Given I erase all 'command' with 'name' starting with 'mixed'
    Given I create a 'device' with body '{"name":"mixed device", "icon": "save"}' and store it to 'created'
    Given I create a 'plugin' with body '{"name":"mixed plugin", "icon": "drop"}' and store it to 'created'
    Given I create a 'command' with body '{"name":"mixed command", "icon": "save", "body": "return 100;", "type": "LUA"}' and store it to 'created'
    # check device attributes
    When I search a 'device' with 'name' equals to 'mixed device' and store it to 'dev001'
    Then the store 'dev001'.'name' must be 'mixed device'
    Then the store 'dev001'.'icon' must be 'save'
    # check plugin attributes
    When I search a 'plugin' with 'name' equals to 'mixed plugin' and store it to 'plug001'
    Then the store 'plug001'.'name' must be 'mixed plugin'
    Then the store 'plug001'.'icon' must be 'drop'
    # check command attributes
    When I search a 'command' with 'name' equals to 'mixed command' and store it to 'cmd001'
    Then the store 'cmd001'.'name' must be 'mixed command'
    Then the store 'cmd001'.'type' must be 'LUA'
    # Assert we can continue testing
    Then count 'device' with 'name' equals to 'mixed device' is 1
    Then count 'plugin' with 'name' equals to 'mixed plugin' is 1
    Then count 'command' with 'name' equals to 'mixed command' is 1
    # Execute the command
    When I apply 'execute' on 'command' identified by 'cmd001'.'id' and store result to 'exc'
    Then the store 'exc'.'object.result' must be '100'
    # Now we have to link them
    Given I link 'device' with name 'mixed device' to 'plugin' with name 'mixed plugin'
    Given I link 'plugin' with name 'mixed plugin' to 'command' with name 'mixed command'

    
