# DetectDee
<font color="red">*Disclaimer: This article and this tool are for technical discussion and sharing only. Illegal use is strictly prohibited.</font>

DetectDee: Hunt down social media accounts by **username, email or phone** across [social networks](site.md)
![example.gif](https://s2.loli.net/2023/04/30/FZ1QtKoGud4xVPW.gif)
## Feat
- Includes sites frequently used by **CyberSecurity practitioners**
- Hunt down social media accounts by **username, email or phone**
- Precise thread control and custom request headers are used to prevent WAF recognition
- Extensible, simple, and easy-to-use template
- Integration of mobile versions of social networking sites

## Install
### Download(recommend)
https://github.com/piaolin/DetectDee/releases
### Compile
```shell
git clone https://github.com/piaolin/DetectDee.git
cd DetectDee
go mod tidy
go run .
```
## Usage
[English](README.md)

[中文文档](README_ZH.md)
### Detect
```text
Usage:
  DetectDee detect [flags]

Flags:
  -c, --check          self-check
  -h, --help           help for detect
  -n, --name strings   name[s]
      --nsfw           Include checking of NSFW sites from default list.
      --precisely      Check precisely
  -p, --proxy string   Make requests over a proxy. e.g. socks5://127.0.0.1:1080
  -s, --site strings   Limit analysis to just the listed sites. Add multiple opt
ions to specify more than one site.
  -t, --timeout int    Time (in seconds) to wait for response to requests (defau
lt 60)

Global Flags:
  -v, --verbose   verbose output
```

To search for only one user:
```shell
./DetectDee detect -n piaolin
```

To search for more than one user:
```shell
./DetectDee detect -n piaolin,blue
```

To search in specified site:
```shell
./DetectDee detect -n piaolin -s github,v2ex
```

## Contributing
We would love to have you help us with the development of DetectDee. Each and every contribution is greatly valued!

Here are some things we would appreciate your help on:

- Addition of new site support, You can notify me that a site has an interface available, or you can write [JSON](t.json) directly
- Bringing back site support of sites that have been removed in the past due to false positives

## Todo
- **Credential Stuffing** for result
- More site
- Secret

## Thanks
https://github.com/sherlock-project/sherlock