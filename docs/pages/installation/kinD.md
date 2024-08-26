---
layout: default
title: kinD
permalink: installation/kubernetes/kind
type: installation
category: kubernetes
redirect_from:
- installation/platforms/kind
display-title: "false"
language: en
list: include
image: /assets/img/platforms/kind.png
abstarct: Install Meshplay on kinD. Deploy Meshplay in kinD in-cluster or outside of kinD out-of-cluster.
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

Manage your kinD clusters with Meshplay. Deploy Meshplay in your [kinD cluster](#in-cluster-installation).

<div class="prereqs"><h4>Prerequisites</h4>
<ol>
<li>Install the Meshplay command line client, <a href="{{ site.baseurl }}/installation/meshplayctl" class="meshplay-light">meshplayctl</a>.</li>
<li>Install <a href="https://kubernetes.io/docs/tasks/tools/">kubectl</a> on your local machine.</li>
<li>Install <a href="https://kind.sigs.k8s.io/docs/user/quick-start/#installation">kinD</a>, on your local machine.</li>
</ol>
</div>

Also see: [Install Meshplay on Kubernetes]({{ site.baseurl }}/installation/kubernetes)

## Available Deployment Methods

- [In-cluster Installation](#in-cluster-installation)
  - [Preflight Checks](#preflight-checks)
    - [Preflight: Cluster Connectivity](#preflight-cluster-connectivity)
    - [Preflight: Plan your access to Meshplay UI](#preflight-plan-your-access-to-meshplay-ui)
  - [Installation: Using `meshplayctl`](#installation-using-meshplayctl)
  - [Installation: Using Helm](#installation-using-helm)
- [Post-Installation Steps](#post-installation-steps)
  - [Access Meshplay UI](#access-meshplay-ui)

# In-cluster Installation

Follow the steps below to install Meshplay in your kinD cluster.

## Preflight Checks

Read through the following considerations prior to deploying Meshplay on kinD.

### Preflight: Cluster Connectivity

Start the kinD, if not started using the following command:
{% capture code_content %}kind create cluster{% endcapture %}
{% include code.html code=code_content %}
Check up on your kinD cluster :
{% capture code_content %}kind get clusters{% endcapture %}
{% include code.html code=code_content %}
Verify your kubeconfig's current context.
{% capture code_content %}kubectl config current-context{% endcapture %}
{% include code.html code=code_content %}

### Preflight: Plan your access to Meshplay UI

1. If you are using port-forwarding, please refer to the [port-forwarding](/tasks/accessing-meshplay-ui) guide for detailed instructions.
2. Customize your Meshplay Provider Callback URL. Meshplay Server supports customizing authentication flow callback URL, which can be configured in the following way:

{% capture code_content %}$ MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

Meshplay should now be running in your kinD cluster and Meshplay UI should be accessible at the `INTERNAL IP` of `meshplay` service.

## Installation: Using `meshplayctl`

Once kinD cluster is configured as current cluster-context, execute the below command.

Before executing the below command, go to `~/.meshplay/config.yaml` and ensure that current platform is set to kubernetes.
{% capture code_content %}$ meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

If you encounter any authentication issues, you can use `meshplayctl system login`. For more information, click [here](/guides/meshplayctl/authenticate-with-meshplay-via-cli) to learn more.

## Installation: Using Helm

For detailed instructions on installing Meshplay using Helm V3, please refer to the [Helm Installation](/installation/kubernetes/helm) guide.

# Post-Installation Steps

## Access Meshplay UI

To access Meshplay's UI, please refer to the [instruction](/tasks/accessing-meshplay-ui) for detailed guidance.

{% include suggested-reading.html language="en" %}

{% include related-discussions.html tag="meshplay" %}