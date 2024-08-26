---
layout: default
title: meshplayctl-app-list
permalink: reference/meshplayctl/app/list
redirect_from: reference/meshplayctl/app/list/
type: reference
display-title: "false"
language: en
command: app
subcommand: list
---

# meshplayctl app list

List applications

## Synopsis

Display list of all available applications.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl app list [flags]

</div>
</pre> 

## Examples

List all the applications
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl app list

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help      help for list
  -v, --verbose   Display full length user and app file identifiers

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string   path to config file (default "/home/runner/.meshplay/config.yaml")
  -t, --token string    Path to token file default from current context

</div>
</pre>

## Screenshots

Usage of meshplayctl app list
![app-list-usage](/assets/img/meshplayctl/app-list.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
