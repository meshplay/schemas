---
layout: default
title: meshplayctl-exp-filter-delete
permalink: reference/meshplayctl/exp/filter/delete
redirect_from: reference/meshplayctl/exp/filter/delete/
type: reference
display-title: "false"
language: en
command: exp
subcommand: filter
---

# meshplayctl exp filter delete

Delete filter file

## Synopsis

delete filter file will trigger deletion of the filter file

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp filter delete [flags]

</div>
</pre> 

## Examples

Delete the specified WASM filter file using name or ID
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp filter delete [filter-name | ID]

</div>
</pre> 

Delete using the file name
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp filter delete test-wasm

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for delete

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
