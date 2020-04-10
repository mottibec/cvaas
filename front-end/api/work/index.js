(function (){
    'use strict'; 

    var async = require("async");
    var express   = require("express");
    var request   = require("request");
    var helpers   = require("../../helpers");
    var endpoints = require("../endpoints");
    var app       = express();

    app.get("/work", (req, res, next) => {
        request(`${endpoints.workUrl}/items`, (error, response, body) => {
            if (error) {
              return next(error);
            }
            helpers.respondStatusBody(res, response.statusCode, body)
          });
    })

    module.exports = app;
}());