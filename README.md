![CI](https://github.com/jonascheng/stellar-feedback-fakeit/actions/workflows/ci.yaml/badge.svg)
![CD](https://github.com/jonascheng/stellar-feedback-fakeit/actions/workflows/cd.yaml/badge.svg)
![codecov](https://codecov.io/gh/jonascheng/stellar-feedback-fakeit/branch/main/graph/badge.svg)

# stellar-feedback-fakeit

## Usage

```bash
$ ./bin/fakeit-go --help
usage: fakeit-go [<flags>]

Flags:
  --help                Show context-sensitive help (also try --help-long and
                        --help-man).
  --server-info         Random generate server-telemetry.
  --agent-info          Random generate agent-telemetry.
  --agent-system-env    Random generate agent-telemetry-system-environment.
  --agent-software-env  Random generate agent-telemetry-software-environment.
  --agent-cert          Random generate agent-telemetry-cert.
  --threat-info         Random generate agent-telemetry-threat.
  --all-info            Random generate all telemetry above.
  --benchmark           Benchmark performance for agent-system-env,
                        agent-software-env and agent-cert.
  --size=1              Random size
  --cleanup             Remove all json and gzip files
  --debug               Debug output results in json format
  --version             Show application version.

```

## LICENSE

[MIT](https://github.com/jonascheng/stellar-feedback-fakeit/blob/master/LICENSE)
