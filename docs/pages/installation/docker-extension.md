---
layout: default
title: Docker Extension
permalink: installation/docker/docker-extension
type: installation
category: docker
redirect_from:
- installation/platforms/docker-extension
display-title: "false"
language: en
list: include
image: /assets/img/platforms/docker.svg
abstract: Install Docker Extension for Meshplay
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

## Docker Extension for Meshplay

The Docker Extension for Meshplay extends Docker Desktop’s position as the developer’s go-to Kubernetes environment with easy access to full the capabilities of Meshplay's collaborative cloud native management features.

### Using Docker Desktop

1. Navigate to the Extensions Marketplace of Docker Desktop.

2. From the Dashboard, select Add Extensions in the menu bar or open the Extensions Marketplace from the menu options.

<a href="{{ site.baseurl }}/assets/img/platforms/docker-extension-marketplace-1.png">
  <img style="width:350px;" src="{{ site.baseurl }}/assets/img/platforms/docker-extension-marketplace-1.png">
</a>

3. Navigate to Meshplay in the Marketplace and press install.

<a href="{{ site.baseurl }}/assets/img/platforms/docker-extension.png">
  <img style="width:90%" src="{{ site.baseurl }}/assets/img/platforms/docker-extension.png">
</a>

OR

You can visit the [Docker Hub](https://hub.docker.com/extensions/khulnasoft/docker-extension-meshplay) marketplace to directly install Meshplay extension in your Docker Desktop.

### Using Docker CLI

Meshplay runs as a set of one or more containers inside your Docker Desktop virtual machine.

<!--
{% capture code_content %}docker extension install khulnasoft/docker-extension-meshplay{% endcapture %} -->
<!-- {% include code.html code=code_content %} -->

<pre class="codeblock-pre" style="padding: 0; font-size:0px;"><div class="codeblock" style="display: block;">
 <div class="clipboardjs" style="padding: 0">
 <span style="font-size:0;">docker extension install khulnasoft/docker-extension-meshplay</span> 
 </div>
 <div class="window-buttons"></div>
 <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
            <span data-ty="input">docker extension install khulnasoft/docker-extension-meshplay</span>
            <span data-ty="progress"></span>
            <span data-ty="">Successfully installed Meshplay</span>
            <span data-ty="input">meshplayctl system dashboard</span>
  </div>
 </div>
</pre>

## Remove Meshplay as a Docker Extension

If you want to remove Meshplay as a Docker extension from your system, follow these steps:

1. **Stop Meshplay Container:**

   - First, stop the running Meshplay container (if it's currently running) using the following Docker command:

   <pre class="codeblock-pre" style="padding: 0; margin-top: 2px; font-size:0px;"><div class="codeblock" style="display: block;">
    <div class="clipboardjs" style="padding: 0">
    <span style="font-size:0;">docker stop meshplay-container</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
      <span data-ty="input">docker stop meshplay-container</span>
    </div>
    </div>
   </pre>
    
2. **Remove Meshplay Container:**

   - After stopping the container, you can remove it using the following command:

   <pre class="codeblock-pre" style="padding: 0; margin-top: 2px; font-size:0px;"><div class="codeblock" style="display: block;">
    <div class="clipboardjs" style="padding: 0">
    <span style="font-size:0;">docker rm meshplay-container</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
      <span data-ty="input">docker rm meshplay-container</span>
    </div>
    </div>
   </pre>

3. **Remove Meshplay Images:**

   - Meshplay might have pulled Docker images for its components. You can remove these images using the `docker rmi` command. Replace the image names with the actual ones you want to remove:

   <pre class="codeblock-pre" style="padding: 0; margin-top: 2px; font-size:0px;"><div class="codeblock" style="display: block;">
    <div class="clipboardjs" style="padding: 0">
    <span style="font-size:0;">docker rmi khulnasoft/meshplay:latest</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
      <span data-ty="input">docker rmi khulnasoft/meshplay:latest</span>
    </div>
    </div>
   </pre>

   <pre class="codeblock-pre" style="padding: 0; margin-top: 2px; font-size:0px;"><div class="codeblock" style="display: block;">
    <div class="clipboardjs" style="padding: 0">
    <span style="font-size:0;">docker rmi meshplay/adapters:latest</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
      <span data-ty="input">docker rmi meshplay/adapters:latest</span>
    </div>
    </div>
   </pre>

      ...and so on for other Meshplay-related images

4. **Remove Meshplay Volumes (if necessary):**

   - Meshplay may have created Docker volumes to persist data. You can list and remove these volumes using the `docker volume ls` and `docker volume rm` commands. For example:

   <pre class="codeblock-pre" style="padding: 0; margin-top: 2px; font-size:0px;"><div class="codeblock" style="display: block;">
    <div class="clipboardjs" style="padding: 0">
    <span style="font-size:0;">docker volume ls</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
      <span data-ty="input">docker volume ls</span>
    </div>
    </div>
   </pre>

   <pre class="codeblock-pre" style="padding: 0; margin-top: 2px; font-size:0px;"><div class="codeblock" style="display: block;">
    <div class="clipboardjs" style="padding: 0">
    <span style="font-size:0;">docker volume rm meshplay-data-volume</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
      <span data-ty="input">docker volume rm meshplay-data-volume</span>
    </div>
    </div>
   </pre>

    ...remove other Meshplay-related volumes if present

5. **Remove Docker Network (if necessary):**

   - If Meshplay created a custom Docker network, you can remove it using the `docker network rm` command. For example:

   <pre class="codeblock-pre" style="padding: 0; margin-top: 2px; font-size:0px;"><div class="codeblock" style="display: block;">
    <div class="clipboardjs" style="padding: 0">
    <span style="font-size:0;">docker network rm meshplay-network</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width:100%; height:200px; max-width:100%;" data-termynal="">
      <span data-ty="input">docker network rm meshplay-network</span>
    </div>
    </div>
   </pre>

6. **Clean Up Configuration (optional):**
   - If Meshplay created configuration files or directories on your host machine, you can remove them manually if you no longer need them.

<script src="{{ site.baseurl }}/assets/js/terminal.js" data-termynal-container="#termynal2"></script>

{% include suggested-reading.html language="en" %}

{% include related-discussions.html tag="meshplay" %}
