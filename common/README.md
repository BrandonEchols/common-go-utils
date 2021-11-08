# A Collection Of Common Go Utilities

## Purpose of this module
This module should contain the classes that are used in nearly every go micro-service. The following are descriptions of
the current files/classes that are available in this module.

#### ConfigGetter.go
This contains a ready to use, three tier configuration system that loads string config values from environment variables,
and a configuration json file. For information on how to use it, see ConfigGetter.go. For a mock interface to test with,
see the 'mocks' module.

#### ZapLogger.go
This module is a wrapper class to the https://github.com/uber-go/zap SugaredLogger. The wrapper provides an
easy-to-use-for-testing interface that can be initialized to either a Production or Development Logger.
