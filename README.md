cqcontent
=========
##Under Development!

cqcontent is a cq content package handler written in Go (i.e. Golang). No shelling out
means no extra dependencies. 

See go-curl for more insight into Go's libcurl bindings if
you care to hack on this project. Additionally, it relies on CQ's 
curl API, which is documented here:

[CQ Curl API](http://docs.adobe.com/docs/en/crx/current/how_to/package_manager.html#Managing Packages on the Command Line)
```
Usage:
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
