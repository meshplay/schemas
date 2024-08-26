---
layout: default
title: Quick Start Guide
permalink: installation/quick-start
# redirect_from: installation/quick-start/
language: en
type: installation
abstract: Getting Meshplay up and running locally on a Docker-enabled system or in Kubernetes is easy. Meshplay deploys as a set of Docker containers, which can be deployed to either a Docker host or Kubernetes cluster.
---

<a name="getting-started"></a>

Getting Meshplay up and running locally on a Docker-enabled system or in Kubernetes is easy. Meshplay deploys as a set of Docker containers, which can be deployed to either a Docker host or Kubernetes cluster.

{% include alert.html type="info" title="All Supported Platforms" content="Download, install, and run Meshplay in a single command. See all <a href='/installation'>supported platforms</a>." %}

## Install Meshplay

To install and start Meshplay, begin by [installing meshplayctl]({{ site.baseurl }}/installation/meshplayctl). If you are on macOS or Linux system, you can download, install, and run both `meshplayctl` and Meshplay Server with the command shown in the figure.

 <pre class="codeblock-pre" style="padding: 0; font-size:0px;"><div class="codeblock" style="display: block;">
 <div class="window-buttons"></div>
  <div id="termynal0" style="width:100%; height:150px; max-width:100%;" data-termynal="">
            <span data-ty="input">curl -L https://khulnasoft.com/install | PLATFORM=kubernetes bash -</span>
            <span data-ty="progress"></span>
            <span data-ty="">Successfully installed Meshplay</span>
            <span data-ty="input">meshplayctl system dashboard</span>
  </div>
  </div>

 </pre>

{% include alert.html type="info" title="Meshplay CLI Package Managers" content="In addition to <a href='/installation/linux-mac/bash'>Bash</a>, you can also use <a href='/installation/linux-mac/brew'>Brew</a> or <a href='/installation/windows/scoop'>Scoop</a> to install <code>meshplayctl</code>. Alternatively, <code>meshplayctl</code> is also available <a href='https://github.com/khulnasoft/meshplay/releases/latest'>direct download</a>." %}

## Install using Docker Meshplay Extension

You can find the [Docker Meshplay Extension in Docker Hub](https://hub.docker.com/extensions/khulnasoft/docker-extension-meshplay) marketplace or use the Extensions browser in Docker Desktop to install the Docker Meshplay Extension.

[![Docker Meshplay Extension]({{site.baseurl}}/assets/img/platforms/docker-desktop-meshplay-extension.png)]({{site.baseurl}}/assets/img/platforms/docker-desktop-meshplay-extension.png)

## Access Meshplay

Your default browser will be opened and directed to Meshplay's web-based user interface typically found at `http://localhost:9081`.

{% include alert.html type="info" title="Accessing Meshplay UI" content="Meshplay's web-based user interface is embedded in Meshplay Server and is available as soon as Meshplay starts. The location and port that Meshplay UI is exposed varies depending upon your mode of deployment. See <a href='/tasks/accessing-meshplay-ui'>accessing Meshplay UI</a> for deployment-specific details." %}

### Select a Provider

Select from the list of [Providers]({{ site.baseurl }}/extensibility/providers) in order to login to Meshplay. Authenticate with your chosen Provider.

<a href="/assets/img/meshplay-server-page.png">
  <img class="center" style="width:min(100%,650px)" src="/assets/img/meshplay-server-page.png" />
</a>

### Configure Connections to your Kubernetes Clusters

If you have deployed Meshplay out-of-cluster, Meshplay will automatically connect to any available Kubernetes clusters found in your kubeconfig (under `$HOME/.kube/config`). If you have deployed Meshplay out-of-cluster, Meshplay will automatically connect to the Kubernetes API Server availabe in the control plane. Ensure that Meshplay is connected to one or more of your Kubernetes clusters.

Visit <i class="fas fa-cog"></i> Settings:

<a href="/assets/img/platforms/meshplay-settings.png">
  <img class="center" style="width:min(100%,650px);" src="/assets/img/platforms/meshplay-settings.png" />
</a>

If your config has not been autodetected, you can manually upload your kubeconfig file (or any number of kubeconfig files). By default, Meshplay will attempt to connect to and deploy Meshplay Operator to each reachable context contained in the imported kubeconfig files. See Managing Kubernetes Clusters for more information.

### Verify Deployment

Run connectivity tests and verify the health of your Meshplay system. Verify Meshplay's connection to your Kubernetes clusters by clicking on the connection chip. A quick connectivity test will run and inform you of Meshplay's ability to reach and authenticate to your Kubernetes control plane(s). You will be notified of your connection status. You can also verify any other connection between Meshplay and either its components (like [Meshplay Adapters]({{ site.baseurl }}/concepts/architecture/adapters)) or other managed infrastructure by clicking on any of the connection chips. When clicked, a chip will perform an ad-hoc connectivity test.

<a href="{{site.baseurl}}/assets/img/platforms/k8s-context-switcher.png" alt="Meshplay Kubernetes Context Switcher">
  <img class="center" style="width:min(100%,350px);" src="{{site.baseurl}}/assets/img/platforms/k8s-context-switcher.png" />
</a>

### Design and operate Kubernetes clusters and their workloads

You may now proceed to managed any cloud native infrastructure supported by Meshplay. See all integrations for a complete list of supported infrastructure.

<a href="{{site.baseurl}}/assets/img/platforms/meshplay-designs.png">
  <img class="center" style="width:min(100%,650px);" src="{{site.baseurl}}/assets/img/platforms/meshplay-designs.png" />
</a>

## Additional Guides

<div class="section">
    <ul>
        <li><a href="{{ site.baseurl }}/guides/troubleshooting/installation">Troubleshooting Meshplay Installations</a></li>
        <li><a href="{{ site.baseurl }}/reference/error-codes">Meshplay Error Code Reference</a></li>
        <li><a href="{{ site.baseurl }}/reference/meshplayctl/system/check">Meshplayctl system check</a></li> 
    </ul>
</div>
<script src="/assets/js/terminal.js" data-termynal-container="#termynal0|#termynal1|#termynal2"></script>

