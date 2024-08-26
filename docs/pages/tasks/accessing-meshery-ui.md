---
layout: default
title: Accessing Meshplay UI
permalink: tasks/accessing-meshplay-ui
type: tasks
language: en
list: include
abstract: "How to find where Meshplay UI is exposed from your Kubernetes cluster."
---

To access Meshplay’s UI via port-forwarding, please refer to the port-forwarding guide for detailed instructions.

Use `meshplayctl system dashboard` to open your default browser to Meshplay UI, [click here](/reference/meshplayctl/system/dashboard) to see the reference.


## Docker Desktop

Your default browser will be opened and directed to Meshplay's web-based user interface typically found at `http://localhost:9081`.


## Kubernetes

Access Meshplay UI by exposing it as a Kubernetes service or by port forwarding to Meshplay UI.

#### [Optional] Port Forward to Meshplay UI

{% capture code_content %}kubectl port-forward svc/meshplay 9081:9081 --namespace meshplay{% endcapture %}
{% include code.html code=code_content %}
