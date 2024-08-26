---
layout: default
title: Meshplay Adapter for Cilium Service Mesh
name: Meshplay Adapter for Cilium Service Mesh
mesh_name: Cilium
earliest_version: v1.10.6
port: 10012/gRPC
project_status: stable
lab: cilium-meshplay-adapter
github_link: https://github.com/khulnasoft/meshplay-cilium
image: /assets/img/service-meshes/cilium.svg
white_image: /assets/img/service-meshes/cilium-white.svg
permalink: extensibility/adapters/cilium
language: en
---

{% assign sorted_tests_group = site.compatibility | group_by: "meshplay-component" %}
{% for group in sorted_tests_group %}
      {% if group.name == "meshplay-cilium" %}
        {% assign items = group.items | sort: "meshplay-component-version" | reverse %}
        {% for item in items %}
          {% if item.meshplay-component-version != "edge" %}
            {% if item.overall-status == "passing" %}
              {% assign adapter_version_dynamic = item.meshplay-component-version %}
              {% break %}
            {% elsif item.overall-status == "failing" %}
              {% continue %}
            {% endif %}
          {% endif %}
        {% endfor %} 
      {% endif %}
{% endfor %}

{% include compatibility/adapter-status.html %}

### Features

1. {{page.mesh_name}} Lifecycle Management
1. Workload Lifecycle Management
   1. Using Service Mesh Standards
      1. Service Mesh Performance (SMP)
         1. Prometheus and Grafana connections
      1. Service Mesh Interface (SMI)
1. Configuration Analysis, Patterns, and Best Practices
   1. Custom Service Mesh Configuration

## Lifecycle management

The {{page.name}} can install **{{page.earliest_version}}** of {{page.mesh_name}} service mesh. A number of sample applications can be installed using the {{page.name}}.

The {{ page.name }} is currently under construction ({{ page.project_status }} state), which means that the adapter is not functional and cannot be interacted with through the <a href="{{ site.baseurl }}/installation#6-you-will-now-be-directed-to-the-meshplay-ui"> Meshplay UI </a>at the moment. Check back here to see updates.

Want to contribute? Check our [progress]({{page.github_link}}).

### Install {{ page.mesh_name }}

##### Choose the Meshplay Adapter for {{ page.mesh_name }}

<a href="{{ site.baseurl }}/assets/img/adapters/linkerd/linkerd-adapter.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/linkerd/linkerd-adapter.png" />
</a>

##### Click on (+) and choose the {{page.earliest_version}} of the {{page.mesh_name}} service mesh.

<a href="{{ site.baseurl }}/assets/img/adapters/linkerd/linkerd-install.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/linkerd/linkerd-install.png" />
</a>

