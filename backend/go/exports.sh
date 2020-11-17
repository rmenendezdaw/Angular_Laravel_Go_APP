#!/bin/bash
#GOROOT is the location where Go package is installed on your system
echo 'export GOROOT=/usr/local/go'

#GOPATH is the location of your work directory

echo 'export GOPATH=$HOME/Escritorio/Desarrollo_servidor/go_projects'


#Now set the PATH variable to access go binary system wide
echo 'export PATH=$PATH:/home/raul/Escritorio/Desarrollo_servidor/go_projects/bin'

echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH'

