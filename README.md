![Build Status](https://api.cirrus-ci.com/github/as27/reorder.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/as27/reorder)](https://goreportcard.com/report/github.com/as27/reorder)

# reorder
Renames ordered file and folder names

## Ordered files and folders

Such files starts with digits. For example `012_myFolder`  oder `056_myFile.txt`. This little programm scans the whole folder for such files and renames the digits. For the renaming you need to define a gap (by default this is 10), which will be between the elements.

## Why?

For what is this program used? If you are working on a project where the order of each element counts. This could be for example a book. Every chaptercan be a file or a folder. When you now want to reorder the structure you just have to change the order digits. To get at every element the same gap just run `reorder`. 
