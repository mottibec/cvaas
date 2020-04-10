var request      = require("request")
  , express      = require("express")
  , morgan       = require("morgan")
  , path         = require("path")
  , bodyParser   = require("body-parser")
  , async        = require("async")
  , config       = require("./config")
  , app          = express()


  app.use(bodyParser.json());
  app.use(morgan("dev", {}));

  var server = app.listen(process.env.PORT || 8079, () =>  {
    var port = server.address().port;
    console.log("App now running in %s mode on port %d", app.get("env"), port);
  });