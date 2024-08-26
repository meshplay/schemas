---
layout: enhanced
title: Extensibility
permalink: extensibility
type: Extensibility
abstract: 'Meshplay architecture is extensible. Meshplay provides several extension points for working with different cloud native projects via <a href="extensibility#adapters">adapters</a>, <a href="extensibility#load-generators">load generators</a> and <a href="extensibility/providers">providers</a>.'
# redirect_from:
#   - reference/extensibility
#   - extensibility/
language: en
list: exclude
---

Meshplay has an extensible architecture with several extension points. Meshplay provides several extension points for working with different cloud and cloud native infrastructure via [adapters]({{site.baseurl}}/extensibility/adapters), different [load generators]({{site.baseurl}}/extensibility/load-generators) and different [providers]({{site.baseurl}}/extensibility/providers). Meshplay also offers a REST API.

## Extensibility Topics

{% assign sorted = site.pages | sort: "extensibility" %}

<ul>
    {% for item in sorted %}
    {% if item.type=="Extensibility" and item.list!="exclude" and item.language !="es"  -%}
      <li><a href="{{ site.baseurl }}{{ item.url }}">{{ item.title }}</a>
      {% if item.abstract != " " %}
              -  {{ item.abstract }}
            {% endif %}
            </li>
            {% endif %}
    {% endfor %}
</ul>

**Guiding Principles for Extensibility**

The following principles are upheld in the design of Meshplay's extensibility.

1. Recognize that different deployment environments have different systems to integrate with.
1. Offer a default experience that provides the optimal user experience.

## Extension Points

Meshplay is not just an application. It is a set of microservices where the central component is itself called Meshplay. Integrators may extend Meshplay by taking advantage of designated Extension Points. Extension points come in various forms and are available through Meshplay’s architecture.

![Meshplay Extension Points]({{site.baseurl}}/assets/img/architecture/meshplay_extension_points.svg)

_Figure: Extension points available throughout Meshplay_

The following points of extension are currently incorporated into Meshplay.

## Types of Extension Points

1. [Adapters]({{site.baseurl}}/extensibility/adapters)
   - Messaging Framework (CloudEvents and NATS)
1. [GraphQL API](/extensibility/api#graphql)
1. [Load Generators]({{site.baseurl}}/extensibility/load-generators)
1. [Providers]({{site.baseurl}}/extensibility/providers)
1. [REST API](/extensibility/api#rest)
1. [UI Plugins](extensibility/ui)
1. [Integrations](/extensibility/integrations)

