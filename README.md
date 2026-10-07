# Header to Filename Converter

This is my first attempt at using Go to write a simple file conversion program.

I am switching from Zettlr to Obsidian and need to rename all my files so they more closely match
the H1 headers in my Zettlr .md files. Otherwise, many of my links will look nonsensical and be
difficult to audit.

This program recursively and non-destructively copies most text files, namely .md and .txt files,
taking the content up to the first newline character and then stripping away all the awkward characters
and emojis that appear in my H1s, which are always the first line of my Zettlr files.

This makes the filename mostly readable and circumvents issues that odd characters can create for
an operating system, whether it's Windows, macOS, or Linux ("/" really wreaks havoc on everything).

All files are copied into a `renamed_copies` folder so that you can doublecheck the results without
deleting or moving the files inapprorpriately.

# Basic Use

This thing is a sketch of a program at the moment. But it works like this:

1. `go run .` from inside the installation directory OR run `go install` to run it from anywhere
2. At the prompt, enter the directory in which you want to rename files
    - the program currently assumes these files are somewhere in your HOME directory.
    - that means you type `Documents/MyFolder` if you want to target `~/Documents/MyFolder`
    - if you want to target a file somewhere else, install the program and `cd` into the target directory
    - at the prompt, simply type a `.` (period) and hit enter
    - this tells the program you want to begin in the current working directory
3. At present, you'll get a bunch of messy messages about what's happening.
4. When the program is done, type `exit`. This kills the program.
5. Go check your folder at `Documents/MyFolder/renamed_copies`
6. If your folder contained a bunch of sub-folders, each sub-folder will have its own `renamed_copies`
7. Profit.

This program benefits immensely from [Gomoji](https://github.com/forPelevin/gomoji).
