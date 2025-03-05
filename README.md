# Starter Kits Monorepo

A monorepo containing a Next.js website and Go CLI tool for managing framework starter kits.

## Project Structure

```bash
bun install
```

To run:

```bash
bun run index.ts
```

This project was created using `bun init` in bun v1.0.1. [Bun](https://bun.sh) is a fast all-in-one JavaScript runtime.

## Prerequisites

- Node.js (v18 or later)
- Go (v1.21 or later)
- npm (v9 or later)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/latreon/starter-kits-monorepo.git
cd starter-kits-monorepo
```

2. Install dependencies:
```bash
npm install
```

3. Build all projects:
```bash
npm run build
```

## Development

### Website

To run the website locally:
```bash
cd apps/website
npm run dev
```

The website will be available at `http://localhost:3000`.

### CLI

To run the CLI:
```bash
cd apps/cli
go run main.go generate
```

## Building

To build all projects:
```bash
npm run build
```

This will:
- Build the Next.js website
- Build the Go CLI

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
