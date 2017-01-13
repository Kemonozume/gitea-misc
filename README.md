# gitea-misc

gitea-misc is a mergable extension for gitea to allow login with student informations. (for hs-weingarten.de) 

# Anleitung 

Installation von gitea 
```$ go get code.gitea.io/gitea/```

Wechsel in das gitea Verzeichnis innerhalb des GOPATHS 
```$ cd $GOPATH/src/code.gitea.io/gitea/```

Hinzufügen des gitea-misc-Repositories 
```$ git remote add gitea-misc https://github.com/Kemonozume/gitea-misc.git```

Quelldateien fetchen 
```$ git fetch gitea-misc```

Mergen der gitea-misc-Dateien in das gitea-Repository 
```$ git merge -X theirs --allow-unrelated-histories gitea-misc/master```

Nun kann über ```$ go build ``` gitea kompiliert werden. 

Sollte bei ```$ go build ``` ein Fehler funktioniert das automatische Überschreiben der Dateien nicht mehr und sollte angepasst werden. 
Hierbei reicht es in der Datei login_source.go die Funktion [UserSignIn](https://github.com/Kemonozume/gitea-misc/blob/master/models/login_source.go#L548) zu bearbeiten. 

License
----

MIT
