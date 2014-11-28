cqcontent
=========
##Under Development!

cqcontent is a cq content package handler written in Go (i.e. Golang). No shelling out
means no extra dependencies. Each sub-command does what you would expect. 

The file subcommand reads a json configuration file and loops through each specified job 
in the array of jobs. It'll first validate json syntax. Next, it'll validate the configuration. 
Lastly, it'll loop through each job and execute. Since cqcontent is idempotent, you can use
configuration management tools to manage deployment as state declarations. Alternatively, 
you may use remote execution tools to deploy the configuration file and call cqcontent 
on the file. 

See go-curl for more insight into Go's libcurl bindings if
you care to hack on this project. Additionally, it relies on CQ's 
curl API, which is documented here:

[CQ Curl API](http://docs.adobe.com/docs/en/crx/current/how_to/package_manager.html#Managing Packages on the Command Line)
```
Usage:
  cqcontent file     CONFIG_FILE
  cqcontent xml      (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent list     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent download (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent upload   (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent install  (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
                     [--autosave NUMNODES]
                     [--recursive]
                     [--acIgnore]
                     [--acOverwrite]
                     [--acClear]
  cqcontent          [--help|-h|--version]
```
