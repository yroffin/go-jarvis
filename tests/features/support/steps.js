const { Given, When, Then } = require('cucumber')
const { expect } = require('chai')
var _ = require('lodash');

Given('I create a command with body {string}', function(body, done) {
  let command = JSON.parse(body);
  this.command().post(command, (command) => {
    this.storage.post = command;
    done();
  });
})

When('I read last created command', function(done) {
  this.command().get(this.storage.post.id, (command) => {
    this.storage.get = command;
    done();
  });
})

When('I search a command with name {string}', function(name, done) {
  this.command().findAll((commands) => {
    _.each(commands, (command) => {
      if(command.name === name) {
        this.storage.get = command;
        done();
      }
    });
  });
})

Then('the command name must be {string}', function(name) {
  expect(this.storage.get.name).to.eql(name)
})