# FILE-STAT API & CLI

A basic app to list files and folders. The project contains a CLI app that can list the files of a given directory. 
It also contains a REST API that use the same logic, but return the results in a JSON format.


## Run it with Docker


#### Building the docker image

In the root of the project directory, we'll start by building the Docker image that will run our code, to proceed, 
run the following command

    docker build -t file-stat .

To summarize what the Docker files does in background, we'll start from a Go image to install the dependencies and 
create the executable file. When this is done, we start building with an Ubuntu image that will host the files and run the executable. 
The executable file that have been created during the building phase, will be copied in the `/usr/local/bin/` 
directory, so it can be used easily.

Now that we have our Docker image, we can run it with the following command

    docker run -t file-stat
    
The container is now running, we'll go through the two main functionalities that the container make available. The way the Dockerfile is made, 
when the container is launched, the API is automatically lauched as well.

As mentioned above the executable is located in `/usr/local/bin/` and is named `file_stat`. The executable is used for both 
the API and the CLI command. We'll explain the nuances to run those two below


#### Run the CLI

To list files via the CLI, the argument to give is `list`. The option `directory` gives the ability to provide the CLI the directory you want information
from. When not provided, the default directory to list information is `/tmp`. So a concrete example to list a folder will look like this

    file_stat list --directory /home/appuser
    
Note for the container version, either you "login" in the running container or you can also run something like that

    docker exec -it <container-id> file_stat list --directory /home/appuser 


#### Run the REST API

To list files via the CLI, the argument to give is `serve`. By default, the API listen on port `5150`. 
You will be able to access the API on `127.0.0.1:5150` or `<docker-ip>:5150`

    file_stat serve

There's two(2) path available

`/`  the "home" path, mostly used to know if the API is running properly in the health checks


`/files`  this path is the one used to lists the files of directory. A query parameter can be provide to choose the path to list. 
To do so, use the parameter `dirPath`. If not given, the default directory is `/tmp`. So let's we'll use the same a similar example than the 
CLI ealier, it would look like this to list a directory

    http://127.0.0.1:5150/files?dirPath=/home/appuser

And it should return something similar to this

    {
      "files": [
        {
          "name": "config-file-UVjkyC",
          "size": 174,
          "is_directory": false,
          "last_modification": "2020-11-09T19:39:36-05:00"
        },
        {...}
      ],
      "stats": {
        "total_size": 22,
        "total_elements": 28,
        "total_files": 3
      }
    }


## Run it on Linux

To build the executable without docker, you will need to have Go installed locally, version greater or equal to `1.14`.
When you have this installed, we'll be ready to build our executable.

First we'll download and install the dependencies (Assuming you are at the root directory of the project)

    go get -d -v ./ 
    go install -v ./ 
    
When this is done, we can now build our executable

    go build -i -o ./file_stat ./main.go
    
Now we have the executable and interact with it. Just see the section **Run the CLI** and **Run the REST API** above.