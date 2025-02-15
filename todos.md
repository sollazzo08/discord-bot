# TODO List

- [ ] Expand on User Commands:
  - [ ] Add !role for listing out all the available server roles

- [ ] Add Event Listeners:
  - [ ] Add !role for listing out all the available server roles
  - [X] Role assignment -> reaction roles

- [ ] Add concurrency support when making calls out to my weather service
- [X] Add rate limiting
- [X] Reset the rate limit block after 24 hours
- [X] Add DockerFile
- [ ] Research Discord rich embeds for more visual responses
- [X] Restructure the Discord bot file directory to support different features instead of having everything in main.go.
- [X] Reformat the weather service response to Discord users by converting raw JSON into a more user-friendly, readable message.
- [ ] Create instructions on main readme
- [X] Fix sunrise and sunset time bug: sunset and sunrise are not accurate during late EST hours
- [ ] Add a weatherCommand handler
- [ ] Use different API endpoints based on whether it is running locally or in Docker.
- [ ] Add zip-code caching (would require a DB, could look into using sql-lite)
- [X] Update Rate Limit to 25 request per day per user
- [ ] Handle invalid zip codes
- [X] Create a parse command for the movies-shows channel
- [ ] Lock down the parse command to only me
- [ ] integrate openAI to sift through movie-show data and return a clean data in the form of user profiles with critic data points
- [ ]





