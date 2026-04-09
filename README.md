# ASCII-ART-FS

A Go implementation of a graphic generator that converts strings into ASCII art using external banner templates.

## 📁 File System (FS) Logic
The "FS" in the title signifies the project's reliance on the Go **File System API**. The program dynamically loads banner templates from the disk based on user input, ensuring modularity and support for multiple font styles.

## 🚀 Usage
`go run . [STRING] [BANNER]`

**Example:**
`go run . "Hello" shadow`

## 🧠 Core Algorithm
- **Offset Mapping:** Calculates character positions in `.txt` files using: `(ASCII - 32) * 9 + Row`.
- **Normalization:** Converts all line endings to `\n` to ensure cross-platform mathematical consistency.
- **Efficient Assembly:** Uses `strings.Builder` to minimize memory allocations.

## 🧪 Testing
Run `go test -v .` to verify the mathematical engine.