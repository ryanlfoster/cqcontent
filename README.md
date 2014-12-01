cqcontent
=========
##Under Development!

cqcontent is a cq content package handler written in Go (i.e. Golang). No shelling out
means no extra dependencies. Each sub-command does what you would expect. 

The file subcommand reads a json configuration file and loops through each specified job 
in the array of jobs. It'll first validate json syntax. Next, it'll validate the configuration. 
Lastly, it'll loop through each job and execute. 

###Goals
Since cqcontent is idempotent, configuration management tools should be able to manage content 
package deployment as declarative state via the json configuration file. Alternatively, users 
should be able to use remote execution tools to deploy the configuration file and call cqcontent 
on the file. 

Allow users of json configuration file to upload and install a content package to an arbitrary 
number of CQ nodes, concurently, in a single job. 

###Current Issues
Downloading zip files via the go-curl libcurl bindings writes more data than is present in the 
original file (in the current implementation). The file can be unzipped with the unzip binary,
but [native OS X unarchiving tools have problems with it](http://osxdaily.com/2013/02/13/open-zip-cpgz-file/). 

See go-curl for more insight into Go's libcurl bindings if
you care to hack on this project. Additionally, it relies on CQ's 
curl API, which is documented here: [CQ Curl API](http://docs.adobe.com/docs/en/crx/current/how_to/package_manager.html#Managing Packages on the Command Line)
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
