var before = require('hooks').before;
var sha1 = require('sha1');

var hooks = require('hooks');
before("/items > POST > 200 > application/json", function (transaction) {
    key = sha1(process.env.DEMOSHASALT + "Carlos")
    transaction.request['headers']['Auth'] = key;
});
before("/nope > GET > 200 > application/json", function (transaction) {
    transaction.skip = true;
});
