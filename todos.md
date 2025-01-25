# TODO List

- [ ] Expand on User Commands:
  - [ ] Add !role for listing out all the available server roles
  - [ ]

- [ ] Add Event Listeners:
  - [ ] Add !role for listing out all the available server roles
  - [ ] Custom Welcome message
  - [ ] Role assignment -> reaction roles





Weather Command Workflow
Command Handling

Parse the !weather <city> command to extract the city name.
Validate the input (ensure the city name is provided and non-empty).
Handle cases where the input is invalid (e.g., reply with an error message).
API Integration

Build the OpenWeatherAPI request URL dynamically using the city name and API key.
Send an HTTP GET request to the API.
Handle potential API errors (e.g., invalid API key, city not found, rate-limiting).
Data Processing

Parse the JSON response from the API.
Extract relevant fields (e.g., temperature, description, humidity).
Format the extracted data into a user-friendly response.
Discord Integration

Reply to the user in the same channel with the formatted weather data.
Handle potential Discord errors (e.g., failed to send a message).
Error Handling and Edge Cases

Handle invalid or missing city names.
Handle rate limits or connection errors from OpenWeatherAPI.
Handle cases where the bot cannot connect to Discord.




Getting content approach

m.Content gets me the a string i.e. !weather New York

We convert the entire string to an array delmieted by spaces so that each index capturs a string

!weather = [0]
Monroe = [1]

How do capture a city properly is the challange here.

A user could enter an invalid city -> We can throw and error calling out invalid city name
A user could enter two spaces after the ! weather commmand - > We can use the string.Fields command
A user could not enter a space after the weather command -> we can throw can error out can call out inproper input format


I also want to accept lowercase and upper case

To capture the city properly after the !weather command

I first need to split the entire string including the command by spaces

Then I check the length of the slice

If len is 2 then i take the city as index 1

If len is 3 then I take both index 1 and 2 as the city concated by a space index 1 + " " index 2

if len is 1 then its an error

