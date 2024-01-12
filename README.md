# Sledge 🛷

A small, simple, minimalistic, and fast Redis TUI client written in Go.

<p align="center">
  <img src="./assets/demo.gif" width="700"/>
</p>

## Usage

- It is always good to wrap the connection string in quotes to avoid shell expansion.
- Example below shows how to connect to a local Redis server:

```sh
sledge "redis://localhost:6379"
```

## Installation

You can install the appropriate binary for your operating system by visiting the [Release page](https://github.com/somnek/sledge/releases/).

**Note**:  
If you're on macOS, you may need to run:

```sh
xattr -c ./Sledge\ Darwin\ x86\ .gz
```

to (to avoid "unknown developer" warning)

## Controls

| Key        | Description |
| ---------- | ----------- |
| `j`        | Down        |
| `k`        | Up          |
| `g`        | To Top      |
| `G`        | To Bottom   |
| `Ctrl + C` | Quit        |

## Features

Currently, **Sledge** supports the following types:

- String
- List
- Set
- Hash
