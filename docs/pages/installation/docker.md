---
layout: default
title: Docker
permalink: installation/docker
type: installation
category: docker
redirect_from:
- installation/platforms/docker
display-title: "false"
language: en
list: include
image: /assets/img/platforms/docker.svg
abstract: Install Meshplay on Docker
---

{% include installation_prerequisites.html %}

## Deploying Meshplay on Docker

Follow these installation steps to use Docker and Docker Compose to run Meshplay. Users often choose this installation approach in order to run Meshplay on their local machine. If you need to install *docker*, see [Getting Started with Docker](https://docs.docker.com/get-started/) and if you need to install *docker-compose*, see [Installing Docker Compose](https://docs.docker.com/compose/install/). 

Start Meshplay by executing:

{% capture code_content %}meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}
 - Meshplay server supports customizing authentication flow callback URL, which can be configured in the following way
{% capture code_content %}MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

Once you have verified that all the services are up and running, Meshplay UI will be accessible on your local machine on port 9081. Open your browser and access Meshplay at [`http://localhost:9081`](http://localhost:9081). Log into the [Provider](/extensibility/providers) of your choice.

Upon starting Meshplay successfully, instructions to access Meshplay will be printed on the screen. See these [guides]({{ site.baseurl }}/guides) for usage tips.

<pre class="codeblock-pre" style="padding: 0; font-size:0px;"><div class="codeblock" style="display: block;">
 <div class="clipboardjs" style="padding: 0">
 <span style="font-size:0;">curl -L https://khulnasoft.com/install | PLATFORM=docker bash -</span>  
 </div>
 <div class="window-buttons"></div>
 <div id="termynal1" style="width:100%; height:150px; max-width:100%;" data-termynal="">
            <span data-ty="input">curl -L https://khulnasoft.com/install | PLATFORM=docker bash -</span>
            <span data-ty="progress"></span>
            <span data-ty="">Successfully installed Meshplay</span>
  </div>
 </div>
 </pre>

 <script src="/assets/js/terminal.js" data-termynal-container="#termynal1"></script>

 {% include suggested-reading.html language="en" %}

{% include related-discussions.html tag="meshplay" %}
