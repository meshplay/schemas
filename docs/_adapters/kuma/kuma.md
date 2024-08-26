---
layout: page
title: Meshplay Adapter for Kuma
name: Meshplay Adapter for Kuma
mesh_name: Kuma
earliest_version: v1.2.2
port: 10007/gRPC
project_status: stable
lab: kuma-meshplay-adapter
github_link: https://github.com/khulnasoft/meshplay-kuma
image: /assets/img/service-meshes/kuma.svg
white_image: /assets/img/service-meshes/kuma-white.svg
permalink: extensibility/adapters/kuma
language: en
---
{% assign sorted_tests_group = site.compatibility | group_by: "meshplay-component" %}
{% for group in sorted_tests_group %}
      {% if group.name == "meshplay-kuma" %}
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

{% include adapter-labs.html %}

## Lifecycle management

The {{page.name}} can install **{{page.earliest_version}}** of {{page.mesh_name}} service mesh. A number of sample applications can be installed using the {{page.name}}.

### Install {{ page.mesh_name }}

Choose the Meshplay Adapter for {{ page.mesh_name }}.

<a href="{{ site.baseurl }}/assets/img/adapters/kuma/kuma-adapter.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/kuma/kuma-adapter.png" />
</a>

Click on (+) and choose the {{page.earliest_version}} of the {{page.mesh_name}} service mesh.

<a href="{{ site.baseurl }}/assets/img/adapters/kuma/kuma-install.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/kuma/kuma-install.png" />
</a>

## Workload Management

The following sample applications are available in this adapter.

- [Bookinfo]({{site.baseurl}}/guides/sample-apps#bookinfo)
  - The sample BookInfo application displays information about a book, similar to a single catalog entry of an online book store.

### SMI Conformance Capability

Meshplay allows you to analyze the compliance status and functional capabilities of your service mesh. This allows you to compare high-level functional differences between service meshes and verify whether your service mesh is conformant with the SMI specification.

Learn more about the SMI specification and [Meshplay's conformance test suite]({{ site.baseurl }}/tasks/service-mesh-interface).
