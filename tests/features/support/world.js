const { setWorldConstructor } = require('cucumber')
var request = require('request');
var _ = require('lodash');

class JarvisClient {
  constructor() {
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

  // Simple GET
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

setWorldConstructor(JarvisClient)