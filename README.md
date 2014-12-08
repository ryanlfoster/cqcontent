cqcontent
=========

cqcontent is a cq content package handler written in Go (i.e. Golang). No shelling out
means no extra dependencies. Each sub-command does what you would expect. 

The file subcommand reads a json configuration file and loops through each specified job 
in the array of jobs. It'll first validate json syntax. Next, it'll validate the configuration. 
Lastly, it'll loop through each job and execute. 

###Usage
```
Usage:
  cqcontent file     CONFIG_FILE
  cqcontent xml      [options]
                     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent list     [options]
                     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent download [options]
                     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent upload   [options]
                     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent install  [options]
                     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
                     [--autosave NUMNODES]
                     [--recursive]
                     [--acIgnore]
                     [--acOverwrite]
                     [--acClear]
  cqcontent delete   [options]
                     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent          [--help|-h|--version]

Options:
  --port PORT        Specify the port. [default: 8080]
  --vt TIMEOUT       Verification timeout in seconds for the upload, download, 
                     or install actions [default: 60]
```
###Sample Configuration file
```json
[
    {
        "mode": "upload",
        "node": "my.node.net",
        "username": "admin",
        "password": "admin",
        "package": "test.zip"
    },
    {
        "mode": "install",
        "node": "my.node.net",
        "username": "admin",
        "password": "admin",
        "package": "test.zip",
        "acIgnore": true
    }
]
```

###Developing cqcontent

See go-curl for more insight into Go's libcurl bindings if
you care to hack on this project. Additionally, it relies on CQ's 
curl API, which is documented here: [CQ Curl API](http://docs.adobe.com/docs/en/crx/current/how_to/package_manager.html#Managing Packages on the Command Line)

Since go-curl makes use of cgo, you can't take shortcuts with cross-compilation. If you
need to build for environments other than your development machine, try Vagrant or Docker. 
