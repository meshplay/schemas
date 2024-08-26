---
layout: default
title: Importing Applications
abstract: Learn how to import your existing application definitions and your existing infrastructure configruations into Meshplay as you to manage, operate, and observe your cloud native infrastructure more effectively.
permalink: guides/configuration-management/importing-apps
category: configuration
type: guides
language: en
---

Import your existing application definitions and your existing infrastructure configruations into Meshplay
Meshplay supports a number of different application definition formats. You can import apps into Meshplay using either the Meshplay CLI or the Meshplay UI.

## Supported Application Definition Formats

Meshplay supports the following application definition formats:

- [Kubernetes Manifests](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/)
- [Helm Charts](https://helm.sh/docs/topics/charts/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Meshplay Designs](/concepts/logical/designs)

## Import Apps Using Meshplay CLI

**Step 1: Install Meshplay CLI**

Before you can use the Meshplay CLI to import a Docker Compose app, you must first install it. You can install Meshplay CLI by [following the instructions]({{site.baseurl}}/installation#install-meshplayctl).

**Step 2: Import the App Manifest**

Once you have created your App Definition file, you can use the Meshplay CLI to import your Docker Compose app into Meshplay. To do this, run the following command:

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">meshplayctl app import -f [file/url] -s [source-type]</div></div>
</pre>

This command enable users to import their existing applications from sources as

- Helm Charts
- Kubernetes Manifests
- Docker Compose

**Example :**

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">meshplayctl app import -f ./SampleApplication.yml -s "Kubernetes Manifest"</div></div>
</pre>

## Import Apps Using Meshplay UI

**Step 1: Access the Meshplay UI**

To import a Docker Compose app into Meshplay using the Meshplay UI, you must first [install Meshplay](../installation/quick-start.md)

**Step 2: Navigate to the Application section in the Configuration**

Once you have accessed the Meshplay UI, navigate to the App Import page. This page can be accessed by clicking on the "Applications" menu item and then selecting "Import Application".

<a href="{{ site.baseurl }}/assets/img/applications/Menu.png"><img alt="Application-Navigation" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/applications/Menu.png" /></a>

**Step 3: Upload the Application**

On the App Import page, you can upload your application by select File Type from the options and clicking on the "Browse" button and selecting the file from your local machine or uploading in through URL. Once you have selected the file, click on the "Import" button to import app into Meshplay.

When you import an app into Meshplay, it will create a Meshplay App based on definition. This Meshplay App will include all of the services, ports, and other parameters defined in the File.

<a href="{{ site.baseurl }}/assets/img/applications/ImportApp.png"><img alt="Import-Application" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/applications/ImportApp.png" /></a>

Once the Meshplay App has been created, you can use Meshplay to manage, operate and observe your cloud native infrastructure. You can also use Meshplay to deploy your Meshplay App to any of your connected kubernetes clusters. For more information, see [connections](/installation/kubernetes)
