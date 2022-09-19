Read a Kubernetes [ConfigMap][] (or [Secret][]) on standard input. Extract values from `.data`, writing them to files named after the corresponding key. Automatically base64 decode Secret values.

[configmap]: https://kubernetes.io/docs/concepts/configuration/configmap/
[secret]: https://kubernetes.io/docs/concepts/configuration/secret/
