---
layout: default
title: Meshplay Adapter for Tanzu Service Mesh
name: Meshplay Adapter for Tanzu Service Mesh
mesh_name: Tanzu Service Mesh
earliest_version: pre-GA
port: 10011/gRPC
project_status: alpha
github_link: https://github.com/khulnasoft/meshplay-tanzu-sm
image: /assets/img/service-meshes/tanzu.svg
white_image: /assets/img/service-meshes/tanzu.svg
permalink: extensibility/adapters/tanzu-sm
language: en
---

{% assign sorted_tests_group = site.compatibility | group_by: "meshplay-component" %}
{% for group in sorted_tests_group %}
      {% if group.name == "meshplay-tanzu-mesh" %}
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
## Features

1. {{page.mesh_name}} Lifecycle Management
1. Workload Lifecycle Management
   1. Using Service Mesh Standards
      1. Service Mesh Performance (SMP)
         1. Prometheus and Grafana connections
      1. Service Mesh Interface (SMI)
1. Configuration Analysis, Patterns, and Best Practices
   1. Custom Service Mesh Configuration

## Lifecycle management

The {{page.name}} can install **{{page.earliest_version}}** of {{page.mesh_name}} service mesh. A number of sample applications for {{page.mesh_name}} can also be installed using Meshplay.

The {{ page.name }} is currently under construction ({{ page.project_status }} state), which means that the adapter is not functional and cannot be interacted with through the <a href="{{ site.baseurl }}/installation#6-you-will-now-be-directed-to-the-meshplay-ui"> Meshplay UI </a>at the moment. Check back here to see updates.

Want to contribute? Check our [progress]({{page.github_link}}).
## Workload Management

The Meshplay Adapter for {{ page.name }} includes some sample applications operations. Meshplay can be used to deploy any of these sample applications.  

- [BookInfo](https://github.com/khulnasoft/istio-service-mesh-workshop/blob/master/lab-2/README.md#what-is-the-bookinfo-application)
    - This application is a polyglot composition of microservices are written in different languages and sample BookInfo application displays information about a book, similar to a single catalog entry of an online book store.
- [httpbin](https://httpbin.org)
    - This is a simple HTTP Request & Response Service.
- [hipster](https://github.com/GoogleCloudPlatform/microservices-demo)
    - Hipster Shop Application is a web-based, e-commerce demo application from the Google Cloud Platform.

### Suggested Reading

- Examine [Meshplay's architecture]({{ site.baseurl }}/architecture) and how adapters fit in as a component.
- Learn more about [Meshplay Adapters]({{ site.baseurl }}/architecture/adapters).
