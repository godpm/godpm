## Godpm
A golang based command line tool for deploying and manage programs.

#### Deploy
Support features:
- Removte deploy
- Version management
- Machine group

##### Remote deploy
> godpm deploy -f file -start_secs 5 -env 'GOPATH=/xxx/xx;GOROOT=/xxx/xxx' -run_dir '/xxx/xx' - pre_command 'tar -zxvf file' -auto_start true -auto_restart true -retry_times 3 -remote_addr http://xxx.xxx.x.x:10086

##### Check version
> godpm version xxxxx

##### Switch version
> godpm switch xxxxx verion

#### Process Manager
A supervisor like program manager

##### Check status
> godpm status

##### Stop process
> godpm stop xxxx

##### Start process
> godpm start xxxxx

##### Restart process
> godpm restart xxxx

##### Reread
> godpm reread

