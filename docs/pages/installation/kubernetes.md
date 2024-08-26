---
layout: default
title: Kubernetes
permalink: installation/kubernetes
type: installation
category: kubernetes
redirect_from:
- installation/platforms/kubernetes
display-title: "false"
language: en
list: include
image: /assets/img/platforms/kubernetes.svg
abstract: Install Meshplay on Kubernetes. Deploy Meshplay in Kubernetes in-cluster or outside of Kubernetes out-of-cluster.
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

Manage your kubernetes clusters with Meshplay. Deploy Meshplay in kubernetes [in-cluster](#in-cluster-installation) or outside of kubernetes [out-of-cluster](#out-of-cluster-installation). **_Note: It is advisable to [Install Meshplay in your kubernetes clusters](#install-meshplay-into-your-kubernetes-cluster)_**

<div class="prereqs"><h4>Prerequisites</h4>
  <ol>
    <li>Install the Meshplay command line client, <a href="{{ site.baseurl }}/installation/meshplayctl" class="meshplay-light">meshplayctl</a>.</li>
    <li>Install <a href="https://kubernetes.io/docs/tasks/tools/">kubectl</a> on your local machine.</li>
    <li>Access to an active kubernetes cluster.</li>
  </ol>
</div>

## Available Deployment Methods

- [In-cluster Installation](#in-cluster-installation)
  - [Preflight Checks](#preflight-checks)
    - [Preflight: Cluster Connectivity](#preflight-cluster-connectivity)
    - [Preflight: Plan your access to Meshplay UI](#preflight-plan-your-access-to-meshplay-ui)
  - [Installation: Using `meshplayctl`](#installation-using-meshplayctl)
  - [Installation: Using Helm](#installation-using-helm)
- [Post-Installation Steps](#post-installation-steps)
  - [Access Meshplay UI](#access-meshplay-ui)
- [Out-of-cluster Installation](#out-of-cluster-installation)
  - [Installation: Upload Config File in Meshplay Web UI](#installation-upload-config-file-in-meshplay-web-ui)

# In-cluster Installation

Follow the steps below to install Meshplay in your kubernetes cluster.

## Preflight Checks

Read through the following considerations prior to deploying Meshplay on kubernetes.

### Preflight: Cluster Connectivity

Verify your kubeconfig's current context is set the kubernetes cluster you want to deploy Meshplay.
{% capture code_content %}kubectl config current-context{% endcapture %}
{% include code.html code=code_content %}

### Preflight: Plan your access to Meshplay UI

1. If you are using port-forwarding, please refer to the [port-forwarding](/tasks/accessing-meshplay-ui) guide for detailed instructions.
2. Customize your Meshplay Provider Callback URL. Meshplay Server supports customizing authentication flow callback URL, which can be configured in the following way:

{% capture code_content %}$ MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

Meshplay should now be running in your Kubernetes cluster and Meshplay UI should be accessible at the `EXTERNAL IP` of `meshplay` service.

## Installation: Using `meshplayctl`

Once configured, execute the following command to start Meshplay.

Before executing the below command, go to ~/.meshplay/config.yaml and ensure that current platform is set to kubernetes.
{% capture code_content %}$ meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

If you encounter any authentication issues, you can use `meshplayctl system login`. For more information, click [here](/guides/meshplayctl/authenticate-with-meshplay-via-cli) to learn more.

## Installation: Using Helm

For detailed instructions on installing Meshplay using Helm V3, please refer to the [Helm Installation](/installation/kubernetes/helm) guide.

# Post-Installation Steps

## Access Meshplay UI

To access Meshplay's UI, please refer to the [instruction](/tasks/accessing-meshplay-ui) for detailed guidance.

# Out-of-cluster Installation

Install Meshplay on Docker (out-of-cluster) and connect it to your Kubernetes cluster.

## Installation: Upload Config File in Meshplay Web UI

- Run the below command to generate the _"config_minikube.yaml"_ file for your cluster:

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">kubectl config view --minify --flatten > config_minikube.yaml</div></div>
 </pre>

- Upload the generated config file by navigating to _Settings > Environment > Out of Cluster Deployment_ in the Web UI and using the _"Upload kubeconfig"_ option.

{% include suggested-reading.html language="en" %}

{% include related-discussions.html tag="meshplay" %}
