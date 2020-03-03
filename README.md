<h1 align="center">Welcome to injectorctl 👋</h1>
<p>
  <a href="https://stash.trendyol.com/scm/plat/injectorctl.git" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
</p>

> This project provide simple inject command to make it easy to work with our Mutating Admission Controller Webhook , in a nutshell it modifies given valid Deployment YAML file and gives output which is the modified YAML file

### 🏠 [Homepage](https://stash.trendyol.com/scm/plat/injectorctl.git)

### ✨ [Demo](https://asciinema.org/a/304863)

## Install

```sh
You can directly install this project, clone this project and run :
$ go install

Or you can use docker image to build this project :

$ docker image build -t <your_id>/injectorctl .

```

## Usage

```sh
injectorctl inject -f <file_path> or <stdin>

```

## Run tests

```sh
$ injectorctl inject -f $HOME/hello.yaml

or directly from image in two ways

$ docker container run --interactive <your_id>/injectorctl:latest -<./examples/pod.yaml

$ docker container run --interactive trendyoltech/injectorctl:latest -<<EOF
apiVersion: v1
kind: Pod
metadata:
  labels:
    pod: busybox
spec:
  containers:
    - name: busybox-container
      image: busybox
      command: ["/bin/sh"]
      args: ["-c", "while true; do cat /var/busybox/config.txt; sleep 2; done"]
  serviceAccountName: busybox-sa
EOF
```

## Author

👤 **Trendyol**

* Website: https://stash.trendyol.com/scm/plat/injectorctl.git
* Github: [@TrendyolTech](https://github.com/TrendyolTech)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://stash.trendyol.com/scm/plat/injectorctl.git). You can also take a look at the [contributing guide](https://stash.trendyol.com/scm/plat/injectorctl.git).

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_

<a href="https://asciinema.org/a/304863"><img src="https://asciinema.org/a/304863.png" width="836"/></a>
