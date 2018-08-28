const { Given, When, Then } = require('cucumber')
const { expect } = require('chai')
var _ = require('lodash');

var { After } = require('cucumber');
var async = require('async');

Given('I create a {string} with body {string}', function (api, body, done) {
  let command = JSON.parse(body);
  eval('this.' + api + '()').post(command, (command) => {
    this.attach(JSON.stringify(command, null, 4));
    this.storage.post = command;
    done();
  });
})

Given('I erase all {string} with {string} starting with {string}', function (api, property, value, done) {
  var api = eval('this.' + api + '()');
  api.findAll((elements) => {
    var results = _.filter(elements, (element) => {
      return eval('element.' + property).startsWith(value);
    });
    // async ops
    async.filterSeries(results, (result, callback) => {
      api.delete(result.id, (found) => {
        this.attach('[FOUND]', JSON.stringify(found, null, 4));
        callback(null, true);
      });
    }, (err, results) => {
      done();
    });
  });
})

Given('I erase all {string} with {string} is empty', function (api, property, done) {
  var api = eval('this.' + api + '()');
  api.findAll((elements) => {
    var results = _.filter(elements, (element) => {
      return eval('element.' + property) === '';
    });
    // async ops
    async.filterSeries(results, (result, callback) => {
      api.delete(result.id, (found) => {
        this.attach('[FOUND]', JSON.stringify(found, null, 4));
        callback(null, true);
      });
    }, (err, results) => {
      done();
    });
  });
})

When('I read last created {string}', function (api, done) {
  eval('this.' + api + '()').get(this.storage.post.id, (found) => {
    this.storage.get = found;
    done();
  });
})

When('I search a {string} with name {string}', function (api, name, done) {
  eval('this.' + api + '()').findAll((commands) => {
    let find = _.findIndex(commands, (command) => {
      return command.name === name;
    });
    if (find >= 0) {
      this.attach(JSON.stringify(commands[find], null, 4));
      this.storage.get = commands[find];
      done();
    }
  });
})

Then('the {string} name must be {string}', function (api, name) {
  this.attach(JSON.stringify(this.storage.get, null, 4));
  expect(this.storage.get.name).to.eql(name)
})

Then('the {string} search {string} = {string} count is {int}', function (api, property, value, count, done) {
  eval('this.' + api + '()').findAll((elements) => {
    let results = _.filter(elements, (element) => {
      return eval('element.' + property) === value;
    });
    this.attach('count: ' + results.length);
    expect(results.length).to.eql(count)
    done();
  });
});
