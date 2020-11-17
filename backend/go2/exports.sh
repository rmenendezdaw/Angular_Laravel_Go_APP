#!/bin/bash
#GOROOT is the location where Go package is installed on your system
echo 'export GOROOT=/usr/local/go'

#GOPATH is the location of your work directory

echo 'export GOPATH=$HOME/Escritorio/Desarrollo_servidor/Go_Laravel_Angular/Angular_Laravel_Go_APP/backend/go2'


#Now set the PATH variable to access go binary system wide
echo 'export PATH=$PATH:/home/raul/Escritorio/Desarrollo_servidor/Go_Laravel_Angular/Angular_Laravel_Go_APP/backend/go2/bin'

echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH'

