const { setWorldConstructor } = require('cucumber')
var request = require('request');
var _ = require('lodash');

class JarvisClient {
  constructor(prm) {
    this.attach = prm.attach;
    this.storage = {
      "default": "default"
    };
  }

  // Simple GET
  get(uri, id, callback) {
    let options = {
      url: 'http://localhost:8080/api/' + uri + '/' + id,
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }

    request(options, (error, response, body) => {
      if (!error && response.statusCode == 200) {
        callback(JSON.parse(body));
      } else {
        console.error("ERROR", error, response.statusCode, body);
      }
    });
  }

  // Simple DELETE
  delete(uri, id, callback) {
    let options = {
      url: 'http://localhost:8080/api/' + uri + '/' + id,
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }
    }

    request(options, (error, response, body) => {
      if (!error && response.statusCode == 200) {
        callback(JSON.parse(body));
      } else {
        console.error("ERROR", error, response.statusCode, body);
      }
    });
  }

  // Simple GET ALL
  findAll(uri, callback) {
    let options = {
      url: 'http://localhost:8080/api/' + uri,
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }

    request(options, (error, response, body) => {
      if (!error && response.statusCode == 200) {
        callback(JSON.parse(body));
      } else {
        console.error("ERROR", error, response.statusCode, body);
      }
    });
  }

  // Simple POST
  post(uri, body, callback) {
    let options = {
      url: 'http://localhost:8080/api/' + uri,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      form: JSON.stringify(body)
    }

    request(options, (error, response, body) => {
      if (!error && response.statusCode == 201) {
        callback(JSON.parse(body));
      } else {
        console.error("ERROR", error, response.statusCode, options.url);
      }
    });
  }

  api(name, json, callback) {
    return {
      delete: (id, callback) => {
        return this.delete(name, id, callback);
      },
      get: (id, callback) => {
        return this.get(name, id, callback);
      },
      findAll: (callback) => {
        return this.findAll(name, callback);
      },
      post: (json, callback) => {
        return this.post(name, json, callback);
      }
    };
  }

  // command
  command(json, callback) {
    return this.api('commands', json, callback);
  }

  // device
  device(json, callback) {
    return this.api('devices', json, callback);
  }

  // plugin
  plugin(json, callback) {
    return this.api('plugins/scripts', json, callback);
  }
}

setWorldConstructor(JarvisClient);

var reporter = require('cucumber-html-reporter');

process.on('exit', function () {
  var options = {
    theme: 'hierarchy',
    jsonFile: 'test/report/cucumber_report.json',
    output: 'test/report/cucumber_report.html',
    reportSuiteAsScenarios: true,
    launchReport: true,
    columnLayout: 1,
    metadata: {
      "App Version": "0.3.5",
      "Test Environment": "STAGING",
      "Browser": "Chrome  54.0.2840.98",
      "Platform": "Windows 10",
      "Parallel": "Scenarios",
      "Executed": "Remote"
    }
  };

  reporter.generate(options);
});