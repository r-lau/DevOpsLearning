## Day 04 - YAML

More in-depth about YAML than Day03.

To start, YAML stands for YAML Ain't Markup Language(used to be: Yet Another Markup Language)

It's a data format used to exchange data and is case sensitive
*Can only store data and not commands*

Two extensions: .yaml and .yml

YAML is a superset of JSON; any valid JSON file is also a valid YAML file

Compared to XML and JSON, YAML uses line seperate and indentation, but is more human readable and intuitive.

YAML is used in Docker compose files, Ansible, Prometheus, K8s, and more

YAML Examples:

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
*Tip: Make sure everything is aligned correctly after starting the list*

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


Actual Sample from K8s:

```
apiVersion: 1
kind: Pod
metadata: 
  name: nginx
  labels:
    app: nginx
  spec: 
    containers:
    - name: nginx-container
      image: nginx
      ports:
      - containerPort: 80
        volumeMounts:
      - name: nginx-vol
        mountPath: /usr/nginx/html
   - name: sidecar-container
     image: curlimages/curl
     command: ["bin/sh"]
     args: ["-c", "echo Hello from sidecar container; sleep 300"]
```

In the above example, there are key-value pairs, metadata(object), labels(object), spec(object), containers(list of objects), ports(list), volumeMounts(lis of objects)

**Multiline String:**

```
multilinestring: |
  this is a string
  another string
  next line
``` 

```
multilinesingle: >
  this is a string
  that should be on 
  a single line
```

Sample from K8s:
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: mosquito-config-file
data:
  mosquito.conf: |
    log_dest stdout
    log_type all
    log_timestamp true
    listener 9001
```

**Environmental Variables:**

Represented by "$"

MySQL Example
```
command:
- /bin/sh
- -ec
- >-
  mysql -h 127.0.0.1 -u root -p$MYSQL_ROOT_PASSWORD -e 'SELECT 1'
```

**Placeholders:**

Example from Helm
```
apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.service.name }}
spec:
  selector:
    app: {{ .Values.service.app }}
  ports:
    - protocol: TCP
      port: {{ .Values.service.port }}
      targetPort: {{ .Values.service.targetPort }}
```      
Syntax for placeholders is {{ templategenerator }} 

**Multiple YAML Files:***

Can separate the components with three dashes, "---"

**Specifying Data Types:**

Integers
```
positiveNum: !!int 5
commaValue: !!int 540_000 #540,000
```
Floats
```
grade: !!float 91.5
infinity: !!float .inf
not a num: .nan
```
Boolean:
```
thisTrue: !!bool true
```
String:
```
message: !!str hello world
```
Null
```
dataOne: !!null Null
~: this is a null key
```
Date and Time
```
date: !!timestamp 2022-09-19
```
*Tip: Defaults to UTC time if no timezone is provided*

**Advanced Data Types:**

Sequence:

```
student: !!seq
- marks
- name
- stu_id
```

Spare Sequence:
```
sparse:
- hey
- how
- 
- Null
```

Nested Sequence:
```
-
  - banana
  - apple
-
  -
```

Maps:
Key-value pairs are called maps
!!map

Nested Maps:
```
name: Ray
role:
  age: 40
  job: engineer
```
Can also be written as ```role: { age: 40, job: engineer}```

Pairs:
Keys may have duplicate values
!!pairs

```
pair example: !!pairs
- job: engineer
- job: contractor
```
Can also be written as ```pair example: !!pairs [job: engineer, job: contractor]```

An array of hash tables

Set:
Allows you to have unique values
```
names: !!set
  ? Tristan
  ? Greg
  ? Leo
```

Dictionary:
!!omap
```
people: !!omap
  - Ray:
      name: Ray L
      age: 90
  - Greg:
      name: Greg N
      age: 81
```

Reusing properties using anchors:
```
likes: &likes
  fruit: apple
  brand: bmw
  
person:
  name: Ray
  <<: *likes

person2:
  name: Dylan
  <<: *likes
  brand: toyota
```

**Tools:**
Validating YAML files and working with K8s

Lens: https://k8slens.dev/
Monokle: https://monokle.kubeshop.io/
Datree: https://datree.io/
