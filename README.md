# Web Analyzer

This is a web application written in Go using the Echo framework that analyzes a web page and provides the following information:

- HTML version of the document
- Page title
- Number of headings (by level)
- Number of internal and external links
- Number of inaccessible links
- Whether the page contains a login form

## Features

- **HTML Version Detection**: Identifies the HTML version used by the document.
- **Title Extraction**: Retrieves the title of the web page.
- **Headings Count**: Counts the number of headings (h1, h2, h3, h4, h5, h6) on the page.
- **Link Analysis**: Counts the internal and external links and identifies any inaccessible links.
- **Login Form Detection**: Checks if the page contains a login form.

## Prerequisites

- Go 1.21.6
- Git

## Getting Started

### Helpful Commands
- `make run`: Starts the web analyzer.
- `make test`: Executes the test suite.

### Clone the Repository

```bash
git https://github.com/mtmibjas/analyze-web.git
cd analyze-web.

