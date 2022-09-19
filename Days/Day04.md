## Day 04

## YAML

More in-depth about YAML than Day03.

To start, YAML stands for YAML Ain't Markup Language

Two extensions: .yaml and .yml

YAML is a superset of JSON; any valid JSON file is also a valid YAML file

Compared to XML and JSON, YAML uses line spearate and indentation

YAML is used in Docker compose files, Ansible, Prometheus, K8s, and more

Syntax of YAML Examples:

**Key-Value Pairs**

app: user-authentication
port: 9000
version: 1.7

**Comments:**

Uses "#" 

**Objects:**

Indent key-value pairs and enclosing it in an object
```
microservice:
  app: user-authentication
  port: 9000
  version: 1.7
```
*Tip: Good idea to use a YAML validator to make sure the spacings are correct*

**Lists:**

Uses a "-"
```
microservices:
- app: blah
  port: 1000
  version: 1.0
```
Tip: Make sure everything is aligned correctly after starting the list

Can add multiple lists in an object
Can have list inside a list
```
microservices:
  - app: 1
    port: 1000
    version: 1.0
  - app: 2
    port: 1002
    versions: [1.5, 1.6] <-- this method can be used as well
    - 1.5
    - 1.6
```
**Boolean:**

Can be expressed as ```yes``` or ```no``` along with ```on``` or ```off``` besides ```true``` or ```false```


