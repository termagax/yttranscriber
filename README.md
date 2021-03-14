# Simple Youtube Transcriber

## Run in Docker

Clone the repository

```
git clone https://github.com/termagax/yttranscriber.git
```

Change directory into the cloned repository

```
cd yttranscriber
```

Build your docker image (name it whatever you want)

```
docker build -t <image> .
```

Run the application with this command, running on your desired port and container name

```
docker run --publish <port>:8080 -it --rm --name <container> <image>
```
## Set up as a cron job

```
crontab -e
```

```
@reboot docker run --publish <port>:8080 -it --rm --name <container> <image>
```

## How to use:

1. Append "/transcriber" to your url

2. Paste in youtube url and hit "transcribe" button
