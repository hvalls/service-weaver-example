## Service Weaver example

This is a "hello world" example to illustrate how [Service Weaver](https://serviceweaver.dev) framework works.

![Diagram](diagram.png)

Run locally:

```
$ weaver generate
$ go run .
```

Deploy in Google Kubernetes Service:

```
$ weaver gke deploy weaver.toml
```

This code belongs to my blog post [Service Weaver: Write monolith, deploy microservices](https://hvalls.dev/posts/service-weaver).
