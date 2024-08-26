---
layout: default
title: Catalog
permalink: concepts/catalog
redirect_from: catalog
type: components
abstract: Browsing and using cloud native patterns
language: en
list: include
---

[Meshplay Catalog](https://khulnasoft.com/catalog) functions much like a cloud marketplace, providing a user-friendly interface for browsing, discovering, and sharing configurations and patterns for cloud native infrastructure. With Meshplay Catalog, you can easily find and deploy Kubernetes-based infrastructure and tools, making it easy to streamline your cloud native development and operations processes. A Catolog is based on the Meshplay's [Catalog Schema](https://github.com/khulnasoft/schemas/blob/master/openapi/schemas/catalog.yml) with defined attributes.

### Simplify Your Cloud Native Infrastructure Deployment and Management

Meshplay Catalog functions much like a cloud marketplace, providing a user-friendly interface for browsing, discovering, and sharing configurations and patterns for cloud native infrastructure. With Meshplay Catalog, you can easily find and deploy Kubernetes-based infrastructure and tools, making it easy to streamline your cloud native development and operations processes.

It also supports a collaborative environment, where DevOps engineers can share their experiences, feedback, and best practices with others in the community. Import cloud native patterns published by others into your Meshplay Server. Benefit from and build upon each pattern by incorporating your own tips and tricks, then publish and share with the community at-large. This facilitates knowledge-sharing and helps to build a strong ecosystem of cloud native infrastructure experts.


### To create a design pattern using Meshplay UI

1. Open the [Meshplay UI](https://docs.khulnasoft.com/installation/quick-start) in your web browser.
2. Navigate to the configuration section, usually located in the main navigation menu.
3. Head over to Designs and click on import or create design.
4. Select the category and Model as per your need and configure the application.
5. Voilà, You can publish or deploy you design.


### To create design pattern using Meshplay CLI

1. Ensure that you have [Meshplay CLI](https://docs.khulnasoft.com/installation/meshplayctl) installed on your machine and it is configured to connect to your desired Meshplay instance.
2. Open a terminal or command prompt.
3. Use the Meshplay CLI commands to interact with the catalog. `meshplayctl pattern`
4. Follow the prompts or instructions provided by the Meshplay CLI help.
* Apply [pattern file](https://docs.khulnasoft.com/guides/configuration-management):  `meshplayctl pattern apply --file [path to pattern file | URL of the file]`
* Delete pattern file:  `meshplayctl pattern delete --file [path to pattern file]`
* View pattern file:  `meshplayctl pattern view [pattern name | ID]`
* List all patterns: `meshplayctl pattern list`
5. [Onboarding](managing-applications-through-meshplay-cli) an application. `meshplayctl app onboard -f [file-path]`
6. Applying [WASM Filter](https://docs.khulnasoft.com/guides/configuration-management#wasm-filters). `meshplayctl exp filter apply --file [GitHub Link]`


{% include alert.html
    type="info"
    title="Help with Meshplay Catalog"
    content="If you have any questions or need assistance, please refer to the [Meshplay Documentation](https://docs.khulnasoft.com/) or reach out to our discussion form [khulnasoft.com](http://discuss.khulnasoft.com/)." %}
