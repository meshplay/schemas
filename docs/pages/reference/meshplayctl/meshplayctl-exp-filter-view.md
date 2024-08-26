---
layout: default
title: meshplayctl-exp-filter-view
permalink: reference/meshplayctl/exp/filter/view
redirect_from: reference/meshplayctl/exp/filter/view/
type: reference
display-title: "false"
language: en
command: exp
subcommand: filter
---

# meshplayctl exp filter view

Display filters(s)

## Synopsis

Displays the contents of a specific filter based on name or id

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp filter view [filter name] [flags]

</div>
</pre> 

## Examples

View the specified WASM filter file
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp filter view [filter-name | ID]	

</div>
</pre> 

View using filter name
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp filter view test-wasm

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -a, --all                    (optional) view all filters available
  -h, --help                   help for view
  -o, --output-format string   (optional) format to display in [json|yaml] (default "yaml")

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string   path to config file (default "/home/runner/.meshplay/config.yaml")
  -t, --token string    Path to token file default from current context
  -v, --verbose         verbose output

</div>
</pre>

## See Also

Go back to [command reference index](/reference/meshplayctl/) 
