# Usage

```bash
# to get info (temp) of those cities concurrently
  ./app <city name> <city name2> ... etc
```

- it send a get request (query) to `json-server` which expose data in the form:
  `{id:1, name:city1, temp:23}`
