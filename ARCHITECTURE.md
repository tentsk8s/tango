# Tango Architecture

Tango is a client and a server. The server runs in the cluster that Tango works with and is part of the "control plane" (TODO: fill in more about the control plane later), and you install it (and the rest of the control plane) with the CLI using `tango install`.

The CLI talks to the control plane, and you can configure it by doing one of these:

1. Using a Kubernetes config file (the same one that `kubectl` uses) - with this method, `tango` will create a tunneling proxy
>This is similar to the method that [Helm](https://helm.sh) version 2 uses to communicate with Tiller, its server-side component
2. Connecting to a public endpoint that the Tango server listens on
>By default, `tango install` will install the server so that it listens on a public endpoint, so that other `tango` users don't need to have a Kubernetes configuration on their machine at all
