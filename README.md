# Hackclub Mail
A simple tool allowing you to track your HackClub packages via your terminal.

## Installation
### Prerequisites
- Golang

### Installation
1. Install with Go
```sh
go install github.com/oriionn/hackclub-mail
```

2. Edit your configuration file with your API Key
    **Windows:** `%APPDATA%\hackclub-mail\config.ini`
    **MacOS:** `~/Library/Application Support/hackclub-mail/config.ini`
    **Linux:** `~/.config/hackclub-mail/config.ini`

```ini
[general]
api_key=YOUR_API_KEY
```

## License
This projet is under [MIT License](LICENSE)
