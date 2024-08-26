---
layout: enhanced
title: Components
permalink: concepts/logical/components
type: concepts
abstract: "Meshplay Components identify and characterize infrastructure under management."
language: en
list: include
---

[Components](https://github.com/khulnasoft/meshplay/tree/master/server/meshmodel/components) represent entities in the Meshplay ecosystem, exposing capabilities of the underlying platform. They can be registered, created, and used by users and operators. Components have definitions, instances, and associated metadata. Components having the same `kind`, `apiVersion` and `model.name` attributes are considered duplicates.

[![Meshplay Components]({{ site.baseurl }}/assets/img/architecture/meshplay-components.svg
)]({{ site.baseurl }}/assets/img/architecture/meshplay-components.svg)

## Component Status

Components have a status that is represented as a `Connection` object. Both the administrative and real-time status of a component is a normalized representation of the connection's state. The status is represented as a `Connection` object because the status of a component is a *connection* to the component. For example, the status of a Kubernetes cluster is a direct reflection of a Meshplay Server's connection to the cluster.

Normalizing and extracting the status of a component as a direct property of the component and putting it into a connection allows multiple systems to share the same component with different states. For example, different Meshplay Servers can access the same Kubernetes cluster, but each Meshplay Server has its own connection to the cluster with its own status.

Learn more about [Connections](/concepts/logical/connections).
