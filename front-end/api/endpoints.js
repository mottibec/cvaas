(function () {
    'use strict';

    var util = require('util');

    var domain = "";
    process.argv.forEach(function (val, index, array) {
        var arg = val.split("=");
        if (arg.length > 1) {
            if (arg[0] == "--domain") {
                domain = "." + arg[1];
                console.log("Setting domain to:", domain);
            }
        }
    });

    module.exports = {
        workUrl: util.format("http://work%s", domain),
        contactUrl: util.format("http://contact%s", domain),
        cvUrl: util.format("http://cv%s", domain),
    };
}());