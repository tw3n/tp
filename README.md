# tp

As it happens often, you’re not the one who wrote the app but you still have to Dockerize it.

## Requirements

- Create an OpenWeatherMap account and grab an API key
- Write a Dockerfile for the API
    - Find a way to not redownload the dependencies if the code changes
    - Find a way to expose the author / maintainer of the image
    - Find the most suitable base image (it should be small)
    - Find a way to indicate the port exposed
- Write a `docker-compose` file
    - The command `docker-compose up` must start the API
    - The API should be reachable on its default port
    - Find a way to configure the API using environment variables
