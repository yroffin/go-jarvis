const { setWorldConstructor } = require('cucumber')
var request = require('request');
var async = require('async');
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

  // Simple TASK
  task(uri, id, task, parameters, callback) {
    let options = {
      url: 'http://localhost:8080/api/' + uri + '/' + id,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      qs: {
        'task': task
      },
      form: JSON.stringify(parameters)
    }

    request(options, (error, response, body) => {
      if (!error && response.statusCode == 202) {
        callback(JSON.parse(body));
      } else {
        console.error("ERROR", error, response.statusCode, options.url);
      }
    });
  }

  // Simple LINK
  link(uri, id, to, tid, callback) {
    let options = {
      url: 'http://localhost:8080/api/' + uri + '/' + id + '/' + to + '/' + tid,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      form: '{}'
    }

    request(options, (error, response, body) => {
      if (!error && response.statusCode == 200) {
        callback(JSON.parse(body));
      } else {
        console.error("ERROR", error, response.statusCode, options.url);
      }
    });
  }

  // Apply task
  taskExec(api, id, task, prm, commit) {
    this.task(eval('this.' + api + '()').name(), id, task, prm, (result) => {
      this.attach(JSON.stringify(result, null, 4));
      commit(result);
    });
  }

  // Create element
  create(api, body, commit) {
    let tocreate = JSON.parse(body);
    eval('this.' + api + '()').post(tocreate, (created) => {
      this.attach(JSON.stringify(created, null, 4));
      commit(created);
    });
  }

  // Create a link
  createLink(from, fname, to, tname, commit) {
    var afrom = eval('this.' + from + '()');
    var ato = eval('this.' + to + '()');
    // Find from element
    afrom.findAll((elements) => {
      var froms = _.filter(elements, (element) => {
        return eval('element.' + 'name') == fname;
      });
      // Find to element
      ato.findAll((elements) => {
        var tos = _.filter(elements, (element) => {
          return eval('element.' + 'name') == tname;
        });
        // Link them
        afrom.link(froms[0].id, ato.name(), tos[0].id, (link) => {
          this.attach(JSON.stringify(link, null, 4));
          commit(link);
        });
      });
    });
  }

  // Search element with property
  drop(api, property, value, commit) {
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
        commit();
      });
    });
  }

  // Drop element with emty property
  dropEmptyValue(api, property, commit) {
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
        commit();
      });
    });
  }

  // Search element with property
  search(api, property, value, callback) {
    eval('this.' + api + '()').findAll((elements) => {
      let results = _.filter(elements, (element) => {
        return eval('element.' + property) === value;
      });
      this.attach('count: ' + results.length);
      callback(results);
    });
  }

  api(name, json, callback) {
    return {
      name: () => {
        return name;
      },
      delete: (id, callback) => {
        return this.delete(name, id, callback);
      },
      get: (id, callback) => {
        return this.get(name, id, callback);
      },
      findAll: (callback) => {
        return this.findAll(name, callback);
      },
      link: (id, to, tid, callback) => {
        return this.link(name, id, to, tid, callback);
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