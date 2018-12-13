# Tango

>Kubernetes for App Devs

Tango is an opinionated toolchain that helps application developers deploy easily to Kubernetes. Its goals are simple:

- Configure a Kubernetes cluster with the best technologies for running apps, with one command
- Build and deploy an application according to best practices, with one command

Tango combines a lot of software from the Kubernetes & cloud native ecosystems and follows best practices for your app. Check out some of these practices and technologies:

- [Draft](https://draft.sh) for a great developer workflow
- [Traefik](https://traefik.io/) for the ingress controller
- [Linkerd](https://linkerd.io/) for a service mesh
- [Prometheus](https://prometheus.io/) for metrics

>Tango was first explained at [@arschles](https://twitter.com/arschles)'s [talk](https://arschles.com/kubecon2018) at KubeCon North America 2018.

# Commands

The `tango` CLI gives app developers access to their Kubernetes cluster in a really opinionated way. If you're used to using `kubectl` to access your cluster, `tango` will probably feel super restrictive to you.

Here are the commands that `tango` gives you:

```console
tango install
```

This command sets up a "vanilla" Kubernetes cluster to work with Tango. You only need to run this once per cluster, before you initialize any apps.

```console
tango init
```

This command initializes an app for Tango. It creates all the files you need in your local repository, and also tells the cluster about your app so that you can deploy it in the future.

```console
tango dev
```

This command builds your app, deploys one instance of it to the cluster. Then, it:

- Creates a tunneling proxy to your instance for the port (or ports) that you specify
- Tails the logs

It's a great command to use for dev and testing.

```console
tango deploy
```

This command builds and deploys your app to production. It guides you step by step through the deployment process and asks you for confirmation at each step, so you don't accidentally break something :grinning:. As it goes, it follows best practices along the way (immutable docker image tags, metrics, traffic splitting, ...)

```console
tango fork
```

This command takes your entire app's structure and deploys it to a different tango environment. The application's configurations won't come along with it, though, so you don't accidentally share a database with another tango environment.

```console
tango login
```

This command logs you into the Tango server. TODO: more on how logins and user creation works.
