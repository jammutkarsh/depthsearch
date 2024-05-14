# Depth Search

An API wrapper around [depth](https://github.com/JammUtkarsh/depth) with a UI.

> Depth is a library and a CLI tool to resolve internal, external and standard libraries of a Go based project.
> It does so by collecting the build information like dependency graph and returns the head node.

This project is part of our [Project Work Research Paper](https://links.utkarshchourasia.in/rpaper) for University.

## Usage

The current setup designed to only work locally i.e. in `localhost`.

### Using Docker Compose

Required: [Docker](https://www.docker.com)

```bash
docker compose up --build
```

The UI will be accessible on localhost:3000.

---

> NOTE: This project has only been positive tested and not negative tested i.e. the response and status code replies for incorrect methods may not provide the full reason of why it is not resolving.
