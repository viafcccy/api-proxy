# api-proxy
Go Proxy Forwarding Tool

## usage
- use cmd line
1. run `./api-proxy -t http://your-target-host` or `nohup ./api-proxy -t http://your-target-host forward &`

- use config
1. modify your target host in ./config/config.yaml or `nohup ./api-proxy forward &`
2. run `./api-proxy`

if cannot execute binary file, please run `go build`.