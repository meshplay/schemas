
<p style="text-align:center;" align="center"><a href="https://meshplay.khulnasoft.com"><picture>
 <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/meshplay/meshplay/master/.github/assets/images/readme/meshplay-logo-light-text-side.svg">
 <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/meshplay/meshplay/master/.github/assets/images/readme/meshplay-logo-dark-text-side.svg">
<img src="https://raw.githubusercontent.com/meshplay/meshplay/master/.github/assets/images/readme/meshplay-logo-dark-text-side.svg"
alt="Meshplay Logo" width="70%" /></picture></a><br /><br /></p>
<p align="center">
<a href="https://hub.docker.com/r/khulnasoft/meshplay" alt="Docker pulls">
  <img src="https://img.shields.io/docker/pulls/khulnasoft/meshplay.svg" /></a>
<a href="https://github.com/issues?q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Akhulnasoft+org%3Ameshplay+org%3Aservice-mesh-performance+org%3Aservice-mesh-patterns+org%3A+label%3A%22help+wanted%22+" alt="GitHub issues by-label">
  <img src="https://img.shields.io/github/issues/khulnasoft/meshplay/help%20wanted.svg?color=informational" /></a>
<a href="https://github.com/meshplay/meshplay/blob/master/LICENSE" alt="LICENSE">
  <img src="https://img.shields.io/github/license/meshplay/meshplay?color=brightgreen" /></a>
<a href="https://artifacthub.io/packages/helm/meshplay/meshplay" alt="Artifact Hub Meshplay">
  <img src="https://img.shields.io/endpoint?color=brightgreen&label=Helm%20Chart&style=plastic&url=https%3A%2F%2Fartifacthub.io%2Fbadge%2Frepository%2Fartifact-hub" /></a>  
<a href="https://goreportcard.com/report/github.com/meshplay/meshplay" alt="Go Report Card">
  <img src="https://goreportcard.com/badge/github.com/meshplay/meshplay" /></a>
<a href="https://github.com/meshplay/meshplay/actions" alt="Build Status">
  <img src="https://img.shields.io/github/actions/workflow/status/meshplay/meshplay/release-drafter.yml" /></a>
<a href="https://bestpractices.coreinfrastructure.org/projects/3564" alt="CLI Best Practices">
  <img src="https://bestpractices.coreinfrastructure.org/projects/3564/badge" /></a>
<a href="http://discuss.meshplay.khulnasoft.com" alt="Discuss Users">
  <img src="https://img.shields.io/discourse/users?label=discuss&logo=discourse&server=https%3A%2F%2Fdiscuss.khulnasoft.com" /></a>
<a href="https://slack.meshplay.khulnasoft.com" alt="Join Slack">
  <img src="https://img.shields.io/badge/Slack-@khulnasoft.svg?logo=slack" /></a>
<a href="https://twitter.com/intent/follow?screen_name=meshplayio" alt="Twitter Follow">
  <img src="https://img.shields.io/twitter/follow/meshplayio.svg?label=Follow+Meshplay&style=social" /></a>
<a href="https://github.com/meshplay/meshplay/releases" alt="Meshplay Downloads">
  <img src="https://img.shields.io/github/downloads/meshplay/meshplay/total" /></a>  
<!-- <a href="https://app.fossa.com/projects/git%2Bgithub.com%2Fmeshplay%2Fmeshplay?ref=badge_shield" alt="License Scan Report">
  <img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmeshplay%2Fmeshplay.svg?type=shield"/></a>  
  -->
</p>

<h5><p align="center"><i>If you’re using Meshplay or if you like the project, please <a href="https://github.com/meshplay/meshplay/stargazers">★</a> this repository to show your support! 🤩</i></p></h5>

# Meshplay Schemas

Meshplay follows schema-driven development. As a project, Meshplay has different types of schemas. Some schemas are external facing, and some internal to Meshplay itself. This repository serves as a central location for storing schemas from which all Meshplay components can take reference.


<!-- The schema.go emabeds the openapi schema which gets packaged & released  used for purpose like validation

We can refer the unresolved schemas, but
1. It increases the size of the pkg as it will then embed multiple dirs.
2. Resolution of refs at run time is ineffective. And because every request will be valiated, it is better to pre-process the schema.
-->

<!-- For code generation (schema to golang structs), unresolved schemas should be used, and proper import mappings needs to be provided.
(In some cases, first level resolution of schemas might be required.)
 -->
### External

Meshplay schemas file structure is defined based on definitions and schemas, checkout [docs.meshplay.khulnasoft.com](https://docs.meshplay.khulnasoft.com/concepts/logical) to learn more about definitions and schemas.

Definitions
- model
  - version
    - model.definition
    - components
      - component-1.definition
      - component-2.definition
    - policy.definition
    - relationship.definition

Schemas
- constructs
  - schema.version // Schema version
    - component.schema
    - model.schema
    - policy.schema
    - relationship.schema

REST API
 - swagger.yaml

Adapters
- meshes.proto
 
<p style="clear:both;">&nbsp;</p>

## Join the Meshplay community!

<a name="contributing"></a><a name="community"></a>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://khulnasoft.com/community/newcomers">Contributor Journey Map</a> and <a href="https://khulnasoft.com/community/handbook">Community Handbook</a> for a tour of resources available to you and the <a href="https://khulnasoft.com/community/handbook/repository-overview">Repository Overview</a> for a cursory description of repository by technology and programming language. Jump into community <a href="https://slack.meshplay.khulnasoft.com">Slack</a> or <a href="http://discuss.meshplay.khulnasoft.com">discussion forum</a> to participate.

<p style="clear:both;">
<a href ="https://khulnasoft.com/community"><img alt="MeshMates" src=".github/assets/images/readme/khulnasoft-community-sign.png" style="margin-right:36px; margin-bottom:7px;" width="140px" align="left" /></a>
<h3>Find your MeshMate</h3>

<p>MeshMates are experienced KhulnaSoft community members, who will help you learn your way around, discover live projects, and expand your community network. Connect with a Meshmate today!</p>

Find out more on the <a href="https://khulnasoft.com/community#meshmate">KhulnaSoft community</a>. <br />

</p>
<br /><br />
<div style="display: flex; justify-content: center; align-items:center;">
<div>
<a href="https://meshplay.khulnasoft.com/community"><img alt="KhulnaSoft Cloud Native Community" src="https://docs.meshplay.khulnasoft.com/assets/img/readme/community.png" width="140px" style="margin-right:36px; margin-bottom:7px;" width="140px" align="left"/></a>
</div>
<div style="width:60%; padding-left: 16px; padding-right: 16px">
<p>
✔️ <em><strong>Join</strong></em> any or all of the weekly meetings on <a href="https://meshplay.khulnasoft.com/calendar">community calendar</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/playlist?list=PL3A-A6hPO2IMPPqVjuzgqNU5xwnFFn3n0">meeting recordings</a>.<br />
✔️ <em>Fill-in</em> a <a href="https://khulnasoft.com/newcomers">community member form</a> to gain access to community resources.
<br />
✔️ <em><strong>Discuss</strong></em> in the <a href="http://discuss.meshplay.khulnasoft.com">Community Forum</a>.<br />
✔️ <em><strong>Explore more</strong></em> in the <a href="https://khulnasoft.com/community/handbook">Community Handbook</a>.<br />
✔️ <em><strong>Not sure where to start?</strong></em> Grab an open issue with the <a href="https://github.com/issues?q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Akhulnasoft+org%3Ameshplay+org%3Aservice-mesh-performance+org%3Aservice-mesh-patterns+label%3A%22help+wanted%22+">help-wanted label</a>.
</p>
</div>
<div>&nbsp;</div>

## Contributing

Please do! We're a warm and welcoming community of open source contributors. Please join. All types of contributions are welcome. Be sure to read the [Contributor Guides](https://docs.meshplay.khulnasoft.com/project/contributing) for a tour of resources available to you and how to get started.

<div>&nbsp;</div>

### License

This repository and site are available as open-source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).

<p align="center" >
MESHPLAY IS A CLOUD NATIVE COMPUTING FOUNDATION PROJECT
</p>
