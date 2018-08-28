const { Given, When, Then } = require('cucumber')
const { expect } = require('chai')
var _ = require('lodash');

Given('I create a {string} with body {string}', function(api, body, done) {
  let command = JSON.parse(body);
  eval('this.'+api+'()').post(command, (command) => {
    this.storage.post = command;
    done();
  });
})

When('I read last created {string}', function(api, done) {
  eval('this.'+api+'()').get(this.storage.post.id, (command) => {
    this.storage.get = command;
    done();
  });
})

When('I search a {string} with name {string}', function(api, name, done) {
  eval('this.'+api+'()').findAll((commands) => {
    _.each(commands, (command) => {
      if(command.name === name) {
        this.storage.get = command;
        done();
      }
    });
  });
})

Then('the {string} name must be {string}', function(api, name) {
  expect(this.storage.get.name).to.eql(name)
})