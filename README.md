# YqFlat

yqflat provides functionality to read a YAML file 
and flatten its nested structure into a flat map of string key-value pairs. 
The main abstraction is the YqFlat struct, which encapsulates the separator used for joining nested keys.

Example input yaml
```yaml
application:
  spring:
    port: 8080
    profiles:
      active: dev
server:
  port: 9090
```
Example flatterned output

```json
// separator=/
{"application/spring/port": "8080"}
{"application/spring/profiles/active": "dev"}
{"server/port": "9090"}

// separator=.
{"application.spring.port": "8080"}
{"application.spring.profiles.active": "dev"}
{"server.port": "9090"}
```


## cmd/main.go

Entry point for a CLI application that:

- Loads configuration from a YAML file (using Viper).
- Reads and flattens another YAML file (using the internal yqflat package).
- Logs the flattened key-value pairs (using Zap logger).

## config/config.yaml

The file config/config.yaml is used for
- setting up the logger and
- defining the separator for flatterned entries (e.g.: "/" or ".").

As alternative, the flag ```-config=./path/to/config``` can also be used.

## example/example.yaml

As default ```example/example.yaml``` is used as file to parse,
but the flag ```-yaml=path/to/yaml``` can be used as well.
