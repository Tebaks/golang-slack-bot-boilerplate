package configs

type Configs struct {
	Channels ChannelsConfig `mapstructure:"channels"`
}

type ChannelsConfig struct {
	SlackBotTestChannel SlackBotTestChannelConfig `mapstructure:"slack-bot-test"`
}

type SlackBotTestChannelConfig struct {
	Enabled     bool                       `mapstructure:"enabled"`
	ID          string                     `mapstructure:"id"`
	Description string                     `mapstructure:"description"`
	Messages    SlackBotTestMessagesConfig `mapstructure:"messages"`
}

type SlackBotTestMessagesConfig struct {
	HiTestBotMessage MessageConfig `mapstructure:"hi-test-bot"`
}

type MessageConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	Description string `mapstructure:"description"`
	Text        string `mapstructure:"text"`
}

type Credentials struct {
	SlackCredentials SlackCredentials `mapstructure:"slack"`
}

type SlackCredentials struct {
	SlackAuthToken string `mapstructure:"slack_auth_token"`
	SlackAppToken  string `mapstructure:"slack_app_token"`
}
