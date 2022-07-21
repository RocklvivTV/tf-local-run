# tf-local-run
Simple binary to run terraform commands with different credentials.

### Prerequisites
Before using application configuration file should be set.
Default path which application is looking for: `~/.tf_executor/config.yaml`

Example of config:

```yaml
profiles:
  - name: test
    envs:
      AWS_ACCESS_KEY: ""
      AWS_SECRET_KEY: ""
  - name: test2
    envs:
      SOME_SECRET_VAR: ""
      SOME_API_KEY_VAR: ""
```

### Usage

`tf-local-run --profile test --path <path to terraform directory> <init|plan|apply|state>`