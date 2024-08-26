---
layout: default
title: SOFAmesh
name: Meshplay Adapter for SOFAmesh
project_status: 
visibility: hidden
language: en
---

{% assign sorted_tests_group = site.compatibility | group_by: "meshplay-component" %}
{% for group in sorted_tests_group %}
      {% if group.name == "meshplay-sofamesh" %}
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

### Lifecycle management

The {{page.name}} can install **{{page.earliest_version}}** of the {{page.mesh_name}} service mesh. A number of sample applications for {{page.mesh_name}} can also be installed using Meshplay.

### Suggested Topics

- Examine [Meshplay's architecture]({{ site.baseurl }}/architecture) and how adapters fit in as a component.
- Learn more about [Meshplay Adapters]({{ site.baseurl }}/architecture/adapters).
