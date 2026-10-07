# Rename files using your H1 headers

This program recursively and non-destructively copies text files, renaming them according to the first line
of the document. More specifically, it takes all text up to the first line break, replaces or removes 
a small handful of problematic characters, and uses that text to rename the file. It also removes emojis.

I wrote this because I am switching from [Zettlr](https://www.zettlr.com/) to [Obsidian](https://obsidian.md/)
and all of my filenames use a date template that would make managing wikilinks and references difficult. 
As a result, this program assumes you are probably working with `.md` files that begin with "# ".

But the following format types are currently supported:
  - .doc
  - .docx
  - .md
  - .odt
  - .rtf
  - .txt

All files are copied into a `renamed_copies` auv-folder so that you can doublecheck the results without
worrying about deleting or permanently renaming files you would rather not rename.

# Basic Use

This thing is a sketch of a program at the momenr, but it works like this:

1. `go run .` from inside the installation directory OR run `go build && go install`
    - if you use `go install` the program can be called from anywhere with `header2filename`
    - otherwise, you'll need to work from the directory where you saved the program
2. At the prompt, enter the directory in which you want to rename files
    - the program currently assumes these files are somewhere in your HOME directory.
    - that means you type `Documents/MyFolder` if you want to target `~/Documents/MyFolder`
    - if you want to target a directory outside HOME, install the program and `cd` into the target directory
    - at the prompt, simply type a `.` (period) and hit enter
    - this tells the program you want to begin in the current working directory
3. At present, results are printed to `stdout`, so you'll see a bunch of messages in your terminal.
4. When the program is done, type `exit`. This kills the program.
5. Go check your folder at `Documents/MyFolder/renamed_copies`
6. If your folder contained a bunch of sub-folders, each sub-folder will have its own `renamed_copies` folder.
7. Make sure you're happy with the results and then do what you need to do with the new files.

# Future changes

1. There's no point having a prompt inside the program. Future versions will take arguments from the command line.
2. Future versions will also allow you to specify a target folder.
3. I'd like to make the recursive functionality optional.
4. It shouldn't be too hard to take all those terminal messages and put them into a log instead.
5. And I have some ideas for how you could speficy character replacements.

This program benefits from the very easy-to-use [Gomoji](https://github.com/forPelevin/gomoji). Thanks forPelevin!
