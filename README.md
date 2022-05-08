# helm-etcd

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](./LICENSE)
[![GitHub release](https://img.shields.io/github/tag-date/QQGoblin/helm-etcd.svg)](https://github.com/QQGoblin/helm-etcd/releases)

The Helm downloader plugin that provides etcd protocol support.

This plugin is used to get the values of the --values parameter from etcd and configurations stored on etcd need to be base64 encoded.

# Install 

```shell
git clone https://github.com/QQGoblin/helm-etcd
cd helm-etcd
```

Edit config.json and fill in the TLS certificate file address of the scope etcd.

```json
{
  "caFile": "/path/ca.crt",
  "keyFile": "/path/tls.key",
  "certFile": "/path/tls.crt"
}
```
Execute the make command to install.

```shell
make
helm plugin list 
```

# Use

Specify values for the --values parameter

```shell 
helm install release-name repo/charts --values etcd:///<key-path>
```