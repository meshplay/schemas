---
layout: default
title: Command Line Reference
abstract: "A guide to Meshplay's CLI: meshplayctl"
permalink: reference/meshplayctl
redirect_from:
  - reference/meshplayctl/commands/
  - reference/meshplayctl/commands
  - reference/meshplayctl/
type: Reference
language: en
abstract: "A guide to Meshplay's CLI: meshplayctl"
---

## Categories and Command Structure

Meshplay CLI commands are categorized by function, which are:

- `meshplayctl` - Global flags and CLI configuration
- `meshplayctl system` - Meshplay Lifecycle and Troubleshooting
- `meshplayctl mesh` - Cloud Native Lifecycle & Configuration Management: provisioning and configuration best practices
- `meshplayctl perf` - Cloud Native Performance Management: Workload and cloud native performance characterization
- `meshplayctl pattern` - Cloud Native Pattern Configuration & Management: cloud native patterns and Open Application Model integration
- `meshplayctl app` - Cloud Native Application Management
- `meshplayctl filter` - Data Plane Intelligence: Registry and configuration of WebAssembly filters for Envoy

## Global Commands and Flags

<table>
<thead>
  <tr>
    <th>Command</th>
    <th>Subcommand</th>
    <th>Flag</th>
    <th>Function</th>
  </tr>
  {% assign command1 = site.data.meshplayctlcommands.cmds.global %}
    <tr>
      <td rowspan=6><a href="{{ site.baseurl }}/reference/meshplayctl/main">{{ command1.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command1.description }}</td>
    </tr>
  {% for flag_hash in command1.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td></td>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
  {% endfor %}
  {% assign subcommand1 = command1.subcommands.version %}
    <tr>
      <td><a href="{{ site.baseurl }}/reference/meshplayctl/version">{{ subcommand1.name }}</a></td>
      <td></td>
      <td>{{ subcommand1.description }}</td>
    </tr>
  {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
  {% endfor %}
  {% assign subcommand2 = command1.subcommands.completion %}
  <tr>
    <td><a href="{{ site.baseurl }}/reference/meshplayctl/completion">{{ subcommand2.name }}</a></td>
    <td></td>
    <td>{{ subcommand2.description }}</td>
  </tr>
  {% for flag_hash in subcommand2.flag %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
  {% endfor %}
</thead>
</table>

## Meshplay Lifecycle Management and Troubleshooting

Installation, troubleshooting and debugging of Meshplay and its adapters

<table>
<thead>
  <tr>
    <th>Main Command</th>
    <th>Arguments</th>
    <th>Flag</th>
    <th>Function</th>
  </tr>
  {% assign command2 = site.data.meshplayctlcommands.cmds.system %}
    <tr>
      <td rowspan=34><a href="{{ site.baseurl }}/reference/meshplayctl/system">{{ command2.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command2.description }}</td>
    </tr>
    {% for flag_hash in command2.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td></td>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand1 = command2.subcommands.start %}
      <tr>
        <td rowspan=5><a href="{{ site.baseurl }}/reference/meshplayctl/system/start">{{ subcommand1.name }}</a></td>
        <td></td>
        <td>{{ subcommand1.description }}</td>
      </tr>
    {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand2 = command2.subcommands.stop %}
    <tr>
      <td rowspan=4><a href="{{ site.baseurl }}/reference/meshplayctl/system/stop">{{ subcommand2.name }}</a></td>
      <td></td>
      <td>{{ subcommand2.description }}</td>
    </tr>
    {% for flag_hash in subcommand2.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand4 = command2.subcommands.update %}
    <tr>
      <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/system/update">{{ subcommand4.name }}</a></td>
      <td></td>
      <td>{{ subcommand4.description }}</td>
    </tr>
    {% for flag_hash in subcommand4.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand5 = command2.subcommands.config %}
    <tr>
      <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/system/config">{{ subcommand5.name }}</a></td>
      <td></td>
      <td>{{ subcommand5.description }}</td>
    </tr>
    {% for flag_hash in subcommand5.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand7 = command2.subcommands.reset %}
    <tr>
      <td rowspan="1"><a href="{{ site.baseurl }}/reference/meshplayctl/system/reset">{{ subcommand7.name }}</a></td>
      <td></td>
      <td>{{ subcommand7.description }}</td>
    </tr>
    {% for flag_hash in subcommand7.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand7 = command2.subcommands.logs %}
    <tr>
      <td rowspan="2"><a href="{{ site.baseurl }}/reference/meshplayctl/system/logs">{{ subcommand7.name }}</a></td>
      <td></td>
      <td>{{ subcommand7.description }}</td>
    </tr>
    {% for flag_hash in subcommand7.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand8 = command2.subcommands.restart %}
    <tr>
      <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/system/restart">{{ subcommand8.name }}</a></td>
      <td></td>
      <td>{{ subcommand8.description }}</td>
    </tr>
    {% for flag_hash in subcommand8.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand9 = command2.subcommands.status %}
    <tr>
      <td rowspan="2"><a href="{{ site.baseurl }}/reference/meshplayctl/system/status">{{ subcommand9.name }}</a></td>
      <td></td>
      <td>{{ subcommand9.description }}</td>
    </tr>
    {% for flag_hash in subcommand9.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand13 = command2.subcommands.dashboard %}
    <tr>
      <td rowspan="3"><a href="{{ site.baseurl }}/reference/meshplayctl/system/dashboard">{{ subcommand13.name }}</a></td>
      <td></td>
      <td>{{ subcommand13.description }}</td>
    </tr>
    {% for flag_hash in subcommand13.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand10 = command2.subcommands.login %}
    <tr>
      <td rowspan="2"><a href="{{ site.baseurl }}/reference/meshplayctl/system/login">{{ subcommand10.name }}</a></td>
      <td></td>
      <td>{{ subcommand10.description }}</td>
    </tr>
    {% for flag_hash in subcommand10.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
    {% assign subcommand11 = command2.subcommands.logout %}
    <tr>
      <td rowspan="1"><a href="{{ site.baseurl }}/reference/meshplayctl/system/logout">{{ subcommand11.name }}</a></td>
      <td></td>
      <td>{{ subcommand11.description }}</td>
    </tr>
    {% for flag_hash in subcommand11.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign subcommand12 = command2.subcommands.check %}
    <tr>
      <td rowspan=6><a href="{{ site.baseurl }}/reference/meshplayctl/system/check">{{ subcommand12.name }}</a></td>
      <td></td>
      <td>{{ subcommand12.description }}</td>
    </tr>
    {% for flag_hash in subcommand12.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign command3 = site.data.meshplayctlcommands.cmds.system-channel %}
    <tr>
      <td rowspan=5><a href="{{ site.baseurl }}/reference/meshplayctl/system/channel">{{ command3.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command3.description }}</td>
    </tr>
    {% assign subcommand1 = command3.subcommands.set %}
    <tr>
      <td rowspan=1><a href="{{ site.baseurl }}/reference/meshplayctl/system/channel/set">{{ subcommand1.name }}</a></td>
      <td></td>
      <td>{{ subcommand1.description }}</td>
    </tr>
    {% assign subcommand2 = command3.subcommands.switch %}
    <tr>
      <td rowspan=1><a href="{{ site.baseurl }}/reference/meshplayctl/system/channel/switch">{{ subcommand2.name }}</a></td>
      <td></td>
      <td>{{ subcommand2.description }}</td>
    </tr>
    {% assign subcommand3 = command3.subcommands.view %}
    <tr>
      <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/system/channel/view">{{ subcommand3.name }}</a></td>
      <td></td>
      <td>{{ subcommand3.description }}</td>
    </tr>
    {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign command4 = site.data.meshplayctlcommands.cmds.system-context %}
    <tr>
      <td rowspan=12><a href="{{ site.baseurl }}/reference/meshplayctl/system/context">{{ command4.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command4.description }}</td>
    </tr>
    {% assign subcommand1 = command4.subcommands.create %}
    <tr>
      <td rowspan=5><a href="{{ site.baseurl }}/reference/meshplayctl/system/context/create">{{ subcommand1.name }}</a></td>
      <td></td>
      <td>{{ subcommand1.description }}</td>
    </tr>
    {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign subcommand2 = command4.subcommands.delete %}
    <tr>
      <td><a href="{{ site.baseurl }}/reference/meshplayctl/system/context/delete">{{ subcommand2.name }}</a></td>
      <td></td>
      <td>{{ subcommand2.description }}</td>
    </tr>
    {% assign subcommand3 = command4.subcommands.view %}
    <tr>
      <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/system/context/view">{{ subcommand3.name }}</a></td>
      <td></td>
      <td>{{ subcommand3.description }}</td>
    </tr>
    {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign subcommand4 = command4.subcommands.switch %}
    <tr>
      <td><a href="{{ site.baseurl }}/reference/meshplayctl/system/context/switch">{{ subcommand4.name }}</a></td>
      <td></td>
      <td>{{ subcommand4.description }}</td>
    </tr>
    {% for flag_hash in subcommand4.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign subcommand5 = command4.subcommands.list %}
    <tr>
      <td><a href="{{ site.baseurl }}/reference/meshplayctl/system/context/list">{{ subcommand5.name }}</a></td>
      <td></td>
      <td>{{ subcommand5.description }}</td>
    </tr>
    {% for flag_hash in subcommand5.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign command5 = site.data.meshplayctlcommands.cmds.system-provider %}
    <tr>
      <td rowspan=8><a href="{{ site.baseurl }}/reference/meshplayctl/system/provider">{{ command5.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command5.description }}</td>
    </tr>
    {% assign subcommand2 = command5.subcommands.list %}
    <tr>
      <td><a href="{{ site.baseurl }}/reference/meshplayctl/system/provider/list">{{ subcommand2.name }}</a></td>
      <td></td>
      <td>{{ subcommand2.description }}</td>
    </tr>
    {% assign subcommand3 = command5.subcommands.reset %}
    <tr>
      <td><a href="{{ site.baseurl }}/reference/meshplayctl/system/provider/reset">{{ subcommand3.name }}</a></td>
      <td></td>
      <td>{{ subcommand3.description }}</td>
    </tr>
    {% assign subcommand4 = command5.subcommands.switch %}
    <tr>
      <td><a href="{{ site.baseurl }}/reference/meshplayctl/system/provider/switch">{{ subcommand4.name }}</a></td>
      <td></td>
      <td>{{ subcommand4.description }}</td>
    </tr>
    {% assign subcommand5 = command5.subcommands.view %}
    <tr>
      <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/system/provider/view">{{ subcommand5.name }}</a></td>
      <td></td>
      <td>{{ subcommand5.description }}</td>
    </tr>
    {% for flag_hash in subcommand5.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
    {% assign subcommand1 = command5.subcommands.set %}
    <tr>
      <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/system/provider/set">{{ subcommand1.name }}</a></td>
      <td></td>
      <td>{{ subcommand1.description }}</td>
    </tr>
    {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
    <tr>
      <td>{{ flag.name }}</td>
      <td>{{ flag.description }}</td>
    </tr>
    {% endfor %}
</thead>
</table>

## Cloud Native Performance Management

<table>
<thead>
  <tr>
    <th>Main Command</th>
    <th>Arguments</th>
    <th>Flag</th>
    <th>Function</th>
  </tr>
  {% assign command5 = site.data.meshplayctlcommands.cmds.perf %}
    <tr>
      <td rowspan=20><a href="{{ site.baseurl }}/reference/meshplayctl/perf">{{ command5.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command5.description }}</td>
    </tr>
    {% for flag_hash in command5.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td></td>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
    {% endfor %}
  {% assign subcommand1 = command5.subcommands.apply %}
      <tr>
        <td rowspan=11><a href="{{ site.baseurl }}/reference/meshplayctl/perf/apply">{{ subcommand1.name }}</a></td>
        <td></td>
        <td>{{ subcommand1.description }}</td>
      </tr>
  {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
  {% endfor %}
  {% assign subcommand2 = command5.subcommands.profile %}
      <tr>
        <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/perf/profile">{{ subcommand2.name }}</a></td>
        <td></td>
        <td>{{ subcommand2.description }}</td>
      </tr>
  {% for flag_hash in subcommand2.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
  {% endfor %}
  {% assign subcommand3 = command5.subcommands.result %}
      <tr>
        <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/perf/result">{{ subcommand3.name }}</a></td>
        <td></td>
        <td>{{ subcommand3.description }}</td>
      </tr>
  {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
      <tr>
        <td>{{ flag.name }}</td>
        <td>{{ flag.description }}</td>
      </tr>
  {% endfor %}
</thead>
</table>

## Cloud Native Lifecycle and Configuration Management

<table>
<thead>
  <tr>
    <th>Main Command</th>
    <th>Command</th>
    <th>Flag</th>
    <th>Function</th>
  </tr>
  {% assign command7 = site.data.meshplayctlcommands.cmds.mesh %}
    <tr>
      <td rowspan=13><a href="{{ site.baseurl }}/reference/meshplayctl/mesh">{{ command7.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command7.description }}</td>
    </tr>
    {% assign subcommand1 = command7.subcommands.validate %}
      <tr>
        <td rowspan=5><a href="{{ site.baseurl }}/reference/meshplayctl/mesh/validate">{{ subcommand1.name }}</a></td>
        <td></td>
        <td>{{ subcommand1.description }}</td>
      </tr>
      {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand2 = command7.subcommands.remove %}
      <tr>
        <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/mesh/remove">{{ subcommand2.name }}</a></td>
        <td></td>
        <td>{{ subcommand2.description }}</td>
      </tr>
      {% for flag_hash in subcommand2.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand3 = command7.subcommands.deploy %}
      <tr>
        <td rowspan=4><a href="{{ site.baseurl }}/reference/meshplayctl/mesh/deploy">{{ subcommand3.name }}</a></td>
        <td></td>
        <td>{{ subcommand3.description }}</td>
      </tr>
      {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
</thead>
</table>

## Cloud Native Pattern Configuration and Management

<table>
<thead>
  <tr>
    <th>Main Command</th>
    <th>Command</th>
    <th>Flag</th>
    <th>Function</th>
  </tr>
  {% assign command7 = site.data.meshplayctlcommands.cmds.pattern %}
    <tr>
      <td rowspan=10><a href="{{ site.baseurl }}/reference/meshplayctl/pattern">{{ command7.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command7.description }}</td>
    </tr>
    {% assign subcommand1 = command7.subcommands.apply %}
      <tr>
        <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/pattern/apply">{{ subcommand1.name }}</a></td>
        <td></td>
        <td>{{ subcommand1.description }}</td>
      </tr>
      {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand2 = command7.subcommands.view %}
      <tr>
        <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/pattern/view">{{ subcommand2.name }}</a></td>
        <td></td>
        <td>{{ subcommand2.description }}</td>
      </tr>
      {% for flag_hash in subcommand2.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand3 = command7.subcommands.list %}
      <tr>
        <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/pattern/list">{{ subcommand3.name }}</a></td>
        <td></td>
        <td>{{ subcommand3.description }}</td>
      </tr>
      {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand3 = command7.subcommands.delete %}
      <tr>
        <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/pattern/delete">{{ subcommand3.name }}</a></td>
        <td></td>
        <td>{{ subcommand3.description }}</td>
      </tr>
      {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
</thead>
</table>

## Cloud Native Application Management

<table>
<thead>
  <tr>
    <th>Main Command</th>
    <th>Command</th>
    <th>Flag</th>
    <th>Function</th>
  </tr>
  {% assign command8 = site.data.meshplayctlcommands.cmds.app %}
    <tr>
      <td rowspan=15><a href="{{ site.baseurl }}/reference/meshplayctl/app">{{ command8.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command8.description }}</td>
    </tr>
    {% assign subcommand1 = command8.subcommands.import %}
      <tr>
        <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/app/import">{{ subcommand1.name }}</a></td>
        <td></td>
        <td>{{ subcommand1.description }}</td>
      </tr>
      {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand2 = command8.subcommands.onboard %}
      <tr>
        <td rowspan=4><a href="{{ site.baseurl }}/reference/meshplayctl/app/onboard">{{ subcommand2.name }}</a></td>
        <td></td>
        <td>{{ subcommand2.description }}</td>
      </tr>
      {% for flag_hash in subcommand2.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand3 = command8.subcommands.offboard %}
      <tr>
        <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/app/offboard">{{ subcommand3.name }}</a></td>
        <td></td>
        <td>{{ subcommand3.description }}</td>
      </tr>
      {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand4 = command8.subcommands.list %}
      <tr>
        <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/app/list">{{ subcommand4.name }}</a></td>
        <td></td>
        <td>{{ subcommand4.description }}</td>
      </tr>
      {% for flag_hash in subcommand4.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand5 = command8.subcommands.view %}
      <tr>
        <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/app/view">{{ subcommand5.name }}</a></td>
        <td></td>
        <td>{{ subcommand5.description }}</td>
      </tr>
      {% for flag_hash in subcommand5.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
</thead>
</table>

## Data Plane Intelligence

<table>
<thead>
  <tr>
    <th>Main Command</th>
    <th>Command</th>
    <th>Flag</th>
    <th>Function</th>
  </tr>
  {% assign command9 = site.data.meshplayctlcommands.cmds.filter %}
    <tr>
      <td rowspan=10><a href="{{ site.baseurl }}/reference/meshplayctl/filter">{{ command9.name }}</a></td>
      <td></td>
      <td></td>
      <td>{{ command9.description }}</td>
    </tr>
    {% assign subcommand1 = command9.subcommands.import %}
      <tr>
        <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/filter/import">{{ subcommand1.name }}</a></td>
        <td></td>
        <td>{{ subcommand1.description }}</td>
      </tr>
      {% for flag_hash in subcommand1.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand2 = command9.subcommands.delete %}
      <tr>
        <td rowspan=1><a href="{{ site.baseurl }}/reference/meshplayctl/filter/delete">{{ subcommand2.name }}</a></td>
        <td></td>
        <td>{{ subcommand2.description }}</td>
      </tr>
      {% for flag_hash in subcommand2.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand3 = command9.subcommands.list %}
      <tr>
        <td rowspan=2><a href="{{ site.baseurl }}/reference/meshplayctl/filter/list">{{ subcommand3.name }}</a></td>
        <td></td>
        <td>{{ subcommand3.description }}</td>
      </tr>
      {% for flag_hash in subcommand3.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
    {% assign subcommand4 = command9.subcommands.view %}
      <tr>
        <td rowspan=3><a href="{{ site.baseurl }}/reference/meshplayctl/filter/view">{{ subcommand4.name }}</a></td>
        <td></td>
        <td>{{ subcommand4.description }}</td>
      </tr>
      {% for flag_hash in subcommand4.flags %}{% assign flag = flag_hash[1] %}
        <tr>
          <td>{{ flag.name }}</td>
          <td>{{ flag.description }}</td>
        </tr>
      {% endfor %}
</thead>
</table>
{% include related-discussions.html tag="meshplayctl" %}
