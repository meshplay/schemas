---
layout: default
title: Operation Guides
permalink: guides
redirect_from: guides/
language: en
list: exclude
abstract: Operating and troubleshooting Meshplay deployments.
---

Guides to using and troubleshooting Meshplay's various features and components.

{% assign sorted_guides = site.pages | sort: "type" | reverse %}

### General

<ul>
    {% for item in sorted_guides %}
    {% if item.type=="guides" and item.category!="meshplayctl" and item.list!="exclude" and item.language=="en"  -%}
      <li><a href="{{ site.baseurl }}{{ item.url }}">{{ item.title }}</a></li>
      {% endif %}
    {% endfor %}
</ul>

### <a href="{{ site.baseurl }}/guides/meshplayctl" class="text-black">Meshplay CLI</a>

<ul>
  {% for item in sorted_guides %}
  {% if item.type=="guides" and item.category=="meshplayctl" and item.list!="exclude" and item.language=="en" -%}
    <li><a href="{{ site.baseurl }}{{ item.url }}">{{ item.title }}</a>
    </li>
    {% endif %}
  {% endfor %}
</ul>

<!-- {% comment %}
#
#  Change date order by adding '| reversed'
#  To sort by title or other variables use {% assign sorted_posts = category[1] | sort: 'title' %}
#
{% endcomment %}

{% for guide in site.adapter %}
<h2 id="{{guide[0] | uri_escape | downcase }}">{{guide[0] | capitalize}}1</h2>

{% endfor %}

{% assign sorted_guides = site.guides | sort %}
{% for guide in sorted_guides %}
<h2 id="{{guide[0] | uri_escape | downcase }}">{{guide[0] | capitalize}}</h2>

{% endfor %} -->
