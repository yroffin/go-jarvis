const { Given, When, Then } = require('cucumber')
const { expect } = require('chai')
var _ = require('lodash');

var {After} = require('cucumber');

Given('I create a {string} with body {string}', function(api, body, done) {
  let command = JSON.parse(body);
  eval('this.'+api+'()').post(command, (command) => {
    this.attach(JSON.stringify(command, null, 4));
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
    let find = _.findIndex(commands, (command) => {
      return command.name === name;
    });
    if(find >= 0) {
      this.attach(JSON.stringify(commands[find], null, 4));
      this.storage.get = commands[find];
      done();
    }
  });
})

Then('the {string} name must be {string}', function(api, name) {
  this.attach(JSON.stringify(this.storage.get, null, 4));
  expect(this.storage.get.name).to.eql(name)
})