# MockGopher (WIP)

Is a library and CLI/GUI application to make mock server in a ease.

## CLI Usage

You can see some examples in [cli/examples folder](cli/examples).

### Project Structure

```
.
├── resources            # Files that can be served
├── templates            # All templates go here
└── project.toml         # Configuration file, see "The TOML file"
```

### The TOML File

```toml
host = "0.0.0.0"
port = 3000

[[routes]]
  [routes.request]
    path = "/posts"
    method = "GET"
    headers = [
      { key = "Content-Type", value = "application/json.*" }
    ]
  [routes.response]
    headers = [
      { key = "Content-Type", value = "application/json" }
    ]
    template = "get-posts.json"

[[routes]]
  [routes.request]
    path = "/media/avatar/{file}"
    method = "GET"
  [routes.response]
    headers = [
      { key = "Content-Type", value = "image/svg+xml" }
    ]
    template = "get-posts.json"
    resources = [
      "avatars/25789.svg",
      "avatars/527688.svg",
      "avatars/7896451.svg",
      "avatars/9785412.svg",
      "avatars/25678412.svg"
    ]
```

### Serve

```powershell
mockgopher.exe .\path\to\project.toml
```
