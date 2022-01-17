# Video encoder

A golang application to study memory, parallelism and concurrency.

## Objective

The objective of this project is to create a video encoder that process several videos at same time. For this, we are going to work with goroutines and channels. The encoder will convert video from MP4 to MPEG-DASH

## Memory

![Memory management](img/memory-management.png)

Slots L1, L2 and L3 are transient. They communicate with SSD through a bus.

![Memory management](img/virtual-memory.png)

With virtual memory, each process has it own virtual memory, wich is allocated insed a real memory slot, so there is no address collision between processes.

## Parallelism and Concurrency

A process can have n threads, and all threads will share the same Virtual Memory.
Threads can only run simultaneously if there are available cores to run the thread. Sometimes threads need to share data.

### Cooperative Multitask

The process that has the core will realease the core when it wants.

### Preemptive Multitask

Each process has a deliberated time to run into a core

### Race Condition

Severeal threads trying to use same resource. Generates a problem, because one thread impacts the other. Solution: semaphore, mutex, etc...

#### Deadlock

When threads are locked awaiting a resource to be released


# Golang

Runtime -> golang code

Your code -> code you developed

Runtime + Your code = Runable binary code

## Runtime

Runtime has a scheduler that works with cooperative multitask. Inside a thread there can be several goroutines (Green threads, or fake threads) wich are faster and consumes less memory.

### Enviroment Variables

#### GOMAXPROCS

Amount of cores your program can use.

## Go routines

Sometimes goroutines need to share a resource. This is done through channels. A routine sends a message through a channel, another routine emptys the channel and uses the data. When first routine won't write data in the channel anymore, it can close the channel. Data can only be writen in the channel when the channel is empty. Otherwise, the writer func will be locked awaiting the channel to be read.

## Go Mod

Gerenciador de pacotes. Comandos:

go mod init video-encoder -> gera arquivo go.mod

to add a new dependency, you just need to import it

go mod tidy -> before running, all dependencies will be downloaded

## Go Race

Go has a flag that allows the go runtime to detect race conditions. 

To use it, run your application with the flag: -race

With this flag go runtime will show warnings into application log with the taga 'WARNING: DATA RACE'
whenever it finds a race condition problem

## Service Architecture

### Success

![Success](img/architecture.png)

### Failure

![Failure](img/arch-failure.png)

### Software Arch

![Software Arch](img/software-arch.png)

### Running project

To run this project you need to create a bucket into GCP (Google Cloud Plataform) and a service account with access to that bucket.

#### GCP

The GCP config is into 'bucket-credential.json' in the root dir. You have to change this, and add your own config.

#### Enviroment variables

The variables are into the file '.env' in the project root dir.

You have to change this variables to set up your buckets

inputBucketName=codeeducationtest-amr
outputBucketName=codeeducationtest-amr

#### Local Test
To test locally you have to upload a video into your bucket

You have to run docker-compose up

Access rabbitMQ into http://localhost:15672/

Create an Exchange with type fanout and name 'dlx'

Create two queues, one with name 'videos-failed' and other with 'videos-result'

Associate the dlx with the videos-failed queue

Access the docker container

Run the following command: go run framework/cmd/server/server.go or go run -race framework/cmd/server/server.go (verify race conditions)

Send a message to videos queue through rabbitMQ web interface following this format:

''''
{
  "resource_id": "id-client-1",
  "file_path": "teste.mp4"
}
''''

## Problem resolution

### Not able to ping google inside container

Hardcode DNS server in docker daemon.json

Edit /etc/docker/daemon.json

{
    "dns": ["10.1.2.3", "8.8.8.8"]
}
Restart the docker daemon for those changes to take effect:
sudo systemctl restart docker

Now when you run/start a container, docker will populate /etc/resolv.conf with the values from daemon.json.
