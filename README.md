# lavalyrics-plugin

[DisGoLink](https://github.com/disgoorg/disgolink) plugin for [LavaLyrics](https://github.com/topi314/LavaLyrics).

## Installation

```shell
$ go get github.com/disgoorg/lavalyrics-plugin
```

## Usage

```go
var client disgolink.Client
var sessionID string
var guildID snowflake.ID
lyrics, err := lavalyrics.GetLyrics(context.TODO(), client.BestNode().Rest(), sessionID, guildID)
if err != nil {
    // handle error
}
```

## Troubleshooting

For help feel free to open an issue or reach out on [Discord](https://discord.gg/TewhTfDpvW)

## Contributing

Contributions are welcomed but for bigger changes we recommend first reaching out via [Discord](https://discord.gg/TewhTfDpvW) or create an issue to discuss your problems, intentions and ideas.

## License

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE). See LICENSE for more information.
