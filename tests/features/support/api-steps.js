const { Given, When, Then } = require('cucumber')
const { expect } = require('chai')
var _ = require('lodash');

var { After } = require('cucumber');
var async = require('async');

Given('I create a {string} with body {string} and store it to {string}', function (api, body, store, done) {
  this.create(api, body, (created) => {
    eval('this.storage.' + store + ' = created');
    done();
  });
});

Given('I link {string} with name {string} to {string} with name {string}', function (from, fname, to, tname, done) {
  this.createLink(from, fname, to, tname, (linked) => {
    done();
  });
})

Given('I erase all {string} with {string} starting with {string}', function (api, property, value, done) {
  this.drop(api, property, value, () => {
    done();
  });
})

Given('I erase all {string} with {string} is empty', function (api, property, done) {
  this.dropEmptyProperty(api, property, () => {
    done();
  });
})

When('I apply {string} on {string} with name {string} and store result to {string}', function (task, api, name, store, done) {
  done();
})

When('I search a {string} with {string} equals to {string} and store it to {string}', function (api, property, value, store, done) {
  this.search(api, property, value, (founds) => {
    eval('this.storage.' + store + ' = founds[0]');
    done();
  });
})

Then('the store {string}.{string} must be {string}', function (store, property, value) {
  expect(eval('this.storage.' + store + '.' + property)).to.eql(value);
});

Then('count {string} with {string} equals to {string} is {int}', function (api, property, value, count, done) {
  this.search(api, property, value, (founds) => {
    expect(founds.length).to.eql(count)
    done();
  });
});
