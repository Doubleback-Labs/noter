# noter

noter is a simple app to open an editor of your choice assuming it has a 'app filename' command.
Files are stored in a central repo of your choosing but will default to $HOME/.noter/notes

## Install 
```
go install github.com/Doubleback-Labs/noter@latest
```

## Building
```
mage linux
```
[Mage](https://github.com/magefile/mage) 

##  Running 
```
noter                           // creates new note with DAyOnly time
noter --name new_note           // creates new note with name provided
noter ls                        // lists all notes in repo
```