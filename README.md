# go-ha

User interfaced CLI tool for quick actions in Home assistant.

Allows toggling of lights & automations.
Allows viewing of sensors.

Automatically detects structure from the Home Assistant REST API.

## Usage

The CLI needs the `HA_URL`, and a token to work.
You can get a Long-Lived Access Token from the Home Assistant frontend.

Export these variables into the environment before running the CLI, as shown below.

```
export HA_URL=192.168.1.101:8123
export HA_TOKEN=<secret-token>
```

Then simply run the CLI and follow the instructions on there.