---
layout: default
title: Helm
permalink: installation/kubernetes/helm
type: installation
category: kubernetes
redirect_from:
- installation/platforms/helm
display-title: "false"
language: en
list: include
abstract: Install Meshplay on Kubernetes using Helm. Deploy Meshplay in Kubernetes in-cluster.
---
# Install Meshplay on Kubernetes Using Helm

<div class="prereqs"><h4>Prerequisites</h4>
<ol>
<li><a href="https://helm.sh/docs/intro/install/" class="meshplay-light">Helm</a> should be installed on your local machine.</li>
<li>You should have access to the cluster/platform where you want to deploy Meshplay.</li>
<li>Ensure that the kubeconfig file has the correct current context/cluster configuration.</li>
</ol>
</div>

## Install Meshplay on Your Kubernetes Cluster Using Helm

{% capture code_content %}helm repo add meshplay https://khulnasoft.com/charts/
helm install meshplay khulnasoft/meshplay --namespace meshplay --create-namespace
helm install meshplay-operator khulnasoft/meshplay-operator{% endcapture %}
{% include code.html code=code_content %}

Optionally, Meshplay Server supports customizing the callback URL for your remote provider, like so:

{% capture code_content %}helm install meshplay khulnasoft/meshplay --namespace meshplay --set env.MESHPLAY_SERVER_CALLBACK_URL=https://custom-host --create-namespace{% endcapture %}
{% include code.html code=code_content %}

### Customizing Meshplay's Installation with values.yaml

Meshplay's Helm chart supports a number of configuration options. Please refer to the [Meshplay Helm chart](https://github.com/khulnasoft/meshplay/tree/master/install/kubernetes/helm/meshplay#readme) and [Meshplay Operator Helm Chart](https://github.com/khulnasoft/meshplay/tree/master/install/kubernetes/helm/meshplay-operator#readme) for more information.

### Accessing Meshplay UI for Clusters

To access Meshplay's UI , please refer to the [accessing-ui](/tasks/accessing-meshplay-ui) guide for detailed instructions.

{% include suggested-reading.html language="en" %}

{% include related-discussions.html tag="meshplay" %}
