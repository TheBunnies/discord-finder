# About
This is the backend application for **Discord Finder**, it allows you to retrive information about people on discord just like the [discord lookup](https://discord.id/) website but without captcha and you can host it on your own machine.

![discord finder](https://i.imgur.com/9Un160l.png)

# Front-end
The front-end for this application can be found [here](https://github.com/TheBunnies/discord-finder-frontend).

## Running the back-end with docker
1. Populate the `config.json` with your bot or user token.
2. Install [Docker](https://www.docker.com/) on your OS.
3. Build docker image `docker build -t discord_finder_backend .`.
4. Run the container `docker run --rm -d -p 8000:8000 discord_finder_backend`.