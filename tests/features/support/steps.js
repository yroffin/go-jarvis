const { Given, When, Then } = require('cucumber')
const { expect } = require('chai')

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

Then('the command name must be {string}', function(name) {
  console.log("Checking command ", this.storage.get.name);
  expect(this.storage.get.name).to.eql(name)
})