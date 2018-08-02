# Onfig Webapp Example

This is a simple webapp example. You'll use the `onfig` CLI to develop,
test and build a binary that can do some configuration against your
cluster.

```console
onfig build .
```

The resulting binary will be a CLI that will have a few commands:

- `run` - actually run the config against a live Kubernetes cluster
- `test` - test your config against a mock Kubernetes server

You can also run `onfig dev` to run your tests whenever any of your
code changes.
