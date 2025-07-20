# mrkdown_to_latex

## Project Overview

The goal of this project is to create a command-line tool in Go that converts a Markdown file into a LaTeX document. The tool will parse the Markdown file, identify its structure (headings, paragraphs, lists, code blocks, etc.), and generate a corresponding LaTeX document.

## Key Features

  1) Input: Accept a Markdown file as input.
  2) Output: Generate a LaTeX file as output.
  3) Supported Markdown Elements:

    - Headings (#, ##, ###, etc.)

    - Paragraphs

    - Bold (**bold**) and Italic (*italic*)

    - Lists (ordered and unordered)

    - Code blocks (inline and multiline)

    - Links and images

    - Horizontal rules (---)

  4) Customization: Allow users to specify LaTeX document class, title, author, and other metadata.
  5) Error Handling: Handle invalid Markdown syntax gracefully.
