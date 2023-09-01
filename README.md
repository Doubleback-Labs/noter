# noter

noter is a simple app to open a GUI (atm) editor of your choice assuming it has a 'app filename' command.
Files are stored in a central repo of your choosing.

## Install 
```
go install github.com/Doubleback-Labs/noter@v0.2.0
```

## Building
```
mage buildlinux
```

##  Running 
```
noter                           // creates new note with current time
noter --contentName new_note    // creates new note with name provided
noter ls                        // lists all notes in repo
```