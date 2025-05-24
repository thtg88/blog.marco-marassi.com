# Personal Blog - blog.marco-marassi.com

You can find the app in action [here](https://blog.marco-marassi.com).

## Requirements

### Go v1.24

Install it using [gvm](https://github.com/moovweb/gvm) with:

```bash
gvm install go1.24
```

Use it with:

```bash
gvm use go1.24
```

### NodeJS v22.14.0

An `.nvmrc` is provided. Using [nvm](https://github.com/nvm-sh/nvm) with:

```bash
nvm use
```

Install the dependencies with:

```bash
npm install
```

## Server

You can start the server locally with:

```bash
go run main.go
```

You can also build the server with:

```bash
go build -o wedding main.go
```

Make sure the server is executable with:

```bash
chmod +x wedding
```

And execute it with:

```bash
./wedding
```

## CSS

You can run a dev server with auto-reload for the CSS styling with:

```bash
npm run dev:css
```

You can also build a minified `style.css` file with:

```bash
npm run build:css
```
