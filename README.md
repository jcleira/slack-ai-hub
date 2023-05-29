# Slack AI Hub

![Slack AI Hub](https://i.imgur.com/wGH6t0P.png)

> This is a WIP

This is a project for creating a Slack AI Bot using Go. It provides a basic structure and integration with the OpenAI API for generating answers based on a given context provided from different sources.

## Prerequisites

Before running the Slack bot, ensure you have the following:

- [Go](https://golang.org/) installed on your machine
- Slack API token for your bot (refer to the [Slack API documentation](https://api.slack.com/authentication/basics#creating) for instructions on creating a new app and obtaining a token)
- OpenAI API key for generating answers using the OpenAI API (refer to the OpenAI website for instructions on obtaining an API key)

## Getting Started

1. Clone the repository:

```bash
git clone git@github.com:jcleira/slack-ai-hub.git
cd slack-ai-hub
```

2. Set the required environment variables:

`SLACK_TOKEN`: Slack API token for your bot.
`OPENAI_API_KEY`: OpenAI API key for generating answers.

```bash
export SLACK_TOKEN=your_slack_token_here
export OPENAI_API_KEY=your_openai_api_key_here
```

3. Provide any extra context for the bot.

> TODO

4. Build and run the Slack bot:

```
make run
```

Your Slack bot is now up and running!

It will connect to the Slack API and start listening for incoming messages. When a message is received, it will use the OpenAI API to generate an answer based on the message context and post the answer back to the Slack channel.

## Customizing the Bot
You can customize the bot's behavior by modifying the generateAnswer function in the main.go file. This function is responsible for calling the OpenAI API to generate answers. Feel free to experiment with different models, parameters, and logic to suit your specific requirements.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.
