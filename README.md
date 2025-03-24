# tempo-demo

This project demonstrates how [`tempo`](https://github.com/indaco/tempo) works in a real-world Go + [Templ](https://templ.guide) setup.

It covers:

- ✅ Default project structure
- ⚙️ A typical development workflow
- 🎨 Component scaffolding with CSS and JS support
- 💡 Live reloading and code generation

We'll walk through building a **Button** component with a variant (`outline`) and some JavaScript to showcase how `tempo` handles styles and scripts together.

## Prerequisites

Make sure you have the following tools installed:

```bash
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/indaco/tempo/cmd/tempo@latest
go install github.com/air-verse/air@latest
```

## 🧪 Getting Started

### 1. Create a new Go module

```bash
mkdir tempo-demo
cd tempo-demo

go mod init github.com/indaco/tempo-demo
go get github.com/a-h/templ@latest
```

### 2. Initialize tempo

```bash
tempo init
```

This creates a `tempo.yaml` config file. You can customize it by adding metadata like:

```yaml
templates:
  user_data:
    author: John Doe
    year: 2025
```

### 3. Define your templates

```bash
tempo component define --js
```

This scaffolds template files under `.tempo-files/templates`.
Optionally update `.tempo-files/templates/component/templ/component.templ.gotxt` with metadata comments for reference.

### 4. Create your first component

```bash
tempo component new --name button --js
```

This will create:

- `assets/button` – CSS and JS files
- `components/button` – Templ code

Write styles and scripts in the assets folder to bring your component to life.

### 5. Sync asset files with templ components

```bash
tempo sync
```

This injects the contents of your CSS and JS into their corresponding `.templ` components.

### 6. Generate code with templ

```bash
templ fmt . && templ generate
```

Then clean up dependencies:

```bash
go mod tidy
```

### 7. Run the dev server

```bash
make dev # or: task dev
```

Then visit: [http://localhost:7331](http://localhost:7331)

This starts:

- `templ` in watch mode
- Asset syncing with live reload
- A hot-reloading Go server using air

### 8. Edit CSS or JS file

Just edit and save any CSS or JS file in assets/ — `tempo` syncs your changes instantly, and `templ`’s proxy takes care of reloading the browser automatically.

### 9. Build

Test your production-ready setup.

```bash
make build # or: task build
```

This will:

- Run `tempo sync --prod` (minifies and injects assets)
- Run `templ fmt` and `templ generate`
- Compile the Go binary into `tmp/bin/tempo-demo`

## 🛠️ Running Tasks

This project supports both a `Makefile` and a `Taskfile.yaml`. You can use whichever you prefer:

### View available tasks

- _Makefile_:

    ```bash
    make help
    ```

- _Taskfile_:

    ```bash
    task --list-all
    ```

#### Common tasks

```bash
build:       # Build for production with minified asset files
dev:         # Run the dev server with live reload
```

## 🆓 License

This project is licensed under the MIT License – see the [LICENSE](./LICENSE) file for details.
